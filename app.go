package main

import (
	"github.com/urfave/cli/v2"
)

func (ccli *CCLI) getApp() *cli.App {

	app := &cli.App{
		Name:                 "CloverDB CLI",
		Usage:                "clover [flags] [command] [subcommand]",
		EnableBashCompletion: true,
		// Before:               altsrc.InitInputSourceWithContext(flags, NewYamlSourceFromProfileFunc("profile")),
		Before: func(c *cli.Context) error {
			var err error
			ccli.output, ccli.printer, err = setWriter(c)
			return err
		},
		Flags: geRootFlags(),
		Commands: []*cli.Command{
			{
				Name:    "open",
				Aliases: []string{"op"},
				Usage:   "opens cloverdb directory to test",
				Action:  ccli.open,
			},
			{
				Name:    "create",
				Aliases: []string{"cr"},
				Usage:   "creates a new cloverdb directory",
				Action:  ccli.create,
			},
			{
				Before: func(c *cli.Context) error {
					var err error
					ccli.cdb, err = openDB(c)
					return err
				},
				Name:    "collection",
				Aliases: []string{"coll", "collections"},
				Usage:   "perform actions against a cloverdb collection",
				Subcommands: []*cli.Command{
					{
						Name:     "list",
						Usage:    "list all cloverdb collections",
						Action:   ccli.ListCollections,
						Category: "list",
					},
					{
						Name:     "exists",
						Usage:    "[--name <collection name>]",
						Action:   ccli.HasCollection,
						Category: "list",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "name",
								Usage:       "[--name <collection name>]",
								DefaultText: "",
								Required:    true,
							},
						},
					},
					{
						Name:     "create",
						Usage:    "[--name <collection name>]",
						Action:   ccli.CreateCollection,
						Category: "new",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "name",
								Usage:       "[--name <collection name>]",
								DefaultText: "",
								Required:    true,
							},
						},
					},
					{
						Name:     "drop",
						Usage:    "[--name <collection name>]",
						Action:   ccli.DropCollection,
						Category: "delete",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "name",
								Usage:       "[--name <collection name>]",
								DefaultText: "",
								Required:    true,
							},
						},
					},
					{
						Name:     "import",
						Usage:    "[--path <path to json file> --name <collection name>]",
						Action:   ccli.ImportCollection,
						Category: "import",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "name",
								Usage:       "[--name <collection name>]",
								DefaultText: "",
								Required:    true,
							},
							&cli.StringFlag{
								Name:        "path",
								Usage:       "[--path <path to json file>]",
								DefaultText: "",
								Required:    true,
							},
						},
					},
					{
						Name:     "export",
						Usage:    "[--name <collection name> --path <destination file>]",
						Action:   ccli.ExportCollection,
						Category: "export",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "name",
								Usage:       "[--name <collection name>]",
								DefaultText: "",
								Required:    true,
							},
							&cli.StringFlag{
								Name:        "path",
								Usage:       "[--path <path to json file>]",
								DefaultText: "",
								Required:    true,
							},
						},
					},
				},
			},
			{
				Before: func(c *cli.Context) error {
					var err error
					ccli.cdb, err = openDB(c)
					return err
				},
				Name:    "doc",
				Aliases: []string{"document", "documents"},
				Usage:   "insert, query, update, delete documents",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "collection-name",
						Aliases:     []string{"coll-name", "collection"},
						Usage:       "[--name <collection name>]",
						DefaultText: "",
						Required:    true,
					},
				},
				Subcommands: []*cli.Command{
					{
						Name:     "list",
						Usage:    "retrieve all documents",
						Action:   ccli.ListDocuments,
						Category: "list",
					},
					{
						Name:     "insert",
						Usage:    "[--file-name <filename>]",
						Action:   ccli.InsertDocument,
						Category: "create",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "file-name",
								Usage:       "[--file-name <filename>]",
								DefaultText: "",
								Required:    true,
							},
						},
					},
					{
						Name:     "insert-batch",
						Usage:    "[--file-name <filename>]",
						Action:   ccli.InsertBatchDocument,
						Category: "create",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "file-name",
								Usage:       "[--file-name <filename>]",
								DefaultText: "",
								Required:    true,
							},
						},
					},
					{
						Name:     "query",
						Usage:    "[--doc-query <field:value>]",
						Action:   ccli.Query,
						Category: "query",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "doc-query",
								Usage:       "[--doc-query <field=value>]",
								DefaultText: "",
								Required:    true,
							},
						},
					},
					{
						Name:     "query-one",
						Usage:    "[--doc-id <document id>] || [--doc-query <field:value>]",
						Action:   ccli.QueryOne,
						Category: "query",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:        "doc-id",
								Usage:       "[--doc-id <document id>]",
								DefaultText: "",
							},
							&cli.StringFlag{
								Name:        "doc-query",
								Usage:       "[--doc-query <field:value>]",
								DefaultText: "",
							},
						},
					},
				},
			},
		},
	}

	return app
}
