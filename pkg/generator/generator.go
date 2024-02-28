package generator

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"

	"github.com/broderick-westrope/goenums/pkg/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//go:embed template/enum.gotmpl
var fs embed.FS

type EnumTemplateData struct {
	Package        string
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

type Generator struct {
	outputPath string
	data       []EnumTemplateData
}

func New(cfg *config.Config) *Generator {
	return &Generator{
		data:       parseConfig(cfg),
		outputPath: cfg.OutputPath,
	}
}

func (g *Generator) Generate() error {
	for _, d := range g.data {
		err := generateEnum(d, g.outputPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func generateEnum(etd EnumTemplateData, outPath string) error {
	f, fp, err := setupFiles(outPath, etd.Package, etd.TypeName)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Error closing file: ", err)
		}
	}()
	t := template.Must(template.ParseFS(fs, "template/enum.gotmpl"))
	if err != nil {
		return err
	}
	err = t.Execute(f, etd)
	if err != nil {
		return err
	}
	// TODO: optionally skip formatting
	cmd := exec.Command("gofmt", "-w", fp)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func parseConfig(cfg *config.Config) []EnumTemplateData {
	data := make([]EnumTemplateData, len(cfg.Configs))
	for i, c := range cfg.Configs {
		typeName, packageName, enums := configToVars(c)
		c := cases.Title(language.English)
		typeNameTitle := c.String(typeName)
		typeNamePlural := typeNameTitle + "s"
		if typeNameTitle[len(typeNameTitle)-1] == 's' {
			typeNamePlural = typeNameTitle + "es"
		}
		typeNameLower := strings.ToLower(typeName)
		etd := EnumTemplateData{
			Package:        packageName,
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

func configToVars(cfg *config.EnumConfig) (typ, pkg string, enums []Enum) {
	typ = cfg.Type
	pkg = cfg.Package
	enumStrs := cfg.Enums
	enums = make([]Enum, len(enumStrs))
	c := cases.Title(language.English)
	typeTitle := c.String(typ)
	for i, enumStr := range enumStrs {
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

func setupFiles(outpath, pkg, typ string) (*os.File, string, error) {
	if err := makeDirIfNotExist(outpath); err != nil {
		return nil, "", err
	}

	dir := path.Join(outpath, pkg)
	if err := makeDirIfNotExist(dir); err != nil {
		return nil, "", err
	}

	fName := fmt.Sprintf("%s.go", strings.ToLower(typ))
	fPath := path.Join(dir, fName)
	f, err := os.Create(fPath)
	if err != nil {
		return nil, "", err
	}
	return f, fPath, nil
}

func makeDirIfNotExist(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}
