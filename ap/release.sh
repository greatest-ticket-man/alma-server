#!/bin/sh

# ラズパイ用にコンパイル
echo "build"
GOOS=linux GOARCH=arm go build

# kill
ssh -t www@221.170.118.102 'sudo ./stop-alma.sh'

# バイナリ
scp -p ./ap www@221.170.118.102:/home/www

# configとasset
scp -pr ./asset www@221.170.118.102:/home/www
scp -pr ./config www@221.170.118.102:/home/www

# start
# ssh -t www@221.170.118.102 'sudo ./start-alma.sh'
echo "直接Serverでstartは実行する必要があります"

# バイナリを削除
rm ap
