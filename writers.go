package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

type TableWriter struct {
	table  tablewriter.Table
	writer io.Writer
}

type JSONWriter struct {
	writer io.Writer
}

func setWriter(c *cli.Context) (OutputWriter, error) {
	if c.String("format") == "json" {
		return &JSONWriter{writer: os.Stdout}, nil
	} else {
		tr := tablewriter.NewWriter(os.Stdout)
		return &TableWriter{table: *tr, writer: os.Stdout}, nil
	}
}

func (jr JSONWriter) Write(data [][]string) {

	headers := len(data[0])
	if len(data) == 2 {
		outMap := make(map[string]interface{})
		for i := 0; i < headers; i++ {
			outMap[data[0][i]] = data[1][i]
		}
		jsonByte, _ := json.Marshal(outMap)
		fmt.Println(string(jsonByte))
		return
	} else if len(data) > 2 {
		outMap := make([]map[string]interface{}, 0, 0)
		dataHeaders := data[0]
		for x := 1; x < len(data); x++ {
			outMapInt := make(map[string]interface{})
			for i := 0; i < len(dataHeaders); i++ {
				outMapInt[dataHeaders[i]] = data[x][i]
			}
			outMap = append(outMap, outMapInt)
		}

		jsonByte, _ := json.Marshal(outMap)
		fmt.Println(string(jsonByte))
		return
	}

	jsonByte, _ := json.Marshal(data)
	fmt.Println(string(jsonByte))

	return
}

func (tr *TableWriter) Write(data [][]string) {
	tr.table.SetHeader(data[0])
	tr.table.SetAutoWrapText(false)
	tr.table.AppendBulk(data[1:])
	tr.table.Render()
	tr.table.ClearRows()
	tr.table = *tablewriter.NewWriter(os.Stdout)
}
