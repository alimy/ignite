package cmd

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	confPath string
)

var (
	optConfPath string
	optCmd      string
	optGui      bool
	optSoft     bool
	optWs       bool

	errInvalideClusterName = errors.New("invalide cluster name")
)

func argsFixed() (string, string, string) {
	argGui := "nogui"
	argHard := "hard"
	argFusion := "fusion"
	if optGui {
		argGui = "gui"
	}
	if optSoft {
		argHard = "sort"
	}
	if optWs {
		argFusion = "ws"
	}
	return argGui, argHard, argFusion
}

func clusterInfos(cmd *cobra.Command) []*cluster {
	clusters, err := parseFrom(optConfPath)
	if err != nil {
		logrus.Fatal(err)
	}

	// set target cluster name
	if cmd.Flags().NArg() > 0 {
		name := cmd.Flags().Arg(0)
		c, exist := clusters[name]
		if !exist {
			logrus.Fatal(errInvalideClusterName)
		}
		return []*cluster{
			c,
		}
	}

	// process all cluster
	cs := make([]*cluster, 0, len(clusters))
	for _, c := range clusters {
		cs = append(cs, c)
	}
	return cs
}
