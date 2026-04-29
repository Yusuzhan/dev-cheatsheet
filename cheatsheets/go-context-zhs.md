---
title: Go Context
icon: fa-rotate
primary: "#00ADD8"
lang: go
locale: zhs
---

## fa-circle-dot context.Background / TODO

```go
ctx := context.Background()

func handleRequest(ctx context.Context) {
    if ctx == nil {
        ctx = context.TODO()
    }
}
```

## fa-ban WithCancel

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func() {
    select {
    case <-ctx.Done():
        fmt.Println("cancelled:", ctx.Err())
    case <-time.After(2 * time.Second):
        fmt.Println("done")
    }
}()

cancel()
```

## fa-clock WithTimeout / WithDeadline

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

deadline := time.Now().Add(3 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
```

## fa-tag WithValue

```go
type key string

ctx := context.WithValue(context.Background(), key("userID"), 42)
ctx = context.WithValue(ctx, key("requestID"), "abc-123")

uid := ctx.Value(key("userID")).(int)
rid := ctx.Value(key("requestID")).(string)
```

## fa-arrow-right-from-bracket Done() Channel

```go
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            doWork()
        }
    }
}

func poll(ctx context.Context) {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            check()
        }
    }
}
```

## fa-triangle-exclamation Err() Checking

```go
select {
case <-ctx.Done():
    if ctx.Err() == context.Canceled {
        fmt.Println("cancelled by caller")
    }
    if ctx.Err() == context.DeadlineExceeded {
        fmt.Println("timeout exceeded")
    }
}
```

## fa-sitemap Propagation Pattern

```go
func Handler(ctx context.Context) {
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    result, err := fetchFromDB(ctx)
    if err != nil {
        log.Fatal(err)
    }
    process(ctx, result)
}

func fetchFromDB(ctx context.Context) (string, error) {
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()
    return query(ctx)
}
```

## fa-server HTTP Server Context

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    ctx = context.WithValue(ctx, key("user"), "alice")

    result, err := db.QueryContext(ctx, "SELECT * FROM users")
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    json.NewEncoder(w).Encode(result)
}
```

## fa-globe HTTP Client Context

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.example.com", nil)
if err != nil {
    log.Fatal(err)
}

resp, err := http.DefaultClient.Do(req)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()
```

## fa-database Database Context

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

row := db.QueryRowContext(ctx, "SELECT name FROM users WHERE id = $1", 1)
var name string
if err := row.Scan(&name); err != nil {
    log.Fatal(err)
}

_, err := db.ExecContext(ctx, "UPDATE users SET name = $1 WHERE id = $2", "Bob", 1)
```

## fa-power-off Graceful Shutdown

```go
srv := &http.Server{Addr: ":8080"}

go func() {
    if err := srv.ListenAndServe(); err != http.ErrServerClosed {
        log.Fatal(err)
    }
}()

quit := make(chan os.Signal, 1)
signal.Notify(quit, os.Interrupt)
<-quit

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
srv.Shutdown(ctx)
```

## fa-arrows-split-x-y Pipeline Cancellation

```go
func gen(ctx context.Context, nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            select {
            case <-ctx.Done():
                return
            case out <- n:
            }
        }
    }()
    return out
}

func sq(ctx context.Context, in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            select {
            case <-ctx.Done():
                return
            case out <- n * n:
            }
        }
    }()
    return out
}

ctx, cancel := context.WithCancel(context.Background())
defer cancel()
for v := range sq(ctx, gen(ctx, 1, 2, 3)) {
    fmt.Println(v)
}
```

## fa-bug Anti-patterns

```go
// BAD: 不要将 Context 存储在结构体中
type Service struct {
    ctx context.Context
}

// BAD: 不要传递 nil 作为 Context
func process(ctx context.Context) { }
process(nil)

// GOOD: 始终将 Context 作为第一个参数传递
func DoSomething(ctx context.Context, arg string) error { return nil }

// BAD: 不要用 WithValue 传递可选参数
context.WithValue(ctx, "timeout", 30)

// GOOD: 使用类型化 key 存储上下文值
type requestIDKey struct{}
context.WithValue(ctx, requestIDKey{}, "abc")
```
