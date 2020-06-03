// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package vmware

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/alimy/ignite/internal/process"
	"github.com/alimy/ignite/internal/provision"
)

type vmwareFusion struct {
	execRun     string
	displayMode string
	stateMode   string
}

func (vm *vmwareFusion) init(config provision.ProviderConfig) {
	rootDir := config.RootDir()
	if fs, err := os.Stat(rootDir); err == nil && fs.IsDir() {
		dir, _ := filepath.EvalSymlinks(rootDir)
		vm.execRun = path.Join(dir, "Contents/Public/vmrun")
	}

	display := config.Feature("display")
	if display == "gui" || display == "nogui" {
		vm.displayMode = display
	}

	state := config.Feature("state")
	if state == "hard" || state == "soft" {
		vm.stateMode = state
	}
}

func (vm *vmwareFusion) Start(unit *provision.Unit) error {
	exec := &process.ExecRun{
		Describe: vm.actionDescribe("start", unit),
		Cmd:      vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"start",
			unit.Path,
			vm.displayMode,
		},
	}
	return exec.Run()
}

func (vm *vmwareFusion) Stop(unit *provision.Unit) error {
	exec := &process.ExecRun{
		Describe: vm.actionDescribe("stop", unit),
		Cmd:      vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"stop",
			unit.Path,
			vm.stateMode,
		},
	}
	return exec.Run()
}

func (vm *vmwareFusion) Reset(unit *provision.Unit) error {
	exec := &process.ExecRun{
		Describe: vm.actionDescribe("reset", unit),
		Cmd:      vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"reset",
			unit.Path,
			vm.stateMode,
		},
	}
	return exec.Run()
}

func (vm *vmwareFusion) Suspend(unit *provision.Unit) error {
	exec := &process.ExecRun{
		Describe: vm.actionDescribe("suspend", unit),
		Cmd:      vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"suspend",
			unit.Path,
			vm.stateMode,
		},
	}
	return exec.Run()
}

func (vm *vmwareFusion) Pause(unit *provision.Unit) error {
	exec := &process.ExecRun{
		Describe: vm.actionDescribe("pause", unit),
		Cmd:      vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"pause",
			unit.Path,
		},
	}
	return exec.Run()
}

func (vm *vmwareFusion) Unpause(unit *provision.Unit) error {
	exec := &process.ExecRun{
		Describe: vm.actionDescribe("unpause", unit),
		Cmd:      vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"unpause",
			unit.Path,
		},
	}
	return exec.Run()
}

func (vm *vmwareFusion) actionDescribe(action string, unit *provision.Unit) string {
	return fmt.Sprintf("%s tier: %s", action, unit.Name)
}

func newVMwareFusion() *vmwareFusion {
	return &vmwareFusion{
		execRun:     "/Applications/VMware Fusion.app/Contents/Public/vmrun",
		displayMode: "nogui",
		stateMode:   "hard",
	}
}
