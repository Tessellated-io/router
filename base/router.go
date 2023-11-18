package base

import "github.com/tessellated-io/router/types"

// baseRouter encapsulates basic functionality for all routers. You probably shouldn't use this directly.
type BaseRouter struct {
	// Map of chain-id to chain.
	chains map[string]types.Chain
}

// NewRouter returns a new BaseRouter

func NewRouter() (BaseRouter, error) {
	chainMap := make(map[string]types.Chain)
	return BaseRouter{
		chains: chainMap,
	}, nil
}

// Router Interace

func (br *BaseRouter) GrpcEndpoint(chainID string) (string, error) {
	chain := br.chains[chainID]
	if chain == nil {
		return "", types.ErrNoChainWithID
	}

	return chain.GrpcEndpoint()
}

// Helper methods

func (br *BaseRouter) AddChains(chains []types.Chain) {
	for _, chain := range chains {
		br.chains[chain.ChainID()] = chain
	}
}
