package parser

import (
	"reflect"
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
		// NOTE: output Enums are not tested.
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
					EnumConfigs: []*config.EnumConfig{
						{
							Package: "Test Pkg",
							Type:    "test Type",
							Enums:   []string{"One value", "Value two", "and.a_Third"},
						},
					},
				},
			},
			expected: output{
				etd: []EnumTemplateData{
					{
						SnakePackage:    "test_pkg",
						SnakeFileName:   "test_type",
						LowerCamelType:  "testType",
						CamelType:       "TestType",
						CamelTypePlural: "TestTypes",
						MethodReceiver:  "t",
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			etd := ParseConfig(tc.input.config)

			for i, e := range etd {
				if i >= len(tc.expected.etd) {
					t.Errorf("more elements than expected: want %d, got %d", len(tc.expected.etd), len(etd))
					return
				}

				if e.SnakePackage != tc.expected.etd[i].SnakePackage {
					t.Errorf("etd[%d].SnakePackage: want %s, got %s", i, tc.expected.etd[i].SnakePackage, e.SnakePackage)
				}
				if e.SnakeFileName != tc.expected.etd[i].SnakeFileName {
					t.Errorf("etd[%d].SnakeFileName: want %s, got %s", i, tc.expected.etd[i].SnakeFileName, e.SnakeFileName)
				}
				if e.LowerCamelType != tc.expected.etd[i].LowerCamelType {
					t.Errorf("etd[%d].LowerCamelType: want %s, got %s", i, tc.expected.etd[i].LowerCamelType, e.LowerCamelType)
				}
				if e.CamelType != tc.expected.etd[i].CamelType {
					t.Errorf("etd[%d].CamelType: want %s, got %s", i, tc.expected.etd[i].CamelType, e.CamelType)
				}
				if e.CamelTypePlural != tc.expected.etd[i].CamelTypePlural {
					t.Errorf("etd[%d].CamelTypePlural: want %s, got %s", i, tc.expected.etd[i].CamelTypePlural, e.CamelTypePlural)
				}
				if e.MethodReceiver != tc.expected.etd[i].MethodReceiver {
					t.Errorf("etd[%d].CamelTypePlural: want %s, got %s", i, tc.expected.etd[i].MethodReceiver, e.MethodReceiver)
				}
				if _, _, enums := extractVariables(tc.input.config.EnumConfigs[i]); !reflect.DeepEqual(e.Enums, enums) {
					t.Errorf("etd[%d].Enums: didn't receive expected result", i)
				}
			}
		})
	}
}
