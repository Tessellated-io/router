package types

// Chain defines an abstraction around a Chain that Validator tooling or other blockchain applications may use.
type Chain interface {
	ChainID() string
	GrpcEndpoint() (string, error)
}

// private implementation
type chain struct {
	chainID      string
	grpcEndpoint *string
}

// Ensure chain is a Chain
var _ Chain = (*chain)(nil)

// Create a new Chain
func NewChain(chainID string, grpcEndpoint *string) Chain {
	return &chain{
		chainID:      chainID,
		grpcEndpoint: grpcEndpoint,
	}
}

// Chain Interface

func (c *chain) ChainID() string { return c.chainID }

func (c *chain) GrpcEndpoint() (string, error) {
	if c.grpcEndpoint == nil {
		return "", ErrNoEndpointValueProvided
	} else {
		return *c.grpcEndpoint, nil
	}
}
