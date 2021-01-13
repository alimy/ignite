// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package sh

import "sync"

var (
	_sshCmd string
	_scpCmd string

	onceSsh sync.Once
	onceScp sync.Once
)

func findSshCmd() string {
	onceSsh.Do(func() {
		// TODO: find ssh abs path
		_sshCmd = "/usr/bin/ssh"
	})
	return _sshCmd
}

func findScpCmd() string {
	onceScp.Do(func() {
		// TODO: find scp abs path
		_scpCmd = "/usr/bin/scp"
	})
	return _scpCmd
}
