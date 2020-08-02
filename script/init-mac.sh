#!/bin/sh


# # VoltDB
# ## xcode 
# xcode-select --install

# ## java install
# brew update

# brew cask install adoptopenjdk/openjdk/adoptopenjdk8

# ## ccache
# brew install ccache

# ## ant
# brew install ant

# ## cmake
# brew install cmake

# ## pull 
# rm -rf ~/voltdb
# cd ~

# git clone https://github.com/VoltDB/voltdb.git
# cd voltdb

# # 9.3.1
# git checkout refs/tags/voltdb-9.3.1
# # ant
# # performance 
# ant -Djmemcheck=NO_MEMCHECK

# path
# PWD=$(pwd)
# echo "Pathを通してください"
# echo "\$PATH:$PWD/bin/"
# PATH="$PATH:$PWD/bin/"


cd /tmp
 go get github.com/golang/protobuf/{proto,protoc-gen-go}
 go get google.golang.org/grpc
 go install github.com/golang/protobuf/protoc-gen-go

# protoc main
# wget https://github.com/protocolbuffers/protobuf/releases/download/v3.11.4/protoc-3.11.4-osx-x86_64.zip

wget --no-check-certificate https://github.com/protocolbuffers/protobuf/releases/download/v3.12.4/protoc-3.12.4-osx-x86_64.zip

# unzip protoc-3.11.4-osx-x86_64.zip
unzip protoc-3.12.4-osx-x86_64.zip
sudo mv ./bin/protoc /usr/local/bin/
/usr/local/bin/protoc --version
rm -f protoc-3.12.4-osx-x86_64.zip

# protoc gen-doc
go get github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
# protoc --doc_out=/tmp/aaa --doc_opt=html,index.html `find proto -name "*.proto"`

# grpcurl
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
grpcurl --help

## static check tools

# golint install
go get golang.org/x/lint/golint
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow

# staticchek install
cd /tmp
git clone https://github.com/dominikh/go-tools
cd go-tools/cmd/staticcheck
GOOS=darwin GOARCH=amd64 go build -o staticcheck staticcheck.go
mv ./staticcheck $HOME/
cd /tmp
rm -rf /tmp/go-tools

