package parser

import (
	"testing"

	"github.com/broderick-westrope/goenums/pkg/config"
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
					Enums:   []string{"One", "Value two", "Three"},
				},
			},
			expected: output{
				typ: "TestType",
				pkg: "test_pkg",
				enums: []Enum{
					{
						CamelValue:          "One",
						LowerCamelValue:     "one",
						ScreamingSnakeValue: "ONE",
						CamelType:           "TestType",
						LowerCamelType:      "testType",
					},
					{
						CamelValue:          "ValueTwo",
						LowerCamelValue:     "valueTwo",
						ScreamingSnakeValue: "VALUE_TWO",
						CamelType:           "TestType",
						LowerCamelType:      "testType",
					},
					{
						CamelValue:          "Three",
						LowerCamelValue:     "three",
						ScreamingSnakeValue: "THREE",
						CamelType:           "TestType",
						LowerCamelType:      "testType",
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
				t.Errorf("pkg: want %q, got %q", tc.expected.pkg, pkg)
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
				if e.CamelType != tc.expected.enums[i].CamelType {
					t.Errorf("enum.CamelType: want %q, got %q", tc.expected.enums[i].CamelType, e.CamelType)
				}
				if e.LowerCamelType != tc.expected.enums[i].LowerCamelType {
					t.Errorf("enum.LowerCamelType: want %q, got %q", tc.expected.enums[i].LowerCamelType, e.LowerCamelType)
				}
			}
		})
	}
}
