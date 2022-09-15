#!/bin/bash
env=$1
server_path=/data/code

if [[ -z "${env}" ]]; then
    echo "Please input environment argument Env=[prod|test]"
    exit 1
fi

chmod +x ${server_path}/super-view-server
sudo  chown -R root:root /data/

${server_path}/super-view-server --config=${env} >> /data/logs/super-view-server_start.log 2>&1
