// This is a submodule to isolate github.com/Kotodian/code-generator from k8s.io/{api,apimachinery,client-go} dependencies in generated code

module github.com/Kotodian/code-generator/examples

go 1.16

require (
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/client-go v0.0.0
	k8s.io/kube-openapi v0.0.0-20210421082810-95288971da7e
)

replace (
	k8s.io/api => ../../api
	k8s.io/apimachinery => ../../apimachinery
	k8s.io/client-go => ../../client-go
)
