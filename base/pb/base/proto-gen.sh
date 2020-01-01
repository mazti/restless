#!/bin/sh

PROTO_FILES=$(find pb -name "*.proto")
for PROTO_FILE in ${PROTO_FILES}; do
    protoc -I=/usr/local/include -I. \
        -I${GOPATH}/src \
        -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --go_out=plugins=grpc:. \
        ${PROTO_FILE}
    protoc -I/usr/local/include -I. \
        -I=${GOPATH}/src \
        -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --grpc-gateway_out=logtostderr=true:. \
        ${PROTO_FILE}
    protoc -I/usr/local/include -I. \
        -I${GOPATH}/src \
        -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
        --swagger_out=logtostderr=true:. \
        ${PROTO_FILE}
done