#!/bin/sh


# VoltDB
## xcode 
xcode-select --install

## java install
brew update
brew cask install adoptopenjdk/openjdk/adoptopenjdk8

## ccache
brew install ccache

## ant
brew install ant

## cmake
brew install cmake

## pull 
rm -rf /tmp/voltdb
cd /tmp
git clone https://github.com/VoltDB/voltdb.git
cd voltdb

# 9.3.1
git checkout refs/tags/voltdb-9.3.1
ant -Djmemcheck=NO_MEMCHECK

cd /tmp
rm -rf /usr/local/voltdb
sudo mv ./voltdb /usr/local/

# push path
echo 'export PATH=$PATH:/usr/local/voltdb/bin' >> $HOME/.zshrc

source $HOME/.zshrc

voltdb init --dir $HOME/voltdb_data
