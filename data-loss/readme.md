## Readme

In this exercise we will be building the client that will interact with the etcd server 

So what I am going to do is to build the client which will update the labels for the pods in the cluster via the etcd server


### Creating the certpool

```go

	ca, err := os.ReadFile(c.etcdCA)
	if err != nil {
		return nil, err
	}
	keypair, err := tls.LoadX509KeyPair(c.etcdCert, c.etcdKey)
	if err != nil {
		return nil, err
	}
	certpool := x509.NewCertPool()
	certpool.AppendCertsFromPEM(ca)
```

## ETCD client 

```bash
ETCDCTL_API=3 etcdctl --cacert=etcd/ca.crt --key=etcd/server.key --cert=etcd/server.crt --endpoints localhost:2379 get "" --prefix --keys-only
```

```bash
kubectl port-forward etcd-kind-control-plane -n kube-system --address 0.0.0.0 2379:2379
```