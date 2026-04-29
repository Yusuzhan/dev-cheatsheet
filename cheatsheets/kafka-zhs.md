---
title: Kafka
icon: fa-arrows-split-up-and-left
primary: "#231F20"
lang: bash
locale: zhs
---

## fa-terminal CLI 基础

```bash
kafka-topics.sh --bootstrap-server localhost:9092 --list
kafka-topics.sh --bootstrap-server localhost:9092 --describe --topic my-topic
kafka-configs.sh --bootstrap-server localhost:9092 --describe --entity-type brokers
kafka-cluster-cluster-id.sh --bootstrap-server localhost:9092
kafka-log-dirs.sh --bootstrap-server localhost:9092 --describe
```

## fa-layer-group 主题 (Topics)

```bash
kafka-topics.sh --bootstrap-server localhost:9092 \
  --create --topic orders --partitions 3 --replication-factor 2

kafka-topics.sh --bootstrap-server localhost:9092 \
  --alter --topic orders --partitions 6

kafka-topics.sh --bootstrap-server localhost:9092 \
  --delete --topic orders

kafka-topics.sh --bootstrap-server localhost:9092 \
  --describe --topic orders --under-replica-partitions
```

## fa-upload 生产者 (Producer)

```bash
kafka-console-producer.sh --bootstrap-server localhost:9092 --topic orders

kafka-console-producer.sh --bootstrap-server localhost:9092 \
  --topic orders --property "parse.key=true" --property "key.separator=:"

echo "order-123:payload" | kafka-console-producer.sh \
  --bootstrap-server localhost:9092 --topic orders --property "parse.key=true" --property "key.separator=:"

kafka-producer-perf-test.sh --topic orders --num-records 100000 \
  --record-size 256 --throughput -1 --producer-props bootstrap.servers=localhost:9092
```

## fa-download 消费者 (Consumer)

```bash
kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic orders --from-beginning

kafka-console-consumer.sh --bootstrap-server localhost:9092 \
  --topic orders --property "print.key=true" --property "key.separator=:" \
  --group my-consumer-group

kafka-console-consumer.sh --bootstrap-server localhost:9092 \
  --topic orders --offset 5 --partition 0 --timeout-ms 10000

kafka-consumer-perf-test.sh --topic orders --messages 10000 \
  --bootstrap-server localhost:9092
```

## fa-users 消费者组 (Consumer Groups)

```bash
kafka-consumer-groups.sh --bootstrap-server localhost:9092 --list
kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe --group my-group
kafka-consumer-groups.sh --bootstrap-server localhost:9092 \
  --delete --group my-group

kafka-consumer-groups.sh --bootstrap-server localhost:9092 \
  --group my-group --topic orders --reset-offsets --to-earliest --execute

kafka-consumer-groups.sh --bootstrap-server localhost:9092 \
  --group my-group --topic orders --reset-offsets --shift-by -10 --execute

kafka-consumer-groups.sh --bootstrap-server localhost:9092 \
  --group my-group --delete-offsets --topic orders
```

## fa-sliders 分区与偏移量 (Partitions & Offsets)

```bash
kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list localhost:9092 --topic orders

kafka-consumer-groups.sh --bootstrap-server localhost:9092 \
  --describe --group my-group --verbose

kafka-verify-consumer-rebalance.sh --zookeeper localhost:2181 \
  --topic orders --group my-group

kafka-preferred-replica-election.sh --bootstrap-server localhost:9092
```

## fa-arrows-rotate 副本复制 (Replication)

```bash
kafka-topics.sh --bootstrap-server localhost:9092 \
  --describe --topic orders --under-replica-partitions

kafka-reassign-partitions.sh --bootstrap-server localhost:9092 \
  --reassignment-json-file reassign.json --execute

kafka-reassign-partitions.sh --bootstrap-server localhost:9092 \
  --reassignment-json-file reassign.json --verify

kafka-replica-verification.sh --broker-list localhost:9092 \
  --topic-white-list orders
```

## fa-gear 配置管理

```bash
kafka-configs.sh --bootstrap-server localhost:9092 \
  --entity-type topics --entity-name orders --describe

kafka-configs.sh --bootstrap-server localhost:9092 \
  --entity-type topics --entity-name orders \
  --alter --add-config retention.ms=86400000

kafka-configs.sh --bootstrap-server localhost:9092 \
  --entity-type topics --entity-name orders \
  --alter --delete-config retention.ms

kafka-configs.sh --bootstrap-server localhost:9092 \
  --entity-type brokers --entity-name 0 --describe
```

## fa-plug Kafka Connect

```bash
connect-standalone.sh config/connect-standalone.properties connector.properties
connect-distributed.sh config/connect-distributed.properties

curl http://localhost:8083/connector-plugins
curl http://localhost:8083/connectors

curl -X POST http://localhost:8083/connectors \
  -H "Content-Type: application/json" \
  -d '{"name":"file-source","config":{"connector.class":"FileStreamSourceConnector","file":"/tmp/input.txt","topic":"connect-test"}}'

curl -X DELETE http://localhost:8083/connectors/file-source
```

## fa-water Kafka Streams 基础

```bash
kafka-run-class.sh org.apache.kafka.streams.examples.wordcount.WordCountDemo

kafka-streams-application-reset.sh --application-id my-app \
  --bootstrap-server localhost:9092 --input-topics orders
```

## fa-shield-halved Schema Registry

```bash
curl http://localhost:8081/subjects
curl http://localhost:8081/subjects/orders-value/versions
curl http://localhost:8081/subjects/orders-value/versions/latest

curl -X POST http://localhost:8081/subjects/orders-value/versions \
  -H "Content-Type: application/vnd.schemaregistry.v1+json" \
  -d '{"schema":"{\"type\":\"record\",\"name\":\"Order\",\"fields\":[{\"name\":\"id\",\"type\":\"string\"}]}"}'

curl -X POST http://localhost:8081/subjects/orders-value \
  -H "Content-Type: application/vnd.schemaregistry.v1+json" \
  -d '{"schema":"{\"type\":\"record\",\"name\":\"Order\",\"fields\":[{\"name\":\"id\",\"type\":\"string\"}]}"}'
```

## fa-chart-line 监控

```bash
kafka-consumer-groups.sh --bootstrap-server localhost:9092 \
  --describe --all-groups

kafka-broker-api-versions.sh --bootstrap-server localhost:9092

kafka-metrics.sh --bootstrap-server localhost:9092

# JMX 指标（启动时设置）
# KAFKA_JMX_OPTS="-Dcom.sun.management.jmxremote -Dcom.sun.management.jmxremote.port=9999"
# JMX_PORT=9999 kafka-server-start.sh config/server.properties
```

## fa-gauge-high 性能调优

```bash
# 生产者配置
# batch.size=16384
# linger.ms=5
# compression.type=lz4
# buffer.memory=33554432
# acks=all

# 消费者配置
# fetch.min.bytes=1048576
# fetch.max.wait.ms=500
# max.partition.fetch.bytes=1048576
# max.poll.records=500

# Broker 配置
# num.io.threads=8
# num.network.threads=3
# socket.send.buffer.bytes=102400
# socket.receive.buffer.bytes=102400
# log.segment.bytes=1073741824
```

## fa-lock 安全 (SASL/TLS)

```bash
kafka-topics.sh --bootstrap-server localhost:9092 --list \
  --command-config client.properties

# client.properties
# security.protocol=SASL_SSL
# sasl.mechanism=PLAIN
# sasl.jaas.config=org.apache.kafka.common.security.plain.PlainLoginModule required username="user" password="pass";

# Broker server.properties
# listeners=SASL_SSL://:9094
# ssl.keystore.location=/var/kafka/keystore.jks
# ssl.keystore.password=changeit
# ssl.truststore.location=/var/kafka/truststore.jks
# sasl.enabled.mechanisms=PLAIN,SCRAM-SHA-256

kafka-configs.sh --bootstrap-server localhost:9092 \
  --entity-type users --entity-name alice \
  --alter --add-config "SCRAM-SHA-256=[password=alice-secret]"
```
