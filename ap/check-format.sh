#!/bin/sh -e

go fmt main.go
go vet main.go
go vet -vettool=$(which shadow) main.go
go mod tidy
golint main.go | grep -v "don't use MixedCaps" | grep -v "don't use ALL_CAPS in Go names" && echo && exit 1 || true

${HOME}/staticcheck ./...

echo "========= static check ok ==========="
