package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	config clientv3.Config
	client *clientv3.Client
	err    error
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

	fmt.Println("连接成功")
}
