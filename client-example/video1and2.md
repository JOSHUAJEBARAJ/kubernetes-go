## Goal 

To write the client application that list the pods in the cluster 

We will be using the following API to list the pods in the cluster
- client-go 
- api 
- api-machinery 


### client-go
Exposes all the interfaces that are required to interact with the kubernetes cluster.

https://github.com/kubernetes/client-go

When you are using the client go please check for the version compatibility with the kubernetes cluster. 

```bash
go get k8s.io/client-go@compatibility-version
```


### Api 

all the resources that are exposed by the kubernetes cluster are defined in the api package.

https://github.com/kubernetes/api


### Api machinery 

It has the lot of the utility methods 


Example https://github.com/kubernetes/client-go/blob/master/examples/create-update-delete-deployment/main.go 


## Goal -2 

In this example we will be using the client-go to list the pods in the cluster as the pods 


Add the below  line 

```go
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Println("error getting the config")
		}
```

and give the service account permission to list the pods 

```bash
