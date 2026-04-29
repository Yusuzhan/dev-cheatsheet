---
title: Redis Streams
icon: fa-water
primary: "#DC382D"
lang: bash
---

## fa-plus XADD

```bash
XADD mystream * name Alice age 30
XADD mystream * sensor-id 1 temperature 23.5
XADD mystream 1640995200000-0 event login user 42
```

Use `*` to auto-generate the entry ID (timestamp-sequence format).

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

`$` means only new entries. `BLOCK` makes it wait up to ms for new data.

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

`-` is the oldest ID, `+` is the newest. `(` denotes an exclusive bound.

```bash
XREVRANGE mystream + - COUNT 3
XREVRANGE mystream + 1640995200000-0 COUNT 10
```

## fa-ruler XLEN

```bash
XLEN mystream
```

Returns the number of entries in the stream.

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

`~` enables approximate trimming for performance. `MINID` removes entries older than the given ID.

## fa-users XREADGROUP (Consumer Groups)

```bash
XREADGROUP GROUP mygroup consumer1 COUNT 1 STREAMS mystream >
XREADGROUP GROUP mygroup consumer1 BLOCK 5000 COUNT 10 STREAMS mystream >
XREADGROUP GROUP mygroup consumer1 COUNT 1 STREAMS mystream 0
```

`>` receives only new messages. `0` receives pending (unacknowledged) messages for this consumer.

```bash
XREADGROUP GROUP mygroup consumer1 COUNT 10 STREAMS s1 s2 > >
```

## fa-check XACK

```bash
XACK mystream mygroup 1640995200000-0
XACK mystream mygroup id1 id2 id3
```

Acknowledges successful processing of a message, removing it from the pending entries list.

```bash
XACK mystream mygroup $(XREADGROUP GROUP mygroup consumer1 COUNT 1 STREAMS mystream > | grep -oP '\d+-\d')
```

## fa-layer-group XGROUP

```bash
XGROUP CREATE mystream mygroup $
XGROUP CREATE mystream mygroup 0 MKSTREAM
XGROUP CREATE mystream mygroup 1640995200000-0
```

`$` starts from new messages only. `0` starts from the beginning. `MKSTREAM` creates the stream if missing.

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

Shows pending (delivered but unacknowledged) messages. The summary form returns total count, min/max ID, and consumers.

```bash
XPENDING mystream mygroup 1640995200000-0 + 20 consumer1
```

## fa-hand-holding XCLAIM

```bash
XCLAIM mystream mygroup consumer1 3600000 id1 id2
XCLAIM mystream mygroup consumer1 0 id1
XCLAIM mystream mygroup consumer1 60000 id1 JUSTID
```

Transfers pending messages from one consumer to another. Time is min idle time in ms.

```bash
XAUTOCLAIM mystream mygroup consumer1 3600000 0-0 COUNT 10
```

`XAUTOCLAIM` automatically claims messages idle longer than the threshold.

## fa-circle-info XINFO

```bash
XINFO STREAM mystream
XINFO GROUPS mystream
XINFO CONSUMERS mystream mygroup
```

Returns detailed metadata about the stream, its consumer groups, and consumers within a group.

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

Use `~` (approximate) for performance — Redis removes whole macros nodes instead of exact trimming. Best paired with `XADD`.

```bash
XADD logstream MAXLEN ~ 50000 * level error msg "disk full"
```

## fa-gears Consumer Group Patterns

```bash
XGROUP CREATE orders shipping 0 MKSTREAM
XGROUP CREATE orders billing 0 MKSTREAM
XGROUP CREATE orders notifications 0 MKSTREAM
```

Multiple groups on the same stream: each group receives all messages independently (broadcast).

```bash
XADD orders * type new user 42 total 99.99
```

## fa-lightbulb Practical Patterns

Reliable queue with retry:

```bash
XGROUP CREATE tasks workers 0 MKSTREAM
XREADGROUP GROUP workers worker1 BLOCK 5000 COUNT 1 STREAMS tasks >
XACK tasks workers 1640995200000-0
```

Fan-out / broadcast:

```bash
XGROUP CREATE events service-a $ MKSTREAM
XGROUP CREATE events service-b $ MKSTREAM
XADD events * type user-created data '{"id":1}'
```

Pending recovery:

```bash
XPENDING tasks workers - + 100 worker1
XAUTOCLAIM tasks workers worker2 300000 0-0 COUNT 50
```
