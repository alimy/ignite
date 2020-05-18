// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package internal

import (
	"github.com/alimy/ignite/internal/config"
	"github.com/alimy/ignite/internal/provision"

	_ "github.com/alimy/ignite/internal/vmware"
)

func Initialize(config *config.IgniteConfig) {
	provision.InitProviderWith(config)
}
