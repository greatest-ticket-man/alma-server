#!/bin/sh

# ラズパイ用にコンパイル
echo "build"
GOOS=linux GOARCH=arm go build

# バイナリ
scp -p ./ap root@192.168.2.201:/home/www

# configとasset
scp -pr ./asset root@192.168.2.201:/home/www
scp -pr ./config root@192.168.2.201:/home/www
