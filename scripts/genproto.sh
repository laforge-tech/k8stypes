#!/bin/bash

go install k8s.io/code-generator/cmd/go-to-protobuf@v0.26.0-alpha.1
go install k8s.io/code-generator/cmd/go-to-protobuf/protoc-gen-gogo@v0.26.0-alpha.1

APIMACHINERY_PKGS=(
    +k8s.io/apimachinery/pkg/util/intstr
    +k8s.io/apimachinery/pkg/api/resource
    +k8s.io/apimachinery/pkg/runtime/schema
    +k8s.io/apimachinery/pkg/runtime
    k8s.io/apimachinery/pkg/apis/meta/v1
    k8s.io/api/core/v1
    k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1
)

PKG_BASE="github.com/laforge-tech/k8stypes"

echo $(IFS=, ; echo "${APIMACHINERY_PKGS[*]}")

go-to-protobuf -h scripts/header.go.txt \
    -p ${PKG_BASE}/pkg/api/demo/v1 \
    --apimachinery-packages $(IFS=, ; echo "${APIMACHINERY_PKGS[*]}") \
     --proto-import=$PWD/vendor \
    --proto-import=/usr/include \
    -v 5
    
#    -o api/proto -v 5
