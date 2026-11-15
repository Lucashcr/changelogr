package config

import (
	"encoding/json"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

func serializeYAML(cfg Config) ([]byte, error) {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func serializeTOML(cfg Config) ([]byte, error) {
	data, err := toml.Marshal(cfg)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func serializeJSON(cfg Config) ([]byte, error) {
	data, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return nil, err
	}

	return data, nil
}
