---
title: Go Testing & Benchmarking
icon: fa-flask
primary: "#00ADD8"
lang: go
locale: zhs
---

## fa-play 基础测试

```go
func TestAdd(t *testing.T) {
    got := add(2, 3)
    want := 5
    if got != want {
        t.Errorf("add(2, 3) = %d, want %d", got, want)
    }
}

func TestError(t *testing.T) {
    _, err := divide(10, 0)
    if err == nil {
        t.Fatal("expected error for division by zero")
    }
}
```

## fa-table-cells 表驱动测试

```go
func TestToUpper(t *testing.T) {
    tests := []struct {
        input string
        want  string
    }{
        {"hello", "HELLO"},
        {"go", "GO"},
        {"", ""},
    }
    for _, tt := range tests {
        got := strings.ToUpper(tt.input)
        if got != tt.want {
            t.Errorf("ToUpper(%q) = %q, want %q", tt.input, got, tt.want)
        }
    }
}
```

## fa-list 子测试

```go
func TestMath(t *testing.T) {
    t.Run("add", func(t *testing.T) {
        if add(1, 2) != 3 {
            t.Error("add failed")
        }
    })
    t.Run("subtract", func(t *testing.T) {
        if subtract(3, 1) != 2 {
            t.Error("subtract failed")
        }
    })
}

go test -run TestMath/add   // 只运行某个子测试
```

## fa-gauge-high 基准测试

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        add(1, 2)
    }
}

func BenchmarkJSONEncode(b *testing.B) {
    data := map[string]string{"key": "value"}
    b.ResetTimer()  // 跳过初始化时间
    for i := 0; i < b.N; i++ {
        json.Marshal(data)
    }
}

go test -bench=. -benchmem   // 运行所有基准测试并显示内存分配
```

## fa-chart-line 基准测试进阶

```go
func BenchmarkSprintf(b *testing.B) {
    b.ReportAllocs()  // 报告内存分配次数
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello %s", "world")
    }
}

func BenchmarkParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            add(1, 2)
        }
    })
}
```

## fa-server TestMain

```go
func TestMain(m *testing.M) {
    setup()          // 全局初始化
    code := m.Run()  // 运行所有测试
    teardown()       // 全局清理
    os.Exit(code)
}

func setup() {
    db = connectDB()
}

func teardown() {
    db.Close()
}
```

## fa-masks-theater 模拟对象（接口）

```go
type Store interface {
    GetUser(id int) (*User, error)
}

type MockStore struct {
    Users map[int]*User
}

func (m *MockStore) GetUser(id int) (*User, error) {
    u, ok := m.Users[id]
    if !ok {
        return nil, fmt.Errorf("not found")
    }
    return u, nil
}

func TestGetUser(t *testing.T) {
    mock := &MockStore{Users: map[int]*User{1: {Name: "Alice"}}}
    u, err := mock.GetUser(1)
    if err != nil || u.Name != "Alice" {
        t.Fail()
    }
}
```

## fa-globe httptest

```go
func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/api/users", nil)
    w := httptest.NewRecorder()
    handler(w, req)                  // 测试 Handler
    resp := w.Result()
    body, _ := io.ReadAll(resp.Body)
    if resp.StatusCode != 200 {
        t.Errorf("expected 200, got %d", resp.StatusCode)
    }
}

func TestHTTPClient(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(200)
        w.Write([]byte(`{"status":"ok"}`))
    }))
    defer server.Close()
    resp, _ := http.Get(server.URL + "/test")  // 测试 HTTP 客户端
    if resp.StatusCode != 200 {
        t.Fail()
    }
}
```

## fa-toolbox 测试辅助工具

```go
func TestWithHelper(t *testing.T) {
    assertEqual(t, add(1, 2), 3)
}

func assertEqual(t *testing.T, got, want int) {
    t.Helper()  // 标记为辅助函数，错误定位到调用处
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}

func TestTempDir(t *testing.T) {
    dir := t.TempDir()  // 自动创建和清理临时目录
    os.WriteFile(filepath.Join(dir, "test.txt"), []byte("hello"), 0644)
}

func TestCleanup(t *testing.T) {
    f, _ := os.Create("temp")
    t.Cleanup(func() { os.Remove("temp") })  // 注册清理函数
}
```

## fa-shuffle 模糊测试

```go
func FuzzReverse(f *testing.F) {
    f.Add("hello")  // 添加种子语料
    f.Add("go")
    f.Fuzz(func(t *testing.T, orig string) {
        rev := reverse(orig)
        double := reverse(rev)
        if orig != double {
            t.Errorf("reverse(reverse(%q)) = %q", orig, double)
        }
    })
}

go test -fuzz=FuzzReverse   // 运行模糊测试
```

## fa-forward 跳过与并行

```go
func TestSkip(t *testing.T) {
    if runtime.GOOS == "windows" {
        t.Skip("skipping on windows")  // 跳过测试
    }
}

func TestParallel(t *testing.T) {
    t.Parallel()  // 标记为可并行执行
    if add(1, 2) != 3 {
        t.Error("failed")
    }
}

func TestParallelTable(t *testing.T) {
    tests := []struct{ a, b, want int }{
        {1, 2, 3}, {4, 5, 9},
    }
    for _, tt := range tests {
        tt := tt  // 捕获变量
        t.Run(fmt.Sprintf("%d+%d", tt.a, tt.b), func(t *testing.T) {
            t.Parallel()
            if got := add(tt.a, tt.b); got != tt.want {
                t.Errorf("got %d, want %d", got, tt.want)
            }
        })
    }
}
```

## fa-chart-pie 覆盖率

```bash
go test -cover ./...                          # 显示覆盖率概览
go test -coverprofile=coverage.out ./...      # 生成覆盖率文件
go tool cover -func=coverage.out              # 按函数查看覆盖率
go tool cover -html=coverage.out -o coverage.html  # 生成 HTML 报告
go test -covermode=atomic ./...               # 并发安全的覆盖率统计
```

## fa-file Golden 文件

```go
func TestOutput(t *testing.T) {
    got := generateReport()
    golden := filepath.Join("testdata", t.Name()+".golden")
    if *update {  // -update 标志更新 golden 文件
        os.WriteFile(golden, []byte(got), 0644)
    }
    want, err := os.ReadFile(golden)
    if err != nil {
        t.Fatal(err)
    }
    if got != string(want) {
        t.Errorf("output mismatch\n got:\n%s\n want:\n%s", got, string(want))
    }
}
```
