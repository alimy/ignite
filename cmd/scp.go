// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/alimy/ignite/internal/provision"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	sshCmd := &cobra.Command{
		Use:   "scp",
		Short: "copy files between host and tier",
		Long:  "copy files between host and tier",
		Run:   scpRun,
	}

	// flags inflate
	sshCmd.Flags().StringVarP(&confPath, "file", "f", "", "config file path")
	sshCmd.Flags().StringVarP(&userName, "user", "u", "", "user name for ssh")
	sshCmd.Flags().Int16VarP(&sshPort, "port", "p", 22, "ssh port")

	// register startCmd as sub-command
	register(sshCmd)
}

func scpRun(cmd *cobra.Command, _args []string) {
	err := checkConfigFile()
	if err != nil {
		logrus.Fatal(err)
	}

	var (
		src, dst []string
		staging  *provision.Staging
	)
	isSrcInTier, isDstInTier := false, true
	flags := cmd.Flags()
	if flags.NArg() == 2 {
		srcFiles, dstFiles := flags.Arg(0), flags.Arg(1)
		if src, isSrcInTier, err = parseUri(srcFiles); err != nil {
			logrus.Error(err)
			goto EXIT
		}
		if dst, isDstInTier, err = parseUri(dstFiles); err != nil {
			logrus.Error(err)
			goto EXIT
		}
		if isSrcInTier == isDstInTier {
			logrus.Error("just have one target in tier")
			goto EXIT
		}
	} else {
		goto EXIT
	}

	staging = prepareStaging()
	if err = staging.ScpTier(src, dst, sshPort); err != nil {
		logrus.Fatal(err)
	}
	return

EXIT:
	cmd.Help()
	os.Exit(1)
}

func parseUri(target string) ([]string, bool, error) {
	if !strings.ContainsAny(target, "@:") {
		return []string{target}, false, nil
	}
	// handle ignite scp user@tiername:path
	ts := strings.Split(target, ":")
	if len(ts) != 2 {
		return nil, true, errors.New("target uri not correct")
	}
	tierName := ts[0]
	if userIdx := strings.Index(tierName, "@"); userIdx > 0 {
		if userName == "" {
			userName = tierName[:userIdx]
		}
		tierName = tierName[userIdx+1:]
	}
	if userName == "" {
		return nil, true, errors.New("not set user name")
	}
	return []string{userName, tierName, ts[1]}, true, nil
}
