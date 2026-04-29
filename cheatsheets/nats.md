---
title: NATS
icon: fa-tower-broadcast
primary: "#27AAE1"
lang: bash
---

## fa-terminal CLI Basics

```bash
nats --help                               # general help
nats server --help                        # server subcommands
nats pub --help                           # publish help
nats sub --help                           # subscribe help
nats req --help                           # request help
nats account info                        # account information
nats context save local --server nats://localhost:4222  # save context
nats context select local                 # switch context
```

## fa-server Server Setup

```bash
nats-server                              # start server (default :4222)
nats-server -p 5222                      # custom port
nats-server -a 0.0.0.0                   # bind address
nats-server -c nats.conf                 # config file
nats-server -js                          # enable JetStream
nats-server -sd /data/nats               # store directory
nats-server --cluster nats://0.0.0.0:6222 --routes nats://node2:6222  # cluster
```

```bash
docker run -d --name nats -p 4222:4222 -p 8222:8222 nats:latest -js
```

## fa-paper-plane Publish / Subscribe

```bash
nats pub subject.hello "Hello World"     # publish message
nats pub subject.data '{"key":"value"}'  # publish JSON
nats sub subject.hello                   # subscribe to subject
nats sub "subject.>"                     # subscribe with wildcard
nats sub subject.hello --queue workers   # queue subscribe
```

## fa-layer-group Queue Groups

```bash
nats sub orders.new --queue order-workers   # subscriber in queue group
nats sub orders.new --queue order-workers   # another member (load balanced)
nats pub orders.new "order-123"             # message delivered to one member
```

Messages published to a subject are delivered to only one subscriber within the same queue group, enabling load balancing.

## fa-arrows-left-right Request / Reply

```bash
nats reply service.time "Received at $(date)"   # register replier
nats req service.time ""                          # send request, wait for reply
nats req service.echo "ping" --timeout 5s         # request with timeout
```

## fa-water JetStream Streams

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

## fa-users JetStream Consumers

```bash
nats consumer add ORDERS pull-cons --pull --deliver all --replay instant
nats consumer add ORDERS push-cons --target push.ords --deliver last --ack explicit
nats consumer info ORDERS pull-cons
nats consumer list ORDERS
nats consumer next ORDERS pull-cons          # fetch next message
nats consumer delete ORDERS pull-cons
```

## fa-box-archive JetStream KV Store

```bash
nats kv add config                          # create KV bucket
nats kv put config db.host "db.internal"    # put key-value
nats kv get config db.host                  # get value
nats kv del config db.host                  # delete key
nats kv history config db.host              # revision history
nats kv status config                       # bucket info
nats kv list                                # list all buckets
```

## fa-asterisk Wildcards

```bash
nats sub "orders.*"                        # single token wildcard
nats sub "orders.>"                        # multi-token wildcard (catch-all)
nats sub ">"                               # subscribe to everything
nats sub "orders.*.eu"                     # match orders.new.eu, orders.update.eu
nats sub "orders.>"                        # match orders.new, orders.new.eu, ...
```

Wildcard tokens:
- `*` matches exactly one token
- `>` matches one or more tokens (must be last)

## fa-thumbtack Retained Messages / Durables

```bash
nats consumer add ORDERS durable-1 --pull --deliver all --ack explicit --durable
nats consumer next ORDERS durable-1 --wait 5s
```

```bash
nats stream add STATUS --subjects "status.>" --storage file --max-msgs-per-subject 1
```

Durable consumers survive client disconnections and resume from where they left off.

## fa-network-wired Clustering

```bash
nats-server --cluster nats://0.0.0.0:6222 --cluster_name my-cluster --routes nats://node1:6222,nats://node2:6222
```

```conf
# nats.conf (cluster block)
cluster {
  name: "my-cluster"
  listen: "0.0.0.0:6222"
  routes: [
    nats://node1:6222
    nats://node2:6222
  ]
}
```

## fa-chart-line Monitoring

```bash
nats-server -m 8222                       # enable monitoring port
curl http://localhost:8222/varz            # server variables
curl http://localhost:8222/connz           # connection info
curl http://localhost:8222/routez          # route info
curl http://localhost:8222/subsz           # subscription info
curl http://localhost:8222/jsz             # JetStream info
```

## fa-lock TLS & Auth

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

## fa-code Go Client Examples

```go
nc, _ := nats.Connect(nats.DefaultURL)
defer nc.Close()

nc.Publish("subject", []byte("hello"))
```

```go
sub, _ := nc.Subscribe("subject", func(m *nats.Msg) {
    fmt.Printf("Received: %s\n", string(m.Data))
})
defer sub.Unsubscribe()
```

```go
msg, _ := nc.Request("service.time", nil, 2*time.Second)
fmt.Printf("Reply: %s\n", string(msg.Data))
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
