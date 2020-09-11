#!/bin/sh

FILENAME='alma-ap'

# statik
go get github.com/rakyll/statik

rm -rf ./statik

# static file
statik -src=./asset/static -ns=static -dest=./statik -p=static
statik -src=./asset/template -ns=template -dest=./statik -p=template
statik -src=./asset -ns=asset

GOOS='linux'
GOARCH='amd64'

LDFLAGS="-w -s"
LDFLAGS="${LDFLAGS} -X \"main.hash=`git rev-parse --verify HEAD`\""
LDFLAGS="${LDFLAGS} -X \"main.builddate=$(date '+%Y/%m/%d %H:%M:%S(%Z)')\""
LDFLAGS="${LDFLAGS} -X \"main.goversion=$(go version)\""
LDFLAGS="${LDFLAGS} -X \"main.goos=$GOOS\""
LDFLAGS="${LDFLAGS} -X \"main.goarch=$GOARCH\""

# delete 
if [ -e "${FILENAME}" ]; then
  rm -rf ./$FILENAME
fi

# build
export GOOS=$GOOS
export GOARCH=$GOARCH
go build -ldflags "${LDFLAGS}" -o $FILENAME main.go
