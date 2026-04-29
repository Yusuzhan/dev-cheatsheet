---
title: Kibana
icon: fa-chart-bar
primary: "#005571"
lang: yaml
locale: zhs
---

## fa-search KQL 语法

```yaml
status: 200
response: 200 OR response: 404
level: "error" AND service: "api"
message: "connection timeout"
not status: 500
@timestamp >= "2024-01-01" AND @timestamp < "2024-02-01"
host.name: "web-*"
```

## fa-code Lucene 语法

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

## fa-magnifying-glass Discover 发现

```yaml
workflow: |
  1. 选择数据视图 / 索引模式
  2. 设置时间范围（右上角）
  3. 在搜索栏输入 KQL 或 Lucene 查询
  4. 点击字段名过滤（+/- 图标）
  5. 展开文档查看所有字段
  6. 保存搜索供后续复用

tips:
  - 使用时间选择器预设：最近 15m、1h、24h、7d、30d
  - 点击直方图柱子缩放到该时间范围
  - 在直方图上拖拽选择自定义范围
  - 使用"检查"查看底层 Elasticsearch 查询
```

## fa-chart-pie 可视化库

```yaml
visualization_types:
  lens: 拖拽式可视化构建器
  aggregation-based: 传统图表构建器
  vega/json: 自定义 vega-lite 可视化
  maps: 坐标和区域地图
  markdown: 静态文本/Markdown 面板
  canvas: 像素级精美演示
  tsvb: 时间序列可视化构建器

creating: |
  1. 导航到可视化库
  2. 点击"创建可视化"
  3. 选择可视化类型
  4. 选择数据视图
  5. 配置聚合和指标
  6. 保存到库中供仪表板复用
```

## fa-table-columns 仪表板

```yaml
features:
  - 基于网格的面板布局（拖拽和缩放）
  - 时间范围应用于所有面板
  - 全局和面板级过滤器
  - 仪表板下钻和链接
  - 输入控件（下拉框、滑块）
  - 可嵌入的保存搜索和可视化
  - 导入/导出为 NDJSON
  - 分享为 PDF/PNG 报告

api_create: |
  POST /api/dashboards/dashboard
  {
    "dashboard": { ... },
    "overwrite": true
  }
```

## fa-floppy-disk 保存的搜索

```yaml
create: |
  1. 在 Discover 中构建查询
  2. 选择要显示的列
  3. 点击"保存" > "保存搜索"
  4. 命名并可选添加到仪表板

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

## fa-layer-group 索引模式

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

## fa-filter 字段过滤器

```yaml
positive_filter: |
  点击字段值"+"图标过滤该值
negative_filter: |
  点击字段值"-"图标排除该值
pin_filter: |
  点击固定图标将过滤器应用到所有仪表板
custom_filter: |
  点击"添加过滤器" > 选择字段 > 运算符 > 值
  运算符：是、不是、属于、不属于、存在、不存在
```

## fa-calculator 聚合

```yaml
metric: avg, sum, min, max, cardinality, value_count
bucket:
  date_histogram: 按时间间隔分组
  terms: 按字段热门值分组
  histogram: 按数值范围分组
  range: 自定义数值范围
  filters: 自定义桶过滤器
pipeline:
  moving_avg: 桶上的移动平均
  derivative: 桶间变化率
  cumulative_sum: 累计总和
  serial_diff: 滞后间隔的桶值差
```

## fa-palette Canvas 画布

```yaml
workpad: 像素级精美的报告和演示
elements: 表格、图表、图片、形状、文本
datasource: Elasticsearch SQL、ES|QL 或原始 ES 查询
expression_language: |
  filters
  | essql query="SELECT * FROM logs" count=100
  | mapColumn "time" fn={getCell "timestamp" | formatdate "YYYY-MM-DD"}
  | table
  | render
```

## fa-chart-line Lens 可视化

```yaml
drag_and_drop: |
  1. 从可视化库打开 Lens
  2. 拖拽字段到水平/垂直轴
  3. 从建议中选择可视化类型
  4. 在同一图表上叠加多个指标
  5. 在柱状、折线、面积、饼图之间切换

formula: |
  count() / kibana_sample_data_flights.AvgTicketPrice
  sum(bytes) / unique_count(clientip)
  last_value(cpu.usage, kql="host.name:web-01")
```

## fa-bell 告警

```yaml
alert_rule: |
  1. Stack Management > Alerts and Insights > Rules
  2. Create rule > 选择规则类型
  3. 定义条件（阈值、库存等）
  4. 设置检查间隔
  5. 配置动作（邮件、Slack、Webhook）
  6. 设置动作频率（每次告警、摘要）

threshold_rule: |
  WHEN count() OVER all documents
  FOR THE LAST 5 minutes
  IS ABOVE 1000
  THEN send email to oncall@example.com
```

## fa-users 空间与 RBAC

```yaml
spaces: |
  默认空间："Default"
  创建：Management > Spaces > Create space
  用例：按团队、环境或功能分离
  通过导入/导出在空间之间复制对象

rbac: |
  Management > Security > Roles
  kibana_user: 基础访问权限
  kibana_admin: 完整 Kibana 访问权限
  read_only: 仅查看仪表板
  custom: 细粒度索引和功能权限
  Feature-level: 按 Kibana 应用控制访问
```
