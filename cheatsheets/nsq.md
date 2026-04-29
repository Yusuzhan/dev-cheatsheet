---
title: NSQ
icon: fa-tower-broadcast
primary: "#4B0082"
lang: bash
---

## fa-sitemap Architecture Overview

```bash
# NSQ has 3 components: nsqd, nsqlookupd, nsqadmin
# nsqd       - daemon that receives, buffers, delivers messages
# nsqlookupd - service discovery daemon
# nsqadmin   - web UI for monitoring

# typical topology:
# Producer → nsqd → Consumer (via nsqlookupd discovery)
# Multiple nsqd instances register with nsqlookupd
# Consumers query nsqlookupd to find nsqd instances for a topic
```

## fa-server nsqd (Daemon)

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

## fa-magnifying-glass nsqlookupd (Discovery)

```bash
nsqlookupd --tcp-address=0.0.0.0:4160 --http-address=0.0.0.0:4161

curl -s http://localhost:4161/nodes
curl -s http://localhost:4161/topics
curl -s http://localhost:4161/channels?topic=orders
curl -s "http://localhost:4161/lookup?topic=orders"
```

## fa-display nsqadmin (Web UI)

```bash
nsqadmin --lookupd-http-address=localhost:4161

nsqadmin --lookupd-http-address=localhost:4161 \
  --nsqd-http-address=localhost:4151 \
  --http-address=0.0.0.0:4171

curl -s http://localhost:4171/ping
```

## fa-upload Publishing Messages

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

## fa-download Consuming Messages

```bash
curl -s 'http://localhost:4151/channel/create?topic=orders&channel=worker'

nsq_to_nsq --topic=orders --channel=worker \
  --lookupd-http-address=localhost:4161 \
  --destination-nsqd-tcp-address=localhost:4150 \
  --destination-topic=processed_orders

# HTTP endpoint for single message
curl -s 'http://localhost:4151/msg?topic=orders&channel=worker'
```

## fa-layer-group Topics & Channels

```bash
# topics hold messages, each topic can have multiple channels
# each channel receives a copy of every message (fan-out)

curl -s 'http://localhost:4151/stats'                    # all topics/channels
curl -s 'http://localhost:4151/topic/create?topic=events'
curl -s 'http://localhost:4151/topic/delete?topic=events'

curl -s 'http://localhost:4151/channel/create?topic=events&channel=archive'
curl -s 'http://localhost:4151/channel/delete?topic=events&channel=archive'

curl -s 'http://localhost:4151/channel/pause?topic=events&channel=archive'
curl -s 'http://localhost:4151/channel/unpause?topic=events&channel=archive'

curl -s 'http://localhost:4151/topic/pause?topic=events'
curl -s 'http://localhost:4151/topic/unpause?topic=events'
```

## fa-shield-halved Message Guarantees (at-least-once)

```bash
# NSQ guarantees at-least-once delivery
# messages are requeued if consumer does not FIN within timeout
# default timeout: 60000ms (60s)

# consumer protocol flow:
#   RDY 100       # set ready count (prefetch)
#   FIN <id>      # acknowledge message
#   REQ <id> <timeout_ms>  # requeue with delay
#   TOUCH <id>    # extend timeout for this message

# nsqd flags for durability
nsqd --mem-queue-size=0            # all messages go straight to disk
nsqd --sync-every=1                # fsync every message
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

## fa-clock-rotate-left Deferring & Requeuing

```bash
# REQ a message with a delay (requeue back to queue)
# via the TCP protocol:
#   REQ <message_id> <timeout_ms>

# defer processing by requeuing with delay
# timeout_ms=0 sends to default requeue queue

# nsqd requeue configuration
nsqd --max-requeue-delay=15m
nsqd --default-requeue-delay=90s
nsqd --max-attempts=5            # after 5 attempts → DP (discounted pipe / dead letter)

# check depth and requeue count
curl -s 'http://localhost:4151/stats?format=json'
```

## fa-lock TLS & Auth

```bash
nsqd --tls-cert=/etc/nsq/cert.pem \
  --tls-key=/etc/nsq/key.pem \
  --tls-root-ca-file=/etc/nsq/ca.pem \
  --tls-required=tls-verify \
  --tls-client-auth-policy=require-verify

nsqd --auth-http-address=http://auth-service:8080/auth

# nsqauthd example response:
# {
#   "ttl": 3600,
#   "authorizations": [
#     {"topic": "orders", "channels": ["*"], "permissions": ["subscribe","publish"]}
#   ]
# }
```

## fa-chart-bar Statistics & Monitoring

```bash
curl -s http://localhost:4151/stats
curl -s http://localhost:4151/stats?format=json
curl -s http://localhost:4151/stats?topic=orders
curl -s http://localhost:4151/ping

curl -s http://localhost:4161/nodes
curl -s http://localhost:4161/topics

# nsq_stat — live stats in terminal
nsq_stat --lookupd-http-address=localhost:4161

# nsq_stat refreshes every 2s by default
nsq_stat --lookupd-http-address=localhost:4161 --interval=5s
```

## fa-docker Docker Setup

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

## fa-code Go Client (nsq.Producer / nsq.Consumer)

```go
// producer
config := nsq.NewConfig()
p, _ := nsq.NewProducer("localhost:4150", config)
err := p.Publish("orders", []byte(`{"id":"123"}`))
p.Stop()

// consumer
config := nsq.NewConfig()
c, _ := nsq.NewConsumer("orders", "worker", config)
c.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
    process(m.Body)
    return nil
}))
c.ConnectToNSQD("localhost:4150")
// or: c.ConnectToNSQLookupd("localhost:4161")
```
