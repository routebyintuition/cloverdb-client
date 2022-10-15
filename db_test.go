package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/olekukonko/tablewriter"
	"github.com/ostafen/clover"
	"github.com/urfave/cli/v2"
)

func TestCCLI_create(t *testing.T) {
	tr := *tablewriter.NewWriter(ioutil.Discard)
	ccli := &CCLI{}

	dir, err := ioutil.TempDir("", "cdb")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	dir = filepath.Join(dir, "t")

	app := ccli.getApp()
	set := flag.NewFlagSet("dir", 0)
	set.String("dir", dir, "directory name")

	ctx := cli.NewContext(app, set, nil)

	type fields struct {
		cdb    *clover.DB
		output OutputWriter
	}
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Create Database",
			fields: fields{
				output: &TableWriter{table: tr},
			},
			args: args{
				ctx,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ccli := &CCLI{
				cdb:    tt.fields.cdb,
				output: tt.fields.output,
			}
			if err := ccli.create(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CCLI.create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_openDB(t *testing.T) {
	ccli := &CCLI{}

	app := ccli.getApp()
	setIV := flag.NewFlagSet("dir", 0)
	setIV.String("dir", "testopIV", "directory name")
	ctxIV := cli.NewContext(app, setIV, nil)

	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *clover.DB
		wantErr bool
	}{
		{
			name: "Invalid open database check",
			args: args{
				ctxIV,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := openDB(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("openDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("openDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCCLI_open(t *testing.T) {
	tr := *tablewriter.NewWriter(ioutil.Discard)
	ccli := &CCLI{}

	dir, err := ioutil.TempDir("", "cdb")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	dir = filepath.Join(dir, "t")

	app := ccli.getApp()
	set := flag.NewFlagSet("dir", 0)
	set.String("dir", dir, "directory name")

	ctx := cli.NewContext(app, set, nil)

	cdb, err := clover.Open(dir)
	if err != nil {
		t.Errorf("CCLI.open() error = %v, wantErr %v", err, false)
		return
	}
	cdb.Close()

	type fields struct {
		cdb    *clover.DB
		output OutputWriter
	}
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Valid open database",
			fields: fields{
				output: &TableWriter{table: tr},
			},
			args: args{
				ctx,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ccli := &CCLI{
				cdb:    tt.fields.cdb,
				output: tt.fields.output,
			}
			if err := ccli.open(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CCLI.open() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
