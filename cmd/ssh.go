// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	userName string
	sshPort  int16
)

func init() {
	sshCmd := &cobra.Command{
		Use:   "ssh",
		Short: "ssh to a tier",
		Long:  "ssh to a tier",
		Run:   sshRun,
	}

	// flags inflate
	sshCmd.Flags().StringVarP(&confPath, "file", "f", "", "config file path")
	sshCmd.Flags().StringVarP(&userName, "user", "u", "", "user name for ssh")
	sshCmd.Flags().Int16VarP(&sshPort, "port", "p", 22, "ssh port")

	// register startCmd as sub-command
	register(sshCmd)
}

func sshRun(cmd *cobra.Command, _args []string) {
	if err := checkConfigFile(); err != nil {
		logrus.Fatal(err)
	}

	var tierName string
	flags := cmd.Flags()
	if flags.NArg() == 1 {
		tierName = flags.Arg(0)
	} else {
		cmd.Help()
		os.Exit(1)
	}

	// handle ignite ssh user@tiername
	if userIdx := strings.Index(tierName, "@"); userIdx > 0 {
		if userName == "" {
			userName = tierName[:userIdx]
		}
		tierName = tierName[userIdx+1:]
	}

	staging := prepareStaging()
	if err := staging.SshTier(userName, tierName, sshPort); err != nil {
		logrus.Fatal(err)
	}
}
