package main

import (
	"errors"
	"strconv"

	"github.com/ostafen/clover"
	"github.com/urfave/cli/v2"
)

func (ccli *CCLI) open(c *cli.Context) error {
	data := [][]string{}
	data = append(data, []string{"Database", "Status"})
	valid, err := dirValid(c.String("dir"))
	if err != nil {
		data = append(data, []string{c.String("dir"), err.Error()})
		ccli.output.Write(data)

		return err
	}

	if !valid {
		data = append(data, []string{c.String("dir"), "directory location is not valid"})
		ccli.output.Write(data)

		return errors.New("directory location not valid")
	}

	cdb, err := clover.Open(c.String("dir"))
	if err != nil {
		data = append(data, []string{c.String("dir"), err.Error()})
		ccli.output.Write(data)

		return err
	}

	_, err = cdb.ListCollections()
	if err != nil {
		data = append(data, []string{c.String("dir"), err.Error()})
		ccli.output.Write(data)

		return err
	}

	data = append(data, []string{c.String("dir"), "success"})
	ccli.output.Write(data)

	return nil
}

func (ccli *CCLI) create(c *cli.Context) error {
	valid, err := dirValid(c.String("dir"))
	if err != nil {
		return err
	}

	if valid {
		return errors.New("database already exists")
	}

	cdb, err := clover.Open(c.String("dir"))
	if err != nil {
		return err
	}

	listColl, err := cdb.ListCollections()
	if err != nil {
		return err
	}

	data := [][]string{}
	data = append(data, []string{"Database", "Collection Count", "Status"})
	data = append(data, []string{c.String("dir"), strconv.Itoa(len(listColl)), "created"})

	ccli.output.Write(data)

	return nil
}

func openDB(c *cli.Context) (*clover.DB, error) {
	valid, err := dirValid(c.String("dir"))
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, errors.New("directory location not valid")
	}
	cdb, err := clover.Open(c.String("dir"))
	if err != nil {
		return nil, err
	}

	_, err = cdb.ListCollections()
	if err != nil {
		return nil, err
	}

	return cdb, nil
}
