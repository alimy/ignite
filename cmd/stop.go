// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "stop a workspace",
		Long:  "stop a workspace",
		Run:   stopRun,
	}

	// flags inflate
	stopCmd.Flags().StringVarP(&confPath, "file", "f", "", "config file path")
	stopCmd.Flags().BoolVar(&allWorkspace, "all", false, "whether process all workspace")

	// register stopCmd as sub-command
	register(stopCmd)
}

func stopRun(cmd *cobra.Command, _args []string) {
	if err := checkConfigFile(); err != nil {
		logrus.Fatal(err)
	}
	w, t := workspaceTier(cmd)
	staging := prepareStaging()
	if err := staging.Stop(w, t); err != nil {
		logrus.Fatal(err)
	}
}
