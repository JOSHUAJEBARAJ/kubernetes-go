## Informers in k8s 

https://stackoverflow.com/questions/59544139/kubernetes-client-go-watch-interface-vs-cache-newinformer-vs-cache-newsharedi

https://aly.arriqaaq.com/kubernetes-informers/

## Custom controllers 

- Listen to the particular resource type and perform some action


### Why not use watch?

- watch queries the API server for many times and it is not efficient


## Informers

- Uses the watch internally and performs the operation efficiently 
- Informer creates a cache of the resource type and watches the cache for the changes

## SharedInformerFactory

- Let's say you want to watch the pods and deployments in this case you can create two informers for each of them but this is not efficient 
- Instead you can create a shared informer factory and create the informers for each of the resource type and use the same cache for both of them
- If we don't want to get all the resources we can use NewFilteredSharedInformerFactory


## resync period

- How often the cache should be refreshed

-  resync period doesn’t mean how often the apiserver is going to be queried for latest resources.There would be a watch call against the apiserver irrespective of the resync period.Resync period helps handling the cases where some of the list/watch calls to aposerver failed abd we dont have those resources. After resync period we would be able to get those resources as well.

```golang
informerfactory := informers.NewSharedInformerFactory(clientset, 30*time.Second)
```

- Don't update the informer cache data , if we want to make the changes you should use the deep copy 

## Queue

- 