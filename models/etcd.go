package models

import (
	"github.com/coreos/go-etcd/etcd"
	"log"
)

type Etcd struct {
	Key string
	Value string
}

func SetEtcd(k string, v string) error {
	client := etcd.NewClient([]string{"http://127.0.0.1:2379"})
	if _, err := client.Set(k, v, 0); err != nil {
		log.Fatal(err)
    return err
	}
  return nil
}

func GetEtcd(k string) string {
	client := etcd.NewClient([]string{"http://127.0.0.1:2379"})
	resp, err := client.Get(k, false, false)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	log.Printf("Current creds: %s: %s\n", resp.Node.Key, resp.Node.Value)
	return resp.Node.Value
}
