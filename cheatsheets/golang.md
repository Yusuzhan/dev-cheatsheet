---
title: Go
icon: fa-golang
primary: "#00ADD8"
lang: go
---

## fa-rocket Basic Types & Variables

```go
var name string = "Go"
var age int = 18
var flag bool = true
x := 42                  // short declaration (inside functions only)
s := "hello"
f := 3.14

const Pi = 3.14159
const (
    StatusOK    = 200
    StatusError = 500
)
```

## fa-pen-to-square Functions

```go
func add(a, b int) int {
    return a + b
}

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

func greet(name string, greetings ...string) string {
    return fmt.Sprintf("Hello %s", name)
}

n, err := strconv.Atoi("42")
if err != nil {
    log.Fatal(err)
}
```

## fa-layer-group Structs & Methods

```go
type User struct {
    ID    int
    Name  string
    Email string
}

func (u User) Display() string {
    return fmt.Sprintf("%s <%s>", u.Name, u.Email)
}

func (u *User) SetName(name string) {
    u.Name = name
}

user := User{ID: 1, Name: "Alice", Email: "a@b.com"}
user.SetName("Bob")
```

## fa-cubes Interfaces & Embedding

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}

type Animal struct {
    Name string
}

type Dog struct {
    Animal              // embedded struct
    Breed string
}

func main() {
    d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}
    fmt.Println(d.Name)   // access embedded field directly
}
```

## fa-arrows-left-right Slices & Maps

```go
// slices
nums := []int{1, 2, 3}
nums = append(nums, 4, 5)
sub := nums[1:3]
make3 := make([]int, 0, 10)

// maps
m := map[string]int{"a": 1, "b": 2}
m["c"] = 3
val, ok := m["a"]
delete(m, "b")

for k, v := range m {
    fmt.Println(k, v)
}
```

## fa-shuffle Concurrency

```go
// goroutine
go func(msg string) {
    fmt.Println(msg)
}("hello")

// channels
ch := make(chan int, 3)
ch <- 42
val := <-ch

// select
select {
case msg := <-ch1:
    fmt.Println(msg)
case ch2 <- 42:
    fmt.Println("sent")
case <-time.After(time.Second):
    fmt.Println("timeout")
}

// sync.WaitGroup
var wg sync.WaitGroup
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        doWork()
    }()
}
wg.Wait()
```

## fa-circle-half-stroke Error Handling

```go
file, err := os.Open("data.txt")
if err != nil {
    return fmt.Errorf("open file: %w", err)
}
defer file.Close()

// custom error type
type NotFoundError struct {
    Name string
}

func (e *NotFoundError) Error() string {
    return fmt.Sprintf("%s not found", e.Name)
}

// errors.Is / errors.As
if errors.Is(err, os.ErrNotExist) { }
var nfe *NotFoundError
if errors.As(err, &nfe) { }
```

## fa-sitemap Packages & Modules

```bash
go mod init github.com/user/project
go mod tidy
go mod vendor
go get github.com/gin-gonic/gin@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

```go
import (
    "fmt"
    "strings"
    "net/http"

    "github.com/user/project/internal/service"
)
```

## fa-code-branch Control Flow

```go
// if with init statement
if err := doSomething(); err != nil {
    log.Fatal(err)
}

// switch (no break needed)
switch os := runtime.GOOS; os {
case "linux":
    fmt.Println("Linux")
case "darwin":
    fmt.Println("macOS")
default:
    fmt.Println(os)
}

// type switch
switch v := x.(type) {
case string:
    fmt.Println("string:", v)
case int:
    fmt.Println("int:", v)
}
```

## fa-dna Generics

```go
func Map[T any, U any](s []T, f func(T) U) []U {
    result := make([]U, len(s))
    for i, v := range s {
        result[i] = f(v)
    }
    return result
}

nums := []int{1, 2, 3}
strs := Map(nums, func(n int) string {
    return strconv.Itoa(n)
})

type Number interface {
    int | int64 | float64
}

func Max[T Number](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

## fa-toolbox Common Patterns

```go
// context with timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// defer for cleanup
f, _ := os.Open("file.txt")
defer f.Close()

// mutex
var mu sync.Mutex
mu.Lock()
defer mu.Unlock()

// pool of goroutines
jobs := make(chan int, 100)
var wg sync.WaitGroup
for w := 0; w < 3; w++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        for j := range jobs {
            process(j)
        }
    }()
}
```

## fa-flask Testing

```go
func TestAdd(t *testing.T) {
    got := add(2, 3)
    want := 5
    if got != want {
        t.Errorf("add(2, 3) = %d, want %d", got, want)
    }
}

func TestDivide(t *testing.T) {
    _, err := divide(10, 0)
    if err == nil {
        t.Error("expected error for division by zero")
    }
}

// table-driven test
func TestToUpper(t *testing.T) {
    tests := []struct {
        in, want string
    }{
        {"hello", "HELLO"},
        {"Go", "GO"},
    }
    for _, tt := range tests {
        got := strings.ToUpper(tt.in)
        if got != tt.want {
            t.Errorf("ToUpper(%q) = %q, want %q", tt.in, got, tt.want)
        }
    }
}
```

## fa-lightbulb Useful Snippets

```go
// read entire file
data, err := os.ReadFile("config.json")

// write file
err := os.WriteFile("out.txt", []byte("hello"), 0644)

// HTTP server
http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
})
log.Fatal(http.ListenAndServe(":8080", nil))

// enumerate with index
for i, v := range []string{"a", "b", "c"} {
    fmt.Println(i, v)
}

// string conversion
s := strconv.Itoa(42)
n, _ := strconv.Atoi("42")
```
