#!/bin/sh
DIR=`pwd`

# clear
rm -rf $DIR/src/infrastructure/grpc/proto
mkdir -p $DIR/src/infrastructure/grpc/proto

export PATH="$PATH:$HOME/go/bin"