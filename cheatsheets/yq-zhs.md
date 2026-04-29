---
title: yq
icon: fa-file-code
primary: "#4A90D9"
lang: bash
locale: zhs
---

## fa-crosshairs 读取值

```bash
yq '.name' config.yaml
yq '.database.host' config.yaml
yq '.users[0].name' config.yaml
yq '.items[2]' config.yaml
yq '.items | length' config.yaml
yq '.items[] | select(.active == true)' config.yaml
yq '.items[] | .name' config.yaml
yq 'keys' config.yaml
yq '.users[] | select(.age > 25) | .name' config.yaml
```

## fa-pen 写入/更新

```bash
yq '.name = "newvalue"' config.yaml
yq '.database.port = 5433' config.yaml -i
yq '.users[0].role = "admin"' config.yaml -i
yq '.tags += "new-tag"' config.yaml -i
yq '.metadata.version |= . + 1' config.yaml -i
yq '.items[1].price = 29.99' config.yaml -i
yq '(.a.b.c) = "deep"' config.yaml -i
yq '.. | select(tag == "!!str") |= upcase' config.yaml
```

## fa-trash 删除

```bash
yq 'del(.temp)' config.yaml -i
yq 'del(.users[2])' config.yaml -i
yq 'del(.items[] | select(.deprecated == true))' config.yaml -i
yq 'del(.metadata)' config.yaml -i
yq 'del(.tags[0])' config.yaml -i
yq 'del(.empty_field)' config.yaml -i
yq 'del(.users[] | select(.active == false))' config.yaml -i
```

## fa-plus 创建 YAML

```bash
yq -n '.name = "myapp"' > new.yaml
yq -n '.database.host = "localhost" | .database.port = 5432' > db.yaml
cat << 'EOF' | yq '.items = ["a","b","c"]' -
---
EOF
yq -n '.users = [{"name": "alice", "role": "admin"}]' > users.yaml
yq -n '.version = "1.0" | .services = []' > compose.yaml
```

## fa-floppy-disk 原地编辑

```bash
yq -i '.version = "2.0"' config.yaml
yq -i '.metadata.updated = now' config.yaml
yq -i '.items += {"name": "new", "value": 42}' config.yaml
yq -i '... comments=""' config.yaml
yq -i 'sort_keys(..)' config.yaml
```

## fa-code-merge 合并

```bash
yq '. * load("override.yaml")' base.yaml
yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' base.yaml override.yaml
yq -i '. * {"database": {"host": "prod-db"}}' config.yaml
yq '. as $item | $item * {"extra": true}' config.yaml
yq eval-all 'select(fi == 0) * select(fi == 1)' defaults.yaml user.yaml
```

## fa-link 锚点与别名

```bash
yq '.defaults' anchors.yaml
yq '[... | select(anchor == "default")]' anchors.yaml
cat << 'EOF' > anchors.yaml
defaults: &defaults
  timeout: 30
  retries: 3
prod:
  <<: *defaults
  timeout: 60
EOF
yq '.prod' anchors.yaml
yq '.prod.timeout' anchors.yaml
```

## fa-list 数组操作

```bash
yq '.items' data.yaml
yq '.items[0]' data.yaml
yq '.items | length' data.yaml
yq '.items += "new-item"' data.yaml -i
yq '.items = (.items | sort)' data.yaml
yq '.items | unique' data.yaml
yq '.items | reverse' data.yaml
yq '.items | flatten' data.yaml
yq '[.items[] | select(.active)]' data.yaml
yq '.items | map(.price)' data.yaml
yq '.items | group_by(.category)' data.yaml
```

## fa-layer-group 多文档

```bash
yq 'select(di == 0)' multi.yaml         # 第一个文档
yq 'select(di == 1)' multi.yaml         # 第二个文档
yq '.' multi.yaml                        # 所有文档
yq '.[0]' multi.yaml                     # 按索引访问第一个文档
yq '-s' multi.yaml                       # 拆分为多个文件
cat doc1.yaml doc2.yaml | yq eval-all '.'
yq eval-all 'select(fi == 0) * select(fi == 1)' a.yaml b.yaml
```

## fa-arrows-rotate 格式转换

```bash
yq -o json '.' config.yaml               # YAML 转 JSON
yq -p json -o yaml '.' config.json       # JSON 转 YAML
yq -p xml -o yaml '.' data.xml           # XML 转 YAML
yq -p yaml -o xml '.' config.yaml        # YAML 转 XML
yq -p json -o props '.' config.json      # JSON 转 properties
yq -p yaml -o json '.' config.yaml | jq '.'
yq -o tsq '.' config.yaml                # YAML 转类 TOML
cat config.json | yq -p json -o yaml '.' > config.yaml
```

## fa-terminal 求值表达式

```bash
yq '.price * .quantity' order.yaml
yq '.items | map(.price) | add' data.yaml
yq '.users | map(select(.active)) | length' data.yaml
yq '.start + duration("1h")' config.yaml
yq '.items | sort_by(.priority)' data.yaml
yq '.values | map(select(. > 10))' data.yaml
yq '.name | upcase' data.yaml
yq '.path | split("/") | .[-1]' data.yaml
```

## fa-code-compare 比较文件

```bash
diff <(yq -C '.' a.yaml) <(yq -C '.' b.yaml)
diff <(yq -o json '.' a.yaml) <(yq -o json '.' b.yaml)
yq eval-all 'select(fi == 0) - select(fi == 1)' a.yaml b.yaml
yq '. == load("other.yaml")' current.yaml
diff <(yq 'sort_keys(..)' a.yaml) <(yq 'sort_keys(..)' b.yaml)
```

## fa-terminal Shell 集成

```bash
export DB_HOST=$(yq '.database.host' config.yaml)
export DB_PORT=$(yq '.database.port' config.yaml)
for svc in $(yq '.services | keys | .[]' docker-compose.yaml); do
  echo "Service: $svc"
done
yq '.env | to_entries | .[] | .key + "=" + .value' config.yaml | export $(xargs)
HOST=$(yq '.servers[0].host' config.yaml)
PORT=$(yq '.servers[0].port' config.yaml)
curl "http://$HOST:$PORT/health"
```

## fa-lightbulb 实用示例

```bash
# 更新 docker-compose 镜像标签
yq -i '.services.app.image = "myapp:v2.0"' docker-compose.yaml

# 提取所有环境变量
yq '.services.web.environment | to_entries | .[] | "\(.key)=\(.value)"' docker-compose.yaml

# 合并多个 YAML 配置
yq eval-all 'select(fi == 0) * select(fi == 1) * select(fi == 2)' \
  base.yaml dev.yaml local.yaml

# 转换 Kubernetes Secret
yq '.data | to_entries | .[] | .key + ": " + (.value | @base64d)' secret.yaml

# 为所有 K8s 资源添加标签
yq -i '.metadata.labels += {"managed": "true"}' manifest.yaml

# 列出 Helm values 中所有镜像标签
yq '.. | select(tag == "!!map" and has("image")) | .image' values.yaml
```
