// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package provision

import "errors"

var (
	providers map[string]Provider

	errNoRegisteredProvider = errors.New("no registered provider")
)

func Register(name string, provider Provider) {
	providers[name] = provider
}
