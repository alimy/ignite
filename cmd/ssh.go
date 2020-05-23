package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	userName string
	sshPort  int16
)

func init() {
	sshCmd := &cobra.Command{
		Use:   "ssh",
		Short: "ssh to a unit",
		Long:  "ssh to a unit",
		Run:   sshRun,
	}

	// flags inflate
	sshCmd.Flags().StringVarP(&confPath, "file", "f", "Ignitefile", "config file path")
	sshCmd.Flags().StringVarP(&userName, "user", "u", "", "user name for ssh")
	sshCmd.Flags().Int16VarP(&sshPort, "port", "p", 22, "ssh port")

	// register startCmd as sub-command
	register(sshCmd)
}

func sshRun(cmd *cobra.Command, _args []string) {
	var tierName string
	flags := cmd.Flags()
	if flags.NArg() == 1 {
		tierName = flags.Arg(0)
	} else {
		cmd.Help()
		os.Exit(1)
	}
	staging := prepareStaging()
	if err := staging.SshTier(userName, tierName, sshPort); err != nil {
		logrus.Fatal(err)
	}
}
