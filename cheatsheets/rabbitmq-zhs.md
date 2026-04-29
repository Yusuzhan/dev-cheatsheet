---
title: RabbitMQ
icon: fa-envelope
primary: "#FF6600"
lang: bash
locale: zhs
---

## fa-terminal CLI 基础 (rabbitmqctl)

```bash
rabbitmqctl status                        # 代理状态
rabbitmqctl list_vhosts                   # 列出虚拟主机
rabbitmqctl add_vhost myapp               # 创建虚拟主机
rabbitmqctl delete_vhost myapp            # 删除虚拟主机
rabbitmqctl list_users                    # 列出用户
rabbitmqctl add_user admin secret         # 创建用户
rabbitmqctl set_permissions -p myapp admin ".*" ".*" ".*"  # 设置权限
rabbitmqctl set_user_tags admin administrator              # 设置用户角色
rabbitmqctl delete_user guest             # 删除用户
```

```bash
rabbitmqctl list_queues name messages consumers          # 列出队列
rabbitmqctl list_exchanges name type                     # 列出交换机
rabbitmqctl list_bindings                                # 列出绑定
rabbitmqctl close_all_connections --reason "restart"     # 关闭所有连接
rabbitmqctl reset                        # 重置节点（删除所有数据）
```

## fa-arrows-turn-right 交换机

```bash
rabbitmqadmin declare exchange name=orders type=direct durable=true
rabbitmqadmin declare exchange name=events type=topic durable=true
rabbitmqadmin declare exchange name=logs type=fanout durable=true
rabbitmqadmin declare exchange name=rpc type=headers durable=false
rabbitmqadmin delete exchange name=orders
```

交换机类型：
- **direct** — 精确匹配路由键
- **topic** — 模式匹配（`*` 一个词，`#` 零或多个词）
- **fanout** — 广播到所有绑定队列
- **headers** — 根据消息头路由

## fa-inbox 队列

```bash
rabbitmqadmin declare queue name=order.new durable=true
rabbitmqadmin declare queue name=order.process durable=true arguments='{"x-max-length":10000,"x-message-ttl":3600000}'
rabbitmqadmin declare queue name=temp.queue durable=false auto_delete=true
rabbitmqadmin delete queue name=order.new
rabbitmqadmin purge queue name=order.new
```

## fa-link 绑定

```bash
rabbitmqadmin declare binding source=orders destination=order.new routing_key=order.new
rabbitmqadmin declare binding source=events destination=order.process routing_key="order.#"
rabbitmqadmin declare binding source=logs destination=order.new routing_key=""
rabbitmqadmin declare binding source=rpc destination=order.new arguments='{"x-match":"all","format":"pdf"}'
```

## fa-paper-plane 发布（生产者）

```bash
rabbitmqadmin publish exchange=orders routing_key=order.new payload='{"id":1,"item":"book"}'
rabbitmqadmin publish exchange=events routing_key="order.created.eu" payload='{"event":"created"}'
rabbitmqadmin publish exchange=logs routing_key="" payload="log message"
```

```bash
rabbitmq-diagnostics publish_exchange_messages orders 10  # 发布测试消息
```

## fa-inbox-in 消费（消费者）

```bash
rabbitmqadmin get queue=order.new ackmode=ack_requeue_false  # 获取一条消息
rabbitmqadmin get queue=order.new count=5 ackmode=ack_requeue_true   # 预览 5 条消息
```

```bash
rabbitmq-consumer --queue order.new --url amqp://guest:guest@localhost:5672/  # CLI 消费者
```

生产环境建议使用客户端库（Python、Go、Java 等）并正确处理确认机制。

## fa-circle-check 确认机制

```bash
rabbitmqctl list_queues name messages_ready messages_unacknowledged
```

确认模式：
- **自动确认** — 发送后代理自动确认（即发即忘）
- **手动确认** — 消费者处理完毕后显式确认
- **nack + 重新入队** — 消费者拒绝，消息返回队列
- **nack + 丢弃** — 消费者拒绝，消息丢弃或进入死信

```bash
rabbitmqadmin get queue=order.new ackmode=ack_requeue_false   # 手动确认
rabbitmqadmin get queue=order.new ackmode=ack_requeue_true    # 拒绝 + 重新入队
```

## fa-route 路由模式

```
direct  — orders.new      → 绑定 routing_key=orders.new 的队列
topic   — orders.*.eu     → 匹配 orders.new.eu, orders.update.eu
topic   — orders.#        → 匹配 orders.new, orders.new.eu, orders.update.eu.urgent
fanout  — （忽略路由键） → 所有绑定队列都收到消息
headers — x-match: all    → 匹配所有头部键值对
headers — x-match: any    → 匹配任意头部键值对
```

## fa-skull-crossbones 死信队列

```bash
rabbitmqadmin declare exchange name=dlx.exchange type=direct durable=true
rabbitmqadmin declare queue name=dlq.orders durable=true
rabbitmqadmin declare binding source=dlx.exchange destination=dlq.orders routing_key=dlq.orders
rabbitmqadmin declare queue name=order.new durable=true arguments='{"x-dead-letter-exchange":"dlx.exchange","x-dead-letter-routing-key":"dlq.orders"}'
```

消息进入死信的条件：
- 被 `nack` / `reject` 且不重新入队
- TTL 过期
- 队列超过最大长度

## fa-clock TTL 与惰性队列

```bash
rabbitmqadmin declare queue name=short.lived durable=true arguments='{"x-message-ttl":60000}'
rabbitmqadmin declare queue name=expire.queue durable=true arguments='{"x-expires":3600000}'
rabbitmqadmin declare queue name=lazy.queue durable=true arguments='{"x-queue-mode":"lazy"}'
```

```bash
rabbitmqadmin publish exchange=orders routing_key=order.new payload="ttl msg" properties='{"expiration":"5000"}'
```

- `x-message-ttl` — 队列消息 TTL（毫秒）
- `x-expires` — 队列空闲后自动删除（毫秒）
- `x-queue-mode: lazy` — 消息写入磁盘，减少内存占用

## fa-network-wired 集群

```bash
rabbitmqctl join_cluster rabbit@node1            # 加入集群
rabbitmqctl cluster_status                       # 查看集群状态
rabbitmqctl forget_cluster_node rabbit@node3     # 移除节点
rabbitmqctl stop_app && rabbitmqctl join_cluster rabbit@node1 && rabbitmqctl start_app  # 完整加入流程
```

```bash
rabbitmqctl set_policy ha-all ".*" '{"ha-mode":"all","ha-sync-mode":"automatic"}' --apply-to queues
rabbitmqctl set_policy ha-two "^orders\." '{"ha-mode":"exactly","ha-params":2}' --apply-to queues
```

## fa-globe 管理 API

```bash
curl -u guest:guest http://localhost:15672/api/overview
curl -u guest:guest http://localhost:15672/api/queues
curl -u guest:guest http://localhost:15672/api/exchanges
curl -u guest:guest http://localhost:15672/api/connections
curl -u guest:guest http://localhost:15672/api/nodes
curl -u guest:guest -X POST http://localhost:15672/api/users/guest/password -d '{"password":"newpass"}'
```

```bash
rabbitmq-plugins enable rabbitmq_management     # 启用管理界面（端口 :15672）
```

## fa-gear 配置

```conf
# rabbitmq.conf
listeners.tcp.default = 5672
management.tcp.port = 15672
default_user = guest
default_pass = guest
default_vhost = /
log.file.level = info
disk_free_limit.absolute = 2GB
vm_memory_high_watermark.relative = 0.6
```

```conf
# advanced.config — 消费者超时
[
  {rabbit, [
    {consumer_timeout, 600000}
  ]}
].
```

## fa-gauge 性能调优

```bash
rabbitmqctl eval 'application:set_env(rabbit, tcp_listen_options, [{backlog, 4096}]).'
rabbitmqctl eval 'os:setenv("RABBITMQ_IO_THREAD_POOL_SIZE", "128").'
```

```conf
# rabbitmq.conf
queue_master_locator = min-masters
quorum_commands_soft_limit = 256
channel_max = 2047
heartbeat = 60
```

关键调优项：
- 提高文件描述符限制（`ulimit -n 65536`）
- 使用仲裁队列替代经典镜像队列以获得更好的持久性
- 设置 `vm_memory_high_watermark` 在 OOM 前触发流控
- 批量发布提升吞吐量
- 大量积压时使用 `lazy` 队列模式
