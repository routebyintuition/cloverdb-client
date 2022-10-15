package main

import (
	"strings"
	"testing"
)

func Test_geRootFlags(t *testing.T) {
	tString := &strings.Builder{}

	ccli := &CCLI{
		//tr: tablewriter.NewWriter(os.Stdout),
	}

	app := ccli.getApp()
	app.Writer = tString

	app.Flags = geRootFlags()

	err := app.Run([]string{"help"})
	if err != nil {
		t.Errorf("app.Run() Run error: %v", err)
	}

	if !strings.Contains(tString.String(), "CloverDB CLI - clover") {
		t.Errorf("app.Run() Run error: %v", err)
	}
}
