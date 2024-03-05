package generator

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"github.com/broderick-westrope/goenums/internal/config"
	"github.com/broderick-westrope/goenums/internal/config/parser"
)

//go:embed template/enum.gotmpl
var fs embed.FS

type Generator struct {
	outputPath string
	data       []parser.EnumTemplateData
	useGofmt   bool
}

func New(cfg *config.Config, useGofmt bool) *Generator {
	return &Generator{
		data:       parser.ParseConfig(cfg),
		outputPath: cfg.OutputPath,
		useGofmt:   useGofmt,
	}
}

func (g *Generator) Generate() error {
	for _, d := range g.data {
		err := generateEnum(d, g.outputPath, g.useGofmt)
		if err != nil {
			return err
		}
	}
	return nil
}

func generateEnum(templateData parser.EnumTemplateData, outPath string, useGofmt bool) error {
	path := filepath.Join(outPath, templateData.SnakePackage)
	file, filePath, err := setupFiles(path, templateData.SnakeFileName)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file: ", err)
		}
	}()
	t := template.Must(template.ParseFS(fs, "template/enum.gotmpl"))
	err = t.Execute(file, templateData)
	if err != nil {
		return err
	}

	if useGofmt {
		cmd := exec.Command("gofmt", "-w", filePath)
		err = cmd.Run()
		if err != nil {
			return err
		}
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
