// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package vmware

import (
	"github.com/alimy/ignite/internal/process"
	"github.com/alimy/ignite/internal/provision"
	"github.com/alimy/ignite/internal/xerror"
	"github.com/sirupsen/logrus"
)

type vmwareFusion struct {
	execRun     string
	displayMode string
	stateMode   string
}

func (vm *vmwareFusion) init(config provision.ProviderConfig) {

}

func (vm *vmwareFusion) Start(unit *provision.Unit) error {
	exec := &process.ExecRun{
		Cmd: vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"start",
			unit.Path,
			vm.displayMode,
		},
	}
	// TODO
	logrus.Fatal(xerror.ErrNotReady)
	if err := exec.Run(); err != nil {
		return err
	}
	return nil
}

func (vm *vmwareFusion) Stop(unit *provision.Unit) error {
	exec := &process.ExecRun{
		Cmd: vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"stop",
			unit.Path,
			vm.displayMode,
		},
	}
	// TODO
	logrus.Fatal(xerror.ErrNotReady)
	if err := exec.Run(); err != nil {
		return err
	}
	return nil
}

func (vm *vmwareFusion) Reset(unit *provision.Unit) error {
	exec := &process.ExecRun{
		Cmd: vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"reset",
			unit.Path,
			vm.displayMode,
		},
	}
	// TODO
	logrus.Fatal(xerror.ErrNotReady)
	if err := exec.Run(); err != nil {
		return err
	}
	return nil
}

func (vm *vmwareFusion) Suspend(unit *provision.Unit) error {
	exec := &process.ExecRun{
		Cmd: vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"suspend",
			unit.Path,
			vm.displayMode,
		},
	}
	// TODO
	logrus.Fatal(xerror.ErrNotReady)
	if err := exec.Run(); err != nil {
		return err
	}
	return nil
}

func newVMwareFusion() *vmwareFusion {
	return &vmwareFusion{
		execRun:     "/Applications/VMware Fusion.app/Contents/Public/vmrun",
		displayMode: "nogui",
		stateMode:   "hard",
	}
}
