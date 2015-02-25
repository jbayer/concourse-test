#!/bin/bash

set -e -x
set -o pipefail

go build -o app $(dirname $0)/main.go

PORT=8000 ./app &

pid=$!

trap "kill $pid" EXIT

sleep 1

curl http://localhost:8000 | grep "test"
