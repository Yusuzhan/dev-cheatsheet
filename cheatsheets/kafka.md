---
title: Kafka
icon: fa-arrows-split-up-and-left
primary: "#231F20"
lang: bash
---

## fa-terminal CLI Basics

```bash
kafka-topics.sh --bootstrap-server localhost:9092 --list
kafka-topics.sh --bootstrap-server localhost:9092 --describe --topic my-topic
kafka-configs.sh --bootstrap-server localhost:9092 --describe --entity-type brokers
kafka-cluster-cluster-id.sh --bootstrap-server localhost:9092
kafka-log-dirs.sh --bootstrap-server localhost:9092 --describe
```

## fa-layer-group Topics

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

## fa-upload Producer

```bash
kafka-console-producer.sh --bootstrap-server localhost:9092 --topic orders

kafka-console-producer.sh --bootstrap-server localhost:9092 \
  --topic orders --property "parse.key=true" --property "key.separator=:"

echo "order-123:payload" | kafka-console-producer.sh \
  --bootstrap-server localhost:9092 --topic orders --property "parse.key=true" --property "key.separator=:"

kafka-producer-perf-test.sh --topic orders --num-records 100000 \
  --record-size 256 --throughput -1 --producer-props bootstrap.servers=localhost:9092
```

## fa-download Consumer

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

## fa-users Consumer Groups

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

## fa-sliders Partitions & Offsets

```bash
kafka-run-class.sh kafka.tools.GetOffsetShell --broker-list localhost:9092 --topic orders

kafka-consumer-groups.sh --bootstrap-server localhost:9092 \
  --describe --group my-group --verbose

kafka-verify-consumer-rebalance.sh --zookeeper localhost:2181 \
  --topic orders --group my-group

kafka-preferred-replica-election.sh --bootstrap-server localhost:9092
```

## fa-arrows-rotate Replication

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

## fa-gear Configuration

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

## fa-water Kafka Streams Basics

```bash
# run a streams application
kafka-run-class.sh org.apache.kafka.streams.examples.wordcount.WordCountDemo

# reset streams application state
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

## fa-chart-line Monitoring

```bash
kafka-consumer-groups.sh --bootstrap-server localhost:9092 \
  --describe --all-groups

kafka-broker-api-versions.sh --bootstrap-server localhost:9092

kafka-metrics.sh --bootstrap-server localhost:9092

# JMX metrics (start with)
# KAFKA_JMX_OPTS="-Dcom.sun.management.jmxremote -Dcom.sun.management.jmxremote.port=9999"
# JMX_PORT=9999 kafka-server-start.sh config/server.properties
```

## fa-gauge-high Performance Tuning

```bash
# producer config
# batch.size=16384
# linger.ms=5
# compression.type=lz4
# buffer.memory=33554432
# acks=all

# consumer config
# fetch.min.bytes=1048576
# fetch.max.wait.ms=500
# max.partition.fetch.bytes=1048576
# max.poll.records=500

# broker config
# num.io.threads=8
# num.network.threads=3
# socket.send.buffer.bytes=102400
# socket.receive.buffer.bytes=102400
# log.segment.bytes=1073741824
```

## fa-lock Security (SASL/TLS)

```bash
kafka-topics.sh --bootstrap-server localhost:9092 --list \
  --command-config client.properties

# client.properties
# security.protocol=SASL_SSL
# sasl.mechanism=PLAIN
# sasl.jaas.config=org.apache.kafka.common.security.plain.PlainLoginModule required username="user" password="pass";

# broker server.properties
# listeners=SASL_SSL://:9094
# ssl.keystore.location=/var/kafka/keystore.jks
# ssl.keystore.password=changeit
# ssl.truststore.location=/var/kafka/truststore.jks
# sasl.enabled.mechanisms=PLAIN,SCRAM-SHA-256

kafka-configs.sh --bootstrap-server localhost:9092 \
  --entity-type users --entity-name alice \
  --alter --add-config "SCRAM-SHA-256=[password=alice-secret]"
```
