// Copyright 2021 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package sh

import (
	"fmt"
	"strconv"

	"github.com/alimy/ignite/internal/process"
)

func Ssh(user string, host string, port int16) error {
	sshCmd := findSshCmd()
	var endpoint string
	if user != "" {
		endpoint = fmt.Sprintf("%s@%s", user, host)
	} else {
		endpoint = host
	}
	exec := &process.ExecRun{
		Describe: fmt.Sprintf("try ssh to %s on port %d", endpoint, port),
		Cmd:      sshCmd,
		Argv: []string{
			sshCmd,
			"-p",
			strconv.Itoa(int(port)),
			endpoint,
		},
		Attr: process.DefaultProcAttr(true),
	}
	return exec.Run()
}

func Scp(srcUri, dstUri string, port int16) error {
	scpCmd := findScpCmd()
	exec := &process.ExecRun{
		Describe: fmt.Sprintf("try scp %s to %s on port %d", srcUri, dstUri, port),
		Cmd:      scpCmd,
		Argv: []string{
			"-r",
			"-P",
			strconv.Itoa(int(port)),
			srcUri,
			dstUri,
		},
		Attr: process.PwdProcAttr(true),
	}
	return exec.Run()
}
