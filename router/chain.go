package router

// Chain defines an abstraction around a Chain that Validator tooling or other blockchain applications may use.
type Chain interface {
	GetHumanReadableName() string
	GetChainName() string

	GetGrpcEndpoint() (string, error)
}

// private implementation
type chain struct {
	chainName         string
	humanReadableName string

	grpcEndpoint *string
}

// Ensure chain is a Chain
var _ Chain = (*chain)(nil)

// Create a new Chain
func NewChain(chainName string, humanReadableName string, grpcEndpoint *string) (Chain, error) {
	return &chain{
		chainName:         chainName,
		humanReadableName: humanReadableName,

		grpcEndpoint: grpcEndpoint,
	}, nil
}

// Chain Interface

func (c *chain) GetHumanReadableName() string { return c.chainName }
func (c *chain) GetChainName() string         { return c.humanReadableName }

func (c *chain) GetGrpcEndpoint() (string, error) {
	if c.grpcEndpoint == nil {
		return "", ErrNoEndpointValueProvided
	} else {
		return *c.grpcEndpoint, nil
	}
}
