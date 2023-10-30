package router

import "errors"

var (
	ErrNoEndpointValueProvided = errors.New("no value provided for requested endpoint")
	ErrNoChainWithID           = errors.New("no known chain with provided ID")
)
