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
	ChainID      string `yaml:"chain-id"`
	GrpcEndpoint string `yaml:"grpc"`
}

// Parse a file into a UserConfig
func parseConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
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
func InitializeConfigFile(filename string, configurationDirectory string, logger *log.Logger) error {
	networks := []NetworkConfig{
		NetworkConfig{
			ChainID:      "cosmoshub-4",
			GrpcEndpoint: "tcp://cosmos.rpc.directory",
		},
		NetworkConfig{
			ChainID:      "osmosis-1",
			GrpcEndpoint: "tcp://osmosis.rpc.directory",
		},
	}

	routerConfig := Config{
		Networks: networks,
	}

	// Create folder if needed
	config.CreateDirectoryIfNeeded(configurationDirectory, logger)

	routerConfigFile := fmt.Sprintf("%s/%s", configurationDirectory, filename)
	header := "This is a file to route chain IDs to RPC endpoints\n"
	err := config.WriteYamlWithComments(routerConfig, header, routerConfigFile, logger)
	if err != nil {
		return err
	}
	return nil
}
