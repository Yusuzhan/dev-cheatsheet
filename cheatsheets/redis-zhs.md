---
title: Redis
icon: fa-database
primary: "#DC382D"
lang: bash
locale: zhs
---

## fa-plug 连接与服务器

```bash
redis-cli                              # 连接 localhost:6379
redis-cli -h 10.0.0.1 -p 6380 -a secret  # 指定地址和密码连接
redis-cli -u redis://user:pass@host:6379/0  # URI 方式连接
redis-cli -n 2                         # 连接到数据库 2

PING                                    # 测试连接 → PONG
INFO                                    # 服务器信息
INFO memory                             # 内存统计
INFO replication                        # 复制状态
DBSIZE                                  # 当前数据库键数量
SELECT 1                                # 切换到数据库 1
FLUSHDB                                 # 清空当前数据库
FLUSHALL                                # 清空所有数据库
```

## fa-key 字符串 (String)

```bash
SET name "Alice"                        # 设置键值
GET name                                # 获取值 → "Alice"
SETNX key "value"                       # 仅在不存在时设置
SETEX session 3600 "data"               # 设置并指定 3600 秒过期
MSET k1 "v1" k2 "v2" k3 "v3"           # 批量设置
MGET k1 k2 k3                           # 批量获取
INCR counter                             # 自增 (+1)
INCRBY counter 10                        # 自增指定值
DECR counter                             # 自减 (-1)
INCRBYFLOAT price 2.5                    # 浮点数自增
APPEND name " Smith"                     # 追加字符串
STRLEN name                              # 字符串长度
GETRANGE name 0 4                        # 截取子串
```

## fa-hashtag 哈希 (Hash)

```bash
HSET user:1 name "Alice" age 25 email "a@b.com"   # 设置多个字段
HGET user:1 name                        # 获取单个字段 → "Alice"
HMGET user:1 name age                   # 获取多个字段
HGETALL user:1                          # 获取所有字段和值
HDEL user:1 email                       # 删除字段
HEXISTS user:1 name                     # 字段是否存在 → 1
HKEYS user:1                            # 所有字段名
HVALS user:1                            # 所有字段值
HLEN user:1                             # 字段数量
HINCRBY user:1 age 1                    # 字段自增
HSETNX user:1 role "admin"              # 仅在字段不存在时设置
```

## fa-list 列表 (List)

```bash
LPUSH tasks "task3" "task2" "task1"     # 从头部插入
RPUSH tasks "task4" "task5"             # 从尾部插入
LPOP tasks                              # 从头部弹出
RPOP tasks                              # 从尾部弹出
LLEN tasks                              # 列表长度
LRANGE tasks 0 -1                       # 获取所有元素
LRANGE tasks 0 2                        # 获取前 3 个元素
LINDEX tasks 0                          # 按索引获取
LSET tasks 0 "updated"                  # 按索引修改
LREM tasks 2 "task1"                    # 移除 2 个指定值
LTRIM tasks 0 99                        # 只保留前 100 个元素
BLPOP queue 30                          # 阻塞式弹出（超时 30 秒）
```

## fa-circle-dot 集合 (Set)

```bash
SADD tags "redis" "db" "cache"          # 添加成员
SMEMBERS tags                           # 获取所有成员
SISMEMBER tags "redis"                  # 检查是否成员 → 1
SCARD tags                              # 成员数量
SREM tags "cache"                       # 移除成员
SPOP tags                               # 随机弹出成员
SRANDMEMBER tags 2                      # 随机获取 2 个成员

SUNION set1 set2                        # 并集
SINTER set1 set2                        # 交集
SDIFF set1 set2                         # 差集
SUNIONSTORE result set1 set2            # 并集存入 result
```

## fa-arrow-up-1-9 有序集合 (Sorted Set)

```bash
ZADD leaderboard 100 "Alice" 95 "Bob" 88 "Charlie"   # 添加带分数的成员
ZSCORE leaderboard "Alice"              # 获取分数 → "100"
ZRANK leaderboard "Alice"               # 排名（升序，从 0 开始）
ZREVRANK leaderboard "Alice"            # 排名（降序）
ZRANGE leaderboard 0 -1                 # 按分数升序
ZRANGE leaderboard 0 -1 WITHSCORES     # 带分数
ZREVRANGE leaderboard 0 2 WITHSCORES   # 前 3 名降序
ZINCRBY leaderboard 5 "Bob"             # 分数加 5
ZCARD leaderboard                       # 成员数量
ZCOUNT leaderboard 80 100               # 分数 80-100 的成员数
ZRANGEBYSCORE leaderboard 90 +INF      # 分数 >= 90
ZREM leaderboard "Charlie"              # 移除成员
ZREMRANGEBYRANK leaderboard 0 2         # 按排名范围移除
```

## fa-clock 键管理

```bash
EXISTS key                              # 键是否存在 → 1
DEL key1 key2                           # 删除键
UNLINK key1 key2                        # 异步删除（非阻塞）
TYPE key                                # 键的数据类型
RENAME old_key new_key                  # 重命名键
TTL key                                 # 剩余过期秒数 (-1=永不过期, -2=已不存在)
PTTL key                                # 剩余过期毫秒数
EXPIRE key 3600                         # 设置 3600 秒后过期
PERSIST key                             # 移除过期时间（持久化）

SCAN 0 COUNT 100                        # 游标式遍历键
SCAN 0 MATCH user:* COUNT 100           # 按模式遍历
RANDOMKEY                               # 随机返回一个键
```

## fa-paper-plane 发布订阅 (Pub/Sub)

```bash
SUBSCRIBE channel1 channel2             # 订阅频道
PSUBSCRIBE news:*                       # 模式订阅
UNSUBSCRIBE channel1                    # 取消订阅
PUBLISH channel1 "hello world"          # 发布消息
PUBSUB CHANNELS                         # 列出活跃频道
PUBSUB NUMSUB channel1                  # 订阅者数量
```

## fa-bolt 事务与 Lua

```bash
MULTI                                   # 开启事务
SET account:A 500
SET account:B 300
EXEC                                    # 执行所有命令

DISCARD                                 # 取消事务
WATCH key                               # 监视键（乐观锁）
UNWATCH                                 # 取消监视
```

```bash
# Lua 脚本（原子执行）
EVAL "return redis.call('SET', KEYS[1], ARGV[1])" 1 mykey myvalue
EVAL "return redis.call('GET', KEYS[1])" 1 mykey

# 脚本缓存
SCRIPT LOAD "return 1"                  # 加载并缓存脚本
EVALSHA <sha1> 0                        # 执行缓存脚本
```

## fa-hard-drive 持久化

```bash
# RDB 快照
SAVE                                    # 同步保存（阻塞）
BGSAVE                                  # 后台保存
LASTSAVE                                # 上次保存时间戳
CONFIG GET save                         # RDB 保存策略

# AOF 追加文件
CONFIG GET appendonly                   # 查看 AOF 状态
CONFIG SET appendonly yes               # 开启 AOF
CONFIG GET appendfsync                  # AOF 同步策略
# everysec   - 每秒同步（推荐）
# always     - 每次写入同步（最安全，最慢）
# no         - 由操作系统决定（最快）

BGREWRITEAOF                            # 压缩 AOF 文件
```

## fa-sitemap 主从复制

```bash
INFO replication                        # 复制状态
ROLE                                    # master 或 slave

# 在从节点执行
REPLICAOF host 6379                     # 成为 host 的从节点
REPLICAOF NO ONE                        # 提升为主节点

# 从节点配置
# replicaof 10.0.0.1 6379
# masterauth secret
```

## fa-lightbulb 性能与调试

```bash
SLOWLOG GET 10                          # 最近 10 条慢查询
SLOWLOG LEN                             # 慢查询数量
CONFIG SET slowlog-log-slower-than 10000  # 记录超过 10ms 的查询

CLIENT LIST                             # 已连接的客户端
CLIENT KILL ADDR 10.0.0.1:12345        # 断开客户端连接
MONITOR                                 # 实时监控所有命令（仅调试用！）

MEMORY USAGE key                        # 键占用的内存
MEMORY DOCTOR                           # 内存分析建议

# 基准测试
redis-benchmark -t set,get -n 100000    # 测试 SET/GET，10 万次请求
redis-benchmark -t ping -c 50           # 50 个并发连接
```
