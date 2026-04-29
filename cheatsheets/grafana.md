---
title: Grafana
icon: fa-chart-line
primary: "#F46800"
lang: yaml
---

## fa-database Provisioning Datasources

```yaml
apiVersion: 1
datasources:
  - name: Prometheus
    type: prometheus
    access: proxy
    url: http://prometheus:9090
    isDefault: true
  - name: Loki
    type: loki
    access: proxy
    url: http://loki:3100
  - name: Elasticsearch
    type: elasticsearch
    access: proxy
    url: http://elasticsearch:9200
    database: "logstash-*"
    jsonData:
      esVersion: "8.0.0"
      timeField: "@timestamp"
```

## fa-table-columns Provisioning Dashboards

```yaml
apiVersion: 1
providers:
  - name: default
    orgId: 1
    folder: ""
    type: file
    disableDeletion: false
    updateIntervalSeconds: 30
    options:
      path: /var/lib/grafana/dashboards
      foldersFromFilesStructure: true
```

## fa-code Dashboard JSON Model

```yaml
dashboard:
  id: null
  title: My Dashboard
  tags:
    - monitoring
  timezone: utc
  refresh: 30s
  time:
    from: now-6h
    to: now
  panels:
    - title: CPU Usage
      type: timeseries
      gridPos:
        h: 8
        w: 12
        x: 0
        y: 0
      targets:
        - expr: rate(node_cpu_seconds_total{mode!="idle"}[5m])
          legendFormat: "{{instance}} {{mode}}"
```

## fa-chart-area Panel Types

```yaml
timeseries: time-series line/area charts
stat: single big value with optional sparkline
gauge: gauge needle for single value
table: tabular data display
bargauge: horizontal/vertical bar gauges
heatmap: density heatmaps
piechart: pie or donut charts
logs: log viewer with label highlighting
traces: distributed trace visualization
canvas: freeform visual elements
geomap: geographic map panel
text: static markdown content
```

## fa-magnifying-glass PromQL Queries

```yaml
rate(http_requests_total[5m])
sum by (status) (rate(http_requests_total[5m]))
histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[5m]))
topk(10, sum by (path) (rate(http_requests_total[5m])))
avg_over_time(up[1h])
count by (job) (up == 0)
increase(http_requests_total[1h])
label_replace(metric, "dst", "$1", "src", "(.*)")
```

## fa-sliders Variables/Templates

```yaml
templating:
  list:
    - name: datasource
      type: datasource
      query: prometheus
      current:
        text: Prometheus
        value: Prometheus
    - name: job
      type: query
      datasource: $datasource
      query: label_values(up, job)
      sort: 1
    - name: instance
      type: query
      datasource: $datasource
      query: label_values(up{job="$job"}, instance)
      refresh: 2
    - name: interval
      type: interval
      query: 1m,5m,10m,30m,1h
      auto: true
      auto_count: 30
      auto_min: 10s
```

## fa-clock Annotations

```yaml
annotations:
  list:
    - name: Deployments
      datasource: Prometheus
      enable: true
      expr: time() bool on() (changes(deploy_count_total[1m]) > 0)
      titleFormat: Deploy
      tags:
        - deploy
    - name: Outages
      type: tags
      tags:
        - outage
```

## fa-bell Alert Rules

```yaml
apiVersion: 1
groups:
  - orgId: 1
    name: system-alerts
    interval: 1m
    rules:
      - uid: abc123
        title: High CPU Usage
        condition: C
        data:
          - refId: A
            relativeTimeRange:
              from: 600
              to: 0
            datasourceUid: prometheus
            model:
              expr: 100 - (avg by(instance) (rate(node_cpu_seconds_total{mode="idle"}[5m])) * 100)
          - refId: B
            datasourceUid: __expr__
            model:
              type: reduce
              expression: A
              reducer: last
          - refId: C
            datasourceUid: __expr__
            model:
              type: threshold
              expression: B
              conditions:
                - evaluator:
                    params:
                      - 90
                    type: gt
        noDataState: OK
        execErrState: Alerting
        for: 5m
        annotations:
          summary: "CPU usage above 90% on {{ $labels.instance }}"
```

## fa-paper-plane Notification Channels

```yaml
notifiers:
  - name: Slack
    type: slack
    uid: slack-1
    settings:
      url: https://hooks.slack.com/services/XXX/YYY/ZZZ
      recipient: "#alerts"
  - name: Email
    type: email
    uid: email-1
    settings:
      addresses: oncall@example.com
  - name: PagerDuty
    type: pagerduty
    uid: pd-1
    settings:
      integrationKey: YOUR_KEY
      severity: critical
```

## fa-link Dashboard Links

```yaml
links:
  - title: Runbook
    type: link
    url: https://wiki.example.com/runbooks/${__panel.id}
    targetBlank: true
  - title: Trace
    type: link
    url: /d/traces?var_trace_id=${__data.fields.traceID}
    asDropdown: true
```

## fa-arrows-spin Transformation

```yaml
transformations:
  - id: merge
    options: {}
  - id: organize
    options:
      excludeByName:
        time: true
      renameByName:
        value: Requests
  - id: calculateField
    options:
      mode: reduceRow
      reduce:
        reducer: sum
      alias: Total
  - id: filterByValue
    options:
      type: include
      filters:
        - fieldName: Total
          config:
            id: greater
            value:
              - 100
```

## fa-chart-line Common PromQL Patterns

```yaml
http_error_rate: |
  sum(rate(http_requests_total{status=~"5.."}[5m]))
  /
  sum(rate(http_requests_total[5m]))
p95_latency: |
  histogram_quantile(0.95,
    sum(rate(http_request_duration_seconds_bucket[5m])) by (le, path)
  )
memory_usage: |
  (node_memory_MemTotal_bytes - node_memory_MemAvailable_bytes)
  /
  node_memory_MemTotal_bytes * 100
saturation: |
  sum(rate(container_cpu_usage_seconds_total[5m]))
  /
  sum(kube_pod_container_resource_limits{resource="cpu"})
```

## fa-keyboard Useful Shortcuts

```yaml
global_shortcuts:
  "/": open dashboard search
  "Ctrl+K": open command palette
  "Ctrl+S": save dashboard
  "d+r": refresh all panels
  "d+z": toggle kiosk/zen mode
  "d+l": toggle panel legend
  "t+t": open time picker
  "Escape": exit fullscreen/kiosk
  "Ctrl+H": hide panel menu
panel_edit:
  "e": toggle panel edit mode
  "v": toggle panel inspect
  "ps": open panel search
```
