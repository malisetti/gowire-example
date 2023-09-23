#!/usr/bin/env bash

set -ex

./wire.sh
./test.sh
# -transform 1
go run example/cmd/hello "${@:1}"
