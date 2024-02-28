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
	OutputPath string        `json:"outputPath"`
	Configs    []*EnumConfig `json:"enums"`
}

type EnumConfig struct {
	Package string   `json:"package"`
	Type    string   `json:"type"`
	Enums   []string `json:"values"`
}

func Parse(path string) (*Config, error) {
	extension := filepath.Ext(path)

	switch strings.ToLower(extension) {
	case "json":
		return parseJson(path)
	case "yaml", "yml":
		return parseYaml(path)
	default:
		return nil, fmt.Errorf("unsupported file type %q", extension)
	}
}

func parseJson(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return &Config{}, err
	}
	config := &Config{}
	err = json.Unmarshal(file, config)
	if err != nil {
		return &Config{}, err
	}
	return config, nil
}

func parseYaml(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return &Config{}, err
	}
	config := &Config{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return &Config{}, err
	}
	return config, nil
}
