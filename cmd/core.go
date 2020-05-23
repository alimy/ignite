// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/alimy/ignite/internal"
	"github.com/alimy/ignite/internal/config"
	"github.com/alimy/ignite/internal/provision"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	confPath     string
	allWorkspace bool
)

func workspaceTier(cmd *cobra.Command) (string, string) {
	var workspace, tier string
	flags := cmd.Flags()
	if flags.NArg() > 1 {
		workspace, tier = flags.Arg(0), flags.Arg(1)
	} else if flags.NArg() == 1 {
		workspace = flags.Arg(0)
	} else {
		if allWorkspace {
			cmd.Help()
			os.Exit(1)
		}
	}
	return workspace, tier
}

func prepareStaging() *provision.Staging {
	conf, err := config.ParseFrom(confPath)
	if err != nil {
		logrus.Fatal(err)
	}
	internal.Initialize(conf)
	return provision.StagingFrom(conf)
}

func checkConfigFile() error {
	if confPath != "" {
		return nil
	}
	if fi, err := os.Stat("Ignitefile"); err == nil && !fi.IsDir() {
		confPath = "Ignitefile"
	} else {
		homeDir, _ := os.UserHomeDir()
		path := filepath.Join(homeDir, ".ignite/Ignitefile")
		if fi, err := os.Stat(path); err == nil && !fi.IsDir() {
			confPath = path
		} else {
			return errors.New("no exist Ignitefile in ./Ignitefile or ~/.ignite/Ignitefile of -f <Ignitefile>")
		}
	}
	return nil
}
