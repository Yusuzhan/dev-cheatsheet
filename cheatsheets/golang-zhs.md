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

## fa-location-dot 指针

```go
x := 42
p := &x            // p 是指向 x 的指针 (*int)
fmt.Println(*p)    // 解引用：输出 42
*p = 100           // 通过指针修改值
fmt.Println(x)     // 100

var ptr *int       // 空指针 (nil)
if ptr != nil {
    fmt.Println(*ptr)
}

// 函数中的值与指针
func inc(n *int) {
    *n++           // 修改原始变量
}
inc(&x)

// new() 分配零值并返回指针
sp := new(string)  // *string，指向 ""
fmt.Println(*sp)   // ""

// 结构体指针
user := &User{ID: 1, Name: "Alice"}
user.Name = "Bob"  // 结构体字段无需显式解引用
fmt.Println(user.Name)
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

## fa-font 字符串格式化

```go
fmt.Sprintf("name: %s, age: %d", "Alice", 25)  // 格式化字符串
fmt.Printf("value: %v\n", data)                 // 格式化打印
fmt.Fprintf(w, "status: %d", code)              // 写入 io.Writer

// 常用动词
// %s   字符串
// %d   十进制整数
// %f   浮点数（默认6位小数），%.2f 保留2位
// %v   默认格式（任意值）
// %+v  显示结构体字段名
// %#v  Go 语法表示
// %T   值的类型
// %q   带引号的字符串
// %p   指针地址
// %02d 补零整数，宽度2
// %x   十六进制
// %b   二进制
// %%   百分号字面量

fmt.Sprintf("%+v", user)          // 带字段名的结构体
fmt.Sprintf("%.2f", 3.14159)      // "3.14"
fmt.Sprintf("%04d", 7)            // "0007"
fmt.Sprintf("%x", 255)            // "ff"
```

## fa-clock 日期与时间

```go
now := time.Now()
now.Year()                         // 2025
now.Month()                        // time.April
now.Day()                          // 29
now.Hour()                         // 14
now.Weekday()                      // time.Tuesday

// 格式化（参考时间: Mon Jan 2 15:04:05 MST 2006）
now.Format("2006-01-02")           // "2025-04-29"
now.Format("2006-01-02 15:04:05")  // "2025-04-29 14:30:00"
now.Format("01/02/06")             // "04/29/25"
now.Format(time.RFC3339)           // "2025-04-29T14:30:00+08:00"
now.Format(time.RFC1123)           // "Tue, 29 Apr 2025 14:30:00 UTC"

// 解析
t, _ := time.Parse("2006-01-02", "2025-04-29")
t, _ := time.Parse(time.RFC3339, "2025-04-29T14:30:00Z")

// 时间间隔
time.Sleep(2 * time.Second)
diff := time.Since(start)          // 经过的时间 time.Duration
diff.Minutes()                     // float64 分钟
diff.Seconds()                     // float64 秒

// 加减时间
later := now.Add(24 * time.Hour)
deadline := now.AddDate(0, 1, 0)   // 加 1 个月

// 比较
now.Before(deadline)               // true/false
now.After(deadline)
now.Equal(otherTime)
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
