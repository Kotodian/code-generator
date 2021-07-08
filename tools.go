// +build tools

/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This package contains code generation utilities
// This package imports things required by build scripts, to force `go mod` to see them as dependencies
package tools

import (
	_ "github.com/Kotodian/code-generator/cmd/client-gen"
	_ "github.com/Kotodian/code-generator/cmd/conversion-gen"
	_ "github.com/Kotodian/code-generator/cmd/deepcopy-gen"
	_ "github.com/Kotodian/code-generator/cmd/defaulter-gen"
	_ "github.com/Kotodian/code-generator/cmd/go-to-protobuf-mine"
	_ "github.com/Kotodian/code-generator/cmd/import-boss"
	_ "github.com/Kotodian/code-generator/cmd/informer-gen"
	_ "github.com/Kotodian/code-generator/cmd/lister-gen"
	_ "github.com/Kotodian/code-generator/cmd/openapi-gen"
	_ "github.com/Kotodian/code-generator/cmd/register-gen"
	_ "github.com/Kotodian/code-generator/cmd/set-gen"
)
