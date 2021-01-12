// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	pauseCmd := &cobra.Command{
		Use:        "pause",
		Aliases:    []string{"p"},
		SuggestFor: []string{"unpause"},
		Short:      "pause a workspace",
		Long:       "pause a workspace",
		Run:        pauseRun,
	}

	// flags inflate
	pauseCmd.Flags().StringVarP(&confPath, "file", "f", "", "config file path")
	pauseCmd.Flags().BoolVar(&allWorkspace, "all", false, "whether process all workspace")

	// register pauseCmd as sub-command
	register(pauseCmd)
}

func pauseRun(cmd *cobra.Command, _args []string) {
	if err := checkConfigFile(); err != nil {
		logrus.Fatal(err)
	}
	w, t := workspaceTier(cmd)
	staging := prepareStaging()
	if err := staging.Pause(w, t); err != nil {
		logrus.Fatal(err)
	}
}
