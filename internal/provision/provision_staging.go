// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package provision

import (
	"errors"

	"github.com/alimy/ignite/internal/xerror"
)

const (
	ActDone State = iota
	ActFail
)

var (
	errNoExistWorkspace = errors.New("no exist workspace")
)

type State = int8

type StateMsg struct {
	Act      string
	UnitName string
	Error    error
}

func (s *Staging) Start(workspace string, tier string) error {
	// TODO
	return xerror.ErrNotReady
}

func (s *Staging) Stop(workspace string, tier string) error {
	// TODO
	return xerror.ErrNotReady
}

func (s *Staging) Reset(workspace string, tier string) error {
	// TODO
	return xerror.ErrNotReady
}

func (s *Staging) Suspend(workspace string, tier string) error {
	// TODO
	return xerror.ErrNotReady
}
