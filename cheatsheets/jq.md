---
title: jq
icon: fa-brackets-curly
primary: "#C7254E"
lang: bash
---

## fa-crosshairs Basic Selectors

```bash
echo '{"name":"alice","age":30}' | jq '.name'
echo '{"name":"alice","age":30}' | jq '.age'
echo '{"a":{"b":{"c":1}}}' | jq '.a.b.c'
echo '{"user":{"name":"bob"}}' | jq '.user.name'
echo '{"data":[1,2,3]}' | jq '.data'
echo '{"data":[1,2,3]}' | jq '.data[0]'
echo '{"data":[1,2,3]}' | jq '.data[-1]'
echo '.' <<< '{"any":"json"}'
```

## fa-hammer Object Construction

```bash
echo '{"first":"Alice","last":"Smith"}' | jq '{name: .first, full: "\(.first) \(.last)"}'
echo '{"a":1,"b":2}' | jq '{a, b, sum: (.a + .b)}'
echo '[1,2,3]' | jq '{items: ., count: length}'
echo '{"x":1,"y":null}' | jq '{x, y, z: "added"}'
echo '{"a":1}' | jq '{a: .a} + {b: 2}'
```

## fa-list Array Operations

```bash
echo '[1,2,3,4,5]' | jq '.[1:3]'
echo '[1,2,3,4,5]' | jq '.[:3]'
echo '[1,2,3,4,5]' | jq '.[2:]'
echo '[1,2,3,4,5]' | jq 'length'
echo '[{"n":1},{"n":2},{"n":3}]' | jq '.[].n'
echo '[[1,2],[3,4]]' | jq 'flatten'
echo '[1,2,2,3]' | jq 'unique'
echo '[3,1,2]' | jq 'sort'
echo '[1,2,3]' | jq 'reverse'
```

## fa-filter Filters & select

```bash
echo '[{"name":"a","v":1},{"name":"b","v":2},{"name":"c","v":3}]' | \
  jq '.[] | select(.v > 1)'
echo '[{"name":"a","v":1},{"name":"b","v":2}]' | \
  jq '.[] | select(.name == "b")'
echo '[1,2,3,4,5]' | jq '[.[] | select(. % 2 == 0)]'
echo '[{"a":1},{"a":null},{"a":2}]' | jq '[.[] | select(.a != null)]'
echo '[{"t":"a"},{"t":"b"},{"t":"a"}]' | jq '[.[] | select(.t == "a")]'
echo '[{"x":1},{"x":2},{"x":3}]' | jq '[.[] | select(.x >= 2)] | length'
```

## fa-font String Functions

```bash
echo '"hello world"' | jq 'ascii_upcase'
echo '"HELLO"' | jq 'ascii_downcase'
echo '"hello world"' | jq 'split(" ")'
echo '["hello","world"]' | jq 'join(" ")'
echo '"hello"' | jq 'length'
echo '"  hello  "' | jq 'ltrimstr("  ")'
echo '"  hello  "' | jq 'rtrimstr("  ")'
echo '"hello world"' | jq 'split(" ") | .[0]'
echo '"abc123"' | jq '[scan("[0-9]+")]'
echo '"hello"' | jq 'test("he.*")'
echo '"hello world"' | jq 'gsub("world"; "jq")'
```

## fa-calculator Math & Comparison

```bash
echo '10' | jq '. + 5'
echo '10' | jq '. * 3'
echo '10' | jq '. / 3'
echo '10' | jq '. % 3'
echo '2' | jq 'pow(.; 10)'
echo '16' | jq 'sqrt'
echo '-5' | jq 'fabs'
echo '3.14159' | jq 'floor'
echo '3.14159' | jq 'ceil'
echo 'null' | jq 'isnan'
echo '[1,2,3,4,5]' | jq 'add'
echo '[1,2,3,4,5]' | jq 'min'
echo '[1,2,3,4,5]' | jq 'max'
```

## fa-box Variables

```bash
echo '{"items":[1,2,3],"tax":0.1}' | \
  jq '.items as $items | $items | map(. * (1 + .tax))'
echo '{"rate":0.2}' | jq '.rate as $r | {rate: $r, display: ($r * 100 | tostring + "%")}'
echo '5' | jq '. as $x | $x * $x'
echo '{"a":1}' | jq '.a as $val | {$val, double: ($val * 2)}'
echo '[1,2,3]' | jq '[.[] | . as $x | {orig: $x, sq: ($x * $x)}]'
echo '{"env":"prod"}' | jq --arg env "staging" '{given: $env, from_file: .env}'
```

## fa-code-branch Conditionals (if-then-else)

```bash
echo '5' | jq 'if . > 10 then "big" elif . > 3 then "medium" else "small" end'
echo '{"type":"user"}' | jq 'if .type == "admin" then "full access" else "limited" end'
echo '[1,"a",null,true,{"x":1}]' | \
  jq '.[] | if type == "number" then "num: \(.)" elif type == "string" then "str: \(.)" else "other" end'
echo 'null' | jq 'if . then "truthy" else "falsy" end'
echo '""' | jq 'if . == "" then "empty" else . end'
```

## fa-arrows-down-to-arc Map & Reduce

```bash
echo '[1,2,3,4]' | jq 'map(. * 2)'
echo '[1,2,3,4]' | jq 'map(select(. > 2))'
echo '["hello","world"]' | jq 'map(length)'
echo '[1,2,3]' | jq 'reduce .[] as $x (0; . + $x)'
echo '[1,2,3,4,5]' | jq 'reduce .[] as $x (1; . * $x)'
echo '[[1,2],[3,4],[5]]' | jq 'map(add)'
echo '[1,2,3]' | jq '[.[] | . * .]'
echo '[{"k":"a","v":1},{"k":"b","v":2}]' | \
  jq 'map({(.k): .v}) | add'
```

## fa-right-to-bracket Input/Output

```bash
echo '{"a":1}' | jq '.'
echo '{"a":1}' | jq -c '.'
echo '{"a":1}' | jq -r '.a'
echo '{"a":1}' | jq -j '"\(.a) "'
echo '{"a":1}' | jq -e '.missing'
echo '[1,2,3]' | jq '.[]'
echo '{"a":1}' | jq -C '.'
echo '{"a":1}' | jq -S '.'
echo '{"b":2,"a":1}' | jq 'keys'
echo '{"a":1,"b":2}' | jq 'values'
```

## fa-layer-group Slurp & Raw

```bash
echo -e '{"a":1}\n{"b":2}' | jq -s '.'
echo -e '1\n2\n3' | jq -s 'add'
echo -e 'hello\nworld' | jq -R '.'
echo -e 'a,b\n1,2\n3,4' | jq -R 'split(",")' | jq -s '.'
echo '"hello"' | jq -r '.'
echo '42' | jq -r '.'
echo -e 'line1\nline2' | jq -Rn '[inputs]'
echo -e 'a  b\nc  d' | jq -R 'split("  ") | {key:.[0], val:.[1]}' | jq -s '.'
```

## fa-sitemap Advanced (recurse/paths)

```bash
echo '{"a":{"b":1},"c":[2,{"d":3}]}' | jq '[paths]'
echo '{"a":{"b":1},"c":2}' | jq '[paths(type == "number")]'
echo '{"a":{"b":{"c":1}}}' | jq 'recurse(.[]?; . != null)'
echo '{"a":{"b":1},"c":{"d":2}}' | jq 'flatten'
echo '{"a":1,"b":{"c":2}}' | jq '[leaf_paths as $p | {($p | join(".")): getpath($p)}] | add'
echo '{"a":{"b":1},"x":2}' | jq 'delpaths([["a","b"]])'
echo '{"a":1}' | jq 'path(.a)'
echo '{"a":1,"b":2}' | jq 'with_entries(select(.value > 1))'
```

## fa-code-merge Combine Files

```bash
jq -s '.[0] * .[1]' file1.json file2.json
jq -s '.[0] + .[1]' file1.json file2.json
jq -s 'map(select(.active))' *.json
jq -s 'group_by(.name) | map({name: .[0].name, count: length})' *.json
jq -n --argfile a file1.json --argfile b file2.json '$a * $b'
jq -s 'add' part1.json part2.json part3.json
paste -d'\0' <(jq -c '.' left.json) <(jq -c '.' right.json) | jq -s '.'
```

## fa-lightbulb Practical One-liners

```bash
# Pretty-print JSON file
jq '.' data.json

# Extract field from array of objects
jq '.[].email' users.json

# Filter and reshape
jq '[.[] | select(.active) | {name, email}]' users.json

# Count by field value
jq 'group_by(.status) | map({status: .[0].status, count: length})' items.json

# CSV from JSON
jq -r '["name","age"], (.[] | [.name, .age]) | @csv' users.json

# Sum a field
jq '[.[].price] | add' orders.json

# JSON to TSV
jq -r 'keys_unsorted | @tsv, (.[] | [.[]] | @tsv)' data.json

# Extract nested value
jq '.data.results[] | select(.score > 90) | .name' api.json
```
