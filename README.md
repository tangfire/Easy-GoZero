可以在WSL2中通过Docker快速安装etcd并实现Go程序连接，以下是具体步骤（综合多个可靠来源并验证）：

---

### **一、WSL2中通过Docker安装etcd**

#### 1. 启动etcd容器
```bash
# 单节点etcd容器（适配WSL2网络）
docker run -d \
  --name my-etcd \
  -p 2379:2379 \
  -p 2380:2380 \
  quay.io/coreos/etcd:v3.5.0 \
  /usr/local/bin/etcd \
  --name etcd-node \
  --data-dir /etcd-data \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://localhost:2379
```
*解释：通过`0.0.0.0`允许跨网络访问，`localhost:2379`适配WSL2与Windows的共享网络。*

#### 2. 验证etcd服务
```bash
# 进入容器内部操作
docker exec -it my-etcd etcdctl put test-docker "success"
docker exec -it my-etcd etcdctl get test-docker
```

---

### **二、Go程序连接etcd（Docker版）**
#### 1. 编写Go代码
```go
package main

import (
	"context"
	"fmt"
	"time"
	"go.etcd.io/etcd/client/v3"
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
```


*输出示例：`Key: docker-etcd-key -> Value: hello-docker`*

---

