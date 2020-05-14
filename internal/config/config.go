// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package config

import (
	"errors"

	"github.com/alimy/ignite/internal/provision"
)

var (
	errNotExistConfigFile = errors.New("not exist config file")

	igniteConfig *IgniteConfig
)

type IgniteConfig struct {
	// TODO
}

func MyConfig() *IgniteConfig {
	return igniteConfig
}

func (c *IgniteConfig) Staging() *provision.Staging {
	staging := provision.DefaultStaging()
	return staging
}

func ParseFrom(path string) (*IgniteConfig, error) {
	// TODO
	return nil, errNotExistConfigFile
}
