---
title: Redis
icon: fa-database
primary: "#DC382D"
lang: bash
---

## fa-plug Connection & Server

```bash
redis-cli                              # connect to localhost:6379
redis-cli -h 10.0.0.1 -p 6380 -a secret  # connect with auth
redis-cli -u redis://user:pass@host:6379/0  # URI connection
redis-cli -n 2                         # connect to database 2

PING                                    # test connection → PONG
INFO                                    # server info
INFO memory                             # memory stats
INFO replication                        # replication stats
DBSIZE                                  # number of keys in current DB
SELECT 1                                # switch to database 1
FLUSHDB                                 # delete all keys in current DB
FLUSHALL                                # delete all keys in all DBs
```

## fa-key String

```bash
SET name "Alice"                        # set key
GET name                                # get value → "Alice"
SETNX key "value"                       # set only if not exists
SETEX session 3600 "data"               # set with 3600s TTL
MSET k1 "v1" k2 "v2" k3 "v3"           # set multiple keys
MGET k1 k2 k3                           # get multiple values
INCR counter                             # increment (+1)
INCRBY counter 10                        # increment by 10
DECR counter                             # decrement (-1)
INCRBYFLOAT price 2.5                    # increment float
APPEND name " Smith"                     # append to string
STRLEN name                              # string length
GETRANGE name 0 4                        # substring
```

## fa-hashtag Hash

```bash
HSET user:1 name "Alice" age 25 email "a@b.com"   # set fields
HGET user:1 name                        # get single field → "Alice"
HMGET user:1 name age                   # get multiple fields
HGETALL user:1                          # get all fields and values
HDEL user:1 email                       # delete field
HEXISTS user:1 name                     # check field exists → 1
HKEYS user:1                            # all field names
HVALS user:1                            # all field values
HLEN user:1                             # number of fields
HINCRBY user:1 age 1                    # increment field by 1
HSETNX user:1 role "admin"              # set field only if not exists
```

## fa-list List

```bash
LPUSH tasks "task3" "task2" "task1"     # push to head (left)
RPUSH tasks "task4" "task5"             # push to tail (right)
LPOP tasks                              # pop from head
RPOP tasks                              # pop from tail
LLEN tasks                              # list length
LRANGE tasks 0 -1                       # get all elements
LRANGE tasks 0 2                        # get first 3 elements
LINDEX tasks 0                          # get by index
LSET tasks 0 "updated"                  # set by index
LREM tasks 2 "task1"                    # remove 2 occurrences
LTRIM tasks 0 99                        # keep only first 100 elements
BLPOP queue 30                          # blocking pop (timeout 30s)
```

## fa-circle-dot Set

```bash
SADD tags "redis" "db" "cache"          # add members
SMEMBERS tags                           # get all members
SISMEMBER tags "redis"                  # check membership → 1
SCARD tags                              # member count
SREM tags "cache"                       # remove member
SPOP tags                               # remove and return random member
SRANDMEMBER tags 2                      # get 2 random members

SUNION set1 set2                        # union
SINTER set1 set2                        # intersection
SDIFF set1 set2                         # difference
SUNIONSTORE result set1 set2            # union → store in result
```

## fa-arrow-up-1-9 Sorted Set

```bash
ZADD leaderboard 100 "Alice" 95 "Bob" 88 "Charlie"   # add with scores
ZSCORE leaderboard "Alice"              # get score → "100"
ZRANK leaderboard "Alice"               # rank (ascending, 0-based)
ZREVRANK leaderboard "Alice"            # rank (descending)
ZRANGE leaderboard 0 -1                 # ascending by score
ZRANGE leaderboard 0 -1 WITHSCORES     # with scores
ZREVRANGE leaderboard 0 2 WITHSCORES   # top 3 descending
ZINCRBY leaderboard 5 "Bob"             # increment score by 5
ZCARD leaderboard                       # member count
ZCOUNT leaderboard 80 100               # count members with score 80-100
ZRANGEBYSCORE leaderboard 90 +INF      # score >= 90
ZREM leaderboard "Charlie"              # remove member
ZREMRANGEBYRANK leaderboard 0 2         # remove by rank range
```

## fa-clock Key Management

```bash
EXISTS key                              # check if key exists → 1
DEL key1 key2                           # delete keys
UNLINK key1 key2                        # async delete (non-blocking)
TYPE key                                # key data type
RENAME old_key new_key                  # rename key
TTL key                                 # remaining TTL in seconds (-1=none, -2=gone)
PTTL key                                # TTL in milliseconds
EXPIRE key 3600                         # set TTL to 3600 seconds
PERSIST key                             # remove TTL (make persistent)

SCAN 0 COUNT 100                        # iterate keys (cursor-based)
SCAN 0 MATCH user:* COUNT 100           # iterate with pattern
RANDOMKEY                               # return random key
```

## fa-paper-plane Pub/Sub

```bash
SUBSCRIBE channel1 channel2             # subscribe to channels
PSUBSCRIBE news:*                       # subscribe with pattern
UNSUBSCRIBE channel1                    # unsubscribe
PUBLISH channel1 "hello world"          # publish message
PUBSUB CHANNELS                         # list active channels
PUBSUB NUMSUB channel1                  # subscriber count
```

## fa-bolt Transactions & Lua

```bash
MULTI                                   # begin transaction
SET account:A 500
SET account:B 300
EXEC                                    # execute all commands

DISCARD                                 # cancel transaction
WATCH key                               # watch key for changes (optimistic lock)
UNWATCH                                 # unwatch all keys
```

```bash
# Lua script (atomic execution)
EVAL "return redis.call('SET', KEYS[1], ARGV[1])" 1 mykey myvalue
EVAL "return redis.call('GET', KEYS[1])" 1 mykey

# script cache
SCRIPT LOAD "return 1"                  # load and cache script
EVALSHA <sha1> 0                        # execute cached script
```

## fa-hard-drive Persistence

```bash
# RDB snapshot
SAVE                                    # synchronous save (blocks)
BGSAVE                                  # background save
LASTSAVE                                # timestamp of last save
CONFIG GET save                         # RDB schedule rules

# AOF append-only file
CONFIG GET appendonly                   # check AOF status
CONFIG SET appendonly yes               # enable AOF
CONFIG GET appendfsync                  # AOF sync policy
# everysec   - sync every second (recommended)
# always     - sync every write (safest, slowest)
# no         - let OS decide (fastest)

BGREWRITEAOF                            # compact AOF file
```

## fa-sitemap Replication

```bash
INFO replication                        # replication status
ROLE                                    # master or slave

# on replica
REPLICAOF host 6379                     # become replica of host
REPLICAOF NO ONE                        # promote to master

# config for replica
# replicaof 10.0.0.1 6379
# masterauth secret
```

## fa-lightbulb Performance & Debug

```bash
SLOWLOG GET 10                          # last 10 slow queries
SLOWLOG LEN                             # slow log count
CONFIG SET slowlog-log-slower-than 10000  # log queries > 10ms

CLIENT LIST                             # connected clients
CLIENT KILL ADDR 10.0.0.1:12345        # kill client connection
MONITOR                                 # stream all commands (debug only!)

MEMORY USAGE key                        # memory used by key
MEMORY DOCTOR                           # memory analysis tips

# benchmark
redis-benchmark -t set,get -n 100000    # benchmark SET/GET, 100k requests
redis-benchmark -t ping -c 50           # 50 parallel connections
```
