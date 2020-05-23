// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	suspendCmd := &cobra.Command{
		Use:   "suspend",
		Short: "suspend a workspace",
		Long:  "suspend a workspace",
		Run:   suspendRun,
	}

	// flags inflate
	suspendCmd.Flags().StringVarP(&confPath, "file", "f", "", "config file path")
	suspendCmd.Flags().BoolVar(&allWorkspace, "all", false, "whether process all workspace")

	// register suspendCmd as sub-command
	register(suspendCmd)
}

func suspendRun(cmd *cobra.Command, _args []string) {
	if err := checkConfigFile(); err != nil {
		logrus.Fatal(err)
	}
	w, t := workspaceTier(cmd)
	staging := prepareStaging()
	if err := staging.Suspend(w, t); err != nil {
		logrus.Fatal(err)
	}
}
