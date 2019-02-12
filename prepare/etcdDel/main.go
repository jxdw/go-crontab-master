package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"time"
)

var (
	config  clientv3.Config
	client  *clientv3.Client
	err     error
	kv      clientv3.KV
	delResp *clientv3.DeleteResponse
	kvpair  *mvccpb.KeyValue
)

func main() {
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	//建立一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	//用于etcd的键值对
	kv = clientv3.NewKV(client)

	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/job1", clientv3.WithFromKey()); err != nil {
		fmt.Println("del", err)
		return
	}

	if len(delResp.PrevKvs) != 0 {
		for _, kvpair = range delResp.PrevKvs {
			fmt.Println("删除了:", string(kvpair.Key), string(kvpair.Value))
		}
	}
}
