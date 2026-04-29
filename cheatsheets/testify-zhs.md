---
title: Testify
icon: fa-flask-vial
primary: "#00ADD8"
lang: go
locale: zhs
---

## fa-check assert 包

```go
import "github.com/stretchr/testify/assert"

func TestAdd(t *testing.T) {
    assert.Equal(t, 4, 2+2)
    assert.NotEqual(t, 5, 2+2)
    assert.True(t, true)
    assert.Nil(t, nil)
    assert.NotNil(t, &struct{}{})
    assert.Contains(t, "hello world", "world")
}

func TestWithMessages(t *testing.T) {
    assert.Equal(t, 200, resp.StatusCode, "状态码应为 200")
}
```

## fa-bolt require 包

```go
import "github.com/stretchr/testify/require"

func TestDB(t *testing.T) {
    db, err := sql.Open("postgres", dsn)
    require.NoError(t, err, "数据库连接必须成功")
    defer db.Close()

    require.NotNil(t, db)

    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
    require.NoError(t, err)
}
```

## fa-equals Equal / NotEqual

```go
func TestEquality(t *testing.T) {
    assert.Equal(t, "hello", "hello")
    assert.NotEqual(t, "hello", "world")

    assert.EqualValues(t, uint32(42), int32(42))

    assert.EqualExportedValues(t, User{Name: "Alice"}, gotUser)

    type Point struct{ X, Y int }
    assert.Equal(t, Point{1, 2}, Point{1, 2})
}
```

## fa-ban Nil / NotNil

```go
func TestNilChecks(t *testing.T) {
    var ptr *int
    assert.Nil(t, ptr)
    assert.Nil(t, error(nil))

    s := []int{}
    assert.NotNil(t, s)
    assert.NotNil(t, &struct{}{})

    var ch chan int
    assert.Nil(t, ch)

    var m map[string]int
    assert.Nil(t, m)
}
```

## fa-toggle-on True / False

```go
func TestBool(t *testing.T) {
    assert.True(t, 1 == 1)
    assert.True(t, strings.Contains("hello", "ell"))
    assert.True(t, len([]int{1, 2, 3}) == 3)

    assert.False(t, 1 == 2)
    assert.False(t, errors.Is(err, os.ErrNotExist))
}
```

## fa-magnifying-glass Contains / NotContains

```go
func TestContains(t *testing.T) {
    assert.Contains(t, "hello world", "world")
    assert.Contains(t, []string{"a", "b", "c"}, "b")
    assert.Contains(t, map[string]int{"a": 1, "b": 2}, "a")

    assert.NotContains(t, "hello", "xyz")
    assert.NotContains(t, []int{1, 2, 3}, 5)

    assert.ElementsMatch(t, []int{3, 1, 2}, []int{1, 2, 3})

    assert.Subset(t, []int{1, 2, 3, 4}, []int{2, 3})
}
```

## fa-code JSON 与 YAML 相等

```go
func TestJSON(t *testing.T) {
    assert.JSONEq(t,
        `{"name":"Alice","age":30}`,
        `{"age":30,"name":"Alice"}`,
    )

    assert.YAMLEq(t,
        "name: Alice\nage: 30\n",
        "age: 30\nname: Alice\n",
    )
}

func TestStructCompare(t *testing.T) {
    expected := User{Name: "Alice", Age: 30}
    assert.Equal(t, expected, gotUser)
}
```

## fa-masks-theater mock 包

```go
type MockRepo struct {
    mock.Mock
}

func (m *MockRepo) GetByID(id int) (*User, error) {
    args := m.Called(id)
    return args.Get(0).(*User), args.Error(1)
}

func TestGetUser(t *testing.T) {
    repo := new(MockRepo)
    repo.On("GetByID", 42).Return(&User{Name: "Alice"}, nil)

    svc := NewUserService(repo)
    user, err := svc.Get(42)

    assert.NoError(t, err)
    assert.Equal(t, "Alice", user.Name)
    repo.AssertExpectations(t)
    repo.AssertCalled(t, "GetByID", 42)
    repo.AssertNumberOfCalls(t, "GetByID", 1)
}
```

## fa-layer-group testify/suite

```go
type UserSuite struct {
    suite.Suite
    db  *sql.DB
    svc *UserService
}

func (s *UserSuite) SetupSuite() {
    s.db, _ = sql.Open("sqlite3", ":memory:")
    s.svc = NewUserService(s.db)
}

func (s *UserSuite) TearDownSuite() {
    s.db.Close()
}

func (s *UserSuite) TestCreate() {
    user, err := s.svc.Create("Alice")
    s.NoError(err)
    s.Equal("Alice", user.Name)
}

func TestUserSuite(t *testing.T) {
    suite.Run(t, new(UserSuite))
}
```

## fa-arrows-rotate BeforeTest / AfterTest

```go
type DBSuite struct {
    suite.Suite
    db *sql.DB
    tx *sql.Tx
}

func (s *DBSuite) SetupSuite() {
    s.db, _ = sql.Open("postgres", dsn)
}

func (s *DBSuite) TearDownSuite() {
    s.db.Close()
}

func (s *DBSuite) BeforeTest(suiteName, testName string) {
    s.tx, _ = s.db.Begin()
}

func (s *DBSuite) AfterTest(suiteName, testName string) {
    s.tx.Rollback()
}

func (s *DBSuite) TestInsert() {
    _, err := s.tx.Exec("INSERT INTO users (name) VALUES ($1)", "Alice")
    s.NoError(err)
}
```

## fa-globe HTTP 断言

```go
import "github.com/stretchr/testify/assert/http"

func TestHandler(t *testing.T) {
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(200)
        fmt.Fprint(w, `{"status":"ok"}`)
    })

    assert.HTTPSuccess(t, handler, "GET", "/health", nil)
    assert.HTTPStatusCode(t, handler, "GET", "/health", nil, 200)
    assert.HTTPBodyContains(t, handler, "GET", "/health", nil, "ok")
    assert.HTTPBodyNotContains(t, handler, "GET", "/health", nil, "error")
}
```

## fa-clock Eventually / EventuallyWithT

```go
func TestAsync(t *testing.T) {
    state := &atomic.Int32{}
    go func() {
        time.Sleep(100 * time.Millisecond)
        state.Store(42)
    }()

    assert.Eventually(t, func() bool {
        return state.Load() == 42
    }, 2*time.Second, 50*time.Millisecond)

    assert.Never(t, func() bool {
        return state.Load() == 99
    }, 1*time.Second, 50*time.Millisecond)
}

func TestEventuallyWithT(t *testing.T) {
    assert.EventuallyWithT(t, func(c *assert.CollectT) {
        resp, _ := http.Get("http://localhost:8080/health")
        assert.Equal(c, 200, resp.StatusCode)
    }, 5*time.Second, 100*time.Millisecond)
}
```

## fa-triangle-exclamation 错误断言

```go
func TestErrors(t *testing.T) {
    _, err := strconv.Atoi("notanumber")
    assert.Error(t, err)
    assert.ErrorIs(t, err, strconv.ErrSyntax)

    wrapped := fmt.Errorf("parse: %w", err)
    assert.ErrorIs(t, wrapped, strconv.ErrSyntax)

    assert.EqualError(t, err, `strconv.Atoi: parsing "notanumber": invalid syntax`)

    assert.ErrorContains(t, err, "invalid syntax")

    _, noErr := strconv.Atoi("42")
    assert.NoError(t, noErr)
}
```

## fa-table 表驱动测试

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"正数", 2, 3, 5},
        {"负数", -1, -1, -2},
        {"零值", 0, 0, 0},
        {"混合", -5, 10, 5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equal(t, tt.expected, Add(tt.a, tt.b))
        })
    }
}

func TestWithSuite(t *testing.T) {
    tests := []struct {
        input    string
        expected bool
    }{
        {"valid@email.com", true},
        {"invalid", false},
    }
    for _, tt := range tests {
        t.Run(tt.input, func(t *testing.T) {
            assert.Equal(t, tt.expected, IsValidEmail(tt.input))
        })
    }
}
```
