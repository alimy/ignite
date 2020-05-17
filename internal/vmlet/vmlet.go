// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package vmlet

import (
	"github.com/alimy/ignite/internal/config"
	"github.com/alimy/ignite/internal/vmware"
)

func LetStart(path, workspace, tier string) error {
	conf, err := config.ParseFrom(path)
	if err != nil {
		return err
	}

	// TODO: init provider
	initProvider()

	staging := conf.Staging()
	return staging.Start(workspace, tier)
}

func LetStop(path, workspace, tier string) error {
	conf, err := config.ParseFrom(path)
	if err != nil {
		return err
	}

	// TODO: init provider
	initProvider()

	staging := conf.Staging()
	return staging.Stop(workspace, tier)
}

func LetReset(path, workspace, tier string) error {
	conf, err := config.ParseFrom(path)
	if err != nil {
		return err
	}

	// TODO: init provider
	initProvider()

	staging := conf.Staging()
	return staging.Reset(workspace, tier)
}

func LetSuspend(path, workspace, tier string) error {
	conf, err := config.ParseFrom(path)
	if err != nil {
		return err
	}
	// TODO: init provider
	staging := conf.Staging()
	return staging.Suspend(workspace, tier)
}

func initProvider() {
	vmware.Initialize()
}
