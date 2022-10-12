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



