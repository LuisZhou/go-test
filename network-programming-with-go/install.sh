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

#https://stackoverflow.com/questions/10383299/how-do-i-configure-go-to-use-a-proxy
#https://security.stackexchange.com/questions/80853/relationship-between-rsa-diffie-hellman-key-exchange-pki-and-x-509
#http_proxy=127.0.0.1:8087 go get -insecure golang.org/x/crypto/blowfish
go install -v -gcflags "-N -l" $TARGET

echo "build project end  " $(date +"%Y-%m-%d %H:%M:%S")
