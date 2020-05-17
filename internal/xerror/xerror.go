package xerror

import "errors"

var (
	ErrProviderNotSupported = errors.New("provider not supported")
	ErrNotReady             = errors.New("not ready")
)
