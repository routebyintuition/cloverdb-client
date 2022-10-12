package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ostafen/clover"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

type docMap map[string]interface{}

func (ccli *CCLI) ListDocuments(c *cli.Context) error {
	docs, err := ccli.cdb.Query(c.String("collection-name")).FindAll()
	if err != nil {
		log.Error().Err(err).Msg("QueryAll error")
		return err
	}

	for _, doc := range docs {
		itemDoc := ccli.DocToSlice(doc)

		ccli.output.Write(itemDoc)

	}

	return nil
}

func (ccli *CCLI) InsertDocument(c *cli.Context) error {
	path := c.String("file-name")

	insertDoc := make(docMap)

	data := [][]string{}
	data = append(data, []string{"File", "Status", "Details"})

	jsonSrc, err := os.Open(path)
	defer jsonSrc.Close()
	if err != nil {
		data = append(data, []string{path, "Failed", err.Error()})
		ccli.output.Write(data)

		return err
	}
	jsonByte, err := ioutil.ReadAll(jsonSrc)
	if err != nil {
		data = append(data, []string{path, "Failed", err.Error()})
		ccli.output.Write(data)

		return err
	}
	json.Unmarshal(jsonByte, &insertDoc)

	newDoc := clover.NewDocumentOf(&insertDoc)
	if newDoc == nil {
		data = append(data, []string{path, "Failed", "json invalid...must be map[string]interface{}"})
		ccli.output.Write(data)

		return errors.New("could not convert provided json to clover document")
	}
	docId, err := ccli.cdb.InsertOne(c.String("collection-name"), newDoc)
	if err != nil {
		data = append(data, []string{path, "Failed", err.Error()})
		ccli.output.Write(data)

		return err
	}
	data = append(data, []string{path, "Success", docId})
	ccli.output.Write(data)

	return nil
}

func (ccli *CCLI) InsertBatchDocument(c *cli.Context) error {
	path := c.String("file-name")

	insertDocs := make([]docMap, 1)
	line := 0

	data := [][]string{}
	data = append(data, []string{"File", "Line", "Status", "Details"})
	jsonSrc, err := os.Open(path)
	defer jsonSrc.Close()
	if err != nil {
		data = append(data, []string{path, strconv.Itoa(line), "Failed", err.Error()})
		ccli.output.Write(data)

		return err
	}
	jsonByte, err := ioutil.ReadAll(jsonSrc)
	if err != nil {
		data = append(data, []string{path, strconv.Itoa(line), "Failed", err.Error()})
		ccli.output.Write(data)

		return err
	}
	json.Unmarshal(jsonByte, &insertDocs)
	for _, docVal := range insertDocs {
		newDoc := clover.NewDocumentOf(&docVal)
		if newDoc == nil {
			data = append(data, []string{path, strconv.Itoa(line), "Failed", "json invalid...must be map[string]interface{}"})
			ccli.output.Write(data)

			return errors.New("could not convert provided json to clover document")
		}
		docId, err := ccli.cdb.InsertOne(c.String("collection-name"), newDoc)
		if err != nil {
			data = append(data, []string{path, strconv.Itoa(line), "Failed", err.Error()})
			ccli.output.Write(data)

			return err
		}
		data = append(data, []string{path, strconv.Itoa(line), "Success", docId})
		line++
	}

	ccli.output.Write(data)

	return nil
}

func (ccli *CCLI) Query(c *cli.Context) error {
	if c.String("doc-query") == "" {
		return errors.New("query parameter not provided")
	}
	count := 0

	queryData := strings.Split(c.String("doc-query"), "=")
	findQuery := ccli.cdb.Query(c.String("collection")).Where(clover.Field(queryData[0]).Eq(queryData[1]))
	val, err := findQuery.FindAll()
	if err != nil {
		return err
	}

	for _, doc := range val {
		itemDoc := ccli.DocToSlice(doc)
		ccli.output.Write(itemDoc)
		count++
	}

	data := [][]string{}
	data = append(data, []string{"Count", "Status"})
	data = append(data, []string{strconv.Itoa(count), "success"})
	ccli.output.Write(data)

	return nil
}

func (ccli *CCLI) QueryOne(c *cli.Context) error {
	if c.String("doc-id") == "" && c.String("doc-query") == "" {
		return errors.New("neither doc-id or query provided")
	}

	doc, err := ccli.cdb.Query(c.String("collection-name")).FindById(c.String("doc-id"))
	if err != nil {
		return err
	}

	itemDoc := ccli.DocToSlice(doc)

	ccli.output.Write(itemDoc)

	return nil
}

func (ccli *CCLI) DocToSlice(doc *clover.Document) [][]string {
	docMap := make(map[string]interface{})
	doc.Unmarshal(&docMap)
	keyList := []string{}
	for docKey := range docMap {
		keyList = append(keyList, docKey)
	}
	sort.Strings(keyList)
	itemDoc := [][]string{}
	itemDoc = append(itemDoc, keyList)
	docList := []string{}
	for _, itemVal := range keyList {
		docList = append(docList, docMap[itemVal].(string))
	}
	itemDoc = append(itemDoc, docList)

	return itemDoc
}
