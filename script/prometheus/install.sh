#!/bin/sh

mkdir /tmp/prometheus 2> /dev/null

cd /tmp/prometheus

workspace=`pwd`

wget https://github.com/prometheus/prometheus/releases/download/v2.21.0/prometheus-2.21.0.linux-amd64.tar.gz

tar xvfz prometheus-*.tar.gz
cd prometheus-*

# run
# ./prometheus --config.file=prometheus.yml

# node exporter
cd $workspace
wget https://github.com/prometheus/node_exporter/releases/download/v1.0.1/node_exporter-1.0.1.linux-amd64.tar.gz

tar -xzvf node_exporter-*.*.tar.gz

cd node_exporter-*.*

# ./node_exporter --web.listen-address 127.0.0.1:8080
# ./node_exporter --web.listen-address 127.0.0.1:8081
# ./node_exporter --web.listen-address 127.0.0.1:8082

# grafana
cd $workspace
wget https://dl.grafana.com/oss/release/grafana-7.2.0.linux-amd64.tar.gz
tar -zxvf grafana-7.2.0.linux-amd64.tar.gz
