#!/bin/sh

FILENAME='alma-ap'

# statik
go get github.com/rakyll/statik

rm -rf ./statik
rm -rf ./src/infrastructure/file/asset

# static file
statik -src=./asset/static -ns=static -dest=./src/infrastructure/file/asset -p=static
statik -src=./asset -ns=asset -dest=./src/infrastructure/file -p=asset

# statik.goになるので、ファイル名変更
mv ./src/infrastructure/file/asset/static/statik.go ./src/infrastructure/file/asset/static/static.go
mv ./src/infrastructure/file/asset/statik.go ./src/infrastructure/file/asset/asset.go

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
