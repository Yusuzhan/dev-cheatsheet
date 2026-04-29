---
title: redigo
icon: fa-database
primary: "#DC382D"
lang: go
---

## fa-plug Connect & Pool

```go
import "github.com/gomodule/redigo/redis"

pool := &redis.Pool{
    MaxIdle:     10,
    MaxActive:   100,
    IdleTimeout: 5 * time.Minute,
    Dial: func() (redis.Conn, error) {
        return redis.Dial("tcp", "localhost:6379")
    },
}
defer pool.Close()

conn := pool.Get()
defer conn.Close()

conn.Do("PING")
```

## fa-crosshairs Basic GET/SET

```go
conn.Do("SET", "key", "value")
conn.Do("SET", "key", "value", "EX", 600)

val, err := redis.String(conn.Do("GET", "key"))

ok, err := redis.Bool(conn.Do("SETNX", "key", "value"))

conn.Do("DEL", "key")
exists, _ := redis.Bool(conn.Do("EXISTS", "key"))
```

## fa-font Strings

```go
conn.Do("MSET", "k1", "v1", "k2", "v2")
vals, _ := redis.Strings(conn.Do("MGET", "k1", "k2"))

conn.Do("INCR", "counter")
conn.Do("INCRBY", "counter", 10)
conn.Do("DECR", "counter")

n, _ := redis.Int64(conn.Do("INCRBY", "hits", 1))
conn.Do("APPEND", "key", "suffix")
strlen, _ := redis.Int64(conn.Do("STRLEN", "key"))
```

## fa-layer-group Hashes

```go
conn.Do("HSET", "user:1", "name", "Alice", "age", 30)

val, _ := redis.String(conn.Do("HGET", "user:1", "name"))
all, _ := redis.StringMap(conn.Do("HGETALL", "user:1"))
exists, _ := redis.Bool(conn.Do("HEXISTS", "user:1", "name"))

conn.Do("HDEL", "user:1", "age")
conn.Do("HINCRBY", "user:1", "age", 1)

cnt, _ := redis.Int64(conn.Do("HLEN", "user:1"))
```

## fa-list Lists

```go
conn.Do("LPUSH", "queue", "a", "b", "c")
conn.Do("RPUSH", "queue", "x", "y")

val, _ := redis.String(conn.Do("LPOP", "queue"))
val, _ := redis.String(conn.Do("RPOP", "queue"))

vals, _ := redis.Strings(conn.Do("LRANGE", "queue", 0, -1))
conn.Do("LTRIM", "queue", 0, 99)

len, _ := redis.Int64(conn.Do("LLEN", "queue"))
conn.Do("LSET", "queue", 0, "updated")
conn.Do("LREM", "queue", 1, "value")
```

## fa-circle-nodes Sets

```go
conn.Do("SADD", "tags", "go", "redis", "db")

members, _ := redis.Strings(conn.Do("SMEMBERS", "tags"))
ok, _ := redis.Bool(conn.Do("SISMEMBER", "tags", "go"))

conn.Do("SREM", "tags", "db")
cnt, _ := redis.Int64(conn.Do("SCARD", "tags"))

diff, _ := redis.Strings(conn.Do("SDIFF", "set1", "set2"))
inter, _ := redis.Strings(conn.Do("SINTER", "set1", "set2"))
union, _ := redis.Strings(conn.Do("SUNION", "set1", "set2"))

member, _ := redis.String(conn.Do("SPOP", "tags"))
```

## fa-arrow-up-wide-short Sorted Sets

```go
conn.Do("ZADD", "board", 100, "alice", 85, "bob")
conn.Do("ZINCRBY", "board", 10, "bob")

vals, _ := redis.Strings(conn.Do("ZRANGE", "board", 0, -1, "WITHSCORES"))
top, _ := redis.StringMap(conn.Do("ZREVRANGE", "board", 0, 9, "WITHSCORES"))

rank, _ := redis.Int64(conn.Do("ZRANK", "board", "alice"))
score, _ := redis.Float64(conn.Do("ZSCORE", "board", "alice"))

conn.Do("ZREM", "board", "alice")
conn.Do("ZREMRANGEBYRANK", "board", 0, 2)
cnt, _ := redis.Int64(conn.Do("ZCARD", "board"))

cnt, _ = redis.Int64(conn.Do("ZCOUNT", "board", "80", "100"))
```

## fa-key Keys & Expiry

```go
conn.Do("DEL", "key1", "key2")
conn.Do("UNLINK", "key1", "key2")

conn.Do("EXPIRE", "key", 600)
conn.Do("EXPIREAT", "key", time.Now().Add(time.Hour).Unix())
conn.Do("PERSIST", "key")

ttl, _ := redis.Int64(conn.Do("TTL", "key"))
exists, _ := redis.Int64(conn.Do("EXISTS", "key"))

conn.Do("RENAME", "old", "new")
keys, _ := redis.Strings(conn.Do("KEYS", "user:*"))
```

## fa-rotate Transactions

```go
conn.Send("MULTI")
conn.Send("SET", "acc:a", 100)
conn.Send("SET", "acc:b", 200)
conn.Send("EXEC")
_, err := conn.Do("EXEC")

conn.Send("WATCH", "key")
val, _ := redis.String(conn.Do("GET", "key"))
conn.Send("MULTI")
conn.Send("SET", "key", "new_"+val)
_, err := conn.Do("EXEC")
if err != nil {
    conn.Do("UNWATCH")
}
```

## fa-grip-lines Pipelines

```go
conn.Send("SET", "k1", "v1")
conn.Send("SET", "k2", "v2")
conn.Send("GET", "k1")
conn.Send("GET", "k2")
conn.Flush()

r1, _ := conn.Receive()
r2, _ := conn.Receive()
r3, _ := redis.String(conn.Receive())
r4, _ := redis.String(conn.Receive())
```

## fa-tower-broadcast Pub/Sub

```go
psc := redis.PubSubConn{Conn: pool.Get()}
defer psc.Close()

psc.Subscribe("channel1")

go func() {
    for {
        switch v := psc.Receive().(type) {
        case redis.Message:
            fmt.Printf("%s: %s\n", v.Channel, v.Data)
        case redis.Subscription:
            fmt.Printf("%s %s %d\n", v.Kind, v.Channel, v.Count)
        case error:
            return
        }
    }
}()

conn.Do("PUBLISH", "channel1", "hello")
```

## fa-magnifying-glass Scan Helper

```go
import "github.com/gomodule/redigo/redis"

var keys []string
var cursor uint64
for {
    values, err := redis.Values(conn.Do("SCAN", cursor, "MATCH", "user:*", "COUNT", 100))
    if err != nil {
        log.Fatal(err)
    }
    cursor, _ = redis.Uint64(values[0], nil)
    batch, _ := redis.Strings(values[1], nil)
    keys = append(keys, batch...)
    if cursor == 0 {
        break
    }
}

hkeys, _ := redis.Strings(conn.Do("HKEYS", "user:1"))
```

## fa-scroll Lua Scripts

```go
script := redis.NewScript(1, `
local val = redis.call('GET', KEYS[1])
if not val then return 0 end
return tonumber(val) + tonumber(ARGV[1])
`)

result, err := redis.Int64(script.Do(conn, "counter", 5))

limitScript := redis.NewScript(1, `
local count = redis.call('INCR', KEYS[1])
if count == 1 then
    redis.call('EXPIRE', KEYS[1], ARGV[1])
end
if count > tonumber(ARGV[2]) then return 0 end
return 1
`)

ok, _ := redis.Int64(limitScript.Do(conn, "rate:ip:127.0.0.1", 60, 100))
```

## fa-sliders Connection Pool Config

```go
pool := &redis.Pool{
    MaxIdle:     10,
    MaxActive:   100,
    IdleTimeout: 5 * time.Minute,
    Wait:        true,
    MaxConnLifetime: 30 * time.Minute,
    Dial: func() (redis.Conn, error) {
        return redis.Dial("tcp", "localhost:6379",
            redis.DialPassword("secret"),
            redis.DialDatabase(1),
            redis.DialConnectTimeout(5*time.Second),
            redis.DialReadTimeout(3*time.Second),
            redis.DialWriteTimeout(3*time.Second),
        )
    },
    TestOnBorrow: func(c redis.Conn, t time.Time) error {
        _, err := c.Do("PING")
        return err
    },
}
```

## fa-clock Context Support

```go
import "github.com/gomodule/redigo/redis"

conn, err := redis.Dial("tcp", "localhost:6379",
    redis.DialConnectTimeout(5*time.Second),
    redis.DialReadTimeout(3*time.Second),
)

ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

c := conn.(redis.ConnWithTimeout)
_, err = c.DoWithTimeout(5*time.Second, "GET", "key")

reply, err := redis.String(redis.DoContext(conn, ctx, "GET", "key"))
```
