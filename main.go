package main

import (
	"log"
	"os"

	"github.com/ostafen/clover"
)

type CCLI struct {
	cdb *clover.DB
	//tr      *tablewriter.Table
	output OutputWriter
}

type OutputWriter interface {
	Write([][]string)
}

type OutputPrinter interface {
	Write(map[string]interface{})
}

func main() {
	ccli := &CCLI{
		//tr: tablewriter.NewWriter(os.Stdout),
	}

	app := ccli.getApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
