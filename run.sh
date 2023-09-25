#!/usr/bin/env bash

set -ex

./wire.sh
./test.sh
# -port 3333
go run example/cmd/hello "${@:1}"
