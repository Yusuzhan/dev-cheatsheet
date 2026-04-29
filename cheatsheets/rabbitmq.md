---
title: RabbitMQ
icon: fa-envelope
primary: "#FF6600"
lang: bash
---

## fa-terminal CLI Basics (rabbitmqctl)

```bash
rabbitmqctl status                        # broker status
rabbitmqctl list_vhosts                   # list virtual hosts
rabbitmqctl add_vhost myapp               # create vhost
rabbitmqctl delete_vhost myapp            # delete vhost
rabbitmqctl list_users                    # list users
rabbitmqctl add_user admin secret         # create user
rabbitmqctl set_permissions -p myapp admin ".*" ".*" ".*"  # set permissions
rabbitmqctl set_user_tags admin administrator              # set user role
rabbitmqctl delete_user guest             # delete user
```

```bash
rabbitmqctl list_queues name messages consumers          # list queues
rabbitmqctl list_exchanges name type                     # list exchanges
rabbitmqctl list_bindings                                # list bindings
rabbitmqctl close_all_connections --reason "restart"     # close connections
rabbitmqctl reset                        # reset node (delete all data)
```

## fa-arrows-turn-right Exchanges

```bash
rabbitmqadmin declare exchange name=orders type=direct durable=true
rabbitmqadmin declare exchange name=events type=topic durable=true
rabbitmqadmin declare exchange name=logs type=fanout durable=true
rabbitmqadmin declare exchange name=rpc type=headers durable=false
rabbitmqadmin delete exchange name=orders
```

Exchange types:
- **direct** — route by exact match
- **topic** — route by pattern (`*` one word, `#` zero or more)
- **fanout** — broadcast to all bound queues
- **headers** — route by message headers

## fa-inbox Queues

```bash
rabbitmqadmin declare queue name=order.new durable=true
rabbitmqadmin declare queue name=order.process durable=true arguments='{"x-max-length":10000,"x-message-ttl":3600000}'
rabbitmqadmin declare queue name=temp.queue durable=false auto_delete=true
rabbitmqadmin delete queue name=order.new
rabbitmqadmin purge queue name=order.new
```

## fa-link Bindings

```bash
rabbitmqadmin declare binding source=orders destination=order.new routing_key=order.new
rabbitmqadmin declare binding source=events destination=order.process routing_key="order.#"
rabbitmqadmin declare binding source=logs destination=order.new routing_key=""
rabbitmqadmin declare binding source=rpc destination=order.new arguments='{"x-match":"all","format":"pdf"}'
```

## fa-paper-plane Publish (Producer)

```bash
rabbitmqadmin publish exchange=orders routing_key=order.new payload='{"id":1,"item":"book"}'
rabbitmqadmin publish exchange=events routing_key="order.created.eu" payload='{"event":"created"}'
rabbitmqadmin publish exchange=logs routing_key="" payload="log message"
```

```bash
rabbitmq-diagnostics publish_exchange_messages orders 10  # publish test messages
```

## fa-inbox-in Consume (Consumer)

```bash
rabbitmqadmin get queue=order.new ackmode=ack_requeue_false  # get one message
rabbitmqadmin get queue=order.new count=5 ackmode=ack_requeue_true   # peek 5 messages
```

```bash
rabbitmq-consumer --queue order.new --url amqp://guest:guest@localhost:5672/  # CLI consumer
```

For production consumers, use a client library (Python, Go, Java, etc.) with proper ack handling.

## fa-circle-check Acknowledgements

```bash
rabbitmqctl list_queues name messages_ready messages_unacknowledged
```

Ack modes:
- **automatic** — broker acks on send (fire and forget)
- **manual** — consumer acks explicitly after processing
- **nack + requeue** — consumer rejects and message returns to queue
- **nack + discard** — consumer rejects and message drops or goes to DLX

```bash
rabbitmqadmin get queue=order.new ackmode=ack_requeue_false   # manual ack
rabbitmqadmin get queue=order.new ackmode=ack_requeue_true    # reject + requeue
```

## fa-route Routing Patterns

```
direct  — orders.new      → queue bound with routing_key=orders.new
topic   — orders.*.eu     → orders.new.eu, orders.update.eu
topic   — orders.#        → orders.new, orders.new.eu, orders.update.eu.urgent
fanout  — (ignore key)    → all bound queues receive message
headers — x-match: all    → match all header key-value pairs
headers — x-match: any    → match any header key-value pair
```

## fa-skull-crossbones Dead Letter Queues

```bash
rabbitmqadmin declare exchange name=dlx.exchange type=direct durable=true
rabbitmqadmin declare queue name=dlq.orders durable=true
rabbitmqadmin declare binding source=dlx.exchange destination=dlq.orders routing_key=dlq.orders
rabbitmqadmin declare queue name=order.new durable=true arguments='{"x-dead-letter-exchange":"dlx.exchange","x-dead-letter-routing-key":"dlq.orders"}'
```

Messages are dead-lettered when:
- rejected with `nack` / `reject` without requeue
- TTL expires
- queue exceeds max-length

## fa-clock TTL & Lazy Queues

```bash
rabbitmqadmin declare queue name=short.lived durable=true arguments='{"x-message-ttl":60000}'
rabbitmqadmin declare queue name=expire.queue durable=true arguments='{"x-expires":3600000}'
rabbitmqadmin declare queue name=lazy.queue durable=true arguments='{"x-queue-mode":"lazy"}'
```

```bash
rabbitmqadmin publish exchange=orders routing_key=order.new payload="ttl msg" properties='{"expiration":"5000"}'
```

- `x-message-ttl` — per-queue message TTL (ms)
- `x-expires` — queue auto-delete after idle (ms)
- `x-queue-mode: lazy` — keep messages on disk, reduce RAM usage

## fa-network-wired Clustering

```bash
rabbitmqctl join_cluster rabbit@node1            # join cluster
rabbitmqctl cluster_status                       # view cluster status
rabbitmqctl forget_cluster_node rabbit@node3     # remove node
rabbitmqctl stop_app && rabbitmqctl join_cluster rabbit@node1 && rabbitmqctl start_app  # full join flow
```

```bash
rabbitmqctl set_policy ha-all ".*" '{"ha-mode":"all","ha-sync-mode":"automatic"}' --apply-to queues
rabbitmqctl set_policy ha-two "^orders\." '{"ha-mode":"exactly","ha-params":2}' --apply-to queues
```

## fa-globe Management API

```bash
curl -u guest:guest http://localhost:15672/api/overview
curl -u guest:guest http://localhost:15672/api/queues
curl -u guest:guest http://localhost:15672/api/exchanges
curl -u guest:guest http://localhost:15672/api/connections
curl -u guest:guest http://localhost:15672/api/nodes
curl -u guest:guest -X POST http://localhost:15672/api/users/guest/password -d '{"password":"newpass"}'
```

```bash
rabbitmq-plugins enable rabbitmq_management     # enable management UI on :15672
```

## fa-gear Config

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
# advanced.config — consumer timeout
[
  {rabbit, [
    {consumer_timeout, 600000}
  ]}
].
```

## fa-gauge Performance Tuning

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

Key tuning knobs:
- Increase file descriptor limits (`ulimit -n 65536`)
- Use quorum queues for durability over classic mirrored queues
- Set `vm_memory_high_watermark` to trigger flow control before OOM
- Batch publishes for throughput
- Use `lazy` queue mode for large backlogs
