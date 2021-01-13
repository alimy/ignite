// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package internal

import (
	"github.com/alimy/ignite/internal/conf"
	"github.com/alimy/ignite/internal/provision"

	_ "github.com/alimy/ignite/internal/vmware"
)

func Initialize(config *conf.IgniteConfig) {
	provision.InitProviderWith(config)
}
