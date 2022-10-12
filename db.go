package main

import (
	"errors"

	"github.com/ostafen/clover"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func (ccli *CCLI) open(c *cli.Context) error {
	valid, err := dirValid(c.String("dir"))
	if err != nil {
		return err
	}

	if !valid {
		return errors.New("directory location not valid")
	}

	cdb, err := clover.Open(c.String("dir"))
	if err != nil {
		return err
	}

	log.Debug().Msg("attempting to list collections")
	_, err = cdb.ListCollections()
	if err != nil {
		return err
	}

	log.Debug().Msg("success on opening database directory")

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

	log.Debug().Int("collection count", len(listColl)).Msg("successfully created database")

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
