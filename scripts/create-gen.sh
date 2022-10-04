#!/bin/sh

go install  sigs.k8s.io/controller-tools/cmd/controller-gen@v0.10.0

controller-gen crd object paths=./pkg/api/...