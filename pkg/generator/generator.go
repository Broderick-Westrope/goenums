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
	"github.com/broderick-westrope/goenums/pkg/config/parser"
)

//go:embed template/enum.gotmpl
var fs embed.FS

type Generator struct {
	outputPath string
	data       []parser.EnumTemplateData
}

func New(cfg *config.Config) *Generator {
	return &Generator{
		data:       parser.ParseConfig(cfg),
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

func generateEnum(templateData parser.EnumTemplateData, outPath string) error {
	f, fp, err := setupFiles(outPath, templateData.SnakePackage, templateData.CamelType)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Error closing file: ", err)
		}
	}()
	t := template.Must(template.ParseFS(fs, "template/enum.gotmpl"))
	err = t.Execute(f, templateData)
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

func setupFiles(outpath, pkg, typ string) (*os.File, string, error) {
	if err := makeDirIfNotExist(outpath); err != nil {
		return nil, "", err
	}

	dir := path.Join(outpath, pkg)
	if err := makeDirIfNotExist(dir); err != nil {
		return nil, "", err
	}

	typ = strings.ReplaceAll(typ, " ", "_")
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
