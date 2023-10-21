package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"os"

	clientv3 "go.etcd.io/etcd/client/v3"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/protobuf"
)

type etcdconfig struct {
	etcdCA   string
	etcdKey  string
	etcdCert string
	etcdHost string
	etcdPort int
}

func main() {
	etcdCA := flag.String("etcdCA", "etcd/ca.crt", "Default etcd cert CA")
	etcdKey := flag.String("etcdKey", "etcd/server.key", "Default etcd cert CA Key")
	etcdCert := flag.String("etcdCert", "etcd/server.crt", "Default etcd cert")
	etcdHost := flag.String("etcdhost", "localhost", "path on which the the etcd is running")
	etcdPort := flag.Int("port", 2379, "ETCD port")
	c := etcdconfig{
		etcdCA:   *etcdCA,
		etcdKey:  *etcdKey,
		etcdCert: *etcdCert,
		etcdHost: *etcdHost,
		etcdPort: *etcdPort,
	}
	client, err := etcdClient(c)
	if err != nil {
		fmt.Println(err)
	}
	repsonse, err := client.Get(context.Background(), "/registry/pods/kube-system", clientv3.WithPrefix())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	gvk := schema.GroupVersionKind{
		Group:   v1.GroupName,
		Version: "v1",
		Kind:    "Pod",
	}

	runtimeSchema := runtime.NewScheme()
	runtimeSchema.AddKnownTypeWithName(gvk, &v1.Pod{})
	protoSerializer := protobuf.NewSerializer(runtimeSchema, runtimeSchema)
	for _, v := range repsonse.Kvs {
		//fmt.Println(string(v.Value))
		pods := &v1.Pod{}
		_, _, err := protoSerializer.Decode(v.Value, &gvk, pods)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(pods)
	}
}

func etcdClient(c etcdconfig) (*clientv3.Client, error) {

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

	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{fmt.Sprintf("%s:%d", c.etcdHost, c.etcdPort)},
		TLS: &tls.Config{
			RootCAs: certpool,
			Certificates: []tls.Certificate{
				keypair,
			},
			InsecureSkipVerify: true,
		},
	})
	return client, nil
}
