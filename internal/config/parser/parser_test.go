package parser

import (
	"testing"

	"github.com/broderick-westrope/goenums/internal/config"
)

func TestExtractVariables(t *testing.T) {
	type input struct {
		config *config.EnumConfig
	}

	type output struct {
		typ   string
		pkg   string
		enums []Enum
	}

	tt := []struct {
		name     string
		input    input
		expected output
	}{
		{
			name: "basic",
			input: input{
				config: &config.EnumConfig{
					Package: "Test Pkg",
					Type:    "test Type",
					Enums:   []string{"One value", "Value two", "and.a_Third"},
				},
			},
			expected: output{
				typ: "TestType",
				pkg: "test_pkg",
				enums: []Enum{
					{
						CamelValue:          "OneValue",
						LowerCamelValue:     "oneValue",
						ScreamingSnakeValue: "ONE_VALUE",
					},
					{
						CamelValue:          "ValueTwo",
						LowerCamelValue:     "valueTwo",
						ScreamingSnakeValue: "VALUE_TWO",
					},
					{
						CamelValue:          "AndAThird",
						LowerCamelValue:     "andAThird",
						ScreamingSnakeValue: "AND_A_THIRD",
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			typ, pkg, enums := extractVariables(tc.input.config)

			if typ != tc.expected.typ {
				t.Errorf("typ: want %q, got %q", tc.expected.typ, typ)
			}

			if pkg != tc.expected.pkg {
				t.Errorf("internal: want %q, got %q", tc.expected.pkg, pkg)
			}

			if len(enums) != len(tc.expected.enums) {
				t.Errorf("len(enums): want %d, got %d", len(tc.expected.enums), len(enums))
			}

			for i, e := range enums {
				if e.CamelValue != tc.expected.enums[i].CamelValue {
					t.Errorf("enum.CamelValue: want %q, got %q", tc.expected.enums[i].CamelValue, e.CamelValue)
				}
				if e.LowerCamelValue != tc.expected.enums[i].LowerCamelValue {
					t.Errorf("enum.LowerCamelValue: want %q, got %q", tc.expected.enums[i].LowerCamelValue, e.LowerCamelValue)
				}
				if e.ScreamingSnakeValue != tc.expected.enums[i].ScreamingSnakeValue {
					t.Errorf("enum.ScreamingSnakeValue: want %q, got %q", tc.expected.enums[i].ScreamingSnakeValue, e.ScreamingSnakeValue)
				}
			}
		})
	}
}

func TestParseConfig(t *testing.T) {
	type input struct {
		config *config.Config
	}

	type output struct {
		etd []EnumTemplateData
	}

	tt := []struct {
		name     string
		input    input
		expected output
	}{
		{
			name: "basic",
			input: input{
				config: &config.Config{
					OutputPath: "./output/path/",
					//EnumConfigs: ,
				},
			},
			expected: output{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}
