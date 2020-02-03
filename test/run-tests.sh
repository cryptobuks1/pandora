#!/usr/bin/env bash

BASEDIR=$(dirname "$0")
cd "${BASEDIR}"/.. || exit

go test \
  -o ./test/node/node.test \
  -c -covermode=count -coverpkg ./pkg/...,./cmd/node \
  ./cmd/node/ || exit

if eval go test -v -race -timeout 10m -count=1 ./test/node/; then
  echo "TESTS ARE OK"
else
  echo "TESTS AREN'T OK"
fi

rm -rf ./test/node/coverage.out ./test/node/node.test || exit
