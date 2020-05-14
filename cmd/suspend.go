package cmd

import (
	"fmt"

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
	suspendCmd.Flags().StringVarP(&confPath, "config", "c", "ignite.yml", "config file path")

	// register suspendCmd as sub-command
	register(suspendCmd)
}

func suspendRun(cmd *cobra.Command, _args []string) {
	clusters := clusterInfos(cmd)
	_, argHard, argFusion := argsFixed()
	ci := &cmdInfo{
		cmd: optCmd,
		argv: []string{
			optCmd,
			"-T",
			argFusion,
			"suspend",
			"",
			argHard,
		},
	}
	for _, cluster := range clusters {
		logrus.Infof("suspend cluster %s\n", cluster.name)
		for _, node := range cluster.nodes {
			ci.describe = fmt.Sprintf("suspend node %s", node.name)
			ci.argv[4] = node.path
			if err := runCmd(ci); err != nil {
				logrus.Fatal(err)
			}
		}
	}
}
