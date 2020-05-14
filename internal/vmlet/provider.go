package vmlet

import "github.com/alimy/ignite/internal/provision"

var (
	providers map[string]provision.Provider
)

func providerFrom(name string, staging *provision.Staging) (provision.Provider, error) {
	p, exist := providers[name]
	if exist {
		return p, nil
	}
	p, err := provision.InitProvider(name, staging)
	if err != nil {
		return nil, err
	}
	return p, nil
}
