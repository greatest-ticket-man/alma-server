#!/bin/sh

# ubuntu
sudo apt update
sudo apt -y install ant 
sudo apt -y install cmake
sudo apt -y install python
sudo apt -y install ccache

# VoltDB 9.3.1
rm -rf /tmp/voltdb

cd /tmp
git clone https://github.com/VoltDB/voltdb.git
cd /tmp/voltdb
git checkout refs/tags/voltdb-9.3.1
ant -Djmemcheck=NO_MEMCHECK

cd /tmp
sudo mv ./voltdb /usr/local/

# push path
echo 'export PATH=$PATH:/usr/local/voltdb/bin' >> $HOME/.bashrc
