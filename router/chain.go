package router

// Chain defines an abstraction around a Chain that Validator tooling or other blockchain applications may use.
type Chain interface {
	HumanReadableName() string
	ChainID() string
	Bech32Prefix() string

	GrpcEndpoint() (string, error)
}

// private implementation
type chain struct {
	chainID           string
	humanReadableName string
	bech32Prefix      string

	grpcEndpoint *string
}

// Ensure chain is a Chain
var _ Chain = (*chain)(nil)

// Create a new Chain
func NewChain(chainID, humanReadableName, bech32Prefix string, grpcEndpoint *string) (Chain, error) {
	return &chain{
		chainID:           chainID,
		humanReadableName: humanReadableName,
		bech32Prefix:      bech32Prefix,

		grpcEndpoint: grpcEndpoint,
	}, nil
}

// Chain Interface

func (c *chain) HumanReadableName() string { return c.humanReadableName }
func (c *chain) ChainID() string           { return c.chainID }
func (c *chain) Bech32Prefix() string      { return c.bech32Prefix }

func (c *chain) GrpcEndpoint() (string, error) {
	if c.grpcEndpoint == nil {
		return "", ErrNoEndpointValueProvided
	} else {
		return *c.grpcEndpoint, nil
	}
}
