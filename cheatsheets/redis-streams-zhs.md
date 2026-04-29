---
title: Redis Streams
icon: fa-water
primary: "#DC382D"
lang: bash
locale: zhs
---

## fa-plus XADD

```bash
XADD mystream * name Alice age 30
XADD mystream * sensor-id 1 temperature 23.5
XADD mystream 1640995200000-0 event login user 42
```

使用 `*` 自动生成条目 ID（时间戳-序列号格式）。

```bash
XADD mystream MAXLEN 1000 * message hello
XADD mystream MAXLEN ~ 1000 * message hello
```

## fa-book-open XREAD

```bash
XREAD COUNT 10 STREAMS mystream $
XREAD COUNT 5 BLOCK 5000 STREAMS mystream 1640995200000-0
XREAD STREAMS stream1 stream2 0-0 0-0
```

`$` 表示仅读取新条目。`BLOCK` 使命令阻塞等待新数据（毫秒）。

```bash
XREAD COUNT 10 BLOCK 0 STREAMS mystream $
```

## fa-arrows-left-right XRANGE / XREVRANGE

```bash
XRANGE mystream - +
XRANGE mystream 1640995200000-0 1640995260000-0
XRANGE mystream - + COUNT 5
XRANGE mystream 1640995200000-0 (1640995200000-0 COUNT 1
```

`-` 是最旧 ID，`+` 是最新 ID。`(` 表示不包含边界值。

```bash
XREVRANGE mystream + - COUNT 3
XREVRANGE mystream + 1640995200000-0 COUNT 10
```

## fa-ruler XLEN

```bash
XLEN mystream
```

返回流中的条目数量。

```bash
XLEN mystream
XLEN orders
XLEN events
```

## fa-trash XDEL & XTRIM

```bash
XDEL mystream 1640995200000-0
XDEL mystream id1 id2 id3
XTRIM mystream MAXLEN 1000
XTRIM mystream MAXLEN ~ 5000
XTRIM mystream MINID 1640995200000-0
```

`~` 启用近似裁剪以提升性能。`MINID` 移除早于指定 ID 的条目。

## fa-users XREADGROUP (Consumer Groups)

```bash
XREADGROUP GROUP mygroup consumer1 COUNT 1 STREAMS mystream >
XREADGROUP GROUP mygroup consumer1 BLOCK 5000 COUNT 10 STREAMS mystream >
XREADGROUP GROUP mygroup consumer1 COUNT 1 STREAMS mystream 0
```

`>` 仅接收新消息。`0` 接收该消费者的待处理（未确认）消息。

```bash
XREADGROUP GROUP mygroup consumer1 COUNT 10 STREAMS s1 s2 > >
```

## fa-check XACK

```bash
XACK mystream mygroup 1640995200000-0
XACK mystream mygroup id1 id2 id3
```

确认消息处理成功，将其从待处理列表中移除。

```bash
XACK mystream mygroup $(XREADGROUP GROUP mygroup consumer1 COUNT 1 STREAMS mystream > | grep -oP '\d+-\d')
```

## fa-layer-group XGROUP

```bash
XGROUP CREATE mystream mygroup $
XGROUP CREATE mystream mygroup 0 MKSTREAM
XGROUP CREATE mystream mygroup 1640995200000-0
```

`$` 仅从新消息开始消费。`0` 从头开始。`MKSTREAM` 在流不存在时自动创建。

```bash
XGROUP DESTROY mystream mygroup
XGROUP DELCONSUMER mystream mygroup consumer1
XGROUP SETID mystream mygroup $
```

## fa-clock XPENDING

```bash
XPENDING mystream mygroup
XPENDING mystream mygroup - + 10
XPENDING mystream mygroup - + 10 consumer1
```

查看已投递但未确认的待处理消息。摘要模式返回总数、最小/最大 ID 和消费者列表。

```bash
XPENDING mystream mygroup 1640995200000-0 + 20 consumer1
```

## fa-hand-holding XCLAIM

```bash
XCLAIM mystream mygroup consumer1 3600000 id1 id2
XCLAIM mystream mygroup consumer1 0 id1
XCLAIM mystream mygroup consumer1 60000 id1 JUSTID
```

将待处理消息从一个消费者转移到另一个消费者。时间为最小空闲时间（毫秒）。

```bash
XAUTOCLAIM mystream mygroup consumer1 3600000 0-0 COUNT 10
```

`XAUTOCLAIM` 自动认领空闲时间超过阈值的消息。

## fa-circle-info XINFO

```bash
XINFO STREAM mystream
XINFO GROUPS mystream
XINFO CONSUMERS mystream mygroup
```

返回流、消费者组及组内消费者的详细元数据。

```bash
XINFO STREAM mystream FULL
XINFO STREAM mystream FULL COUNT 5
```

## fa-scissors MAXLEN Strategy

```bash
XADD mystream MAXLEN 10000 * field value
XADD mystream MAXLEN ~ 10000 * field value
XTRIM mystream MAXLEN = 10000
XTRIM mystream MAXLEN ~ 10000
```

使用 `~`（近似裁剪）可提升性能——Redis 按宏节点整块删除而非精确裁剪。建议配合 `XADD` 使用。

```bash
XADD logstream MAXLEN ~ 50000 * level error msg "disk full"
```

## fa-gears Consumer Group Patterns

```bash
XGROUP CREATE orders shipping 0 MKSTREAM
XGROUP CREATE orders billing 0 MKSTREAM
XGROUP CREATE orders notifications 0 MKSTREAM
```

同一流上创建多个消费者组：每个组独立接收所有消息（广播模式）。

```bash
XADD orders * type new user 42 total 99.99
```

## fa-lightbulb Practical Patterns

可靠队列与重试：

```bash
XGROUP CREATE tasks workers 0 MKSTREAM
XREADGROUP GROUP workers worker1 BLOCK 5000 COUNT 1 STREAMS tasks >
XACK tasks workers 1640995200000-0
```

扇出 / 广播：

```bash
XGROUP CREATE events service-a $ MKSTREAM
XGROUP CREATE events service-b $ MKSTREAM
XADD events * type user-created data '{"id":1}'
```

待处理消息恢复：

```bash
XPENDING tasks workers - + 100 worker1
XAUTOCLAIM tasks workers worker2 300000 0-0 COUNT 50
```
