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

## fa-location-dot Pointers

```go
x := 42
p := &x            // p is a pointer to x (*int)
fmt.Println(*p)    // dereference: prints 42
*p = 100           // modify value through pointer
fmt.Println(x)     // 100

var ptr *int       // nil pointer
if ptr != nil {
    fmt.Println(*ptr)
}

// pointer vs value in functions
func inc(n *int) {
    *n++           // modify original variable
}
inc(&x)

// new() allocates zero value and returns pointer
sp := new(string)  // *string, points to ""
fmt.Println(*sp)   // ""

// pointer to struct
user := &User{ID: 1, Name: "Alice"}
user.Name = "Bob"  // no need to dereference for struct fields
fmt.Println(user.Name)
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

## fa-font String Formatting

```go
fmt.Sprintf("name: %s, age: %d", "Alice", 25)  // formatted string
fmt.Printf("value: %v\n", data)                 // print formatted
fmt.Fprintf(w, "status: %d", code)              // write to io.Writer

// common verbs
// %s   string
// %d   decimal integer
// %f   float (default 6 decimals), %.2f for 2 decimals
// %v   default format (any value)
// %+v  with struct field names
// %#v  Go syntax representation
// %T   type of value
// %q   quoted string
// %p   pointer address
// %02d zero-padded integer, width 2
// %x   hex encoding
// %b   binary representation
// %%   literal percent sign

fmt.Sprintf("%+v", user)          // struct with field names
fmt.Sprintf("%.2f", 3.14159)      // "3.14"
fmt.Sprintf("%04d", 7)            // "0007"
fmt.Sprintf("%x", 255)            // "ff"
```

## fa-clock Date & Time

```go
now := time.Now()
now.Year()                         // 2025
now.Month()                        // time.April
now.Day()                          // 29
now.Hour()                         // 14
now.Weekday()                      // time.Tuesday

// formatting (reference time: Mon Jan 2 15:04:05 MST 2006)
now.Format("2006-01-02")
now.Format("2006-01-02 15:04:05")
now.Format("01/02/06")
now.Format(time.RFC3339)           // "2025-04-29T14:30:00+08:00"
now.Format(time.RFC1123)           // "Tue, 29 Apr 2025 14:30:00 UTC"

// parsing
t, _ := time.Parse("2006-01-02", "2025-04-29")
t, _ := time.Parse(time.RFC3339, "2025-04-29T14:30:00Z")

// duration
time.Sleep(2 * time.Second)
diff := time.Since(start)          // elapsed time.Duration
diff.Minutes()                     // float64
diff.Seconds()

// add / subtract
later := now.Add(24 * time.Hour)
deadline := now.AddDate(0, 1, 0)   // add 1 month

// compare
now.Before(deadline)
now.After(deadline)
now.Equal(otherTime)
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
