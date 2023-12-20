package types

// Router defines a set of rules that allow you to route to a set of blockchain nodes.
type Router interface {
	Refresh() error

	GrpcEndpoint(ChainID string) (string, error)
}
