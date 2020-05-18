// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package provision

import (
	"sync"

	"github.com/alimy/ignite/internal/config"
)

var (
	providers       = make(map[string]providerInstance)
	providerConfigs = make(map[string]ProviderConfig)
)

type providerInstance struct {
	factory  ProviderFactory
	instance Provider
	once     sync.Once
}

func (p *providerInstance) Get() Provider {
	p.once.Do(func() {
		c := providerConfigs[p.factory.Name()]
		p.instance = p.factory.NewProvider(c)
	})
	return p.instance
}

func FindProviderByName(name string) Provider {
	if p, exist := providers[name]; exist {
		return p.Get()
	}
	return nil
}

func Register(factory ProviderFactory) {
	if factory != nil {
		providers[factory.Name()] = providerInstance{
			factory: factory,
		}
	}
}

func InitProviderWith(config *config.IgniteConfig) {
	// TODO
}
