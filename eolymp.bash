#!/bin/bash

set -e


BASE_DIR=competitions/e-olymp/$1
mkdir -p $BASE_DIR

main="package main"
main_testing="package main"
echo $main > $BASE_DIR/solution.go
echo $main_testing > $BASE_DIR/solution_test.go
touch $BASE_DIR/description.txt
