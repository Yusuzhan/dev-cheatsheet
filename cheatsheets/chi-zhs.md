---
title: Chi
icon: fa-route
primary: "#00ADD8"
lang: go
locale: zhs
---

## fa-rocket 初始化与路由

```go
import "github.com/go-chi/chi/v5"

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    http.ListenAndServe(":3000", r)
}
```

## fa-map-pin 路由定义

```go
r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello"))
})
r.Post("/users", createUser)
r.Put("/users/{id}", updateUser)
r.Delete("/users/{id}", deleteUser)
r.Patch("/users/{id}", patchUser)

r.MethodFunc("GET", "/health", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("ok"))
})
```

## fa-layer-group 路由分组

```go
r.Route("/api", func(r chi.Router) {
    r.Get("/status", getStatus)
    r.Route("/users", func(r chi.Router) {
        r.Get("/", listUsers)
        r.Post("/", createUser)
        r.Route("/{id}", func(r chi.Router) {
            r.Get("/", getUser)
            r.Delete("/", deleteUser)
        })
    })
})

r.Mount("/admin", adminRouter())  // 挂载子路由器
```

## fa-shield 中间件

```go
r.Use(middleware.Logger)          // 请求日志
r.Use(middleware.Recoverer)       // panic 恢复
r.Use(middleware.RequestID)       // 请求 ID
r.Use(middleware.RealIP)          // 真实 IP
r.Use(middleware.Timeout(60 * time.Second))  // 超时控制
r.Use(middleware.Throttle(100))   // 限流
r.Use(middleware.Compress(5))     // gzip 压缩
r.Use(middleware.Heartbeat("/ping"))  // 健康检查
```

## fa-link URL 参数

```go
r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    fmt.Fprintf(w, "user: %s", id)
})

r.Get("/articles/{slug}", func(w http.ResponseWriter, r *http.Request) {
    slug := chi.URLParam(r, "slug")
    fmt.Fprintf(w, "article: %s", slug)
})
```

## fa-hand-pointer 请求绑定

```go
type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()
    json.NewEncoder(w).Encode(req)
})
```

## fa-code JSON 响应

```go
func renderJSON(w http.ResponseWriter, v interface{}, status int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(v)
}

r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
    users := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}
    renderJSON(w, users, http.StatusOK)
})

r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
    renderJSON(w, map[string]string{"error": "not found"}, http.StatusNotFound)
})
```

## fa-folder-open 静态文件服务

```go
r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
    fs := http.StripPrefix("/static", http.FileServer(http.Dir("./static")))
    fs.ServeHTTP(w, r)
})

workDir, _ := os.Getwd()
filesDir := http.Dir(filepath.Join(workDir, "files"))
r.Get("/files/*", http.FileServer(filesDir).ServeHTTP)
```

## fa-database Context 传值

```go
r.Use(func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx := context.WithValue(r.Context(), "userID", 42)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
})

r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID").(int)
    fmt.Fprintf(w, "user: %d", userID)
})
```

## fa-wrench 自定义中间件

```go
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "unauthorized", http.StatusUnauthorized)
            return
        }
        ctx := context.WithValue(r.Context(), "token", token)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

r.With(AuthMiddleware).Get("/protected", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("secret data"))
})

r.Group(func(r chi.Router) {
    r.Use(AuthMiddleware)
    r.Get("/dashboard", getDashboard)
    r.Get("/settings", getSettings)
})
```

## fa-right-left 从标准库迁移

```go
r := chi.NewRouter()
r.Use(middleware.Logger)

r.HandleFunc("/legacy", legacyHandler)  // 兼容 http.HandleFunc
r.Get("/new", newHandler)

http.ListenAndServe(":3000", r)

r.NotFound(func(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "custom 404", http.StatusNotFound)
})

r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
})
```
