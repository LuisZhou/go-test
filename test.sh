#!/bin/bash

export GOPATH=$PWD/
export GOBIN=$PWD/bin

go test ${@:1}

