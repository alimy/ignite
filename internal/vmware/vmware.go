// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package vmware

import "github.com/alimy/ignite/internal/provision"

func Initialize() {
	// TODO init vmwareFusion from config
	vm := &vmwareFusion{}
	provision.Register("vmware-fusion", vm)
}
