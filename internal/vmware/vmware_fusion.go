// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package vmware

import (
	"github.com/alimy/ignite/internal/process"
	"github.com/alimy/ignite/internal/provision"
	"github.com/sirupsen/logrus"
)

type vmwareFusion struct {
	execRun     string
	displayMode string
	stateMode   string
	workspaces  map[string]*provision.Workspace
}

func (vm *vmwareFusion) Init(staging *provision.Staging) {
	// TODO
}

func (vm *vmwareFusion) Start(workspace string, unit string) error {
	exec := &process.ExecRun{
		Cmd: vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"start",
			"",
			vm.displayMode,
		},
	}
	// TODO
	logrus.Fatal(errNotReady)
	if err := exec.Run(); err != nil {
		return err
	}
	return nil
}

func (vm *vmwareFusion) Stop(workspace string, unit string) error {
	exec := &process.ExecRun{
		Cmd: vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"stop",
			"",
			vm.displayMode,
		},
	}
	// TODO
	logrus.Fatal(errNotReady)
	if err := exec.Run(); err != nil {
		return err
	}
	return nil
}

func (vm *vmwareFusion) Reset(workspace string, unit string) error {
	exec := &process.ExecRun{
		Cmd: vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"reset",
			"",
			vm.displayMode,
		},
	}
	// TODO
	logrus.Fatal(errNotReady)
	if err := exec.Run(); err != nil {
		return err
	}
	return nil
}

func (vm *vmwareFusion) Suspend(workspace string, unit string) error {
	exec := &process.ExecRun{
		Cmd: vm.execRun,
		Argv: []string{
			vm.execRun,
			"-T",
			"fusion",
			"suspend",
			"",
			vm.displayMode,
		},
	}
	// TODO
	logrus.Fatal(errNotReady)
	if err := exec.Run(); err != nil {
		return err
	}
	return nil
}

func NewVMwareFusion(staging *provision.Staging) provision.Provider {
	vm := &vmwareFusion{}
	vm.init(staging)
	return vm
}
