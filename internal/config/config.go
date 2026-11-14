package config

import (
	"fmt"
	"os"
	"path/filepath"
)

var ConfigTypeOptions = []string{"yaml", "yml", "toml", "json"}

type Config struct {
	Source SourceConfig
	Target TargetConfig
}

type SourceConfig struct {
	Path             string
	FileNameTemplate string
}

type TargetConfig struct {
	Path     string
	FileName string
	Type     string
}

func NewConfig() Config {
	return Config{
		Source: SourceConfig{
			Path:             "changelog",
			FileNameTemplate: "${date}.${counter:3}.md",
		},
		Target: TargetConfig{
			Type:     "markdown",
			FileName: "CHANGELOG.md",
		},
	}
}

func LoadConfig(cwd string) (Config, error) {
	entries, err := filepath.Glob(filepath.Join(cwd, "changelogr.*"))
	if err != nil {
		return Config{}, err
	}

	var configFile, configType string
	for _, entry := range entries {
		for _, option := range ConfigTypeOptions {
			if filepath.Base(entry) == "changelogr."+option {
				configFile = entry
				configType = option
				break
			}
		}
		if len(configFile) > 0 && len(configType) > 0 {
			break
		}
	}

	// TODO: Parse config

	return Config{}, nil
}

func (cfg Config) Save(cwd string, configType string) error {
	content, err := cfg.serialize(configType)
	if err != nil {
		return err
	}

	if err := os.WriteFile(
		filepath.Join(cwd, fmt.Sprintf("changelogr.%s", configType)),
		content,
		0644,
	); err != nil {
		return err
	}

	return nil
}

func (cfg Config) serialize(configType string) ([]byte, error) {
	switch configType {
	case "yaml", "yml":
		return serializeYAML(cfg)
	case "toml":
		return serializeTOML(cfg)
	case "json":
		return serializeJSON(cfg)
	default:
		return nil, fmt.Errorf("invalid config type: %s", configType)
	}
}
