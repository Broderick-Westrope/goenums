package config

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	OutputPath string        `json:"output_path" yaml:"output_path"`
	Configs    []*EnumConfig `json:"enums" yaml:"enums"`
}

type EnumConfig struct {
	Package string   `json:"package" yaml:"package"`
	Type    string   `json:"type" yaml:"type"`
	Enums   []string `json:"values" yaml:"values"`
}

func Parse(path, format string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return &Config{}, err
	}

	var extension string
	if format == "" {
		extension = filepath.Ext(path)
	} else {
		extension = format
	}
	extension = strings.TrimPrefix(extension, ".")

	var cfg *Config
	switch strings.ToLower(extension) {
	case "json":
		cfg, err = parseJson(file)
	case "yaml", "yml":
		cfg, err = parseYaml(file)
	default:
		return nil, fmt.Errorf("unsupported file format %q", extension)
	}
	if err != nil {
		if format == "" {
			return nil, fmt.Errorf("failed to automatically parse file %q. Try specifying a format with the -format flag: %w", path, err)
		}
		return nil, fmt.Errorf("failed to parse file %q: %w", path, err)
	}
	return cfg, nil
}

func parseJson(data []byte) (*Config, error) {
	config := &Config{}
	err := json.Unmarshal(data, config)
	if err != nil {
		return &Config{}, err
	}
	return config, nil
}

func parseYaml(data []byte) (*Config, error) {
	config := &Config{}
	err := yaml.Unmarshal(data, config)
	if err != nil {
		return &Config{}, err
	}
	return config, nil
}
