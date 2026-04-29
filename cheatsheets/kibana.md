---
title: Kibana
icon: fa-chart-bar
primary: "#005571"
lang: yaml
---

## fa-search KQL Syntax

```yaml
status: 200
response: 200 OR response: 404
level: "error" AND service: "api"
message: "connection timeout"
not status: 500
@timestamp >= "2024-01-01" AND @timestamp < "2024-02-01"
host.name: "web-*"
```

## fa-code Lucene Syntax

```yaml
status:200 AND extension:php
response:[400 TO 499]
message:"out of memory"
NOT status:200
host:web-server-01 OR host:web-server-02
*exception*
logger\:org.elasticsearch*
_exists_:field_name
```

## fa-magnifying-glass Discover

```yaml
workflow: |
  1. Select data view / index pattern
  2. Set time range (top right)
  3. Enter KQL or Lucene query in search bar
  4. Click field names to filter (+/- icons)
  5. Expand documents to view all fields
  6. Save search for later reuse

tips:
  - Use time picker presets: Last 15m, 1h, 24h, 7d, 30d
  - Click histogram bars to zoom into time range
  - Drag-select on histogram for custom range
  - Use "Inspect" to see underlying Elasticsearch query
```

## fa-chart-pie Visualize Library

```yaml
visualization_types:
  lens: drag-and-drop visual builder
  aggregation-based: traditional chart builder
  vega/json: custom vega-lite visualizations
  maps: coordinate and region maps
  markdown: static text/markdown panels
  canvas: pixel-perfect presentations
  tsvb: time series visual builder

creating: |
  1. Navigate to Visualize Library
  2. Click "Create visualization"
  3. Select visualization type
  4. Choose data view
  5. Configure aggregations and metrics
  6. Save to library for dashboard reuse
```

## fa-table-columns Dashboard

```yaml
features:
  - Grid-based panel layout (drag & resize)
  - Time range applies to all panels
  - Global and per-panel filters
  - Dashboard drilldowns and links
  - Input controls (dropdowns, sliders)
  - Embeddable saved searches & visualizations
  - Export/import as NDJSON
  - Share as PDF/PNG reports

api_create: |
  POST /api/dashboards/dashboard
  {
    "dashboard": { ... },
    "overwrite": true
  }
```

## fa-floppy-disk Saved Searches

```yaml
create: |
  1. Build query in Discover
  2. Select columns to display
  3. Click "Save" > "Save search"
  4. Name and optionally add to dashboard

api: |
  GET  /api/saved_objects/_find?type=search
  POST /api/saved_objects/search
  {
    "attributes": {
      "title": "my-search",
      "kibanaSavedObjectMeta": {
        "searchSourceJSON": "{\"indexRefName\":\"kibanaSavedObjectMeta.searchSourceJSON.index\"}"
      }
    },
    "references": [...]
  }
```

## fa-layer-group Index Patterns

```yaml
create: |
  Management > Data Views > Create data view
  Name: logstash-*
  Timestamp field: @timestamp

scripted_field: |
  Management > Data Views > Select view > Add scripted field
  Language: Painless
  Type: number
  Script: doc['response'].value * 100

api: |
  POST /api/saved_objects/data-view
  {
    "attributes": {
      "title": "logs-*",
      "timeFieldName": "@timestamp"
    }
  }
```

## fa-filter Field Filters

```yaml
positive_filter: |
  Click field value "+" icon to filter for that value
negative_filter: |
  Click field value "-" icon to filter out that value
pin_filter: |
  Click pin icon to apply filter across all dashboards
custom_filter: |
  Click "Add filter" > select field > operator > value
  Operators: is, is not, is one of, is not one of, exists, does not exist
```

## fa-calculator Aggregations

```yaml
metric: avg, sum, min, max, cardinality, value_count
bucket:
  date_histogram: group by time interval
  terms: group by field top values
  histogram: group by numeric ranges
  range: custom numeric ranges
  filters: custom bucket filters
pipeline:
  moving_avg: moving average over buckets
  derivative: rate of change between buckets
  cumulative_sum: running total
  serial_diff: difference between bucket values at lag interval
```

## fa-palette Canvas

```yaml
workpad: pixel-perfect reports and presentations
elements: tables, charts, images, shapes, text
datasource: Elasticsearch SQL, ES|QL, or raw ES queries
expression_language: |
  filters
  | essql query="SELECT * FROM logs" count=100
  | mapColumn "time" fn={getCell "timestamp" | formatdate "YYYY-MM-DD"}
  | table
  | render
```

## fa-chart-line Lens

```yaml
drag_and_drop: |
  1. Open Lens from Visualize Library
  2. Drag field to horizontal/vertical axis
  3. Choose visualization type from suggestions
  4. Layer multiple metrics on same chart
  5. Switch between bar, line, area, pie views

formula: |
  count() / kibana_sample_data_flights.AvgTicketPrice
  sum(bytes) / unique_count(clientip)
  last_value(cpu.usage, kql="host.name:web-01")
```

## fa-bell Alerting

```yaml
alert_rule: |
  1. Stack Management > Alerts and Insights > Rules
  2. Create rule > Select rule type
  3. Define conditions (threshold, inventory, etc.)
  4. Set check interval
  5. Configure action (email, Slack, webhook)
  6. Set action frequency (on each alert, summary)

threshold_rule: |
  WHEN count() OVER all documents
  FOR THE LAST 5 minutes
  IS ABOVE 1000
  THEN send email to oncall@example.com
```

## fa-users Spaces & RBAC

```yaml
spaces: |
  Default: "Default" space
  Create: Management > Spaces > Create space
  Use cases: separate by team, environment, or function
  Copy objects between spaces via import/export

rbac: |
  Management > Security > Roles
  kibana_user: basic access
  kibana_admin: full Kibana access
  read_only: view dashboards only
  custom: fine-grained index and feature permissions
  Feature-level: control access per Kibana app
```
