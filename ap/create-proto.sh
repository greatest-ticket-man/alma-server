#!/bin/sh

DIR=`pwd`

cd ../..
OUT_DIR=`pwd`

# clear
rm -rf $DIR/src/infrastructure/grpc/proto
mkdir -p $DIR/src/infrastructure/grpc/proto

export PATH="$PATH:$HOME/go/bin"

cd $DIR/proto
find . -type d -name vendor -prune -o -type f -name "*.proto" | xargs --replace=AlmaProtoFile protoc --proto_path=. -Ivendor/protobuf/src AlmaProtoFile --go_out=plugins=grpc:${OUT_DIR}
