package provision

import "errors"

var (
	providers map[string]Provider

	errNoRegisteredProvider = errors.New("no registered provider")
)

func InitProvider(name string, staging Staging) (Provider, error) {
	provider, exist := providers[name]
	if !exist {
		return nil, errNoRegisteredProvider
	}
	provider.Init(staging)
	return provider, nil
}
