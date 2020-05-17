// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	confPath string
)

func workspaceTier(cmd *cobra.Command) (string, string) {
	var workspace, tier string
	flags := cmd.Flags()
	if flags.NArg() > 1 {
		workspace, tier = flags.Arg(0), flags.Arg(1)
	} else if flags.NArg() == 1 {
		workspace = flags.Arg(0)
	} else {
		cmd.Help()
		os.Exit(1)
	}
	return workspace, tier
}
