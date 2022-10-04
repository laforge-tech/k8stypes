#!/bin/sh

go install k8s.io/code-generator/cmd/applyconfiguration-gen@v0.26.0-alpha.1
go install k8s.io/code-generator/cmd/client-gen@v0.26.0-alpha.1
go install k8s.io/code-generator/cmd/lister-gen@v0.26.0-alpha.1
go install k8s.io/code-generator/cmd/informer-gen@v0.26.0-alpha.1

PKG_BASE="github.com/laforge-tech/k8stypes"

echo Generating applyconfigurations
applyconfiguration-gen -h scripts/header.go.txt \
    -i ${PKG_BASE}/pkg/api/demo/v1 \
    --output-package ${PKG_BASE}/pkg/applyconfigurations \
    -h scripts/header.go.txt \
    -o . \
    --trim-path-prefix ${PKG_BASE}

echo Generating client
client-gen -h scripts/header.go.txt \
    -n client \
    --input-base ${PKG_BASE}/pkg/api \
    --input demo/v1 \
    --apply-configuration-package ${PKG_BASE}/pkg/applyconfigurations \
    --output-package ${PKG_BASE}/pkg \
    -o . \
    --trim-path-prefix ${PKG_BASE}

echo Generating lister
lister-gen -h scripts/header.go.txt \
    --input-dirs ${PKG_BASE}/pkg/api/demo/v1 \
    --output-package ${PKG_BASE}/pkg/listers \
    --trim-path-prefix ${PKG_BASE} \
    -o .

echo Generating informer
informer-gen -h scripts/header.go.txt \
    --versioned-clientset-package ${PKG_BASE}/pkg/client \
    --input-dirs ${PKG_BASE}/pkg/api/demo/v1 \
    --listers-package ${PKG_BASE}/pkg/listers \
    --output-package ${PKG_BASE}/pkg/informers \
    --trim-path-prefix ${PKG_BASE} \
    -o .
