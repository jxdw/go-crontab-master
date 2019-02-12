package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	config         clientv3.Config
	client         *clientv3.Client
	err            error
	lease          clientv3.Lease
	leaseGrantResp *clientv3.LeaseGrantResponse
	leaseId        clientv3.LeaseID
	putResp        *clientv3.PutResponse
	getResp        *clientv3.GetResponse
	keepResp       *clientv3.LeaseKeepAliveResponse
	keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
	kv             clientv3.KV
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

	//申请一个lease租约
	lease = clientv3.NewLease(client)

	//申请一个10秒租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}

	//拿到租约的ID
	leaseId = leaseGrantResp.ID

	//5秒后会取消自动续租
	if keepRespChan, err = lease.KeepAlive(context.TODO(), leaseId); err != nil {
		fmt.Println(err)
		return
	}

	//处理续约应答的协程
	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepRespChan == nil {
					fmt.Println("租约已经失效了")
					goto END
				} else {
					//每秒会续租一次，所有就会收到一次应到
					fmt.Println("收到自动续租应答:", keepResp.ID)
				}

			}
		}
	END:
	}()

	//用于etcd的键值对
	kv = clientv3.NewKV(client)

	//Put一个KV,让它与租约关联起来，从而实现10秒后自动过期
	if putResp, err = kv.Put(context.TODO(), "/cron/lock/job1", "xx", clientv3.WithLease(leaseId)); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("写入成功:", putResp.Header.Revision)

	//定时查看一下key过期没有
	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv过期了")
			break
		}
		fmt.Println("还没过期", getResp.Kvs)
		time.Sleep(2 * time.Second)
	}
}
