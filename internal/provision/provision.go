// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package provision

type Staging struct {
	VMwareConfig *VMwareConfig
	Workspaces   map[string]*Workspace
}

type Workspace struct {
	Name     string
	Provider string
	Workdir  string
	Units    map[string]*Unit
}

type Unit struct {
	Name string
	Path string
}

type VMwareConfig struct {
	Name        string
	RootDir     string
	DisplayMode string
	StateMode   string
}

type Provider interface {
	Init(staging Staging)
	Start(workspace string, unit string) error
	Stop(workspace string, unit string) error
	Reset(workspace string, unit string) error
	Suspend(workspace string, unit string) error
}

func DefaultStaging() *Staging {
	return &Staging{
		VMwareConfig: &VMwareConfig{},
		Workspaces:   make(map[string]*Workspace),
	}
}
