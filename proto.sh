#!/bin/bash

dir="app/grpc/protos"
command="protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative"
find $dir -maxdepth 3 -type f -name "*.go" | while read p; do
    rm -rf $p    
done

find $dir -maxdepth 3 -type f -name "*.proto" | while read p; do
    $command $p    
done