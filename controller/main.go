package main

import (
	"flag"
	"fmt"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/joshua/.kube/config", "Location to the kubeconfig file")

	// create the client

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {

		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Println("error getting the config")
		}

	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		// handle the error
	}
	// channel

	ch := make(chan struct{})

	informers := informers.NewSharedInformerFactory(clientset, 30*time.Second)
	// new controller

	c := NewController(clientset, informers.Apps().V1().Deployments())
	informers.Start(ch)
	c.run(ch)
}
