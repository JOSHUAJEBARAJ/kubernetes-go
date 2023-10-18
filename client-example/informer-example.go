package main

import (
	"flag"
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
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
	// creating the informer

	informerfactory := informers.NewSharedInformerFactory(clientset, 30*time.Second)
	podinformer := informerfactory.Core().V1().Pods()
	podinformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("Add was calle")
		},
	})

	// cache will be initialzed
	informerfactory.Start(wait.NeverStop)
	// wait for cache to sync
	informerfactory.WaitForCacheSync(wait.NeverStop)
	pod, err := podinformer.Lister().Pods("kube-system").Get("default")
	fmt.Println(pod)
}
