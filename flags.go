package main

import (
	"github.com/urfave/cli/v2"
)

func geRootFlags() []cli.Flag {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:        "dir",
			Aliases:     []string{"directory", "d", "D"},
			Usage:       "CloverDB directory",
			DefaultText: "",
			EnvVars:     []string{"CLOVER_DIR", "CLOVER_DIRECTORY"},
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "format",
			Aliases:     []string{"form", "f", "F"},
			Usage:       "output format <json> <table>",
			DefaultText: "table",
			EnvVars:     []string{"CLOVER_FORMAT"},
		},
	}
	return flags
}
