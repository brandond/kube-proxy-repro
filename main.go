package main

import (
	"k8s.io/kubernetes/cmd/kube-proxy/app"
)

func main() {
	_ = app.NewProxyCommand()
}
