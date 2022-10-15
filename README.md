# cloverdb-client

Command line utility for use with the CloverDB embedded document database engine. More information about CloverDB can be found here:

[https://github.com/ostafen/clover/tree/v1.2.0](https://github.com/ostafen/clover/tree/v1.2.0)

This client is used for the current stable 1.2.0 branch of CloverDB as the v2 branch is still Alpha at the time of utility creation.

## Why?
Using CloverDB as embedded persistent storage for other projects has necessited a need to perform read actios against a development environment during application creation and testing. Since I needed this application for read workloads to ensure application functionality, it made sense to also include write capabilities in case others needed it.

## Install
Download the [latest release](https://github.com/routebyintuition/cloverdb-client/releases) specific to your operating system. There are packages for Linux, Mac, and Windows. 

```sh
tar -zxvf cloverdb-client_<version number>_<operating system>_x86_64.tar.gz
./cloverdb-client
```

## Build
The main branch is currently the development branch which will change soon. For now, it is recommended to use the [cloverdb-client releases](https://github.com/routebyintuition/cloverdb-client/releases)
```sh
go get -u github.com/routebyintuition/cloverdb-client
```

## Configuration
cloverdb-client can use the directory path to the local CloverDB storage passed as either command line arguments or set as an environmental variable.

```sh
HOST%> export CLOVER_DIR=/path/to/directory

HOST%> ./cloverdb-client coll list
+----+-----------------+
| ID | COLLECTION NAME |
+----+-----------------+
|  0 | coll1           |
|  1 | coll2           |
+----+-----------------+
```

Would perform the same action as:

Exported command line variables:
```sh
HOST%> ./cloverdb-client --dir /path/to/directory coll list
+----+-----------------+
| ID | COLLECTION NAME |
+----+-----------------+
|  0 | coll1           |
|  1 | coll2           |
+----+-----------------+
```

## Use

### Help

Below is the main help menu.

```bash
HOST%> ./cloverdb-client -h
NAME:
   CloverDB CLI - clover [flags] [command] [subcommand]

USAGE:
   CloverDB CLI [global options] command [command options] [arguments...]

COMMANDS:
   open, op                       opens cloverdb directory to test
   create, cr                     creates a new cloverdb directory
   collection, coll, collections  perform actions against a cloverdb collection
   doc, document, documents       insert, query, update, delete documents
   help, h                        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --dir value, --directory value, -d value, -D value  CloverDB directory (default: "../ttale/cdb") [$CLOVER_DIR, $CLOVER_DIRECTORY]
   --format value, --form value, -f value, -F value    output format <json> <table> (default: table) [$CLOVER_FORMAT]
   --help, -h                                          show help (default: false)
```

Each command also has a dedicated help menu.

```bash

HOST%> ./cloverdb-client open -h
NAME:
   CloverDB CLI open - opens cloverdb directory to test

USAGE:
   CloverDB CLI open [command options] [arguments...]

OPTIONS:
   --help, -h  show help (default: false)

HOST%> ./cloverdb-client collection -h
NAME:
   CloverDB CLI collection - perform actions against a cloverdb collection

USAGE:
   CloverDB CLI collection command [command options] [arguments...]

COMMANDS:
   list     list all cloverdb collections
   exists   [--name <collection name>]
   create   [--name <collection name>]
   drop     [--name <collection name>]
   import   [--path <path to json file> --name <collection name>]
   export   [--name <collection name> --path <destination file>]
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help (default: false)

HOST%> ./cloverdb-client documents -h
NAME:
   CloverDB CLI doc - insert, query, update, delete documents

USAGE:
   CloverDB CLI doc command [command options] [arguments...]

COMMANDS:
   list          retrieve all documents
   insert        [--file-name <filename>]
   insert-batch  [--file-name <filename>]
   query         [--doc-query <field:value>]
   query-one     [--doc-id <document id>] || [--doc-query <field:value>]
   help, h       Shows a list of commands or help for one command

OPTIONS:
   --collection-name value, --coll-name value, --collection value  [--name <collection name>]
   --help, -h                                                      show help (default: false)

```

### Database activities

Open database (basically just check if its valid):

```bash

HOST%> export CLOVER_DIR="/path/to/db"
HOST%> ./cloverdb-client open
+--------------+---------+
|   DATABASE   | STATUS  |
+--------------+---------+
| /path/to/db  | success |
+--------------+---------+

```

Create new database and error out if one already exists:

```bash

HOST%> export CLOVER_DIR="tmp2"
HOST%> ./cloverdb-client create
+----------+------------------+---------+
| DATABASE | COLLECTION COUNT | STATUS  |
+----------+------------------+---------+
| tmp2     |                0 | created |
+----------+------------------+---------+

```

### Collections

List collections:

```bash

HOST%> ./cloverdb-client collections list
+----+-----------------+
| ID | COLLECTION NAME |
+----+-----------------+
|  0 | config          |
|  1 | data            |
+----+-----------------+

```

Check if a collection already exists by name:

```bash

HOST%> ./cloverdb-client collections exists --name data
+-----------------+--------+
| COLLECTION NAME | EXISTS |
+-----------------+--------+
| data            | YES    |
+-----------------+--------+

```

Export a collection:

```bash

HOST%> ./cloverdb-client collections export --name data --path export.json
+-----------------+-------------+---------+
| COLLECTION NAME | DESTINATION | STATUS  |
+-----------------+-------------+---------+
| data            | export.json | SUCCESS |
+-----------------+-------------+---------+

HOST%> 

```

