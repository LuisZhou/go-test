#!/bin/bash

# example:
# cd npwg-chapter3
# ../install.sh
# output file is desided by the folder name.

export GOPATH=/home/ubuntu/code/go/go-test/
export GOBIN=/home/ubuntu/code/go/go-test/bin

echo "start to build project..." $(date +"%Y-%m-%d %H:%M:%S")

go install -v -gcflags "-N -l" ./...

echo "build project end" $(date +"%Y-%m-%d %H:%M:%S")