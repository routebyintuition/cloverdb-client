package main

import (
	"os"

	"github.com/lensesio/tableprinter"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

type TableWriter struct {
	table tablewriter.Table
}

type TablePrinter struct {
	printer tableprinter.Printer
}

type JSONWriter struct {
}

func setWriter(c *cli.Context) (OutputWriter, OutputPrinter, error) {
	if c.String("format") == "table" {
		tr := tablewriter.NewWriter(os.Stdout)
		tp := tableprinter.New(os.Stdout)
		return &TableWriter{table: *tr}, &TablePrinter{printer: *tp}, nil
	} else {
		tr := tablewriter.NewWriter(os.Stdout)
		tp := tableprinter.New(os.Stdout)
		return &TableWriter{table: *tr}, &TablePrinter{printer: *tp}, nil
	}
}

func (tp *TablePrinter) Write(data map[string]interface{}) {
	tp.printer.Print(data)
}

func (tr *TableWriter) Write(data [][]string) {
	tr.table.SetHeader(data[0])
	//tr.table.SetFooter(data[len(data)-1])
	tr.table.SetAutoWrapText(false)
	tr.table.AppendBulk(data[1:])
	tr.table.Render()
	tr.table.ClearRows()
	tr.table.ClearFooter()
	tr.table = *tablewriter.NewWriter(os.Stdout)
}
