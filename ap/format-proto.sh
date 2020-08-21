#!/bin/sh 

# ubuntu
# sudo apt install -y clang-format
# mac
# brew install clang-format

cd proto
find ./ -name "*.proto" | xargs clang-format -i
