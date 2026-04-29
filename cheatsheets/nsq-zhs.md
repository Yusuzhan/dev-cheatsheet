---
title: NSQ
icon: fa-tower-broadcast
primary: "#4B0082"
lang: bash
locale: zhs
---

## fa-sitemap 架构概览

```bash
# NSQ 有 3 个核心组件：nsqd、nsqlookupd、nsqadmin
# nsqd       - 消息守护进程，负责接收、缓冲和投递消息
# nsqlookupd - 服务发现守护进程
# nsqadmin   - 监控 Web 界面

# 典型拓扑：
# 生产者 → nsqd → 消费者（通过 nsqlookupd 发现）
# 多个 nsqd 实例向 nsqlookupd 注册
# 消费者查询 nsqlookupd 来查找 topic 对应的 nsqd 实例
```

## fa-server nsqd（守护进程）

```bash
nsqd --lookupd-tcp-address=localhost:4160 --broadcast-address=127.0.0.1

nsqd --tcp-address=0.0.0.0:4150 \
  --http-address=0.0.0.0:4151 \
  --lookupd-tcp-address=localhost:4160 \
  --data-path=/var/lib/nsq

nsqd --max-msg-size=1048576 \
  --max-rdy-count=2500 \
  --max-output-buffer-size=65536 \
  --sync-every=100

curl -s http://localhost:4151/stats
curl -s http://localhost:4151/ping
```

## fa-magnifying-glass nsqlookupd（服务发现）

```bash
nsqlookupd --tcp-address=0.0.0.0:4160 --http-address=0.0.0.0:4161

curl -s http://localhost:4161/nodes
curl -s http://localhost:4161/topics
curl -s http://localhost:4161/channels?topic=orders
curl -s "http://localhost:4161/lookup?topic=orders"
```

## fa-display nsqadmin（Web 管理界面）

```bash
nsqadmin --lookupd-http-address=localhost:4161

nsqadmin --lookupd-http-address=localhost:4161 \
  --nsqd-http-address=localhost:4151 \
  --http-address=0.0.0.0:4171

curl -s http://localhost:4171/ping
```

## fa-upload 发布消息

```bash
curl -d 'hello world' 'http://localhost:4151/pub?topic=orders'

curl -X POST 'http://localhost:4151/pub?topic=orders' \
  -H 'Content-Type: application/json' \
  -d '{"order_id":"123","amount":99.9}'

curl -X POST 'http://localhost:4151/mpub?topic=orders' \
  -d $'message1\nmessage2\nmessage3'

curl -s 'http://localhost:4151/topic/create?topic=orders'
curl -s 'http://localhost:4151/topic/delete?topic=orders'
```

## fa-download 消费消息

```bash
curl -s 'http://localhost:4151/channel/create?topic=orders&channel=worker'

nsq_to_nsq --topic=orders --channel=worker \
  --lookupd-http-address=localhost:4161 \
  --destination-nsqd-tcp-address=localhost:4150 \
  --destination-topic=processed_orders

# HTTP 端点获取单条消息
curl -s 'http://localhost:4151/msg?topic=orders&channel=worker'
```

## fa-layer-group 主题与通道 (Topics & Channels)

```bash
# topic 保存消息，每个 topic 可以有多个 channel
# 每个 channel 都会收到所有消息的副本（扇出分发）

curl -s 'http://localhost:4151/stats'                    # 所有 topic/channel
curl -s 'http://localhost:4151/topic/create?topic=events'
curl -s 'http://localhost:4151/topic/delete?topic=events'

curl -s 'http://localhost:4151/channel/create?topic=events&channel=archive'
curl -s 'http://localhost:4151/channel/delete?topic=events&channel=archive'

curl -s 'http://localhost:4151/channel/pause?topic=events&channel=archive'
curl -s 'http://localhost:4151/channel/unpause?topic=events&channel=archive'

curl -s 'http://localhost:4151/topic/pause?topic=events'
curl -s 'http://localhost:4151/topic/unpause?topic=events'
```

## fa-shield-halved 消息保证 (at-least-once)

```bash
# NSQ 保证至少一次投递
# 如果消费者未在超时时间内 FIN，消息会被重新入队
# 默认超时：60000ms（60 秒）

# 消费者协议流程：
#   RDY 100       # 设置就绪数量（预取）
#   FIN <id>      # 确认消息
#   REQ <id> <timeout_ms>  # 延迟重新入队
#   TOUCH <id>    # 延长此消息的超时时间

# nsqd 持久化配置
nsqd --mem-queue-size=0            # 所有消息直接写入磁盘
nsqd --sync-every=1                # 每条消息都 fsync
```

## fa-file-export nsq_to_file / nsq_to_http

```bash
nsq_to_file --topic=orders --channel=logger \
  --lookupd-http-address=localhost:4161 \
  --output-dir=/var/log/nsq \
  --gzip

nsq_to_file --topic=events --channel=archive \
  --nsqd-tcp-address=localhost:4150 \
  --output-dir=/tmp/nsq-archive \
  --filename-format=2006-01-02_15

nsq_to_http --topic=webhooks --channel=sink \
  --lookupd-http-address=localhost:4161 \
  --post=http://api.example.com/webhook \
  --content-type=application/json
```

## fa-clock-rotate-left 延迟与重新入队

```bash
# 通过 REQ 延迟重新入队
# 通过 TCP 协议：
#   REQ <message_id> <timeout_ms>

# timeout_ms=0 会发送到默认重入队队列

# nsqd 重入队配置
nsqd --max-requeue-delay=15m
nsqd --default-requeue-delay=90s
nsqd --max-attempts=5            # 超过 5 次尝试后进入错误队列

# 查看深度和重入队计数
curl -s 'http://localhost:4151/stats?format=json'
```

## fa-lock TLS 与认证

```bash
nsqd --tls-cert=/etc/nsq/cert.pem \
  --tls-key=/etc/nsq/key.pem \
  --tls-root-ca-file=/etc/nsq/ca.pem \
  --tls-required=tls-verify \
  --tls-client-auth-policy=require-verify

nsqd --auth-http-address=http://auth-service:8080/auth

# nsqauthd 示例响应：
# {
#   "ttl": 3600,
#   "authorizations": [
#     {"topic": "orders", "channels": ["*"], "permissions": ["subscribe","publish"]}
#   ]
# }
```

## fa-chart-bar 统计与监控

```bash
curl -s http://localhost:4151/stats
curl -s http://localhost:4151/stats?format=json
curl -s http://localhost:4151/stats?topic=orders
curl -s http://localhost:4151/ping

curl -s http://localhost:4161/nodes
curl -s http://localhost:4161/topics

# nsq_stat — 终端实时统计
nsq_stat --lookupd-http-address=localhost:4161

# nsq_stat 默认每 2 秒刷新
nsq_stat --lookupd-http-address=localhost:4161 --interval=5s
```

## fa-docker Docker 部署

```bash
docker run -d --name nsqlookupd -p 4160:4160 -p 4161:4161 \
  nsqio/nsq /nsqlookupd

docker run -d --name nsqd -p 4150:4150 -p 4151:4151 \
  nsqio/nsq /nsqd --lookupd-tcp-address=host.docker.internal:4160 \
  --broadcast-address=host.docker.internal

docker run -d --name nsqadmin -p 4171:4171 \
  nsqio/nsq /nsqadmin --lookupd-http-address=host.docker.internal:4161

docker compose up -d
```

## fa-code Go 客户端 (nsq.Producer / nsq.Consumer)

```go
// 生产者
config := nsq.NewConfig()
p, _ := nsq.NewProducer("localhost:4150", config)
err := p.Publish("orders", []byte(`{"id":"123"}`))
p.Stop()

// 消费者
config := nsq.NewConfig()
c, _ := nsq.NewConsumer("orders", "worker", config)
c.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
    process(m.Body)
    return nil
}))
c.ConnectToNSQD("localhost:4150")
// 或：c.ConnectToNSQLookupd("localhost:4161")
```
