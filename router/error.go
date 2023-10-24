package router

import "errors"

var (
	ErrNoEndpointValueProvided = errors.New("no value provided for requested endpoint")
	ErrNoChainWithName         = errors.New("no known chain with provided name")
)
