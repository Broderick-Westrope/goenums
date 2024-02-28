package main

import (
	"fmt"
	"github.com/broderick-westrope/goenums/pkg/generator"
	"github.com/spf13/pflag"
	"os"

	"github.com/broderick-westrope/goenums/pkg/config"
)

func main() {
	pflag.Usage = func() {
		fmt.Println("Usage: goenums <config file path> <output path>")
		pflag.PrintDefaults()
	}
	pflag.Parse()

	config, err := ParseInput()
	if err != nil {
		return
	}
	g := generator.New(config)
	err = g.Generate()
	if err != nil {
		fmt.Println("Error generating code", err)
		return
	}
}

func ParseInput() (config.Config, error) {
	cfgPath := os.Args[1]
	cfg, err := config.ReadConfig(cfgPath)
	if err != nil {
		return config.Config{}, err
	}
	cfg.OutputPath = os.Args[2]
	return cfg, err
}
