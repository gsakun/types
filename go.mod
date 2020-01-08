module github.com/hd-Li/types

go 1.13

replace (
	github.com/matryer/moq => github.com/rancher/moq v0.0.0-20190404221404-ee5226d43009
	k8s.io/client-go => k8s.io/client-go v9.0.0+incompatible
        github.com/knative/pkg => github.com/rancher/pkg v0.0.0-20190514055449-b30ab9de040e
        k8s.io/api => k8s.io/api v0.0.0-20181004124137-fd83cbc87e76
        k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20181004124836-1748dfb29e8a
        k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20180913025736-6dd46049f395
)

require (
	github.com/rancher/norman v0.0.0-20190319175355-e10534b012b0
	github.com/ghodss/yaml v1.0.0
	github.com/gorilla/websocket v0.0.0-20150714140627-6eb6ad425a89
	github.com/maruel/panicparse v0.0.0-20171209025017-c0182c169410
	github.com/maruel/ut v1.0.0 // indirect
	github.com/matryer/moq v0.0.0-20190312154309-6cfb0558e1bd
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v0.9.1
	github.com/rancher/wrangler v0.1.5
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.3.0
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
	google.golang.org/appengine v1.6.1 // indirect
	k8s.io/gengo v0.0.0-20190822140433-26a664648505
)
