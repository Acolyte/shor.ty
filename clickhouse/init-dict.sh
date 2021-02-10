#!/bin/bash
set -e

if [ ! -d "/etc/clickhouse-server/dicts" ]; then
    cd /etc/clickhouse-server
    tar -xzf dicts.tar.gz
fi
