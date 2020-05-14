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
	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "stop a vmware node",
		Long:  "stop a vmware node",
		Run:   stopRun,
	}

	// flags inflate
	stopCmd.Flags().StringVarP(&confPath, "config", "c", "ignite.yml", "config file path")

	// register stopCmd as sub-command
	register(stopCmd)
}

func stopRun(cmd *cobra.Command, _args []string) {
	clusters := clusterInfos(cmd)
	argGui, _, argFusion := argsFixed()
	ci := &cmdInfo{
		cmd: optCmd,
		argv: []string{
			optCmd,
			"-T",
			argFusion,
			"stop",
			"",
			argGui,
		},
	}
	for _, cluster := range clusters {
		logrus.Infof("stop cluster %s\n", cluster.name)
		for _, node := range cluster.nodes {
			ci.describe = fmt.Sprintf("stop node %s", node.name)
			ci.argv[4] = node.path
			if err := runCmd(ci); err != nil {
				logrus.Fatal(err)
			}
		}
	}
}
