#!/bin/sh


# VoltDB
## xcode 
xcode-select --install

## java install
brew update

brew tap adoptopenjdk/openjdk/adoptopenjdk8
brew cask install adoptopenjdk8

## ccache
brew install ccache

## ant
brew install ant

## cmake
brew install cmake

## pull 
rm -rf /tmp/voltdb_install
mkdir /tmp/voltdb_install
cd /tmp/voltdb_install

git clone https://github.com/VoltDB/voltdb.git
cd voltdb

# 9.3.1
git checkout refs/tags/voltdb-9.3.1
ant
