---
title: Elasticsearch
icon: fa-magnifying-glass-chart
primary: "#005571"
lang: bash
locale: zhs
---

## fa-layer-group 索引管理

```bash
curl -X PUT "localhost:9200/products" -H 'Content-Type: application/json' -d '
{
  "settings": { "number_of_shards": 3, "number_of_replicas": 1 }
}'
curl -X GET "localhost:9200/_cat/indices?v"
curl -X GET "localhost:9200/products/_settings"
curl -X DELETE "localhost:9200/products"
curl -X POST "localhost:9200/products/_close"
curl -X POST "localhost:9200/products/_open"
curl -X PUT "localhost:9200/products/_settings" -H 'Content-Type: application/json' -d '
{ "index": { "number_of_replicas": 2 } }'
```

## fa-pen-to-square 文档增删改查

```bash
curl -X POST "localhost:9200/products/_doc" -H 'Content-Type: application/json' -d '
{ "name": "laptop", "price": 999, "in_stock": true }'
curl -X PUT "localhost:9200/products/_doc/1" -H 'Content-Type: application/json' -d '
{ "name": "phone", "price": 699 }'
curl -X GET "localhost:9200/products/_doc/1"
curl -X POST "localhost:9200/products/_update/1" -H 'Content-Type: application/json' -d '
{ "doc": { "price": 799 } }'
curl -X DELETE "localhost:9200/products/_doc/1"
curl -X POST "localhost:9200/products/_create/1" -H 'Content-Type: application/json' -d '
{ "name": "tablet", "price": 499 }'
```

## fa-magnifying-glass 搜索 (Query DSL)

```bash
curl -X GET "localhost:9200/products/_search" -H 'Content-Type: application/json' -d '
{
  "query": { "match": { "name": "laptop" } },
  "_source": ["name", "price"],
  "sort": [{ "price": "asc" }],
  "from": 0,
  "size": 10
}'
curl -X GET "localhost:9200/products/_search" -H 'Content-Type: application/json' -d '
{
  "query": { "match_all": {} },
  "size": 5
}'
curl -X GET "localhost:9200/products/_search" -H 'Content-Type: application/json' -d '
{
  "query": { "multi_match": { "query": "laptop", "fields": ["name", "description"] } }
}'
```

## fa-filter 布尔查询

```bash
curl -X GET "localhost:9200/products/_search" -H 'Content-Type: application/json' -d '
{
  "query": {
    "bool": {
      "must": [
        { "match": { "name": "laptop" } }
      ],
      "filter": [
        { "range": { "price": { "gte": 500, "lte": 1500 } } },
        { "term": { "in_stock": true } }
      ],
      "must_not": [
        { "term": { "brand": "refurb" } }
      ],
      "should": [
        { "term": { "category": "electronics" } }
      ]
    }
  }
}'
```

## fa-chart-bar 聚合查询

```bash
curl -X GET "localhost:9200/products/_search" -H 'Content-Type: application/json' -d '
{
  "size": 0,
  "aggs": {
    "avg_price": { "avg": { "field": "price" } },
    "by_category": {
      "terms": { "field": "category", "size": 10 },
      "aggs": {
        "max_price": { "max": { "field": "price" } }
      }
    },
    "price_ranges": {
      "range": {
        "field": "price",
        "ranges": [
          { "to": 500 },
          { "from": 500, "to": 1000 },
          { "from": 1000 }
        ]
      }
    }
  }
}'
```

## fa-sitemap 映射

```bash
curl -X PUT "localhost:9200/products/_mapping" -H 'Content-Type: application/json' -d '
{
  "properties": {
    "name": { "type": "text", "fields": { "keyword": { "type": "keyword" } } },
    "price": { "type": "float" },
    "in_stock": { "type": "boolean" },
    "created_at": { "type": "date", "format": "yyyy-MM-dd" },
    "tags": { "type": "keyword" },
    "description": { "type": "text", "analyzer": "english" },
    "rating": { "type": "float", "coerce": true }
  }
}'
curl -X GET "localhost:9200/products/_mapping"
```

## fa-wand-magic-sparkles 分析器

```bash
curl -X POST "localhost:9200/_analyze" -H 'Content-Type: application/json' -d '
{
  "analyzer": "standard",
  "text": "Hello World 2024"
}'
curl -X PUT "localhost:9200/custom-index" -H 'Content-Type: application/json' -d '
{
  "settings": {
    "analysis": {
      "analyzer": {
        "my_analyzer": {
          "type": "custom",
          "tokenizer": "standard",
          "filter": ["lowercase", "stop", "snowball"]
        }
      }
    }
  }
}'
```

## fa-bolt 批量 API

```bash
curl -X POST "localhost:9200/_bulk" -H 'Content-Type: application/json' -d '
{"index": {"_index": "products", "_id": "1"}}
{"name": "laptop", "price": 999}
{"index": {"_index": "products", "_id": "2"}}
{"name": "phone", "price": 699}
{"update": {"_index": "products", "_id": "1"}}
{"doc": {"price": 899}}
{"delete": {"_index": "products", "_id": "3"}}
'
```

## fa-forward 游标 API

```bash
curl -X GET "localhost:9200/products/_search?scroll=5m" -H 'Content-Type: application/json' -d '
{
  "size": 1000,
  "query": { "match_all": {} }
}'
curl -X GET "localhost:9200/_search/scroll" -H 'Content-Type: application/json' -d '
{
  "scroll": "5m",
  "scroll_id": "DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAAAO4W..."
}'
curl -X DELETE "localhost:9200/_search/scroll" -H 'Content-Type: application/json' -d '
{
  "scroll_id": "DXF1ZXJ5QW5kRmV0Y2gBAAAAAAAAAO4W..."
}'
```

## fa-file-lines 索引模板

```bash
curl -X PUT "localhost:9200/_index_template/logs-template" -H 'Content-Type: application/json' -d '
{
  "index_patterns": ["logs-*"],
  "priority": 100,
  "template": {
    "settings": {
      "number_of_shards": 1,
      "number_of_replicas": 1,
      "index.lifecycle.name": "logs-policy"
    },
    "mappings": {
      "properties": {
        "@timestamp": { "type": "date" },
        "message": { "type": "text" },
        "level": { "type": "keyword" }
      }
    }
  }
}'
curl -X GET "localhost:9200/_index_template/logs-template"
```

## fa-link 索引别名

```bash
curl -X POST "localhost:9200/_aliases" -H 'Content-Type: application/json' -d '
{
  "actions": [
    { "add": { "index": "logs-2024.01", "alias": "logs-current" } },
    { "add": { "index": "logs-2024.02", "alias": "logs-current" } },
    { "remove": { "index": "logs-2023.12", "alias": "logs-current" } }
  ]
}'
curl -X GET "localhost:9200/_cat/aliases?v"
curl -X GET "localhost:9200/logs-current/_search" -H 'Content-Type: application/json' -d '
{
  "query": { "match": { "level": "error" } }
}'
```

## fa-heart-pulse 集群健康

```bash
curl -X GET "localhost:9200/_cluster/health?pretty"
curl -X GET "localhost:9200/_cat/nodes?v"
curl -X GET "localhost:9200/_cat/shards?v"
curl -X GET "localhost:9200/_cluster/stats?pretty"
curl -X GET "localhost:9200/_cat/pending_tasks?v"
curl -X GET "localhost:9200/_nodes/stats?filter_path=nodes.*.jvm.mem.heap_used_percent"
curl -X GET "localhost:9200/_cluster/allocation/explain?pretty"
```

## fa-copy 重建索引

```bash
curl -X POST "localhost:9200/_reindex" -H 'Content-Type: application/json' -d '
{
  "source": { "index": "old-index" },
  "dest": { "index": "new-index" }
}'
curl -X POST "localhost:9200/_reindex" -H 'Content-Type: application/json' -d '
{
  "source": {
    "index": "products",
    "query": { "range": { "price": { "gte": 100 } } }
  },
  "dest": { "index": "expensive-products" }
}'
curl -X POST "localhost:9200/_reindex?wait_for_completion=false" -H 'Content-Type: application/json' -d '
{
  "source": { "index": "logs-2024-*", "size": 5000 },
  "dest": { "index": "archive-logs-2024" }
}'
```

## fa-camera 快照备份

```bash
curl -X PUT "localhost:9200/_snapshot/my_backup" -H 'Content-Type: application/json' -d '
{
  "type": "fs",
  "settings": { "location": "/mnt/backups/es" }
}'
curl -X PUT "localhost:9200/_snapshot/my_backup/snapshot_2024?wait_for_completion=true" -H 'Content-Type: application/json' -d '
{
  "indices": "products,logs-*",
  "ignore_unavailable": true,
  "include_global_state": false
}'
curl -X GET "localhost:9200/_snapshot/my_backup/snapshot_2024"
curl -X POST "localhost:9200/_snapshot/my_backup/snapshot_2024/_restore" -H 'Content-Type: application/json' -d '
{
  "indices": "products",
  "rename_pattern": "(.+)",
  "rename_replacement": "restored_$1"
}'
```
