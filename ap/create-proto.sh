#!/bin/sh
DIR=`pwd`

cd ../..
OUT_DIR=`pwd`



# clear
rm -rf $DIR/src/infrastructure/grpc/proto
mkdir -p $DIR/src/infrastructure/grpc/proto

export PATH="$PATH:$HOME/go/bin"

cd $DIR/proto
find . name "*.proto" -exec protoc --proto_path . --go_out=plugins=grpc:$OUT_DIR {} \;