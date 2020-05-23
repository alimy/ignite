// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package provision

import (
	"github.com/alimy/ignite/internal/process/ssh"
	"github.com/alimy/ignite/internal/xerror"
	"github.com/sirupsen/logrus"
)

func (t *Unit) Start() error {
	provider := FindProviderByName(t.Provider)
	if provider == nil {
		return xerror.ErrProviderNotSupported
	}
	return provider.Start(t)
}

func (t *Unit) Stop() error {
	provider := FindProviderByName(t.Provider)
	if provider == nil {
		return xerror.ErrProviderNotSupported
	}
	return provider.Stop(t)
}

func (t *Unit) Reset() error {
	provider := FindProviderByName(t.Provider)
	if provider == nil {
		return xerror.ErrProviderNotSupported
	}
	return provider.Reset(t)
}

func (t *Unit) Suspend() error {
	provider := FindProviderByName(t.Provider)
	if provider == nil {
		return xerror.ErrProviderNotSupported
	}
	return provider.Suspend(t)
}

func (t *Unit) Ssh(user string, port int16) error {
	for _, host := range t.Hosts {
		if err := ssh.Run(user, host.Name, port); err != nil {
			logrus.Warn(err)
			continue
		}
		break
	}
	return nil
}
