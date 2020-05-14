package process

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

type ExecRun struct {
	Describe string
	Cmd      string
	Argv     []string
}

func (r *ExecRun) Run() error {
	logrus.Info(r.Describe)
	process, err := os.StartProcess(r.Cmd, r.Argv, nil)
	if err != nil {
		return err
	}
	ps, err := process.Wait()
	if err != nil {
		return err
	}
	if !ps.Success() {
		return errors.New(ps.String())
	}
	return nil
}
