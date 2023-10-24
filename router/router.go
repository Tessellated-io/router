package router

import "fmt"

// Router defines a way to get addresses and API endpoints for blockchain nodes
type Router interface {
	GetHumanReadableName(chainName string) (string, error)

	GetGrpcEndpoint(chainName string) (string, error)
}

// Private implementing type
type router struct {
	// Map of chainName to chain.
	chains map[string]Chain
}

// Ensure router is a Router
var _ Router = (*router)(nil)

// NewRouter makes a new router with the given chains.
func NewRouter(chains []Chain) (Router, error) {
	chainMap := make(map[string]Chain)
	for _, chain := range chains {
		chainName := chain.GetChainName()

		_, isSet := chainMap[chainName]
		if isSet {
			return nil, fmt.Errorf("duplicate chain name: %s", chainName)
		}

		chainMap[chainName] = chain
	}

	return &router{
		chains: chainMap,
	}, nil
}

// Router Interface

func (r *router) GetHumanReadableName(chainName string) (string, error) {
	chain := r.chains[chainName]
	if chain == nil {
		return "", ErrNoChainWithName
	}

	return chain.GetHumanReadableName(), nil
}

func (r *router) GetGrpcEndpoint(chainName string) (string, error) {
	chain := r.chains[chainName]
	if chain == nil {
		return "", ErrNoChainWithName
	}

	return chain.GetGrpcEndpoint()
}
