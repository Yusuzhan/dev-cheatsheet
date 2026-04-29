---
title: Prometheus
icon: fa-fire
primary: "#E6522C"
lang: yaml
---

## fa-chart-line PromQL Basics

```yaml
up
up == 0
count(up) by (job)
```

```yaml
rate(http_requests_total[5m])
irate(http_requests_total[5m])
deriv(cpu_temp[30m])
```

## fa-clock Instant vs Range Queries

```yaml
http_requests_total
http_requests_total[5m]
```

```yaml
rate(http_requests_total[5m])
avg_over_time(cpu_temp[1h])
```

## fa-crosshairs Selectors & Label Matchers

```yaml
http_requests_total{job="api-server"}
http_requests_total{method!="GET"}
http_requests_total{status=~"5.."}
http_requests_total{job=~"api-.*",env="prod"}
```

## fa-calculator Operators & Functions

```yaml
sum(rate(http_requests_total[5m])) by (job)
histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m]))
topk(5, rate(http_requests_total[5m]))
```

```yaml
label_replace(metric, "dst", "$1", "src", "(.*)")
clamp_max(clamp_min(metric, 0), 100)
vector(1)
```

## fa-chart-bar Histogram & Summary

```yaml
histogram_quantile(0.99,
  rate(http_request_duration_seconds_bucket[5m])
)
```

```yaml
rate(http_request_duration_seconds_sum[5m])
  /
rate(http_request_duration_seconds_count[5m])
```

## fa-file-alt Recording Rules

```yaml
groups:
  - name: api_rules
    interval: 30s
    rules:
      - record: job:http_requests:rate5m
        expr: sum(rate(http_requests_total[5m])) by (job)
      - record: job:http_errors:ratio
        expr: |
          sum(rate(http_errors_total[5m])) by (job)
            /
          sum(rate(http_requests_total[5m])) by (job)
```

## fa-bell Alerting Rules

```yaml
groups:
  - name: api_alerts
    rules:
      - alert: HighErrorRate
        expr: |
          sum(rate(http_errors_total[5m])) by (job)
            /
          sum(rate(http_requests_total[5m])) by (job) > 0.05
        for: 10m
        labels:
          severity: critical
        annotations:
          summary: "High error rate on {{ $labels.job }}"
```

## fa-download scrape_config

```yaml
scrape_configs:
  - job_name: prometheus
    scrape_interval: 15s
    scrape_timeout: 10s
    metrics_path: /metrics
    honor_labels: false
    static_configs:
      - targets: ["localhost:9090"]
        labels:
          env: prod
```

## fa-search Service Discovery

```yaml
scrape_configs:
  - job_name: kubernetes-pods
    kubernetes_sd_configs:
      - role: pod
    relabel_configs:
      - source_labels: [__meta_kubernetes_pod_annotation_prometheus_io_scrape]
        action: keep
        regex: true

  - job_name: consul
    consul_sd_configs:
      - server: localhost:8500
        services: [api, web]
```

## fa-tags Relabeling

```yaml
relabel_configs:
  - source_labels: [__meta_kubernetes_namespace]
    target_label: namespace
  - source_labels: [__meta_kubernetes_pod_name]
    target_label: pod
  - source_labels: [__address__]
    regex: "(.*):8080"
    replacement: "${1}:9090"
    target_label: __address__
  - regex: "__meta_kubernetes_pod_label_(.+)"
    action: labelmap
```

## fa-sitemap Federation

```yaml
scrape_configs:
  - job_name: federate
    scrape_interval: 30s
    honor_labels: true
    metrics_path: /federate
    params:
      match[]:
        - '{job="api-server"}'
        - '{__name__=~"job:.*"}'
    static_configs:
      - targets: ["prometheus-dc1:9090", "prometheus-dc2:9090"]
```

## fa-server HA / Thanos

```yaml
global:
  external_labels:
    cluster: us-east-1
    replica: "1"
```

```yaml
thanos_sidecar:
  obj_store:
    type: S3
    config:
      bucket: thanos-data
      endpoint: s3.amazonaws.com
```

```yaml
remote_write:
  - url: http://thanos-receive:10908/api/v1/receive
remote_read:
  - url: http://thanos-query:10914/api/v1/read
```

## fa-globe HTTP API

```yaml
GET /api/v1/query?query=up&time=2025-01-01T00:00:00Z
GET /api/v1/query_range?query=up&start=...&end=...&step=15s
GET /api/v1/series?match[]=up
GET /api/v1/labels
GET /api/v1/label/__name__/values
GET /api/v1/alerts
GET /api/v1/rules
```

## fa-star Best Practices

```yaml
global:
  scrape_interval: 15s
  evaluation_interval: 15s
  scrape_timeout: 10s
  external_labels:
    cluster: prod
```

```yaml
rule_files:
  - "rules/*.yml"
alerting:
  alertmanagers:
    - static_configs:
        - targets: ["alertmanager:9093"]
```

```yaml
storage:
  tsdb:
    retention.time: 30d
    retention.size: 50GB
```
