package main

import (
	"fmt"
	"os"

	"github.com/broderick-westrope/goenums/internal/config"
	"github.com/broderick-westrope/goenums/internal/generator"
	flag "github.com/spf13/pflag"
)

func main() {
	flag.Usage = func() {
		fmt.Println("Usage: goenums <config path> <output path>")
		flag.PrintDefaults()
		os.Exit(0)
	}

	var format string
	flag.StringVarP(&format, "format", "f", "", "The format of the config file (json, yaml, yml)")
	var gofmt bool
	flag.BoolVarP(&gofmt, "gofmt", "g", true, "run gofmt on the generated code")

	flag.Parse()

	// Check for the config file
	if len(os.Args) < 2 {
		fmt.Println("Error: No config file provided.")
		os.Exit(1)
	}
	cfgPath := os.Args[1]

	// Parse the config file
	cfg, err := config.Parse(cfgPath, format)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if len(cfg.EnumConfigs) == 0 {
		fmt.Printf("Error: No enum configurations found in the config file %q.\n", cfgPath)
		os.Exit(1)
	}

	// Check for the output path in both the config and the arguments
	if len(os.Args) < 3 {
		if cfg.OutputPath == "" {
			fmt.Println("Error: Config file contained no output path and no output path was provided as an argument.")
			os.Exit(1)
		}
	} else {
		// Prioritise the output path provided as an argument, when present
		cfg.OutputPath = os.Args[2]
	}

	// Generate the code
	g := generator.New(cfg, gofmt)
	err = g.Generate()
	if err != nil {
		fmt.Println("Error generating code:", err)
		os.Exit(1)
	}
}
