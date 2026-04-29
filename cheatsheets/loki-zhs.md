---
title: Loki
icon: fa-fire
primary: "#F46800"
lang: yaml
locale: zhs
---

## fa-search LogQL 基础查询

```yaml
{job="varlogs"}
{job="nginx",env="prod"}
{job=~"nginx|apache"}
{job="varlogs"} |= "error"
{job="varlogs"} |~ "error|timeout"
{job="varlogs"} != "debug"
```

## fa-tags 标签过滤

```yaml
{job="nginx",status="500"}
{cluster="us-east",namespace=~"prod-.*"}
{level=~"error|warn",service="api"}
{job="app"} | label_format level={{.severity}}
{job="app"} | drop level, unused_label
{job="app"} | keep job, level, message
```

## fa-filter 行过滤

```yaml
{job="nginx"} |= "GET /api"
{job="nginx"} |~ "GET /api/(users|orders)"
{job="nginx"} != "200"
{job="nginx"} !~ "(200|301|302)"
{job="app"} |= "error" != "timeout"
{job="app"} |~ `(?i)error.*database`
```

## fa-arrows-spin 管道阶段

```yaml
{job="app"}
  |> stage 1
  |> stage 2

common_stages:
  label_extract: |json or |logfmt or |pattern
  line_filter: |= or |~ or != or !~
  label_filter: | label=="value"
  format: | line_format "{{.message}}"
  drop: | drop label
  keep: | keep label1, label2
  unpack: | unpack
```

## fa-chart-line LogQL 指标查询

```yaml
log_count: sum(count_over_time({job="nginx"}[5m]))
error_rate: |
  sum(rate({job="nginx"} |= "error" [5m]))
  /
  sum(rate({job="nginx"}[5m]))
bytes_total: sum(bytes_over_time({job="nginx"}[5m]))
latency_p99: |
  quantile_over_time(0.99,
    {job="app"} | json | unwrap latency_ms [5m]
  ) by (path)
avg_by_path: |
  avg_over_time(
    {job="app"} | json | unwrap latency_ms [5m]
  ) by (path)
top_10: topk(10, sum by (host) (rate({job="nginx"}[5m])))
```

## fa-gear Loki 配置

```yaml
auth_enabled: false
server:
  http_listen_port: 3100
  grpc_listen_port: 9096
common:
  instance_addr: 127.0.0.1
  path_prefix: /loki
  storage:
    filesystem:
      chunks_directory: /loki/chunks
      rules_directory: /loki/rules
  replication_factor: 1
  ring:
    kvstore:
      store: inmemory
schema_config:
  configs:
    - from: 2020-10-24
      store: boltdb-shipper
      object_store: filesystem
      schema: v11
      index:
        prefix: index_
        period: 24h
storage_config:
  filesystem:
    directory: /loki/storage
```

## fa-server 本地部署

```yaml
docker_run: |
  docker run -d --name loki \
    -p 3100:3100 \
    -v $(pwd)/loki-config.yaml:/etc/loki/local-config.yaml \
    grafana/loki:latest \
    -config.file=/etc/loki/local-config.yaml

docker_compose: |
  version: "3"
  services:
    loki:
      image: grafana/loki:latest
      ports:
        - "3100:3100"
      volumes:
        - ./loki-config.yaml:/etc/loki/local-config.yaml
      command: -config.file=/etc/loki/local-config.yaml
```

## fa-network-wired Distributor

```yaml
distributor:
  ring:
    kvstore:
      store: inmemory
  rate_limiting_strategy: local
  max_receive_batch_size: 1048576
  max_line_size: 0
  override_ring_key: distributors
```

## fa-database Ingester

```yaml
ingester:
  lifecycler:
    address: 127.0.0.1
    ring:
      kvstore:
        store: inmemory
      replication_factor: 1
    final_sleep: 0s
  chunk_idle_period: 5m
  chunk_block_size: 262144
  chunk_encoding: snappy
  chunk_retain_period: 1m
  max_chunk_age: 2h
  max_transfer_retries: 0
  wal:
    enabled: true
    dir: /loki/wal
```

## fa-magnifying-glass Querier

```yaml
querier:
  engine:
    max_look_back_period: 0s
    timeout: 3m
  max_concurrent: 2048
  query_ingesters_within: 3h
  query_store_only: false
frontend_worker:
  frontend_address: 127.0.0.1:9095
  parallelism: 10
  match_max_concurrent: true
```

## fa-compress Compactor

```yaml
compactor:
  working_directory: /loki/compactor
  shared_store: filesystem
  compaction_interval: 10m
  retention_enabled: true
  retention_delete_delay: 2h
  retention_delete_worker_count: 150
  delete_request_store: filesystem
```

## fa-clock 数据保留

```yaml
limits_config:
  retention_period: 744h
  max_query_length: 721h
  max_query_parallelism: 32
compactor:
  retention_enabled: true
  retention_delete_delay: 2h
table_manager:
  retention_deletes_enabled: true
  retention_period: 744h
chunk_store_config:
  max_look_back_period: 744h
```

## fa-star 最佳实践

```yaml
label_cardinality: |
  保持标签基数较低（<10k 唯一值）。
  避免使用 user_id、request_id、IP 作为标签。
  高基数据使用 structured metadata。
chunk_size: |
  优化 chunk_idle_period 和 max_chunk_age。
  较小的 chunk = 更快的查询但更多开销。
queries: |
  始终先按 job/app 标签过滤。
  在标签提取前使用行过滤器（|=, |~）。
  rate() 查询使用 [5m] 或 [1m] 时间范围。
storage: |
  生产环境使用对象存储（S3, GCS）。
  启用 WAL 保证持久性。
  配置 compactor 管理数据保留。
```
