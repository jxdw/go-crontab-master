package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	config clientv3.Config
	client *clientv3.Client
	err    error
	kv     clientv3.KV
	putOp  clientv3.Op
	getOp  clientv3.Op
	opResp clientv3.OpResponse
)

func main() {
	//客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	//建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	//用于读写etcd的键值对
	kv = clientv3.NewKV(client)

	//创建Op: operation
	putOp = clientv3.OpPut("/cron/jobs/job8", "12345678")

	//执行Op
	if opResp, err = kv.Do(context.TODO(), putOp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("写入Revision:", opResp.Put().Header.Revision)

	//创建Op
	getOp = clientv3.OpGet("/cron/jobs/job8")

	//执行Op
	if opResp, err = kv.Do(context.TODO(), getOp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("数据Revision:", opResp.Get().Kvs[0].ModRevision)
	fmt.Println("数据value:", string(opResp.Get().Kvs[0].Value))
}
