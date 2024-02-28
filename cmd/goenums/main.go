package main

import (
	"fmt"
	"github.com/broderick-westrope/goenums/pkg/config"
	"github.com/broderick-westrope/goenums/pkg/generator"
	"github.com/spf13/pflag"
	"os"
)

func main() {
	pflag.Usage = func() {
		fmt.Println("Usage: goenums <config path> <output path>")
		pflag.PrintDefaults()
	}
	pflag.Parse()

	// Check for the config file
	if len(os.Args) < 2 {
		fmt.Println("Error: No config file provided.")
		os.Exit(1)
	}
	cfgPath := os.Args[1]

	// Parse the config file
	cfg, err := config.Parse(cfgPath)
	if err != nil {
		fmt.Println("Error:", err)
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
	g := generator.New(cfg)
	err = g.Generate()
	if err != nil {
		fmt.Println("Error generating code:", err)
		return
	}
}
