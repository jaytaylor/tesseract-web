# Tesseract Web

[![Documentation](https://godoc.org/github.com/jaytaylor/tesseract-web?status.svg)](https://godoc.org/github.com/jaytaylor/tesseract-web)
[![Build Status](https://travis-ci.org/jaytaylor/tesseract-web.svg?branch=master)](https://travis-ci.org/jaytaylor/tesseract-web)
[![Report Card](https://goreportcard.com/badge/github.com/jaytaylor/tesseract-web)](https://goreportcard.com/report/github.com/jaytaylor/tesseract-web)

## Supported platforms:

* Linux
* Windows
* Mac OS X / macOs

## Requirements

* [Tesseract OCR](https://github.com/tesseract-ocr/tesseract) must be installed and available on the `$PATH`.

## Get it

    git clone https://github.com/jaytaylor/tesseract-web
    cd tesseract-web
    go get -t ./...
    make
    make install

## Run it

    tesseract-web -bind 127.0.0.1:8080

## Example usage:

Remote URL:

    curl -XPOST localhost:8080/v1/tesseract/https://i.imgur.com/14y5P0u.png

Upload image:

    curl -XPOST localhost:8080/v1/tesseract -d@path/to/some/image.jpg

## Run the tests

    make test
    # or
    go test ./...

or for verbose output:

    make test flags=-v
    # or
    go test -v ./...

## About
```(shell)
NAME:
   tesseract-web - Exposes tesseract image OCR as a set of web APIs

USAGE:
   tesseract-web [global options] command [command options] [arguments...]

VERSION:
   0.1.0
built on 2016-10-12 22:54:38 +0000 UTC
git commit 430c047318453fd6983ffa99630341f0526e9cb9

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --install               Install tesseract-web as a system service (default: false)
   --uninstall             Uninstall tesseract-web as a system service (default: false)
   --user value, -u value  Specifies the user to run the tesseract-web system service as (will default to "jay") (default: "jay")
   --bind value, -b value  Set the web-server bind-address and (optionally) port (default: "0.0.0.0:8080")
   --help, -h              show help (default: false)
   --version, -v           print the version (default: false)
```

### License

[MIT](LICENSE)

