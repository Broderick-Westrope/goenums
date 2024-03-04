package parser

import (
	"github.com/broderick-westrope/goenums/internal/config"
	"github.com/iancoleman/strcase"
)

type EnumTemplateData struct {
	SnakePackage    string // e.g. an_enum_package
	SnakeFileName   string // e.g. some_enum_type
	LowerCamelType  string // e.g. someEnumType
	CamelType       string // e.g. SomeEnumType
	CamelTypePlural string // e.g. SomeEnumTypes
	MethodReceiver  string // e.g. s
	Enums           []Enum // The enum values
}

type Enum struct {
	CamelValue          string // e.g. SomeEnumValue
	ScreamingSnakeValue string // e.g. SOME_ENUM_VALUE
	LowerCamelValue     string // e.g. someEnumValue
	CamelType           string // e.g. SomeEnumType
	LowerCamelType      string // e.g. someEnumType
}

func ParseConfig(cfg *config.Config) []EnumTemplateData {
	data := make([]EnumTemplateData, len(cfg.EnumConfigs))
	for i, c := range cfg.EnumConfigs {
		camelType, snakePackage, enums := extractVariables(c)

		var camelTypePlural string
		if camelType[len(camelType)-1] == 's' {
			camelTypePlural = camelType + "es"
		} else {
			camelTypePlural = camelType + "s"
		}

		lowerCamelType := strcase.ToLowerCamel(camelType)

		templateData := EnumTemplateData{
			SnakePackage:    snakePackage,
			SnakeFileName:   strcase.ToSnake(camelType),
			LowerCamelType:  lowerCamelType,
			CamelType:       camelType,
			CamelTypePlural: camelTypePlural,
			MethodReceiver:  string(lowerCamelType[0]),
			Enums:           enums,
		}
		data[i] = templateData
	}
	return data
}

func extractVariables(cfg *config.EnumConfig) (typ, pkg string, enums []Enum) {
	typ = strcase.ToCamel(cfg.Type)
	pkg = strcase.ToSnake(cfg.Package)
	enumNames := cfg.Enums
	enums = make([]Enum, len(enumNames))
	for i, n := range enumNames {
		enums[i] = Enum{
			CamelValue:          strcase.ToCamel(n),
			LowerCamelValue:     strcase.ToLowerCamel(n),
			ScreamingSnakeValue: strcase.ToScreamingSnake(n),
			CamelType:           strcase.ToCamel(typ),
			LowerCamelType:      strcase.ToLowerCamel(typ),
		}
	}
	return typ, pkg, enums
}
