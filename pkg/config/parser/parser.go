package parser

import (
	"strings"

	"github.com/broderick-westrope/goenums/pkg/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type EnumTemplateData struct {
	Package        string
	FileName       string
	TypeName       string
	TypeNameLower  string
	TypeNameTitle  string
	TypeNamePlural string
	MethodReceiver string
	Enums          []Enum
}

type Enum struct {
	VariableStr      string
	VariableStrUpper string
	VariableStrLower string
	TypeName         string
	TypeNameLower    string
}

func ParseConfig(cfg *config.Config) []EnumTemplateData {
	data := make([]EnumTemplateData, len(cfg.Configs))
	for i, c := range cfg.Configs {
		fileName, packageName, enums := extractVariables(c)
		c := cases.Title(language.English)
		typeNameTitle := c.String(fileName)
		typeNameTitle = strings.ReplaceAll(typeNameTitle, " ", "")
		typeNamePlural := typeNameTitle + "s"
		if typeNameTitle[len(typeNameTitle)-1] == 's' {
			typeNamePlural = typeNameTitle + "es"
		}
		typeNameLower := strings.ToLower(typeNameTitle)
		etd := EnumTemplateData{
			Package:        packageName,
			FileName:       typeNameTitle,
			TypeName:       typeNameTitle,
			TypeNameLower:  typeNameLower,
			TypeNameTitle:  typeNameTitle,
			TypeNamePlural: typeNamePlural,
			MethodReceiver: string(typeNameLower[0]),
			Enums:          enums,
		}
		data[i] = etd
	}
	return data
}

func extractVariables(cfg *config.EnumConfig) (typ, pkg string, enums []Enum) {
	typ = cfg.Type
	caser := cases.Title(language.English)
	typeTitle := caser.String(typ)

	pkg = cfg.Package

	enumNames := cfg.Enums
	enums = make([]Enum, len(enumNames))

	for i, enumStr := range enumNames {
		opU := strings.ReplaceAll(enumStr, " ", "_")
		op := strings.ReplaceAll(enumStr, " ", "")
		enums[i] = Enum{
			VariableStr:      op,
			VariableStrUpper: strings.ToUpper(opU),
			VariableStrLower: strings.ToLower(op),
			TypeName:         typeTitle,
			TypeNameLower:    strings.ToLower(typ),
		}
	}
	return typ, pkg, enums
}
