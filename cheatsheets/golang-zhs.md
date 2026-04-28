---
title: Go
icon: fa-golang
primary: "#00ADD8"
lang: go
locale: zhs
---

## fa-rocket 基础类型与变量

```go
var name string = "Go"
var age int = 18
var flag bool = true
x := 42                  // 短变量声明（仅限函数内）
s := "hello"
f := 3.14

const Pi = 3.14159
const (
    StatusOK    = 200
    StatusError = 500
)
```

## fa-pen-to-square 函数

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

func greet(name string, greetings ...string) string {  // 可变参数
    return fmt.Sprintf("Hello %s", name)
}

n, err := strconv.Atoi("42")  // 多返回值
if err != nil {
    log.Fatal(err)
}
```

## fa-layer-group 结构体与方法

```go
type User struct {
    ID    int
    Name  string
    Email string
}

func (u User) Display() string {       // 值接收者（不修改原值）
    return fmt.Sprintf("%s <%s>", u.Name, u.Email)
}

func (u *User) SetName(name string) {  // 指针接收者（可修改原值）
    u.Name = name
}

user := User{ID: 1, Name: "Alice", Email: "a@b.com"}
user.SetName("Bob")
```

## fa-cubes 接口与嵌入

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {  // 接口组合
    Reader
    Writer
}

type Animal struct {
    Name string
}

type Dog struct {
    Animal              // 嵌入结构体
    Breed string
}

func main() {
    d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}
    fmt.Println(d.Name)   // 直接访问嵌入字段
}
```

## fa-arrows-left-right 切片与映射

```go
// 切片 slice
nums := []int{1, 2, 3}
nums = append(nums, 4, 5)
sub := nums[1:3]                    // 切片 [2, 3]
make3 := make([]int, 0, 10)        // 预分配容量

// 映射 map
m := map[string]int{"a": 1, "b": 2}
m["c"] = 3
val, ok := m["a"]                  // 检查键是否存在
delete(m, "b")

for k, v := range m {
    fmt.Println(k, v)
}
```

## fa-shuffle 并发编程

```go
// goroutine 协程
go func(msg string) {
    fmt.Println(msg)
}("hello")

// channel 通道
ch := make(chan int, 3)  // 带缓冲通道
ch <- 42
val := <-ch

// select 多路复用
select {
case msg := <-ch1:
    fmt.Println(msg)
case ch2 <- 42:
    fmt.Println("sent")
case <-time.After(time.Second):
    fmt.Println("timeout")
}

// WaitGroup 等待一组协程完成
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

## fa-circle-half-stroke 错误处理

```go
file, err := os.Open("data.txt")
if err != nil {
    return fmt.Errorf("open file: %w", err)  // 包装错误
}
defer file.Close()

// 自定义错误类型
type NotFoundError struct {
    Name string
}

func (e *NotFoundError) Error() string {
    return fmt.Sprintf("%s not found", e.Name)
}

// errors.Is / errors.As 判断错误
if errors.Is(err, os.ErrNotExist) { }
var nfe *NotFoundError
if errors.As(err, &nfe) { }
```

## fa-sitemap 包与模块

```bash
go mod init github.com/user/project   # 初始化模块
go mod tidy                            # 整理依赖
go mod vendor                          # 复制依赖到 vendor
go get github.com/gin-gonic/gin@latest  # 添加依赖
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

```go
import (
    "fmt"
    "strings"
    "net/http"

    "github.com/user/project/internal/service"  // 内部包
)
```

## fa-code-branch 控制流

```go
// if 带初始化语句
if err := doSomething(); err != nil {
    log.Fatal(err)
}

// switch 无需 break
switch os := runtime.GOOS; os {
case "linux":
    fmt.Println("Linux")
case "darwin":
    fmt.Println("macOS")
default:
    fmt.Println(os)
}

// 类型 switch
switch v := x.(type) {
case string:
    fmt.Println("string:", v)
case int:
    fmt.Println("int:", v)
}
```

## fa-dna 泛型

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

## fa-toolbox 常用模式

```go
// 带超时的 context
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// defer 用于资源清理
f, _ := os.Open("file.txt")
defer f.Close()

// 互斥锁
var mu sync.Mutex
mu.Lock()
defer mu.Unlock()

// goroutine 工作池
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

## fa-flask 测试

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

// 表驱动测试
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

## fa-lightbulb 实用代码片段

```go
// 读取整个文件
data, err := os.ReadFile("config.json")

// 写入文件
err := os.WriteFile("out.txt", []byte("hello"), 0644)

// HTTP 服务
http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
})
log.Fatal(http.ListenAndServe(":8080", nil))

// 遍历时获取索引
for i, v := range []string{"a", "b", "c"} {
    fmt.Println(i, v)
}

// 字符串与整数互转
s := strconv.Itoa(42)
n, _ := strconv.Atoi("42")
```
