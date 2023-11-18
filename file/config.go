package file

import (
	"os"

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
