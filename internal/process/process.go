// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

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
	Attr     *os.ProcAttr
}

func (r *ExecRun) Run() error {
	logrus.Info(r.Describe)
	r.checkProcAttr()
	process, err := os.StartProcess(r.Cmd, r.Argv, r.Attr)
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

func (r *ExecRun) checkProcAttr() {
	if r.Attr == nil {
		r.Attr = &os.ProcAttr{}
		if homedir, err := os.UserHomeDir(); err == nil {
			r.Attr.Dir = homedir
		}
	}
}

func DefaultProcAttr(inStdio bool) *os.ProcAttr {
	attr := &os.ProcAttr{}
	if inStdio {
		attr.Files = []*os.File{
			os.Stdin, os.Stdout, os.Stderr,
		}
	}
	if homedir, err := os.UserHomeDir(); err == nil {
		attr.Dir = homedir
	}
	return attr
}

func PwdProcAttr(inStdio bool) *os.ProcAttr {
	attr := &os.ProcAttr{
		Dir: os.Getenv("PWD"),
	}
	if inStdio {
		attr.Files = []*os.File{
			os.Stdin, os.Stdout, os.Stderr,
		}
	}
	return attr
}
