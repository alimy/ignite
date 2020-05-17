// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package provision

type Staging struct {
	Fallback   bool
	Workspaces map[string]*Workspace
}

type Workspace struct {
	Description string
	Name        string
	Tiers       map[string]*Tier
}

type Tier struct {
	*Unit
	Parents  []string
	Children []string
}

type Unit struct {
	Description string
	Name        string
	Provider    string
	Path        string
	Hosts       []Host
}

type Host struct {
	Name string
}

type Provider interface {
	Start(*Unit) error
	Stop(*Unit) error
	Reset(*Unit) error
	Suspend(*Unit) error
}

func DefaultStaging() *Staging {
	return &Staging{
		Workspaces: make(map[string]*Workspace),
	}
}
