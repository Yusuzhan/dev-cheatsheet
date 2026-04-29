---
title: Go Concurrency
icon: fa-shuffle
primary: "#00ADD8"
lang: go
---

## fa-bolt Goroutines

```go
go func(msg string) {
    fmt.Println(msg)
}("hello")

go doWork()

func doWork() {
    fmt.Println("working...")
}
```

## fa-arrows-left-right Channels

```go
ch := make(chan int)
ch <- 42
val := <-ch

ch := make(chan int, 3)
ch <- 1
ch <- 2
fmt.Println(len(ch))
```

## fa-shuffle Select

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

## fa-users WaitGroup

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

## fa-lock Mutex & RWMutex

```go
var mu sync.Mutex
mu.Lock()
defer mu.Unlock()
counter++

var rw sync.RWMutex
rw.RLock()
defer rw.RUnlock()
_ = data

rw.Lock()
defer rw.Unlock()
data = newData
```

## fa-map sync.Map

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

## fa-battery-half sync.Pool

```go
pool := &sync.Pool{
    New: func() any {
        return bytes.NewBuffer(nil)
    },
}

buf := pool.Get().(*bytes.Buffer)
buf.Reset()
buf.WriteString("hello")
pool.Put(buf)
```

## fa-gauge Atomic Operations

```go
var counter int64

atomic.AddInt64(&counter, 1)
val := atomic.LoadInt64(&counter)
atomic.StoreInt64(&counter, 0)

atomic.CompareAndSwapInt64(&counter, 0, 1)

var v atomic.Value
v.Store("hello")
s := v.Load().(string)
```

## fa-rotate Context

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

## fa-layer-group Worker Pool

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

## fa-timeline Fan-Out / Fan-In

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

## fa-door-open Channel Patterns

```go
done := make(chan struct{})

go func() {
    defer close(done)
    doWork()
}()
<-done

once := make(chan struct{})
close(once)
<-once

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

## fa-repeat errgroup

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

## fa-stopwatch Rate Limiting

```go
limiter := time.Tick(200 * time.Millisecond)

for _, req := range requests {
    <-limiter
    go process(req)
}

rateLimiter := rate.NewLimiter(rate.Every(100*time.Millisecond), 3)
for _, req := range requests {
    rateLimiter.Wait(context.Background())
    go process(req)
}
```

## fa-tower-broadcast Broadcast & Tee

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
