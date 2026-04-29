---
title: Go net/http
icon: fa-globe
primary: "#00ADD8"
lang: go
locale: zhs
---

## fa-server HTTP Server

```go
http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
})

log.Fatal(http.ListenAndServe(":8080", nil))

srv := &http.Server{
    Addr:         ":8080",
    Handler:      nil,
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  120 * time.Second,
}
log.Fatal(srv.ListenAndServe())
```

## fa-puzzle-piece Handler & HandlerFunc

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello!")
}

http.Handle("/hello", &HelloHandler{})

http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi!")
})
```

## fa-route Routing

```go
mux := http.NewServeMux()
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "home")
})
mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "users")
})
mux.HandleFunc("/api/users/", func(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    fmt.Fprintf(w, "user: %s", id)
})

log.Fatal(http.ListenAndServe(":8080", mux))
```

## fa-magnifying-glass Request Parsing

```go
func handler(w http.ResponseWriter, r *http.Request) {
    method := r.Method
    path := r.URL.Path
    query := r.URL.Query()
    page := query.Get("page")

    r.ParseForm()
    name := r.FormValue("name")

    header := r.Header.Get("Content-Type")
    host := r.Host
    remote := r.RemoteAddr

    contentType := r.Header.Get("Content-Type")
    body, _ := io.ReadAll(r.Body)
    defer r.Body.Close()
}
```

## fa-brackets-curly JSON Request/Response

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func createHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
```

## fa-layer-group Middleware

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}

func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}

mux := http.NewServeMux()
mux.HandleFunc("/", handler)
handler := loggingMiddleware(authMiddleware(mux))
http.ListenAndServe(":8080", handler)
```

## fa-folder Static Files

```go
fs := http.FileServer(http.Dir("./static"))
http.Handle("/static/", http.StripPrefix("/static/", fs))

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "index.html")
})

http.ListenAndServe(":8080", nil)
```

## fa-paper-plane HTTP Client

```go
client := &http.Client{
    Timeout: 10 * time.Second,
}

resp, err := client.Get("https://api.example.com/data")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

body, _ := io.ReadAll(resp.Body)
fmt.Println(string(body))
```

## fa-right-left GET / POST Requests

```go
resp, err := http.Get("https://api.example.com/users")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)

payload := bytes.NewBufferString(`{"name":"Alice"}`)
resp, err = http.Post("https://api.example.com/users", "application/json", payload)
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

formData := url.Values{}
formData.Set("name", "Alice")
formData.Set("email", "alice@example.com")
resp, err = http.PostForm("https://api.example.com/users", formData)
```

## fa-gears Custom Transport

```go
transport := &http.Transport{
    MaxIdleConns:        100,
    MaxIdleConnsPerHost: 10,
    IdleConnTimeout:     90 * time.Second,
    TLSHandshakeTimeout: 10 * time.Second,
}

client := &http.Client{
    Transport: transport,
    Timeout:   30 * time.Second,
}

req, _ := http.NewRequest("GET", "https://api.example.com", nil)
req.Header.Set("Authorization", "Bearer token123")
req.Header.Set("Accept", "application/json")

resp, err := client.Do(req)
```

## fa-clock Timeouts

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.example.com", nil)
if err != nil {
    log.Fatal(err)
}

client := &http.Client{}
resp, err := client.Do(req)
if err != nil {
    if ctx.Err() == context.DeadlineExceeded {
        fmt.Println("request timed out")
    }
    log.Fatal(err)
}
defer resp.Body.Close()
```

## fa-lock TLS / HTTPS

```go
srv := &http.Server{
    Addr: ":8443",
    TLSConfig: &tls.Config{
        MinVersion: tls.VersionTLS12,
    },
}
log.Fatal(srv.ListenAndServeTLS("cert.pem", "key.pem"))

client := &http.Client{
    Transport: &http.Transport{
        TLSClientConfig: &tls.Config{
            InsecureSkipVerify: false,
        },
    },
}
```

## fa-power-off Server Graceful Shutdown

```go
srv := &http.Server{Addr: ":8080"}

go func() {
    log.Fatal(srv.ListenAndServe())
}()

quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit

ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

if err := srv.Shutdown(ctx); err != nil {
    log.Fatal("shutdown error:", err)
}
fmt.Println("server stopped")
```

## fa-upload File Upload

```go
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(10 << 20)

    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()

    dst, _ := os.Create(handler.Filename)
    defer dst.Close()
    io.Copy(dst, file)

    fmt.Fprintf(w, "uploaded: %s (%d bytes)", handler.Filename, handler.Size)
}

http.HandleFunc("/upload", uploadHandler)
```
