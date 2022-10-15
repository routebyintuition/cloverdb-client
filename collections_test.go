package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/olekukonko/tablewriter"
	"github.com/ostafen/clover"
	"github.com/urfave/cli/v2"
)

func TestCCLI_ListCollections(t *testing.T) {
	collName := "testcoll"
	tableString := &strings.Builder{}
	tr := *tablewriter.NewWriter(ioutil.Discard)
	tr2 := *tablewriter.NewWriter(tableString)
	ccli := &CCLI{output: &TableWriter{table: tr}}

	dir, err := ioutil.TempDir("", "cdb")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	dir = filepath.Join(dir, "t")

	app := ccli.getApp()
	set := flag.NewFlagSet("dir", 0)

	set.String("dir", dir, "directory name")
	set.String("name", collName, "collection name")

	ctx := cli.NewContext(app, set, nil)

	cdb, err := clover.Open(dir)
	if err != nil {
		t.Errorf("clover.Open CreateCollection error: %v", err)
	}

	ccli.cdb = cdb

	err = ccli.CreateCollection(ctx)
	if err != nil {
		t.Errorf("CCLI.CreateCollection() CreateCollection error = %v, wantErr %v", err, false)
	}

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
			name: "List Collections",
			args: args{
				ctx,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ccli.output = &TableWriter{table: tr2}
			if err := ccli.ListCollections(ctx); (err != nil) != tt.wantErr {
				t.Errorf("CCLI.ListCollections() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !strings.Contains(tableString.String(), collName) {
				t.Errorf("ListCollections does not contain %s", collName)
			}
		})
	}
}

func TestCCLI_CreateCollection(t *testing.T) {
	collName := "testcollcreate"
	tableString := &strings.Builder{}
	tr := *tablewriter.NewWriter(ioutil.Discard)
	tr2 := *tablewriter.NewWriter(tableString)
	ccli := &CCLI{output: &TableWriter{table: tr}}

	dir, err := ioutil.TempDir("", "cdb")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	dir = filepath.Join(dir, "t")

	app := ccli.getApp()
	set := flag.NewFlagSet("dir", 0)

	set.String("dir", dir, "directory name")
	set.String("name", collName, "collection name")

	ctx := cli.NewContext(app, set, nil)

	cdb, err := clover.Open(dir)
	if err != nil {
		t.Errorf("clover.Open CreateCollection error: %v", err)
	}

	ccli.cdb = cdb
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
			name: "Create Collection",
			args: args{
				ctx,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ccli := &CCLI{
				cdb: cdb,
			}
			ccli.output = &TableWriter{table: tr2}
			if err := ccli.CreateCollection(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CCLI.CreateCollection() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !strings.Contains(tableString.String(), collName) || !strings.Contains(tableString.String(), "created") {
				t.Errorf("CreateCollection() does not contain %s", collName)
			}
		})
	}
}

func TestCCLI_DropCollection(t *testing.T) {
	collName := "testcolldrop"
	tableString := &strings.Builder{}
	tr := *tablewriter.NewWriter(ioutil.Discard)
	tr2 := *tablewriter.NewWriter(tableString)
	ccli := &CCLI{output: &TableWriter{table: tr}}

	dir, err := ioutil.TempDir("", "cdb")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	dir = filepath.Join(dir, "t")

	app := ccli.getApp()
	set := flag.NewFlagSet("dir", 0)

	set.String("dir", dir, "directory name")
	set.String("name", collName, "collection name")

	ctx := cli.NewContext(app, set, nil)

	cdb, err := clover.Open(dir)
	if err != nil {
		t.Errorf("clover.Open CreateCollection error: %v", err)
	}

	errD := cdb.CreateCollection(collName)
	if errD != nil {
		t.Errorf("clover.CreateCollection DropCollection error: %v", errD)
	}

	ccli.cdb = cdb

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
			name: "Drop Collection",
			args: args{
				ctx,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ccli := &CCLI{
				cdb: cdb,
			}
			ccli.output = &TableWriter{table: tr2}

			if err := ccli.DropCollection(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CCLI.DropCollection() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !strings.Contains(tableString.String(), collName) || !strings.Contains(tableString.String(), "dropped") {
				t.Errorf("DropCollection() does not contain %s", collName)
			}
		})
	}
}

func TestCCLI_HasCollection(t *testing.T) {
	collName := "testcollhas"
	tableString1 := &strings.Builder{}
	tr := *tablewriter.NewWriter(ioutil.Discard)
	tr1 := *tablewriter.NewWriter(tableString1)

	ccli := &CCLI{output: &TableWriter{table: tr}}

	dir, err := ioutil.TempDir("", "cdb")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	dir = filepath.Join(dir, "t")

	app := ccli.getApp()
	set := flag.NewFlagSet("dir", 0)

	set.String("dir", dir, "directory name")
	set.String("name", collName, "collection name")

	ctx := cli.NewContext(app, set, nil)

	cdb, err := clover.Open(dir)
	if err != nil {
		t.Errorf("clover.Open CreateCollection error: %v", err)
	}

	errD := cdb.CreateCollection(collName)
	if errD != nil {
		t.Errorf("clover.CreateCollection DropCollection error: %v", errD)
	}

	ccli.cdb = cdb
	type fields struct {
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
			name: "Does Have Collection",
			args: args{
				ctx,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ccli := &CCLI{
				cdb:    cdb,
				output: &TableWriter{table: tr1},
			}

			if err := ccli.HasCollection(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CCLI.HasCollection() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !strings.Contains(tableString1.String(), collName) || !strings.Contains(tableString1.String(), "YES") {
				t.Errorf("HasCollection() should be true for: %s", collName)
			}

		})
	}
}
