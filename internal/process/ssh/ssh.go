package ssh

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/alimy/ignite/internal/process"
)

var (
	sshCmd string
	once   sync.Once
)

func Run(user string, host string, port int16) error {
	once.Do(func() {
		// TODO: find ssh abs path
		sshCmd = "/usr/bin/ssh"
	})
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
