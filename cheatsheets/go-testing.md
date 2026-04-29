---
title: Go Testing & Benchmarking
icon: fa-flask
primary: "#00ADD8"
lang: go
---

## fa-play Basic Test

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

## fa-table-cells Table-Driven Tests

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

## fa-list Subtests

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

go test -run TestMath/add
```

## fa-gauge-high Benchmark

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        add(1, 2)
    }
}

func BenchmarkJSONEncode(b *testing.B) {
    data := map[string]string{"key": "value"}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        json.Marshal(data)
    }
}

go test -bench=. -benchmem
```

## fa-chart-line Benchmark Examples

```go
func BenchmarkSprintf(b *testing.B) {
    b.ReportAllocs()
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
    setup()
    code := m.Run()
    teardown()
    os.Exit(code)
}

func setup() {
    db = connectDB()
}

func teardown() {
    db.Close()
}
```

## fa-masks-theater Mocking (interface)

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
    handler(w, req)
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
    resp, _ := http.Get(server.URL + "/test")
    if resp.StatusCode != 200 {
        t.Fail()
    }
}
```

## fa-toolbox Testing Helpers

```go
func TestWithHelper(t *testing.T) {
    assertEqual(t, add(1, 2), 3)
}

func assertEqual(t *testing.T, got, want int) {
    t.Helper()
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}

func TestTempDir(t *testing.T) {
    dir := t.TempDir()
    os.WriteFile(filepath.Join(dir, "test.txt"), []byte("hello"), 0644)
}

func TestCleanup(t *testing.T) {
    f, _ := os.Create("temp")
    t.Cleanup(func() { os.Remove("temp") })
}
```

## fa-shuffle Fuzzing

```go
func FuzzReverse(f *testing.F) {
    f.Add("hello")
    f.Add("go")
    f.Fuzz(func(t *testing.T, orig string) {
        rev := reverse(orig)
        double := reverse(rev)
        if orig != double {
            t.Errorf("reverse(reverse(%q)) = %q", orig, double)
        }
    })
}

go test -fuzz=FuzzReverse
```

## fa-forward Skip & Parallel

```go
func TestSkip(t *testing.T) {
    if runtime.GOOS == "windows" {
        t.Skip("skipping on windows")
    }
}

func TestParallel(t *testing.T) {
    t.Parallel()
    if add(1, 2) != 3 {
        t.Error("failed")
    }
}

func TestParallelTable(t *testing.T) {
    tests := []struct{ a, b, want int }{
        {1, 2, 3}, {4, 5, 9},
    }
    for _, tt := range tests {
        tt := tt
        t.Run(fmt.Sprintf("%d+%d", tt.a, tt.b), func(t *testing.T) {
            t.Parallel()
            if got := add(tt.a, tt.b); got != tt.want {
                t.Errorf("got %d, want %d", got, tt.want)
            }
        })
    }
}
```

## fa-chart-pie Coverage

```bash
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
go tool cover -html=coverage.out -o coverage.html
go test -covermode=atomic ./...
```

## fa-file Golden Files

```go
func TestOutput(t *testing.T) {
    got := generateReport()
    golden := filepath.Join("testdata", t.Name()+".golden")
    if *update {
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
