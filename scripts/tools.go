//go:build tools
package script

import(
	_ "github.com/gogo/protobuf/protoc-gen-gogofast"
	_ "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

)
