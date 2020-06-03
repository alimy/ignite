// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	workspacesCmd := &cobra.Command{
		Use:   "workspaces",
		Short: "list all workspace info",
		Long:  "list all workspace info",
		Run:   workspacesRun,
	}

	// flags inflate
	workspacesCmd.Flags().StringVarP(&confPath, "file", "f", "", "config file path")

	// register workspacesCmd as sub-command
	register(workspacesCmd)
}

func workspacesRun(cmd *cobra.Command, _args []string) {
	if err := checkConfigFile(); err != nil {
		logrus.Fatal(err)
	}
	staging := prepareStaging()
	if err := staging.WorkspacesInfo(); err != nil {
		logrus.Fatal(err)
	}
}
