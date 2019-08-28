package main

import (
    "flag"
    "path/filepath"
    "os"
    "log"

    "github.com/codetime66/k8s-svc-endpoint/pkg/ephttp"
)

func main() {
    var kubeconfig *string
    if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
    } else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
    }
    var svc = flag.String("service-name", "", "The service name.")

    flag.Parse()

    log.Printf("----------------------")
    log.Printf("getting svc endpoints:")
    ipsep := ephttp.StartUp(*kubeconfig, *svc)

    log.Printf("----------------------")
    log.Printf("listing svc endpoints:")
    for _, ipep := range ipsep {
       log.Printf(ipep)
    }
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
