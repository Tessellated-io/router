package static

import (
	"github.com/tessellated-io/router/base"
	"github.com/tessellated-io/router/types"
)

// staticRouter uses a preconfigured set of routes.
type staticRouter struct {
	base.BaseRouter
}

// Type assertion
var _ types.Router = (*staticRouter)(nil)

// NewRouter makes a new router with the given chains.
func NewRouter(chains []types.Chain) (types.Router, error) {
	baseRouter, err := base.NewRouter()
	if err != nil {
		return nil, err
	}

	baseRouter.AddChains(chains)

	return &staticRouter{
		BaseRouter: baseRouter,
	}, nil
}

// Router Interface

func (sr *staticRouter) Refresh() error {
	// Intentional no-op
	return nil
}
