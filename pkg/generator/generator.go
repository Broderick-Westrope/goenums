package generator

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
	path := filepath.Join(outPath, templateData.SnakePackage)
	f, fp, err := setupFiles(path, templateData.SnakeFileName)
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

// setupFiles creates the file and the path if it doesn't exist
// filename should not include the path or extension
func setupFiles(path, filename string) (*os.File, string, error) {
	if err := makeDirIfNotExist(path); err != nil {
		return nil, "", err
	}

	fPath := filepath.Join(path, filename+".go")
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
