#!/bin/sh -ex

# git submodule
git submodule init
git submodule update

DIR=`pwd`

cd ../..
OUT_DIR=`pwd`

# clear
rm -rf $DIR/src/infrastructure/grpc/proto
mkdir -p $DIR/src/infrastructure/grpc/proto

export PATH="$PATH:$HOME/go/bin"

cd $DIR/proto
# find . -type d -name vendor -prune -o -type f -name "*.proto" | xargs echo # ./vendorも読み込まれてるから変なメッセージでるみたい、どうにかしたい
# find . -type d -name vendor -prune -o -type f -name "*.proto" | xargs --replace=AlmaProtoFile protoc --proto_path=. -Ivendor/protobuf/src AlmaProtoFile --go_out=plugins=grpc:${OUT_DIR}


# find . -type d -name vendor -prune -o -type f -name "*.proto" -exec echo {} \; | xargs --replace=AlmaProtoFile protoc --proto_path=. -Ivendor/protobuf/src AlmaProtoFile --go_out=plugins=grpc:${OUT_DIR}
# find . -type d -name vendor -prune -o -type f -name "*.proto" -exec echo {} \; | xargs --replace=AlmaProtoFile protoc --proto_path=. -Ivendor/protobuf/src AlmaProtoFile --go_out=plugins=grpc:${OUT_DIR}

# macはreplace optionをサポートしていない
find . -type d -name vendor -prune -o -type f -name "*.proto" -exec echo {} \; | xargs -I AlmaProtoFile protoc --proto_path=. -Ivendor/protobuf/src AlmaProtoFile --go_out=plugins=grpc:${OUT_DIR}

echo "==== create proto fin ===="