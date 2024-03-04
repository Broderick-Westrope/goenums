package config

import (
	"reflect"
	"testing"
)

func TestParseJson(t *testing.T) {
	type input struct {
		data []byte
	}
	type expected struct {
		cfg *Config
		err bool
	}

	tt := []struct {
		name     string
		input    input
		expected expected
	}{
		{
			name: "TestParseYaml",
			input: input{
				data: []byte("{\"enums\": [{\"package\": \"validation\",\"type\": \"status\",\"values\": [\"Failed\",\"Passed\",\"Skipped\",\"Scheduled\",\"Running\"]}]}"),
			},
			expected: expected{
				cfg: &Config{
					OutputPath: "",
					EnumConfigs: []*EnumConfig{
						{
							Package: "validation",
							Type:    "status",
							Enums:   []string{"Failed", "Passed", "Skipped", "Scheduled", "Running"},
						},
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cfg, err := parseJson(tc.input.data)

			if err != nil {
				if !tc.expected.err {
					t.Errorf("Error: expected nil, got %v", err)
					return
				}
			} else {
				if tc.expected.err {
					t.Errorf("Error: expected error, got nil")
					return
				}
			}

			if cfg.OutputPath != tc.expected.cfg.OutputPath {
				t.Errorf("Output Path: expected %v, got %v", tc.expected.cfg.OutputPath, cfg.OutputPath)
			}
			if !reflect.DeepEqual(cfg.EnumConfigs, tc.expected.cfg.EnumConfigs) {
				t.Errorf("Enum EnumConfigs: expected %v, got %v", tc.expected.cfg.EnumConfigs, cfg.EnumConfigs)
			}
		})
	}
}

func TestParseYaml(t *testing.T) {
	type input struct {
		data []byte
	}
	type expected struct {
		cfg *Config
		err bool
	}

	tt := []struct {
		name     string
		input    input
		expected expected
	}{
		{
			name: "TestParseYaml",
			input: input{
				data: []byte("enums:\n  - package: validation\n    type: status\n    values:\n      - Failed\n      - Passed\n      - Skipped\n      - Scheduled\n      - Running"),
			},
			expected: expected{
				cfg: &Config{
					OutputPath: "",
					EnumConfigs: []*EnumConfig{
						{
							Package: "validation",
							Type:    "status",
							Enums:   []string{"Failed", "Passed", "Skipped", "Scheduled", "Running"},
						},
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cfg, err := parseYaml(tc.input.data)

			if err != nil {
				if !tc.expected.err {
					t.Errorf("Error: expected nil, got %v", err)
					return
				}
			} else {
				if tc.expected.err {
					t.Errorf("Error: expected error, got nil")
					return
				}
			}

			if cfg.OutputPath != tc.expected.cfg.OutputPath {
				t.Errorf("Output Path: expected %v, got %v", tc.expected.cfg.OutputPath, cfg.OutputPath)
			}
			if !reflect.DeepEqual(cfg.EnumConfigs, tc.expected.cfg.EnumConfigs) {
				t.Errorf("Enum EnumConfigs: expected %v, got %v", tc.expected.cfg.EnumConfigs, cfg.EnumConfigs)
			}
		})
	}
}
