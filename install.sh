#!/bin/bash

# example:
# cd npwg-chapter3
# ../install.sh
# if you compile all in the dir, output file is desided by the folder name.
# if you compile single file, like:
# ../install.sh test.go
# the output exe name is desided by the file name.

#BASEDIR=$(dirname "$0")
#echo "$BASEDIR"

#$PWD
#echo $PWD 

mkdir -p ./bin

export GOPATH=$PWD/
export GOBIN=$PWD/bin

echo "build project.start" $(date +"%Y-%m-%d %H:%M:%S")

TARGET=./...

if [ $1 ]; then
	TARGET=$1
fi

go install -v -gcflags "-N -l" $TARGET

echo "build project end  " $(date +"%Y-%m-%d %H:%M:%S")
