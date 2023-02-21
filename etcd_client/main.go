package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// 创建etcd客户端
	ctx := context.TODO()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"www.tanght.xyz:2379"},
		DialTimeout: 5 * time.Second,
		Context:     ctx,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if e := cli.Close(); e != nil {
			fmt.Printf("cli.Close error, %s\n", e.Error())
		}
	}()

	// 租约操作
	lease := clientv3.NewLease(cli)
	// 创建租约
	grantRes, err := lease.Grant(ctx, 10)
	if err != nil {
		fmt.Printf("lease.Grant error, %s\n", err.Error())
		return
	}
	// 租约保活(这个函数会创建一个新的协程, 在单独的协程中keep-alive)
	keepAliveCh, err := lease.KeepAlive(ctx, grantRes.ID)
	if err != nil {
		fmt.Printf("lease.KeepAlive error, %s\n", err.Error())
		return
	}
	go func() {
		for {
			_, ok := <-keepAliveCh
			if !ok {
				break
			}
		}
	}()
	// 查看所有租约
	leaseRes, err := lease.Leases(ctx)
	if err != nil {
		fmt.Printf("lease.Leases error, %s\n", err.Error())
		return
	}
	fmt.Printf("all lease %v", leaseRes.Leases)

	// k-v操作
	kv := clientv3.NewKV(cli)
	// 创建
	kv.Put(ctx, "/tanght1", "haha1", clientv3.WithLease(grantRes.ID))
	kv.Put(ctx, "/tanght2", "xixi1", clientv3.WithLease(grantRes.ID))
	// 修改
	kv.Put(ctx, "/tanght1", "haha2", clientv3.WithLease(grantRes.ID))
	kv.Put(ctx, "/tanght2", "xixi2", clientv3.WithLease(grantRes.ID))
	// 查找
	getRes, err := kv.Get(ctx, "/", clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("kv.Get error, %s\n", err.Error())
		return
	}
	lastRevision := getRes.Header.Revision
	for _, v := range getRes.Kvs {
		fmt.Println(v)
	}
	fmt.Printf("lastRevision is %d\n", lastRevision)

	// watch操作
	fmt.Println("watch-----------------------------------")
	wh := clientv3.NewWatcher(cli)
	whch := wh.Watch(ctx, "/tanght", clientv3.WithPrefix(), clientv3.WithRev(lastRevision+1))
	for {
		whRes := <-whch
		for _, event := range whRes.Events {
			k := string(event.Kv.Key)
			if event.Type == clientv3.EventTypeDelete {
				fmt.Printf("DEL %s\n", k)
			} else if event.IsCreate() {
				fmt.Printf("PUT %s %s\n", k, string(event.Kv.Value))
			} else if event.IsModify() {
				fmt.Printf("MOD %s %s\n", k, string(event.Kv.Value))
			} else {
				fmt.Printf("ERR %s\n", k)
			}
		}
	}
}
