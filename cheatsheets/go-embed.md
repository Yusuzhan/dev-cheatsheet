---
title: Go Embed
icon: fa-paperclip
primary: "#00ADD8"
lang: go
---

## fa-font embed.String

```go
import _ "embed"

//go:embed message.txt
var message string

func main() {
    fmt.Println(message)
}
```

## fa-code embed.Bytes

```go
import _ "embed"

//go:embed schema.sql
var schema []byte

//go:embed favicon.ico
var favicon []byte

func main() {
    fmt.Println(string(schema))
}
```

## fa-folder-tree embed.FS

```go
import "embed"

//go:embed static/*
var staticFS embed.FS

//go:embed templates/*.html
var templateFS embed.FS

//go:embed configs/*
var configs embed.FS
```

## fa-file Embed Files

```go
import _ "embed"

//go:embed hello.txt
var hello string

//go:embed hello.txt
var helloBytes []byte

//go:embed a.txt b.txt
var multiFile embed.FS
```

## fa-folder Embed Directories

```go
//go:embed assets
var assets embed.FS

//go:embed static/img static/css
var webAssets embed.FS

data, _ := assets.ReadFile("assets/logo.png")
entries, _ := assets.ReadDir("assets")
for _, e := range entries {
    fmt.Println(e.Name(), e.IsDir())
}
```

## fa-globe HTTP File Server

```go
//go:embed static/*
var static embed.FS

func main() {
    fs, _ := fs.Sub(static, "static")
    http.Handle("/", http.FileServer(http.FS(fs)))
    http.ListenAndServe(":8080", nil)
}
```

## fa-book-open ReadFile / ReadDir

```go
//go:embed data/*
var data embed.FS

content, err := data.ReadFile("data/config.json")
entries, err := data.ReadDir("data")

file, err := data.Open("data/config.json")
stat, err := file.Stat()
```

## fa-sitemap Sub FS

```go
//go:embed static/*
var static embed.FS

subFS, err := fs.Sub(static, "static")
if err != nil {
    log.Fatal(err)
}

http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(subFS))))
```

## fa-filter Match Patterns

```go
//go:embed *.txt
var txtFiles embed.FS

//go:embed a.txt b.txt c.txt
var specific embed.FS

//go:embed **/*.json
var jsonFiles embed.FS

//go:embed images/*.png images/*.jpg
var images embed.FS

//go:embed LICENSE README.md
var docs embed.FS
```

## fa-code Embed Templates

```go
//go:embed templates/*.html
var templateFS embed.FS

func main() {
    tmpl, err := template.ParseFS(templateFS, "templates/*.html")
    if err != nil {
        log.Fatal(err)
    }
    tmpl.Execute(os.Stdout, map[string]string{"Title": "Hello"})
}
```

## fa-gear Embed Config

```go
//go:embed config.json
var configData []byte

func main() {
    var cfg struct {
        Port int    `json:"port"`
        Host string `json:"host"`
    }
    if err := json.Unmarshal(configData, &cfg); err != nil {
        log.Fatal(err)
    }
    fmt.Printf("serving on %s:%d\n", cfg.Host, cfg.Port)
}
```

## fa-window-maximize Embed Frontend SPA

```go
//go:embed dist/*
var dist embed.FS

func main() {
    subFS, _ := fs.Sub(dist, "dist")
    http.Handle("/", http.FileServer(http.FS(subFS)))
    http.HandleFunc("/api", apiHandler)
    http.ListenAndServe(":8080", nil)
}
```

## fa-wrench Build Constraints

```go
// +build !dev

//go:embed static/*
var staticFS embed.FS
```

```go
// +build dev

// use filesystem directly in dev mode
var staticFS = os.DirFS("static")
```

## fa-triangle-exclamation Limitations

```go
// only works in package-level vars, not inside functions
// BAD:
func main() {
    //go:embed hello.txt
    var s string
}

// GOOD:
//go:embed hello.txt
var s string

// embed patterns use glob syntax, not regex
// hidden files (starting with .) are excluded by default
// use explicit pattern to include:
//go:embed .hidden
var hidden string

// does not follow symlinks
// paths must be relative (no leading /)
// cannot embed files outside the module
```
