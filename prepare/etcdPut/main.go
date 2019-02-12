package main

import (
"context"
"fmt"
"go.etcd.io/etcd/clientv3"
"time"
)

var (
	config  clientv3.Config
	client  *clientv3.Client
	err     error
	kv      clientv3.KV
	putResp *clientv3.PutResponse
)

func main() {
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"}, //集群列表
		DialTimeout: 5 * time.Second,
	}

	//建立一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	//用于读写etcd的键值对
	kv = clientv3.NewKV(client)

	//kv.Put 带clientv3.WithPrevKV() 获取前一个Value
	if putResp, err = kv.Put(context.TODO(), "/cron/jobs/job1", "bye", clientv3.WithPrevKV()); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Revision:", putResp.Header.Revision)
		if putResp.PrevKv != nil {
			fmt.Println("PrevValue:", string(putResp.PrevKv.Value))
		}
	}
}
