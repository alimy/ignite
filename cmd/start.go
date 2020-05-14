// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"github.com/alimy/ignite/internal/vmlet"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "start a workspace",
		Long:  "start a workspace",
		Run:   startRun,
	}

	// flags inflate
	startCmd.Flags().StringVarP(&confPath, "config", "c", "ignite.yml", "config file path")

	// register agentCmd as sub-command
	register(startCmd)
}

func startRun(cmd *cobra.Command, _args []string) {
	var workspace, unit string
	flags := cmd.Flags()
	if flags.NArg() > 1 {
		workspace, unit = flags.Arg(0), flags.Arg(1)
	} else if flags.NArg() == 1 {
		workspace = flags.Arg(0)
	}
	if err := vmlet.LetStart(confPath, workspace, unit); err != nil {
		logrus.Fatal(err)
	}
}
