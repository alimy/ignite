// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package vmlet

import (
	"errors"

	"github.com/alimy/ignite/internal/config"
	"github.com/alimy/ignite/internal/provision"
	"github.com/sirupsen/logrus"
)

var (
	errNoExistWorkspace = errors.New("no exist workspace")
)

func LetStart(path, workspace, unit string) error {
	conf, err := config.ParseFrom(path)
	if err != nil {
		logrus.Fatal(err)
	}
	staging := conf.Staging()
	if workspace != "" {
		w, exist := staging.Workspaces[workspace]
		if !exist {
			return errNoExistWorkspace
		}
		if err := doStart(w, staging, workspace, unit); err != nil {
			return err
		}
		return nil
	}
	for _, w := range staging.Workspaces {
		if err := doStart(w, staging, workspace, unit); err != nil {
			return err
		}
	}
	return nil
}

func doStart(w *provision.Workspace, staging *provision.Staging, workspace string, unit string) error {
	provider, err := providerFrom(w.Provider, staging)
	if err != nil {
		return err
	}
	return provider.Start(workspace, unit)
}

func LetStop(path string) error {
	// TODO
	return nil
}

func LetReset(path string) error {
	// TODO
	return nil
}

func LetSuspend(path string) error {
	// TODO
	return nil
}
