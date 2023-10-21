package main

import (
	"flag"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	kubeconfig := flag.String("kubeconfig", "/Users/joshua/.kube/config", "Location to the kubeconfig file")

	// create the client

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {

	}
	clusterclient, err := dynamic.NewForConfig(config)

	//

}
