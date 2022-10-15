package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

func (ccli *CCLI) ListCollections(c *cli.Context) error {

	collList, err := ccli.cdb.ListCollections()
	if err != nil {
		return err
	}

	data := [][]string{}
	data = append(data, []string{"ID", "Collection Name"})
	count := 0
	for _, v := range collList {
		data = append(data, []string{strconv.Itoa(count), v})
		count++
	}

	ccli.output.Write(data)

	return nil
}

func (ccli *CCLI) CreateCollection(c *cli.Context) error {
	coll := c.String("name")

	err := ccli.cdb.CreateCollection(coll)
	if err != nil {
		return err
	}

	data := [][]string{}
	data = append(data, []string{"Collection Name", "Status"})
	data = append(data, []string{coll, "created"})
	ccli.output.Write(data)

	return nil
}

func (ccli *CCLI) DropCollection(c *cli.Context) error {
	coll := c.String("name")

	err := ccli.cdb.DropCollection(coll)
	if err != nil {
		return err
	}

	data := [][]string{}
	data = append(data, []string{"Collection Name", "Status"})
	data = append(data, []string{coll, "dropped"})
	ccli.output.Write(data)

	return nil
}

func (ccli *CCLI) HasCollection(c *cli.Context) error {
	coll := c.String("name")

	has, err := ccli.cdb.HasCollection(coll)
	if err != nil {
		return err
	}

	data := [][]string{}
	data = append(data, []string{"Collection Name", "Exists"})
	if has {
		data = append(data, []string{coll, "YES"})
	} else {
		data = append(data, []string{coll, "NO"})
	}

	ccli.output.Write(data)

	return nil
}

func (ccli *CCLI) ImportCollection(c *cli.Context) error {
	coll := c.String("name")
	path := c.String("path")

	data := [][]string{}
	data = append(data, []string{"Collection Name", "Source", "Status"})

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		data = append(data, []string{coll, path, "SOURCE DOES NOT EXIST"})
		ccli.output.Write(data)

		return err
	}

	err := ccli.cdb.ImportCollection(coll, path)

	if err == nil {
		data = append(data, []string{coll, path, "SUCCESS"})
	} else {
		data = append(data, []string{coll, path, "FAILURE"})
	}

	ccli.output.Write(data)

	return err
}

func (ccli *CCLI) ExportCollection(c *cli.Context) error {
	coll := c.String("name")
	path := c.String("path")

	data := [][]string{}
	data = append(data, []string{"Collection Name", "Destination", "Status"})

	err := ccli.cdb.ExportCollection(coll, path)

	if err == nil {
		data = append(data, []string{coll, path, "SUCCESS"})
	} else {
		return err
	}

	ccli.output.Write(data)

	return nil
}
