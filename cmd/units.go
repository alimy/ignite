// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	unitsCmd := &cobra.Command{
		Use:   "units",
		Short: "list all unit info",
		Long:  "list all unit info",
		Run:   unitsRun,
	}

	// flags inflate
	unitsCmd.Flags().StringVarP(&confPath, "file", "f", "", "config file path")

	// register unitsCmd as sub-command
	register(unitsCmd)
}

func unitsRun(cmd *cobra.Command, _args []string) {
	if err := checkConfigFile(); err != nil {
		logrus.Fatal(err)
	}
	staging := prepareStaging()
	if err := staging.UnitsInfo(); err != nil {
		logrus.Fatal(err)
	}
}
