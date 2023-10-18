## Api machinery 

https://kubernetes.io/docs/reference/using-api/api-concepts/#:~:text=Kubernetes%20generally%20leverages%20common%20RESTful,which%20is%20called%20a%20kind
## Kind
GroupVersionKind - eg apps/v1/Deployment

## resource


GroupVersionResource - apis/apps/v1/deployments
eg : deployments 

```bash
kubectl get deployments -v 6
```


## Rest mapping 

Allow to convert the kind to the resource


## Scheme 

