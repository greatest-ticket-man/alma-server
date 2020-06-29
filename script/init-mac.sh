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
rm -rf ~/voltdb
mkdir ~/voltdb
cd ~/voltdb

git clone https://github.com/VoltDB/voltdb.git
cd voltdb

# 9.3.1
git checkout refs/tags/voltdb-9.3.1
ant

# path
PATH="$PATH:$(pwd)/bin/"
