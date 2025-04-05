package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	// 连接配置（Windows和WSL2共享localhost）
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// 写入数据
	ctx := context.Background()
	if _, err := cli.Put(ctx, "docker-etcd-key", "hello-docker"); err != nil {
		panic(err)
	}

	// 读取数据
	resp, err := cli.Get(ctx, "docker-etcd-key")
	if err != nil {
		panic(err)
	}
	for _, kv := range resp.Kvs {
		fmt.Printf("Key: %s -> Value: %s\n", kv.Key, kv.Value)
	}
}
