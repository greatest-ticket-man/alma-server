#!/bin/sh

# set -eu

VERSION='0.0.1'
workspce=`pwd`

mkdir deploy 2> /dev/null

cd ../../../../ap

./build.sh

cp ./alma-ap ${workspce}/deploy/
cp -r ./config ${workspce}/deploy/

cd $workspce

# build image
docker build --tag sunjin110/alma:${VERSION} .

# build push
docker push sunjin110/alma:${VERSION}

# rm
rm -rf ${workspce}/deploy