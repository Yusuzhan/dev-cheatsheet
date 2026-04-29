---
title: go-redis
icon: fa-database
primary: "#DC382D"
lang: go
---

## fa-plug Connect & Client

```go
import "github.com/redis/go-redis/v9"

rdb := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
})

err := rdb.Ping(ctx).Err()
if err != nil {
    log.Fatal(err)
}
defer rdb.Close()
```

## fa-font Strings

```go
rdb.Set(ctx, "key", "value", 0)
rdb.Set(ctx, "key", "value", 10*time.Minute)

val, err := rdb.Get(ctx, "key").Result()
err := rdb.SetNX(ctx, "key", "value", time.Minute).Err()

rdb.MSet(ctx, "k1", "v1", "k2", "v2")
vals, err := rdb.MGet(ctx, "k1", "k2").Result()

rdb.Incr(ctx, "counter")
rdb.IncrBy(ctx, "counter", 5)
rdb.Decr(ctx, "counter")

rdb.Append(ctx, "key", "suffix")
strlen, _ := rdb.StrLen(ctx, "key").Result()
```

## fa-layer-group Hashes

```go
rdb.HSet(ctx, "user:1", "name", "Alice", "age", 30)
rdb.HSet(ctx, "user:1", map[string]interface{}{"email": "a@b.com", "city": "NYC"})

val, _ := rdb.HGet(ctx, "user:1", "name").Result()
all, _ := rdb.HGetAll(ctx, "user:1").Result()
exists, _ := rdb.HExists(ctx, "user:1", "name").Result()

rdb.HDel(ctx, "user:1", "age")
rdb.HIncrBy(ctx, "user:1", "age", 1)

keys, _ := rdb.HKeys(ctx, "user:1").Result()
vals, _ := rdb.HVals(ctx, "user:1").Result()
cnt, _ := rdb.HLen(ctx, "user:1").Result()
```

## fa-list Lists

```go
rdb.LPush(ctx, "queue", "a", "b", "c")
rdb.RPush(ctx, "queue", "x", "y")

val, _ := rdb.LPop(ctx, "queue").Result()
val, _ := rdb.RPop(ctx, "queue").Result()

val, _ := rdb.LIndex(ctx, "queue", 0).Result()
vals, _ := rdb.LRange(ctx, "queue", 0, -1).Result()
rdb.LTrim(ctx, "queue", 0, 99)

len, _ := rdb.LLen(ctx, "queue").Result()
rdb.LSet(ctx, "queue", 0, "updated")
rdb.LRem(ctx, "queue", 1, "value")
```

## fa-circle-nodes Sets

```go
rdb.SAdd(ctx, "tags", "go", "redis", "db")

members, _ := rdb.SMembers(ctx, "tags").Result()
ok, _ := rdb.SIsMember(ctx, "tags", "go").Result()

rdb.SRem(ctx, "tags", "db")
cnt, _ := rdb.SCard(ctx, "tags").Result()

diff, _ := rdb.SDiff(ctx, "set1", "set2").Result()
inter, _ := rdb.SInter(ctx, "set1", "set2").Result()
union, _ := rdb.SUnion(ctx, "set1", "set2").Result()

member, _ := rdb.SPop(ctx, "tags").Result()
members, _ := rdb.SRandMemberN(ctx, "tags", 3).Result()
```

## fa-arrow-up-wide-short Sorted Sets

```go
rdb.ZAdd(ctx, "leaderboard", redis.Z{Score: 100, Member: "alice"})
rdb.ZAdd(ctx, "leaderboard", redis.Z{Score: 85, Member: "bob"})

rdb.ZIncrBy(ctx, "leaderboard", 10, "bob")

vals, _ := rdb.ZRangeByScore(ctx, "leaderboard", &redis.ZRangeBy{
    Min: "0", Max: "100", Offset: 0, Count: 10,
}).Result()

top, _ := rdb.ZRevRangeWithScores(ctx, "leaderboard", 0, 9).Result()
rank, _ := rdb.ZRank(ctx, "leaderboard", "alice").Result()
score, _ := rdb.ZScore(ctx, "leaderboard", "alice").Result()

rdb.ZRem(ctx, "leaderboard", "alice")
rdb.ZRemRangeByRank(ctx, "leaderboard", 0, 2)
cnt, _ := rdb.ZCard(ctx, "leaderboard").Result()
```

## fa-key Keys & Expiry

```go
rdb.Del(ctx, "key1", "key2")
rdb.Unlink(ctx, "key1", "key2")

rdb.Expire(ctx, "key", 10*time.Minute)
rdb.ExpireAt(ctx, "key", time.Now().Add(time.Hour))
rdb.Persist(ctx, "key")

ttl, _ := rdb.TTL(ctx, "key").Result()
exists, _ := rdb.Exists(ctx, "key").Result()

rdb.Rename(ctx, "old", "new")
rdb.Copy(ctx, "src", "dst", 0, false)

keys, _ := rdb.Keys(ctx, "user:*").Result()
```

## fa-grip-lines Pipelines & TxPipeline

```go
pipe := rdb.Pipeline()
setCmd := pipe.Set(ctx, "key", "value", 0)
getCmd := pipe.Get(ctx, "key")
_, err := pipe.Exec(ctx)

val := getCmd.Val()

pipe := rdb.TxPipeline()
pipe.Set(ctx, "acc:a", 100, 0)
pipe.Set(ctx, "acc:b", 200, 0)
_, err := pipe.Exec(ctx)
```

## fa-rotate Transactions (WATCH/MULTI)

```go
err := rdb.Watch(ctx, func(tx *redis.Tx) error {
    val, err := tx.Get(ctx, "key").Result()
    if err != nil && err != redis.Nil {
        return err
    }
    _, err = tx.TxPipeline(func(pipe redis.Pipeliner) error {
        pipe.Set(ctx, "key", "new_"+val, 0)
        return nil
    })
    return err
}, "key")
```

## fa-tower-broadcast Pub/Sub

```go
sub := rdb.Subscribe(ctx, "channel1", "channel2")
defer sub.Close()

ch := sub.Channel()
for msg := range ch {
    fmt.Println(msg.Channel, msg.Payload)
}

rdb.Publish(ctx, "channel1", "hello")

msg, err := sub.ReceiveMessage(ctx)
```

## fa-water Streams (XADD/XREAD)

```go
id, _ := rdb.XAdd(ctx, &redis.XAddArgs{
    Stream: "mystream",
    Values: map[string]interface{}{"name": "Alice", "action": "login"},
}).Result()

entries, _ := rdb.XRead(ctx, &redis.XReadArgs{
    Streams: []string{"mystream", "0"},
    Count:   10,
    Block:   5 * time.Second,
}).Result()

rdb.XGroupCreate(ctx, "mystream", "mygroup", "0")

entries, _ := rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
    Group:    "mygroup",
    Consumer: "c1",
    Streams:  []string{"mystream", ">"},
    Count:    10,
    Block:    0,
}).Result()

rdb.XAck(ctx, "mystream", "mygroup", "123456-0")
rdb.XTrim(ctx, "mystream", 0, 1000)
```

## fa-magnifying-glass Scan & Iterator

```go
var keys []string
var cursor uint64
for {
    var batch []string
    batch, cursor, err = rdb.Scan(ctx, cursor, "user:*", 100).Result()
    keys = append(keys, batch...)
    if cursor == 0 {
        break
    }
}

iter := rdb.Scan(ctx, 0, "user:*", 0).Iterator()
for iter.Next(ctx) {
    fmt.Println(iter.Val())
}
if err := iter.Err(); err != nil {
    log.Fatal(err)
}
```

## fa-scroll Lua Scripts

```go
var incrBy = redis.NewScript(`
local val = redis.call('GET', KEYS[1])
if not val then return 0 end
return tonumber(val) + tonumber(ARGV[1])
`)

result, err := incrBy.Run(ctx, rdb, []string{"counter"}, 5).Int()

var limitScript = redis.NewScript(`
local count = redis.call('INCR', KEYS[1])
if count == 1 then
    redis.call('EXPIRE', KEYS[1], ARGV[1])
end
if count > tonumber(ARGV[2]) then
    return 0
end
return 1
`)

ok, _ := limitScript.Run(ctx, rdb, []string{"rate:ip:127.0.0.1"}, 60, 100).Int()
```

## fa-clock Context & Timeouts

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
val, err := rdb.Get(ctx, "key").Result()

rdb := redis.NewClient(&redis.Options{
    Addr:        "localhost:6379",
    DialTimeout: 5 * time.Second,
    ReadTimeout: 3 * time.Second,
    WriteTimeout: 3 * time.Second,
    PoolTimeout: 4 * time.Second,
})

_, err := rdb.Ping(ctx).Result()
if err == context.DeadlineExceeded {
    log.Println("timeout")
}
```

## fa-sitemap Sentinel & Cluster

```go
rdb := redis.NewFailoverClient(&redis.FailoverOptions{
    MasterName:    "mymaster",
    SentinelAddrs: []string{"127.0.0.1:26379", "127.0.0.1:26380"},
    Password:      "",
    DB:            0,
})

cluster := redis.NewClusterClient(&redis.ClusterOptions{
    Addrs:     []string{"10.0.0.1:6379", "10.0.0.2:6379", "10.0.0.3:6379"},
    Password:  "",
    PoolSize:  10,
})
defer cluster.Close()

cluster.Ping(ctx)
```
