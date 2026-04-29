---
title: TOML / YAML / JSON
icon: fa-file-lines
primary: "#9B59B6"
lang: yaml
locale: zhs
---

## fa-brackets-curly JSON 语法

```json
{
  "string": "hello",
  "number": 42,
  "float": 3.14,
  "boolean": true,
  "null": null,
  "array": [1, 2, 3],
  "object": {
    "nested": "value"
  }
}
```

```json
{
  "users": [
    {"name": "Alice", "age": 30},
    {"name": "Bob", "age": 25}
  ]
}
```

## fa-check-double JSON Schema

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "type": "object",
  "required": ["name", "email"],
  "properties": {
    "name": {"type": "string", "minLength": 1},
    "email": {"type": "string", "format": "email"},
    "age": {"type": "integer", "minimum": 0},
    "role": {"type": "string", "enum": ["admin", "user"]}
  },
  "additionalProperties": false
}
```

## fa-file-code YAML 基础

```yaml
name: Alice
age: 30
active: true
score: 95.5
hobbies:
  - reading
  - coding
address:
  city: Beijing
  zip: "100000"
multiline: |
  这是第一行
  这是第二行
single: >
  折叠为
  一行
```

## fa-link YAML 锚点与别名

```yaml
defaults: &defaults
  timeout: 30
  retry: 3
  log_level: info

development:
  <<: *defaults
  database: dev_db

production:
  <<: *defaults
  database: prod_db
  log_level: warning
  retry: 5
```

## fa-layer-group YAML 多文档

```yaml
apiVersion: v1
kind: Service
metadata:
  name: svc-a
---
apiVersion: v1
kind: Service
metadata:
  name: svc-b
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: deploy-a
```

## fa-toml TOML 基础

```toml
title = "My App"
version = "1.0.0"
debug = false
port = 8080

[database]
host = "localhost"
port = 5432
name = "mydb"

[logging]
level = "info"
format = "json"
```

## fa-table-columns TOML 表与数组

```toml
[servers.alpha]
ip = "10.0.0.1"
role = "frontend"

[servers.beta]
ip = "10.0.0.2"
role = "backend"

[[products]]
name = "Hammer"
sku = 738594937

[[products]]
name = "Nail"
sku = 284758393
color = "gray"

[[fruits.varieties]]
name = "red Delicious"

[[fruits.varieties]]
name = "Gala"
```

## fa-minimize TOML 内联表

```toml
point = { x = 1, y = 2 }
server = { host = "localhost", port = 8080 }
tags = ["web", "api", "v2"]

[config]
database = { url = "postgres://localhost/db", pool = 10 }
cache = { driver = "redis", ttl = 300 }
```

## fa-clock TOML 日期与时间

```toml
dob = 1990-05-15

created = 2024-01-15T10:30:00Z
updated = 2024-01-15T10:30:00+08:00

duration = "2h30m"
local_time = 10:30:00
local_date = 2024-01-15
```

## fa-code-compare YAML vs JSON vs TOML

```
特性            | JSON  | YAML  | TOML
----------------|-------|-------|------
注释            | 否    | 是    | 是
多文档          | 否    | 是    | 否
锚点/引用       | 否    | 是    | 否
日期类型        | 否    | 是    | 是
内联表          | 否    | 否    | 是
严格类型        | 是    | 否    | 是
可读性          | 低    | 高    | 高
解析速度        | 快    | 慢    | 快
```

## fa-wrench 验证工具

```bash
jq empty config.json
yamllint config.yaml
tomljson config.toml

python -c "import json; json.load(open('config.json'))"
python -c "import yaml; yaml.safe_load(open('config.yaml'))"
python -c "import tomllib; tomllib.load(open('config.toml','rb'))"

npx ajv validate -s schema.json -d config.json
```

## fa-triangle-exclamation 常见陷阱

```yaml
yes_no: true
port: "8080"
date_value: "2024-01-15"
```

```yaml
content: "null"
is_true: "true"
number_str: "0777"
```

```json
{
  "note": "JSON 没有注释、没有多行字符串、不允许尾随逗号"
}
```

```toml
invalid = {a = 1, a = 2}
also_bad = "unterminated
```

## fa-list-check 配置最佳实践

```yaml
use_version_control: true
validate_on_load: true
use_schema_validation: true
keep_sensitive_data_in_env_vars: true
prefer_flat_over_nested: true
use_consistent_naming: true
document_all_fields: true
provide_default_values: true
```

## fa-scale-balanced 格式选择

```
JSON:  API、数据交换、需要严格解析
YAML:  K8s、CI/CD、使用锚点的复杂配置
TOML:  Rust/Go 项目、简单扁平配置、注重可读性
```

```json
{
  "recommendation": "JSON 用于数据交换，YAML 用于基础设施，TOML 用于应用配置"
}
```
