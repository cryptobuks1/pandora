#!/usr/bin/env bash

BASEDIR=$(dirname "$0")
cd "${BASEDIR}"/.. || exit

TEST_DIR=./test

go test \
  -o $TEST_DIR/node.test \
  -c -covermode=count -coverpkg ./pkg/...,./cmd/node \
  ./cmd/node/ || exit

TESTS_PASSED=false
if eval "go test -v -race -timeout 10m -count=1 '$TEST_DIR'/"; then
  TESTS_PASSED=true
else
  TESTS_PASSED=false
fi

rm -rf $TEST_DIR/coverage.out $TEST_DIR/node.test || exit

if [ $TESTS_PASSED = true ]; then
  echo "TESTS PASSED"
  exit 0
else
  echo "TESTS DIDN'T PASS"
  exit 1
fi
