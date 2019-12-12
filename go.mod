module go.awsctrl.io/manager

go 1.13

require (
	github.com/appscode/jsonpatch v0.0.0-00010101000000-000000000000 // indirect
	github.com/aws/aws-sdk-go v1.25.36
	github.com/awslabs/goformation/v3 v3.1.0
	github.com/davecgh/go-spew v1.1.1
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/go-logr/logr v0.1.0
	github.com/go-openapi/validate v0.19.5 // indirect
	github.com/gobuffalo/envy v1.8.1 // indirect
	github.com/gobuffalo/flect v0.1.7 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/gregjones/httpcache v0.0.0-20180305231024-9cad4c3443a7 // indirect
	github.com/huandu/xstrings v1.2.1 // indirect
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/onsi/ginkgo v1.10.3
	github.com/onsi/gomega v1.7.1
	github.com/rogpeppe/go-internal v1.5.0 // indirect
	github.com/satori/go.uuid v1.2.0
	go.awsctrl.io/generator v0.0.0-20191212173927-7592c34721bd // indirect
	go.etcd.io/etcd v0.0.0-20191023171146-3cf2f69b5738 // indirect
	go.hein.dev/go-version v0.1.0
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413 // indirect
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	golang.org/x/tools v0.0.0-20191213221258-04c2e8eff935 // indirect
	gopkg.in/yaml.v3 v3.0.0-20190905181640-827449938966 // indirect
	k8s.io/api v0.0.0-20191206001707-7edad22604e1 // indirect
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.0.0-20191114101535-6c5935290e33
	k8s.io/code-generator v0.0.0-20191121015212-c4c8f8345c7e // indirect
	k8s.io/klog v1.0.0
	k8s.io/utils v0.0.0-20191114184206-e782cd3c129f // indirect
	sigs.k8s.io/controller-runtime v0.4.0
	sigs.k8s.io/controller-tools v0.2.1 // indirect
	sigs.k8s.io/structured-merge-diff v1.0.1-0.20191108220359-b1b620dd3f06 // indirect
)

replace github.com/awslabs/goformation/v3 => github.com/christopherhein/goformation/v3 v3.1.1-0.20191116080820-55bd9397137c

replace github.com/appscode/jsonpatch => gomodules.xyz/jsonpatch v1.0.1
