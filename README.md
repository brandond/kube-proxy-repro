Minimal reproduction for issue created by: 
* https://github.com/kubernetes/kubernetes/pull/125866

```
brandond@dev01:~/go/src/github.com/brandond/kube-proxy-repro$ go mod why k8s.io/kubernetes/pkg/proxy/metrics
# k8s.io/kubernetes/pkg/proxy/metrics
github.com/brandond/kube-proxy-repro
k8s.io/kubernetes/cmd/kube-proxy/app
k8s.io/kubernetes/pkg/proxy/metrics
```

`newNFAcctMetricCollector` is used when creating Prometheus metrics during module variable initialization, and calls klog.Errors() if the nfacct client cannot be initialized. This is a regression from previous Kubernetes releases.

```console
brandond@dev01:~/go/src/github.com/brandond/kube-proxy-repro$ go run main.go
E0827 17:30:29.414588 2101747 metrics.go:340] "failed to initialize nfacct client" err="nfacct sub-system not available"
E0827 17:30:29.414652 2101747 metrics.go:340] "failed to initialize nfacct client" err="nfacct sub-system not available"
```
