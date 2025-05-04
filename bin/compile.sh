#!/bin/bash

function exit_if() {
    extcode=$1
    msg=$2
    if [ $extcode -ne 0 ]
    then
        if [ "msg$msg" != "msg" ]; then
            echo $msg >&2
        fi
        exit $extcode
    fi
}

echo $GOPATH

# 检查 protoc-gen-go 是否安装
if [ ! -f "$(go env GOPATH)/bin/protoc-gen-go" ]
then
    echo 'Protocol Buffers plugin for Go is not installed.' >&2
    echo 'Please install it using:' >&2
    echo 'go install google.golang.org/protobuf/cmd/protoc-gen-go@latest' >&2
    echo 'go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest' >&2
    exit 1
else
    echo Compiling go interfaces...
    export GO_PATH=$(go env GOPATH)
    export GOBIN=$GO_PATH/bin
    export PATH=$PATH:$GO_PATH/bin

    protoc -I ./ --go_out=./ --go-grpc_out=require_unimplemented_servers=false:. protobuf/*.proto

    exit_if $?
    echo Done
fi