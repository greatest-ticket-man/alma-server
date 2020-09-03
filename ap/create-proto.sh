#!/bin/sh
# DIR=`pwd`

# cd ../..
# OUT_DIR=`pwd`



# # clear
# rm -rf $DIR/src/infrastructure/grpc/proto
# mkdir -p $DIR/src/infrastructure/grpc/proto

# export PATH="$PATH:$HOME/go/bin"

# cd $DIR/proto
# find . -name "*.proto" -type f -exec protoc --proto_path . --go_out=plugins=grpc:$OUT_DIR {} \;

DIR = `pwd`

cd ../..
OUT_DIR=`pwd`

# clear
rm -rf $DIR/src/infrastructure/grpc/proto
mkdir -p $DIR/src/infrastructure/grpc/proto

export PATH="$PATH:$HOME/go/bin"

cd $DIR/proto

find . -type d -name vendor -prune -o -type f -name "*.proto" -exec protoc --proto_path=. -Ivendor/protobuf/src --go_out=plugins=grpc:${OUT_DIR} {}\;
