module github.com/hd-Li/types

go 1.13

replace (
	github.com/knative/pkg => github.com/rancher/pkg v0.0.0-20190514055449-b30ab9de040e
	github.com/matryer/moq => github.com/rancher/moq v0.0.0-20190404221404-ee5226d43009
	k8s.io/api => k8s.io/api v0.0.0-20181004124137-fd83cbc87e76
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20181004124836-1748dfb29e8a
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20180913025736-6dd46049f395
	k8s.io/client-go => k8s.io/client-go v9.0.0+incompatible
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/gorilla/websocket v0.0.0-20150714140627-6eb6ad425a89 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/knative/pkg v0.0.0-00010101000000-000000000000
	github.com/maruel/panicparse v0.0.0-20171209025017-c0182c169410 // indirect
	github.com/maruel/ut v1.0.0 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/common v0.9.1 // indirect
	github.com/prometheus/procfs v0.0.10 // indirect
	github.com/rancher/norman v0.0.0-20190319175355-e10534b012b0
	github.com/rancher/wrangler v0.1.5 // indirect
	github.com/sirupsen/logrus v1.4.2
	google.golang.org/appengine v1.6.1 // indirect
	k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apiextensions-apiserver v0.0.0-20190409022649-727a075fdec8
	k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/gengo v0.0.0-20190822140433-26a664648505
)
