module github.com/banzaicloud/operator-tools

go 1.13

require (
	emperror.dev/errors v0.4.2
	github.com/MakeNowJust/heredoc v1.0.0
	github.com/andreyvit/diff v0.0.0-20170406064948-c7f18ee00883
	github.com/banzaicloud/k8s-objectmatcher v1.1.0
	github.com/go-logr/logr v0.1.0
	github.com/go-logr/zapr v0.1.0
	github.com/goph/emperror v0.17.2
	github.com/iancoleman/orderedmap v0.0.0-20190318233801-ac98e3ecb4b0
	github.com/sergi/go-diff v1.1.0 // indirect
	go.uber.org/zap v1.9.1
	k8s.io/api v0.16.4
	k8s.io/apiextensions-apiserver v0.16.4
	k8s.io/apimachinery v0.16.4
	sigs.k8s.io/controller-runtime v0.4.0
)
