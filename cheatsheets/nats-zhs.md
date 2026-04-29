---
title: NATS
icon: fa-tower-broadcast
primary: "#27AAE1"
lang: bash
locale: zhs
---

## fa-terminal CLI 基础

```bash
nats --help                               # 通用帮助
nats server --help                        # 服务端子命令
nats pub --help                           # 发布帮助
nats sub --help                           # 订阅帮助
nats req --help                           # 请求帮助
nats account info                        # 账户信息
nats context save local --server nats://localhost:4222  # 保存上下文
nats context select local                 # 切换上下文
```

## fa-server 服务器配置

```bash
nats-server                              # 启动服务器（默认 :4222）
nats-server -p 5222                      # 自定义端口
nats-server -a 0.0.0.0                   # 绑定地址
nats-server -c nats.conf                 # 指定配置文件
nats-server -js                          # 启用 JetStream
nats-server -sd /data/nats               # 存储目录
nats-server --cluster nats://0.0.0.0:6222 --routes nats://node2:6222  # 集群模式
```

```bash
docker run -d --name nats -p 4222:4222 -p 8222:8222 nats:latest -js
```

## fa-paper-plane 发布 / 订阅

```bash
nats pub subject.hello "Hello World"     # 发布消息
nats pub subject.data '{"key":"value"}'  # 发布 JSON
nats sub subject.hello                   # 订阅主题
nats sub "subject.>"                     # 通配符订阅
nats sub subject.hello --queue workers   # 队列订阅
```

## fa-layer-group 队列组

```bash
nats sub orders.new --queue order-workers   # 队列组订阅者
nats sub orders.new --queue order-workers   # 另一个成员（负载均衡）
nats pub orders.new "order-123"             # 消息仅投递给一个成员
```

发布到同一主题的消息在同一个队列组内只投递给一个订阅者，实现负载均衡。

## fa-arrows-left-right 请求 / 回复

```bash
nats reply service.time "Received at $(date)"   # 注册回复者
nats req service.time ""                          # 发送请求并等待回复
nats req service.echo "ping" --timeout 5s         # 带超时的请求
```

## fa-water JetStream 流

```bash
nats stream add ORDERS --subjects "orders.>" --storage file --retention limits --max-msgs 100000 --max-age 72h
nats stream info ORDERS
nats stream list
nats stream purge ORDERS
nats stream delete ORDERS
nats stream edit ORDERS --max-msgs 200000
```

```bash
nats stream add EVENTS --subjects "events.>" --storage memory --retention interest --max-bytes 1GB --replicas 3
```

## fa-users JetStream 消费者

```bash
nats consumer add ORDERS pull-cons --pull --deliver all --replay instant
nats consumer add ORDERS push-cons --target push.ords --deliver last --ack explicit
nats consumer info ORDERS pull-cons
nats consumer list ORDERS
nats consumer next ORDERS pull-cons          # 拉取下一条消息
nats consumer delete ORDERS pull-cons
```

## fa-box-archive JetStream KV 存储

```bash
nats kv add config                          # 创建 KV 存储桶
nats kv put config db.host "db.internal"    # 写入键值
nats kv get config db.host                  # 读取值
nats kv del config db.host                  # 删除键
nats kv history config db.host              # 查看修订历史
nats kv status config                       # 存储桶信息
nats kv list                                # 列出所有存储桶
```

## fa-asterisk 通配符

```bash
nats sub "orders.*"                        # 单层通配符
nats sub "orders.>"                        # 多层通配符（匹配所有子级）
nats sub ">"                               # 订阅所有消息
nats sub "orders.*.eu"                     # 匹配 orders.new.eu, orders.update.eu
nats sub "orders.>"                        # 匹配 orders.new, orders.new.eu, ...
```

通配符规则：
- `*` 匹配恰好一个层级
- `>` 匹配一个或多个层级（必须是最后一个 token）

## fa-thumbtack 持久订阅 / 持久消费者

```bash
nats consumer add ORDERS durable-1 --pull --deliver all --ack explicit --durable
nats consumer next ORDERS durable-1 --wait 5s
```

```bash
nats stream add STATUS --subjects "status.>" --storage file --max-msgs-per-subject 1
```

持久消费者在客户端断开后仍然存在，重新连接后从上次位置继续消费。

## fa-network-wired 集群

```bash
nats-server --cluster nats://0.0.0.0:6222 --cluster_name my-cluster --routes nats://node1:6222,nats://node2:6222
```

```conf
# nats.conf（集群配置块）
cluster {
  name: "my-cluster"
  listen: "0.0.0.0:6222"
  routes: [
    nats://node1:6222
    nats://node2:6222
  ]
}
```

## fa-chart-line 监控

```bash
nats-server -m 8222                       # 启用监控端口
curl http://localhost:8222/varz            # 服务器变量
curl http://localhost:8222/connz           # 连接信息
curl http://localhost:8222/routez          # 路由信息
curl http://localhost:8222/subsz           # 订阅信息
curl http://localhost:8222/jsz             # JetStream 信息
```

## fa-lock TLS 与认证

```bash
nats-server --tls --tlscert cert.pem --tlskey key.pem --tlscacert ca.pem
nats-server --user admin --pass secret
```

```conf
# nats.conf
tls {
  cert_file: "/etc/nats/cert.pem"
  key_file: "/etc/nats/key.pem"
  ca_file: "/etc/nats/ca.pem"
}
authorization {
  users: [
    { user: "admin", password: "$2a$11$..." }
  ]
  token: "my-token"
}
```

## fa-code Go 客户端示例

```go
nc, _ := nats.Connect(nats.DefaultURL)
defer nc.Close()

nc.Publish("subject", []byte("hello"))
```

```go
sub, _ := nc.Subscribe("subject", func(m *nats.Msg) {
    fmt.Printf("收到消息: %s\n", string(m.Data))
})
defer sub.Unsubscribe()
```

```go
msg, _ := nc.Request("service.time", nil, 2*time.Second)
fmt.Printf("回复: %s\n", string(msg.Data))
```

```go
js, _ := jetstream.New(nc)
s, _ := js.CreateStream(context.Background(), jetstream.StreamConfig{
    Name:     "ORDERS",
    Subjects: []string{"orders.>"},
})
```

```go
kv, _ := js.CreateKeyValue(context.Background(), jetstream.KeyValueConfig{
    Bucket: "config",
})
kv.Put(context.Background(), "key", []byte("value"))
entry, _ := kv.Get(context.Background(), "key")
```
