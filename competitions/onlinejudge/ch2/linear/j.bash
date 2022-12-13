#!/bin/bash

set -e


mkdir $1
cd $1
cat ../template.py  > ./'p'$1'.py'
touch input1.txt output1.txt input.txt output.txt

