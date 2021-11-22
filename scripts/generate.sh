#!/usr/bin/env bash

set -ex

function setup_proto_deps {
    local dep dep_dir dep_location proto_deps
    dep_dir="${GOPATH}/deps"

    proto_deps=(
        "github.com/golang/protobuf"
        "github.com/mwitkow/go-proto-validators"
        "github.com/grpc-ecosystem/grpc-gateway"
    )

    for dep in "${proto_deps[@]}"; do
        dep_location="$(go list -m -f '{{ .Dir }}' "${dep}")"

        mkdir -p "${dep_dir}/$(dirname "${dep}")"
        rm -f "${dep_dir}/${dep}"
        ln -sf "${dep_location}" "${dep_dir}/${dep}"
    done

    go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
    go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
    go install github.com/golang/protobuf/protoc-gen-go
    go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators

    PATH="/go/bin":${PATH}
    export PATH
}

setup_proto_deps

#PACKAGE_DIR=${GOPATH}/src/github.com/lubovskiy/crud
#DST_DIR=${PACKAGE_DIR}/pkg/api
#INCLUDE="-I${PACKAGE_DIR}/api \
#-I${GOPATH}/deps/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#-I${GOPATH}/deps/github.com/grpc-ecosystem/grpc-gateway \
#-I/usr/local/include \
#-I${GOPATH}/deps \
#-I. -I${GOPATH}/src"
#
#rm -rf ${DST_DIR}/*.pb.*
#
## generate code
#protoc ${INCLUDE} \
#--go_out=plugins=grpc:${DST_DIR} \
#--grpc-gateway_out=logtostderr=true:${DST_DIR} \
#--govalidators_out=${DST_DIR} \
#${PACKAGE_DIR}/api/*.proto
#
## generate service.swagger.json
#protoc ${INCLUDE} \
#--swagger_out=${PACKAGE_DIR}/api \
#${PACKAGE_DIR}/api/service.proto
