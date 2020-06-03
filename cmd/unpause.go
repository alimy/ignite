// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	unpauseCmd := &cobra.Command{
		Use:   "unpause",
		Short: "unpause a workspace",
		Long:  "unpause a workspace",
		Run:   unpauseRun,
	}

	// flags inflate
	unpauseCmd.Flags().StringVarP(&confPath, "file", "f", "", "config file path")
	unpauseCmd.Flags().BoolVar(&allWorkspace, "all", false, "whether process all workspace")

	// register unpauseCmd as sub-command
	register(unpauseCmd)
}

func unpauseRun(cmd *cobra.Command, _args []string) {
	if err := checkConfigFile(); err != nil {
		logrus.Fatal(err)
	}
	w, t := workspaceTier(cmd)
	staging := prepareStaging()
	if err := staging.Unpause(w, t); err != nil {
		logrus.Fatal(err)
	}
}
