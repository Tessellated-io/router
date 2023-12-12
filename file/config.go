package file

import (
	"fmt"
	"os"

	"github.com/tessellated-io/pickaxe/config"
	"github.com/tessellated-io/pickaxe/log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Networks []NetworkConfig `yaml:"networks"`
}

type NetworkConfig struct {
	ChainID      string `yaml:"chain_id"`
	GrpcEndpoint string `yaml:"grpc"`
}

// Parse a file into a UserConfig
func parseConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(config.ExpandHomeDir(filename))
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// Write an example config file.
func InitializeConfigFile(filename, configurationDirectory string, logger *log.Logger) error {
	networks := []NetworkConfig{
		{
			ChainID:      "cosmoshub-4",
			GrpcEndpoint: "tcp://cosmos.rpc.directory",
		},
		{
			ChainID:      "osmosis-1",
			GrpcEndpoint: "tcp://osmosis.rpc.directory",
		},
	}

	routerConfig := Config{
		Networks: networks,
	}

	// Create folder if needed
	err := config.CreateDirectoryIfNeeded(configurationDirectory, logger)
	if err != nil {
		return err
	}

	routerConfigFile := fmt.Sprintf("%s/%s", configurationDirectory, filename)
	header := "This is a file to route chain IDs to RPC endpoints\n"
	err = config.WriteYamlWithComments(routerConfig, header, routerConfigFile, logger)
	if err != nil {
		return err
	}
	return nil
}
