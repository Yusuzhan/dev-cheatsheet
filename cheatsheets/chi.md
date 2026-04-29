---
title: Chi
icon: fa-route
primary: "#00ADD8"
lang: go
---

## fa-rocket Setup & Router

```go
import "github.com/go-chi/chi/v5"

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    http.ListenAndServe(":3000", r)
}
```

## fa-map-pin Routes

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

## fa-layer-group Route Groups

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

r.Mount("/admin", adminRouter())
```

## fa-shield Middleware

```go
r.Use(middleware.Logger)
r.Use(middleware.Recoverer)
r.Use(middleware.RequestID)
r.Use(middleware.RealIP)
r.Use(middleware.Timeout(60 * time.Second))
r.Use(middleware.Throttle(100))
r.Use(middleware.Compress(5))
r.Use(middleware.Heartbeat("/ping"))
```

## fa-link URL Parameters

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

## fa-hand-pointer Request Binding

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

## fa-code JSON Response

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

## fa-folder-open File Server

```go
r.Get("/static/*", func(w http.ResponseWriter, r *http.Request) {
    fs := http.StripPrefix("/static", http.FileServer(http.Dir("./static")))
    fs.ServeHTTP(w, r)
})

workDir, _ := os.Getwd()
filesDir := http.Dir(filepath.Join(workDir, "files"))
r.Get("/files/*", http.FileServer(filesDir).ServeHTTP)
```

## fa-database Context Values

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

## fa-wrench Custom Middleware

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

## fa-right-left Migrating from stdlib

```go
r := chi.NewRouter()
r.Use(middleware.Logger)

r.HandleFunc("/legacy", legacyHandler)
r.Get("/new", newHandler)

http.ListenAndServe(":3000", r)

r.NotFound(func(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "custom 404", http.StatusNotFound)
})

r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
})
```
