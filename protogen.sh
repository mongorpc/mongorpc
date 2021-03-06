#!/bin/sh
# git submodule foreach git pull origin main
git pull && git submodule init && git submodule update && git submodule status

set -e
# remove lib/mongorpc directory if it exists
if [ -d lib/mongorpc ]; then
    rm -rf lib/mongorpc
fi

protoc --proto_path=proto --go_out=lib --go_opt=paths=source_relative \
    --go-grpc_out=lib --go-grpc_opt=paths=source_relative \
    proto/mongorpc/*.proto