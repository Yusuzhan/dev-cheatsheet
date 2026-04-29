---
title: yq
icon: fa-file-code
primary: "#4A90D9"
lang: bash
---

## fa-crosshairs Read Values

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

## fa-pen Write / Update

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

## fa-trash Delete

```bash
yq 'del(.temp)' config.yaml -i
yq 'del(.users[2])' config.yaml -i
yq 'del(.items[] | select(.deprecated == true))' config.yaml -i
yq 'del(.metadata)' config.yaml -i
yq 'del(.tags[0])' config.yaml -i
yq 'del(.empty_field)' config.yaml -i
yq 'del(.users[] | select(.active == false))' config.yaml -i
```

## fa-plus Create YAML

```bash
yq -n '.name = "myapp"' > new.yaml
yq -n '.database.host = "localhost" | .database.port = 5432' > db.yaml
cat << 'EOF' | yq '.items = ["a","b","c"]' -
---
EOF
yq -n '.users = [{"name": "alice", "role": "admin"}]' > users.yaml
yq -n '.version = "1.0" | .services = []' > compose.yaml
```

## fa-floppy-disk In-place Edit

```bash
yq -i '.version = "2.0"' config.yaml
yq -i '.metadata.updated = now' config.yaml
yq -i '.items += {"name": "new", "value": 42}' config.yaml
yq -i '... comments=""' config.yaml
yq -i 'sort_keys(..)' config.yaml
```

## fa-code-merge Merge

```bash
yq '. * load("override.yaml")' base.yaml
yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' base.yaml override.yaml
yq -i '. * {"database": {"host": "prod-db"}}' config.yaml
yq '. as $item | $item * {"extra": true}' config.yaml
yq eval-all 'select(fi == 0) * select(fi == 1)' defaults.yaml user.yaml
```

## fa-link Anchors & Aliases

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

## fa-list Array Operations

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

## fa-layer-group Multiple Documents

```bash
yq 'select(di == 0)' multi.yaml         # first document
yq 'select(di == 1)' multi.yaml         # second document
yq '.' multi.yaml                        # all documents
yq '.[0]' multi.yaml                     # first doc by index
yq '-s' multi.yaml                       # split into files
cat doc1.yaml doc2.yaml | yq eval-all '.'
yq eval-all 'select(fi == 0) * select(fi == 1)' a.yaml b.yaml
```

## fa-arrows-rotate Format Conversion

```bash
yq -o json '.' config.yaml               # YAML to JSON
yq -p json -o yaml '.' config.json       # JSON to YAML
yq -p xml -o yaml '.' data.xml           # XML to YAML
yq -p yaml -o xml '.' config.yaml        # YAML to XML
yq -p json -o props '.' config.json      # JSON to properties
yq -p yaml -o json '.' config.yaml | jq '.'
yq -o tsq '.' config.yaml                # YAML to TOML-like
cat config.json | yq -p json -o yaml '.' > config.yaml
```

## fa-terminal Evaluation Expressions

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

## fa-code-compare Compare Files

```bash
diff <(yq -C '.' a.yaml) <(yq -C '.' b.yaml)
diff <(yq -o json '.' a.yaml) <(yq -o json '.' b.yaml)
yq eval-all 'select(fi == 0) - select(fi == 1)' a.yaml b.yaml
yq '. == load("other.yaml")' current.yaml
diff <(yq 'sort_keys(..)' a.yaml) <(yq 'sort_keys(..)' b.yaml)
```

## fa-terminal Shell Integration

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

## fa-lightbulb Practical Examples

```bash
# Update docker-compose image tag
yq -i '.services.app.image = "myapp:v2.0"' docker-compose.yaml

# Extract all environment variables
yq '.services.web.environment | to_entries | .[] | "\(.key)=\(.value)"' docker-compose.yaml

# Merge multiple YAML configs
yq eval-all 'select(fi == 0) * select(fi == 1) * select(fi == 2)' \
  base.yaml dev.yaml local.yaml

# Convert Kubernetes secret
yq '.data | to_entries | .[] | .key + ": " + (.value | @base64d)' secret.yaml

# Add label to all K8s resources in file
yq -i '.metadata.labels += {"managed": "true"}' manifest.yaml

# List all image tags in Helm values
yq '.. | select(tag == "!!map" and has("image")) | .image' values.yaml
```
