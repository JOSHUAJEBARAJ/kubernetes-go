package main

import (
	"context"
	"flag"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// get the kubeconfig file

	kubeconfig := flag.String("kubeconfig", "/Users/joshua/.kube/config", "Location to the kubeconfig file")

	// create the client

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {

	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		// handle the error
	}
	pods, err := clientset.CoreV1().Pods("kube-system").List(context.Background(), v1.ListOptions{})
	if err != nil {
		// handle
	}
	for _, pod := range pods.Items {
		fmt.Println(pod)
	}
}
