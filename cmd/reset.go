// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	resetCmd := &cobra.Command{
		Use:   "reset",
		Short: "reset a workspace",
		Long:  "reset a workspace",
		Run:   resetRun,
	}

	// flags inflate
	resetCmd.Flags().StringVarP(&confPath, "config", "c", "ignite.yml", "config file path")

	// register agentCmd as sub-command
	register(resetCmd)
}

func resetRun(cmd *cobra.Command, _args []string) {
	clusters := clusterInfos(cmd)
	_, argHard, argFusion := argsFixed()
	ci := &cmdInfo{
		cmd: optCmd,
		argv: []string{
			optCmd,
			"-T",
			argFusion,
			"reset",
			"",
			argHard,
		},
	}
	for _, cluster := range clusters {
		logrus.Infof("reset cluster %s\n", cluster.name)
		for _, node := range cluster.nodes {
			ci.describe = fmt.Sprintf("reset node %s", node.name)
			ci.argv[4] = node.path
			if err := runCmd(ci); err != nil {
				logrus.Fatal(err)
			}
		}
	}
}
