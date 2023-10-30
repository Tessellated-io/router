package router

import "fmt"

// Router defines a way to get addresses and API endpoints for blockchain nodes
type Router interface {
	HumanReadableName(chainName string) (string, error)
	GrpcEndpoint(chainName string) (string, error)

	AddChain(chain Chain) error
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
	router := &router{
		chains: chainMap,
	}

	for _, chain := range chains {
		err := router.AddChain(chain)
		if err != nil {
			return nil, err
		}
	}

	return router, nil
}

// Router Interface

func (r *router) HumanReadableName(chainID string) (string, error) {
	chain := r.chains[chainID]
	if chain == nil {
		return "", ErrNoChainWithID
	}

	return chain.HumanReadableName(), nil
}

func (r *router) Bech32Prefix(chainID string) (string, error) {
	chain := r.chains[chainID]
	if chain == nil {
		return "", ErrNoChainWithID
	}

	return chain.Bech32Prefix(), nil
}

func (r *router) GrpcEndpoint(chainID string) (string, error) {
	chain := r.chains[chainID]
	if chain == nil {
		return "", ErrNoChainWithID
	}

	return chain.GrpcEndpoint()
}

func (r *router) AddChain(chain Chain) error {
	chainName := chain.ChainID()

	_, isSet := r.chains[chainName]
	if isSet {
		return fmt.Errorf("duplicate chain name: %s", chainName)
	}

	r.chains[chainName] = chain

	return nil
}
