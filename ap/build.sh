#!/bin/sh

# statik
go get github.com/rakyll/statik

rm -rf ./statik

# static file
statik -src=./asset/static -ns=asset/static/ -dest=./statik -p=static
statik -src=./asset/template -ns=asset/template/ -dest=./statik -p=template
