---
title: Go 并发
icon: fa-shuffle
primary: "#00ADD8"
lang: go
locale: zhs
---

## fa-bolt Goroutine 协程

```go
go func(msg string) {
    fmt.Println(msg)
}("hello")

go doWork()

func doWork() {
    fmt.Println("working...")
}
```

## fa-arrows-left-right Channel 通道

```go
ch := make(chan int)
ch <- 42
val := <-ch

ch := make(chan int, 3)  // 带缓冲通道
ch <- 1
ch <- 2
fmt.Println(len(ch))
```

## fa-shuffle Select 多路复用

```go
select {
case msg := <-ch1:
    fmt.Println("received:", msg)
case ch2 <- 42:
    fmt.Println("sent")
case <-time.After(time.Second):
    fmt.Println("timeout")
default:
    fmt.Println("no ready channel")
}
```

## fa-users WaitGroup 等待组

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Println("worker", id)
    }(i)
}
wg.Wait()
```

## fa-lock Mutex 互斥锁与读写锁

```go
var mu sync.Mutex
mu.Lock()
defer mu.Unlock()
counter++

var rw sync.RWMutex
rw.RLock()          // 多个读操作可并发
defer rw.RUnlock()
_ = data

rw.Lock()           // 写操作独占
defer rw.Unlock()
data = newData
```

## fa-map sync.Map 并发安全映射

```go
var m sync.Map

m.Store("key", "value")
val, ok := m.Load("key")
m.Delete("key")

m.Range(func(key, value any) bool {
    fmt.Println(key, value)
    return true
})
```

## fa-battery-half sync.Pool 对象池

```go
pool := &sync.Pool{
    New: func() any {
        return bytes.NewBuffer(nil)
    },
}

buf := pool.Get().(*bytes.Buffer)  // 获取对象
buf.Reset()
buf.WriteString("hello")
pool.Put(buf)                      // 归还对象
```

## fa-gauge Atomic 原子操作

```go
var counter int64

atomic.AddInt64(&counter, 1)
val := atomic.LoadInt64(&counter)
atomic.StoreInt64(&counter, 0)

atomic.CompareAndSwapInt64(&counter, 0, 1)  // CAS

var v atomic.Value
v.Store("hello")
s := v.Load().(string)
```

## fa-rotate Context 上下文控制

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

ctx, cancel := context.WithCancel(context.Background())
defer cancel()

ctx = context.WithValue(ctx, "key", "value")
val := ctx.Value("key")

select {
case <-ctx.Done():
    fmt.Println("cancelled:", ctx.Err())
case result := <-ch:
    fmt.Println(result)
}
```

## fa-layer-group Worker Pool 工作池

```go
jobs := make(chan int, 100)
results := make(chan int, 100)

for w := 0; w < 3; w++ {
    go func() {
        for j := range jobs {
            results <- j * 2
        }
    }()
}

for j := 0; j < 10; j++ {
    jobs <- j
}
close(jobs)

for r := 0; r < 10; r++ {
    fmt.Println(<-results)
}
```

## fa-timeline Fan-Out / Fan-In 扇出扇入

```go
func fanOut(input <-chan int, n int) []<-chan int {
    channels := make([]<-chan int, n)
    for i := 0; i < n; i++ {
        channels[i] = process(input)
    }
    return channels
}

func fanIn(channels ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for v := range c {
                out <- v
            }
        }(ch)
    }
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}
```

## fa-door-open Channel 常用模式

```go
// 用空结构体做信号
done := make(chan struct{})

go func() {
    defer close(done)
    doWork()
}()
<-done

// 一次性信号
once := make(chan struct{})
close(once)
<-once

// 退出通知
quit := make(chan bool)
go func() {
    for {
        select {
        case <-quit:
            return
        default:
            doWork()
        }
    }
}()
quit <- true
```

## fa-repeat errgroup 错误组

```go
g, ctx := errgroup.WithContext(context.Background())

g.Go(func() error {
    return fetchURL(ctx, "https://example.com/a")
})
g.Go(func() error {
    return fetchURL(ctx, "https://example.com/b")
})

if err := g.Wait(); err != nil {
    log.Fatal(err)
}
```

## fa-stopwatch Rate Limiting 速率限制

```go
limiter := time.Tick(200 * time.Millisecond)

for _, req := range requests {
    <-limiter
    go process(req)
}

// 使用 golang.org/x/time/rate
rateLimiter := rate.NewLimiter(rate.Every(100*time.Millisecond), 3)
for _, req := range requests {
    rateLimiter.Wait(context.Background())
    go process(req)
}
```

## fa-tower-broadcast 广播与分流

```go
func broadcast(source <-chan int, receivers ...chan int) {
    for v := range source {
        for _, r := range receivers {
            r <- v
        }
    }
    for _, r := range receivers {
        close(r)
    }
}

func tee(ch <-chan int) (_, _ <-chan int) {
    out1, out2 := make(chan int), make(chan int)
    go func() {
        defer close(out1)
        defer close(out2)
        for v := range ch {
            out1 <- v
            out2 <- v
        }
    }()
    return out1, out2
}
```
