---
title: Promtail
icon: fa-arrow-right-to-bracket
primary: "#F46800"
lang: yaml
---

## fa-gear Basic Config

```yaml
server:
  http_listen_port: 9080
  grpc_listen_port: 0
positions:
  filename: /var/lib/promtail/positions.yaml
clients:
  - url: http://loki:3100/loki/api/v1/push
scrape_configs:
  - job_name: system
    static_configs:
      - targets:
          - localhost
        labels:
          job: varlogs
          __path__: /var/log/*.log
```

## fa-list Static Jobs

```yaml
scrape_configs:
  - job_name: nginx
    static_configs:
      - targets:
          - localhost
        labels:
          job: nginx
          env: prod
          __path__: /var/log/nginx/*.log
  - job_name: app
    static_configs:
      - targets:
          - localhost
        labels:
          job: myapp
          __path__: /opt/app/logs/*.log
```

## fa-folder-open File Discovery

```yaml
scrape_configs:
  - job_name: dynamic_logs
    file_sd_configs:
      - files:
          - /etc/promtail/file_targets/*.yaml
        refresh_interval: 5m
    relabel_configs:
      - source_labels: [__path__]
        target_label: path
```

## fa-dharmachakra Kubernetes Discovery

```yaml
scrape_configs:
  - job_name: kubernetes-pods
    kubernetes_sd_configs:
      - role: pod
    relabel_configs:
      - source_labels:
          - __meta_kubernetes_pod_label_app
        target_label: app
      - source_labels:
          - __meta_kubernetes_namespace
        target_label: namespace
      - source_labels:
          - __meta_kubernetes_pod_name
        target_label: pod
      - source_labels:
          - __meta_kubernetes_pod_container_name
        target_label: container
      - replacement: /var/log/pods/*$1/*.log
        separator: /
        source_labels:
          - __meta_kubernetes_pod_uid
        target_label: __path__
```

## fa-arrows-spin Pipeline Stages

```yaml
pipeline_stages:
  - json:
      expressions:
        level: level
        message: msg
        timestamp: ts
        traceID: trace_id
  - labels:
      level:
      traceID:
  - timestamp:
      source: timestamp
      format: Unix
  - output:
      source: message
```

## fa-tags Label Extraction

```yaml
pipeline_stages:
  - regex:
      expression: '^(?P<time>\S+) (?P<level>\S+) (?P<msg>.*)$'
  - labels:
      level:
  - timestamp:
      source: time
      format: "2006-01-02T15:04:05.000Z"
  - output:
      source: msg
```

## fa-clock Timestamp Parsing

```yaml
pipeline_stages:
  - timestamp:
      source: time
      format: RFC3339
  - timestamp:
      source: ts
      format: Unix
  - timestamp:
      source: ts_ms
      format: UnixMs
  - timestamp:
      source: datetime
      format: "2006-01-02 15:04:05"
```

## fa-layer-group Structured Metadata

```yaml
pipeline_stages:
  - json:
      expressions:
        level: level
        request_id: req_id
        user_id: uid
  - structured_metadata:
      level:
      request_id:
  - labels:
      level:
  - output:
      source: message
```

## fa-chart-line Metrics from Logs

```yaml
pipeline_stages:
  - json:
      expressions:
        status: status
        duration: duration_ms
        path: request_path
  - labels:
      status:
      path:
  - metrics:
      request_total:
        type: Counter
        description: Total requests
        config:
          action: inc
      request_duration:
        type: Histogram
        description: Request duration
        config:
          value: duration
          buckets: [0.1, 0.5, 1, 5, 10]
```

## fa-filter Multi-stage Pipeline

```yaml
pipeline_stages:
  - match:
      selector: '{job="nginx"} |= "error"'
      stages:
        - regex:
            expression: '^(?P<remote_addr>\S+) - - \[(?P<time>[^\]]+)\] "(?P<method>\S+) (?P<path>\S+) (?P<status>\d+) (?P<size>\d+)"'
        - labels:
            method:
            status:
        - timestamp:
            source: time
            format: "02/Jan/2006:15:04:05 -0700"
  - match:
      selector: '{job="app"} |~ "level=(?P<level>\w+)"'
      stages:
        - regex:
            expression: 'level=(?P<level>\w+) msg="(?P<msg>[^"]*)"'
        - labels:
            level:
        - output:
            source: msg
```

## fa-server Client Options

```yaml
clients:
  - url: http://loki:3100/loki/api/v1/push
    tenant_id: team-a
    basic_auth:
      username: admin
      password: secret
    bearer_token: eyJhbGciOiJIUzI1NiIs
    external_labels:
      cluster: us-east-1
      env: production
    timeout: 10s
    batch_wait: 1s
    batch_size: 1048576
    max_backoff: 5s
    max_retries: 10
    min_backoff: 500ms
```

## fa-shuffle Relabeling

```yaml
scrape_configs:
  - job_name: kubernetes
    kubernetes_sd_configs:
      - role: pod
    relabel_configs:
      - source_labels: [__meta_kubernetes_namespace]
        target_label: namespace
      - source_labels: [__meta_kubernetes_pod_name]
        target_label: pod
      - source_labels: [__meta_kubernetes_pod_container_name]
        target_label: container
      - regex: "true"
        source_labels: [__meta_kubernetes_pod_label_promtail_ignore]
        action: drop
      - action: labelmap
        regex: __meta_kubernetes_pod_label_(.+)
```
