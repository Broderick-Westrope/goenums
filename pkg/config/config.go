package config

import (
	"encoding/json"
	"fmt"
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

func Parse(cfgPath string) (*Config, error) {
	extension := filepath.Ext(cfgPath)

	switch strings.ToLower(extension) {
	case "json":
		return parseJson(cfgPath)
	default:
		return nil, fmt.Errorf("unsupported file type %q", extension)
	}
}

func parseJson(path string) (*Config, error) {
	// read config file from the json file and return the config
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
