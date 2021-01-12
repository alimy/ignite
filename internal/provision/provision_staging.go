// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package provision

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/alimy/ignite/internal/config"
	"github.com/alimy/ignite/internal/terminal"
	"github.com/sirupsen/logrus"
)

var (
	errNotExistWorkspace = errors.New("not exist workspace")
	errNotExistTier      = errors.New("not exist tier")

	actionStart   = "start"
	actionStop    = "stop"
	actionReset   = "reset"
	actionSuspend = "suspend"
	actionPause   = "pause"
	actionUnpause = "unpause"
)

type StateMsg struct {
	State TierState
	Tier  *Tier
	Error error
}

type Staging struct {
	retryNum        int
	Fallback        bool
	DefaultProvider string
	Tiers           map[string]*Tier
	Workspaces      map[string]*Workspace
}

type handler struct {
	Name string
	Func func(*Unit) error
}

func handleFun(name string, h func(*Unit) error) *handler {
	return &handler{
		Name: name,
		Func: h,
	}
}

func (s *Staging) Start(workspace string, tier string) error {
	return s.run(actionStart, workspace, tier)
}

func (s *Staging) Stop(workspace string, tier string) error {
	return s.run(actionStop, workspace, tier)
}

func (s *Staging) Reset(workspace string, tier string) error {
	return s.run(actionReset, workspace, tier)
}

func (s *Staging) Suspend(workspace string, tier string) error {
	return s.run(actionSuspend, workspace, tier)
}

func (s *Staging) Pause(workspace string, tier string) error {
	return s.run(actionPause, workspace, tier)
}

func (s *Staging) Unpause(workspace string, tier string) error {
	return s.run(actionUnpause, workspace, tier)
}

func (s *Staging) WorkspacesInfo() error {
	ti := terminal.NewTableInfo("Name", "Tiers", "Description")
	for name, workspace := range s.Workspaces {
		tiers := strconv.Itoa(len(workspace.Tiers))
		ti.Add(name, tiers, workspace.Description)
	}
	terminal.PrintTable(ti)
	return nil
}

func (s *Staging) TiersInfo(workspace string) error {
	ti := terminal.NewTableInfo("Tier", "Hosts", "State", "Description")
	if ws, exist := s.Workspaces[workspace]; exist {
		ti.Infos(
			fmt.Sprintf("Name: %s\t Tiers: %d", workspace, len(ws.Tiers)),
			fmt.Sprintf("Description: %s", ws.Description),
		)
		for name, tier := range ws.Tiers {
			hosts := make([]string, len(tier.Hosts))
			for i, host := range tier.Hosts {
				hosts[i] = host.Name
			}
			ti.Add(name, strings.Join(hosts, ","), tier.ActiveState(), tier.Description)
		}
	} else {
		return errNotExistWorkspace
	}
	terminal.PrintTable(ti)
	return nil
}

func (s *Staging) UnitsInfo() error {
	ti := terminal.NewTableInfo("Name", "Provider", "Description")
	conf := config.MyConfig()
	for _, unit := range conf.Units {
		provider := unit.Provider
		if provider == "" {
			provider = s.DefaultProvider
		}
		ti.Add(unit.Name, provider, unit.Description)
	}
	terminal.PrintTable(ti)
	return nil
}

func (s *Staging) SshTier(userName string, tierName string, sshPort int16) error {
	tier, exist := s.Tiers[tierName]
	if !exist {
		return fmt.Errorf("fail search tier for %s: %w", tierName, errNotExistTier)
	}
	return tier.Ssh(userName, sshPort)
}

func (s *Staging) run(action string, workspace string, tier string) error {
	var h *handler
	switch action {
	case actionStart:
		h = handleFun("start", func(unit *Unit) error {
			return unit.Start()
		})
	case actionStop:
		h = handleFun("stop", func(unit *Unit) error {
			return unit.Stop()
		})
	case actionReset:
		h = handleFun("reset", func(unit *Unit) error {
			return unit.Reset()
		})
	case actionSuspend:
		h = handleFun("suspend", func(unit *Unit) error {
			return unit.Suspend()
		})
	case actionPause:
		h = handleFun("pause", func(unit *Unit) error {
			return unit.Pause()
		})
	case actionUnpause:
		h = handleFun("unpause", func(unit *Unit) error {
			return unit.Unpause()
		})
	}
	if workspace == "" {
		return s.handleAllWorkspace(h)
	}
	return s.handleWorkspace(workspace, tier, h)
}

func (s *Staging) tiersBy(workspace string, tier string) (map[string]*Tier, error) {
	ws, exist := s.Workspaces[workspace]
	if !exist {
		return nil, errNotExistWorkspace
	}
	if tier != "" {
		tier, exist := ws.Tiers[tier]
		if !exist {
			return nil, errNotExistTier
		}
		return map[string]*Tier{tier.Name: tier}, nil
	}
	return ws.Tiers, nil
}

func (s *Staging) handleTiers(tiers map[string]*Tier, h *handler) (err error) {
	switch h.Name {
	case actionStart, actionPause:
		err = s.handleTiersAsc(tiers, h)
	case actionStop, actionReset, actionSuspend, actionUnpause:
		err = s.handleTiersDesc(tiers, h)
	}
	return
}

// handleTiersAsc 正序处理tiers
func (s *Staging) handleTiersAsc(tiers map[string]*Tier, h *handler) error {
	stateChan := make(chan *StateMsg, 8)
	remainTiersCount := len(tiers)
	for _, tier := range tiers {
		if len(tier.Parents) == 0 {
			go s.handleTier(stateChan, tier, h)
		}
	}
	for sc := range stateChan {
		if sc.State == TierStateFailed {
			return fmt.Errorf("%s tier: %s but failed by %w", h.Name, sc.Tier.Name, sc.Error)
		} else if sc.State == TierStateInactive {
			logrus.Warnf("%s tier: %s but skiped by inactive state", h.Name, sc.Tier.Name)
		}
		if remainTiersCount--; remainTiersCount == 0 { // handler finish
			break
		}
		children := sc.Tier.Children
		for name := range children {
			tier := tiers[name]
			tier.SetParentState(sc.Tier.Name, TierStateDone)
			if tier.IsParentsDone() {
				go s.handleTier(stateChan, tier, h)
			}
		}
	}
	return nil
}

// handleTiersAsc 反序处理tiers
func (s *Staging) handleTiersDesc(tiers map[string]*Tier, h *handler) error {
	stateChan := make(chan *StateMsg, 8)
	remainTiersCount := len(tiers)
	for _, tier := range tiers {
		if len(tier.Children) == 0 {
			go s.handleTier(stateChan, tier, h)
		}
	}
	for sc := range stateChan {
		if sc.State == TierStateFailed {
			return fmt.Errorf("%s tier: %s but failed by %w", h.Name, sc.Tier.Name, sc.Error)
		} else if sc.State == TierStateInactive {
			logrus.Warnf("%s tier: %s but skiped by inactive state", h.Name, sc.Tier.Name)
		}
		if remainTiersCount--; remainTiersCount == 0 { // handler finish
			break
		}
		parents := sc.Tier.Parents
		for name := range parents {
			tier := tiers[name]
			tier.SetChildState(sc.Tier.Name, TierStateDone)
			if tier.IsChildrenDone() {
				go s.handleTier(stateChan, tier, h)
			}
		}
	}
	return nil
}

func (s *Staging) handleWorkspace(workspace string, tier string, h *handler) error {
	tiers, err := s.tiersBy(workspace, tier)
	if err != nil {
		return err
	}
	if tier != "" {
		logrus.Infof("%s workspace: %s tier: %s\n", h.Name, workspace, tier)
	} else {
		logrus.Infof("%s workspace: %s\n", h.Name, workspace)
	}
	return s.handleTiers(tiers, h)
}

func (s *Staging) handleAllWorkspace(h *handler) error {
	for name := range s.Workspaces {
		logrus.Infof("%s workspace: %s\n", h.Name)
		tiers, _ := s.tiersBy(name, "")
		if err := s.handleTiers(tiers, h); err != nil {
			return err
		}
	}
	return nil
}

func (s *Staging) handleTier(stateChan chan<- *StateMsg, tier *Tier, h *handler) {
	if tier.Inactive {
		stateChan <- &StateMsg{
			State: TierStateInactive,
			Tier:  tier,
		}
		return
	}

	var err error
	for i := 0; i <= s.retryNum; i++ {
		if err = h.Func(tier.Unit); err == nil {
			break
		} else if i != s.retryNum {
			logrus.Warnf("retry[%d] %s tier of %s because %s", i+1, h.Name, tier.Name, err)
			time.Sleep(500 * time.Millisecond)
		}
	}
	if err != nil {
		stateChan <- &StateMsg{
			State: TierStateFailed,
			Tier:  tier,
			Error: err,
		}
	} else {
		stateChan <- &StateMsg{
			State: TierStateDone,
			Tier:  tier,
		}
	}
}

func DefaultStaging() *Staging {
	return &Staging{
		retryNum:        1,
		DefaultProvider: "vmware-fusion",
		Tiers:           make(map[string]*Tier),
		Workspaces:      make(map[string]*Workspace),
	}
}

func StagingFrom(config *config.IgniteConfig) *Staging {
	staging := DefaultStaging()
	if config == nil {
		return staging
	}
	staging.Fallback = config.Staging.Fallback
	if config.Staging.DefaultProvider != "" {
		staging.DefaultProvider = config.Staging.DefaultProvider
	}
	units := unitsFrom(config, staging.DefaultProvider)
	for _, ws := range config.Workspaces {
		workspace := &Workspace{
			Name:        ws.Name,
			Description: ws.Description,
		}
		tiers := make(map[string]*Tier, len(ws.TierList)+len(ws.Tiers))
		for _, name := range ws.TierList {
			if _, exist := tiers[name]; exist {
				continue
			}
			unit, exist := units[name]
			if !exist {
				logrus.Warnf("not exist unit named %s\n", name)
				continue
			}
			tiers[name] = &Tier{
				Unit: unit,
			}
		}
		for _, ts := range ws.Tiers {
			unit, exist := units[ts.Name]
			if !exist {
				logrus.Warnf("not exist unit named %s\n", ts.Name)
				continue
			}
			tmpTiers := make(map[string]*Tier, len(ts.Dependencies))
			for _, dn := range ts.Dependencies {
				unit, exist := units[dn]
				if !exist {
					logrus.Warnf("not exist unit named %s\n", dn)
					continue
				}
				tier, exist := tiers[dn]
				if !exist {
					tmpTiers[dn] = &Tier{
						Unit: unit,
					}
				}
				if tier.Children == nil {
					tier.Children = make(map[string]TierState)
				}
				tier.Children[ts.Name] = TierStateUnknown
			}
			for _, tmpTier := range tmpTiers {
				tiers[tmpTier.Name] = tmpTier
			}
			tier, exist := tiers[ts.Name]
			if !exist {
				tier = &Tier{
					Unit: unit,
				}
				tiers[ts.Name] = tier
			}
			if tier.Parents == nil {
				tier.Parents = make(map[string]TierState, len(ts.Dependencies))
			}
			for _, name := range ts.Dependencies {
				tier.Parents[name] = TierStateUnknown
			}
			tier.Inactive = ts.Inactive
		}
		workspace.Tiers = tiers
		staging.Workspaces[ws.Name] = workspace
		for name, tier := range tiers {
			staging.Tiers[name] = tier
		}
	}
	return staging
}

func unitsFrom(config *config.IgniteConfig, defaultProvider string) map[string]*Unit {
	units := make(map[string]*Unit, len(config.Units))
	for _, spec := range config.Units {
		unit := &Unit{
			Name:        spec.Name,
			Description: spec.Description,
			Path:        fixedPath(spec.Path),
		}
		if spec.Provider != "" {
			unit.Provider = spec.Provider
		} else {
			unit.Provider = defaultProvider
		}
		unit.Hosts = make([]Host, 0, len(spec.Hosts))
		for _, name := range spec.Hosts {
			unit.Hosts = append(unit.Hosts, Host{
				Name: name,
			})
		}
		units[spec.Name] = unit
	}
	return units
}

func fixedPath(path string) string {
	if path[0] == '~' {
		homedir, err := os.UserHomeDir()
		if err != nil {
			return path
		}
		return filepath.Join(homedir, path[1:])
	}
	if absPath, err := filepath.Abs(path); err == nil {
		return absPath
	}
	return path
}
