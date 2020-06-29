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
rm -rf ~/voltdb
cd ~

git clone https://github.com/VoltDB/voltdb.git
cd voltdb

# 9.3.1
git checkout refs/tags/voltdb-9.3.1
# ant
# performance 
ant -Djmemcheck=NO_MEMCHECK

# path
PWD=$(pwd)
echo "Pathを通してください"
echo "\$PATH:$PWD/bin/"
PATH="$PATH:$PWD/bin/"
