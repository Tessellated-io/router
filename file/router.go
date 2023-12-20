package file

import (
	"github.com/tessellated-io/pickaxe/arrays"
	"github.com/tessellated-io/router/base"
	"github.com/tessellated-io/router/types"
)

// fileRouter is a router that maps to networks in a config file.
type fileRouter struct {
	base.BaseRouter

	configFile string
}

// Type assertion
var _ types.Router = (*fileRouter)(nil)

// NewRouter creates a new file router.
func NewRouter(configFile string) (types.Router, error) {
	baseRouter, err := base.NewRouter()
	if err != nil {
		return nil, err
	}

	fileRouter := &fileRouter{
		BaseRouter: baseRouter,

		configFile: configFile,
	}

	err = fileRouter.loadConfigFile()
	if err != nil {
		return nil, err
	}

	return fileRouter, nil
}

// Router interface

func (fr *fileRouter) Refresh() error {
	return fr.loadConfigFile()
}

// Private methods

func (fr *fileRouter) loadConfigFile() error {
	parsed, err := parseConfig(fr.configFile)
	if err != nil {
		return err
	}

	chains := arrays.Map(parsed.Networks, func(networkConfig NetworkConfig) types.Chain {
		return types.NewChain(networkConfig.ChainID, &networkConfig.GrpcEndpoint)
	})

	fr.BaseRouter.AddChains(chains)
	return nil
}
