
mkdir -p ./bin

export GOPATH=$PWD/
export GOBIN=$PWD/bin

${@:1}