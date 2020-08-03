#!/bin/sh -e

# go test -v -cover -count=1 -failfast alma-server/ap/src/... | grep -i "fail"
go test -v -cover -count=1 -failfast alma-server/ap/src/...
echo '☆★☆★　test fail none　☆★☆★'
