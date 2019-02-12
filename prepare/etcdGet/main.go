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
	getResp *clientv3.GetResponse
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

	//用于读写etcd的键值对
	kv = clientv3.NewKV(client)
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs", clientv3.WithPrefix() /*clientv3.WithCountOnly()*/); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getResp.Kvs, getResp.Count)

		for k, v := range getResp.Kvs {
			fmt.Println(k, v)
		}
		//[key:"/cron/jobs/job1" create_revision:4 mod_revision:8 version:3 value:"bye" ] 1
		//create_revision:创建版本
		//mod_revision: 修改版本
		//version:修改了几个版本
	}
}
