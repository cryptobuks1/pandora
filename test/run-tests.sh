#!/usr/bin/env bash

BASEDIR=$(dirname "$0")
cd "${BASEDIR}"/.. || exit

go test \
  -o ./test/node/node.test \
  -c -covermode=count -coverpkg ./pkg/...,./cmd/node \
  ./cmd/node/ || exit

go test -v -timeout 10m -count=1 ./test/node/ || true

rm -rf ./test/node/coverage.out ./test/node/node.test || exit
