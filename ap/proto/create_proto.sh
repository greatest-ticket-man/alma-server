#!/bin/sh


SERVER_OUTPUT_DIR=server
CLIENT_OUTDIR=client

mkdir -p ${CLIENT_OUTDIR} ${SERVER_OUTPUT_DIR}

protoc --proto_path=. -Ivendor/protobuf/src member/member.proto \
    --go_out=plugins=grpc:${SERVER_OUTPUT_DIR}
