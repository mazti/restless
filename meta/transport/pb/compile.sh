#!/usr/bin/env sh

# Install proto3
# sudo apt-get install -y git autoconf automake libtool curl make g++ unzip
# git clone https://github.com/google/protobuf.git
# cd protobuf/
# ./autogen.sh
# ./configure
# make
# make check
# sudo make install
# sudo ldconfig # refresh shared library cache.
#
# Update protoc Go bindings via
#  go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
#
# See also
#  https://github.com/grpc/grpc-go/tree/master/examples

protoc meta.proto --go_out=plugins=grpc:.