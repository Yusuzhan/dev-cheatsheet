---
title: Gin
icon: fa-glass-water
primary: "#00ADD8"
lang: go
locale: zhs
---

## fa-rocket 初始化与路由

```go
import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })
    r.Run(":8080")
}
```

## fa-map-pin 路由与参数

```go
r.GET("/users/:id", func(c *gin.Context) {
    id := c.Param("id")
    c.String(200, "User: %s", id)
})

r.GET("/files/*filepath", func(c *gin.Context) {
    fp := c.Param("filepath")
    c.String(200, "File: %s", fp)
})

r.POST("/submit", handleSubmit)
r.PUT("/users/:id", handleUpdate)
r.DELETE("/users/:id", handleDelete)
r.PATCH("/users/:id", handlePatch)
```

## fa-filter 查询与表单参数

```go
r.GET("/search", func(c *gin.Context) {
    q := c.Query("q")                          // 获取查询参数
    page := c.DefaultQuery("page", "1")        // 带默认值
    c.String(200, "q=%s, page=%s", q, page)
})

r.POST("/form", func(c *gin.Context) {
    name := c.PostForm("name")
    email := c.DefaultPostForm("email", "none@example.com")
    c.String(200, "name=%s, email=%s", name, email)
})

r.POST("/upload", func(c *gin.Context) {
    form, _ := c.MultipartForm()
    files := form.File["files"]
    for _, file := range files {
        c.SaveUploadedFile(file, "./uploads/"+file.Filename)
    }
})
```

## fa-link JSON 绑定

```go
type LoginRequest struct {
    User     string `json:"user" binding:"required"`
    Password string `json:"password" binding:"required,min=6"`
}

r.POST("/login", func(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"user": req.User})
})

type SearchQuery struct {
    Q    string `form:"q" binding:"required"`
    Page int    `form:"page,default=1"`
}

r.GET("/search", func(c *gin.Context) {
    var q SearchQuery
    if err := c.ShouldBindQuery(&q); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, q)
})
```

## fa-code JSON 响应

```go
c.JSON(200, gin.H{"status": "ok"})
c.JSON(200, user)
c.JSON(201, gin.H{"id": 1, "name": "Alice"})

c.PureJSON(200, gin.H{"html": "<b>hello</b>"})  // 不转义 HTML

c.XML(200, gin.H{"message": "hello"})

c.String(200, "hello %s", name)

c.Data(200, "image/png", imageBytes)  // 原始二进制响应
```

## fa-layer-group 路由分组

```go
api := r.Group("/api")
{
    api.GET("/status", getStatus)
    users := api.Group("/users")
    {
        users.GET("", listUsers)
        users.POST("", createUser)
        users.GET("/:id", getUser)
    }
}

admin := r.Group("/admin", AuthMiddleware())  // 分组级别中间件
{
    admin.GET("/dashboard", getDashboard)
    admin.GET("/settings", getSettings)
}
```

## fa-shield 中间件

```go
r.Use(gin.Logger())
r.Use(gin.Recovery())

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
            return
        }
        c.Set("userID", 42)
        c.Next()
    }
}

r.Use(func(c *gin.Context) {
    start := time.Now()
    c.Next()
    duration := time.Since(start)
    fmt.Printf("%s %s - %v\n", c.Request.Method, c.Request.URL.Path, duration)
})
```

## fa-upload 文件上传

```go
r.POST("/upload", func(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    dst := "./uploads/" + file.Filename
    c.SaveUploadedFile(file, dst)
    c.JSON(200, gin.H{"filename": file.Filename, "size": file.Size})
})

r.POST("/multi", func(c *gin.Context) {
    form, _ := c.MultipartForm()
    for _, file := range form.File["files"] {
        c.SaveUploadedFile(file, "./uploads/"+file.Filename)
    }
    c.JSON(200, gin.H{"count": len(form.File["files"])})
})
```

## fa-palette HTML 模板

```go
r.LoadHTMLGlob("templates/*")
r.LoadHTMLFiles("templates/index.html", "templates/about.html")

r.GET("/", func(c *gin.Context) {
    c.HTML(200, "index.html", gin.H{
        "title": "Home",
        "items": []string{"Go", "Gin", "Web"},
    })
})

r.Static("/assets", "./assets")                    // 静态文件目录
r.StaticFile("/favicon.ico", "./favicon.ico")      // 单个文件
```

## fa-check 数据验证

```go
type User struct {
    Name  string `json:"name" binding:"required,min=2,max=50"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"required,gte=0,lte=150"`
    URL   string `json:"url" binding:"url"`
    IP    string `json:"ip" binding:"ip"`
}

if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
    v.RegisterValidation("custom", func(fl validator.FieldLevel) bool {
        return len(fl.Field().String()) >= 5
    })
}
```

## fa-cookie-bite Cookie 与会话

```go
r.GET("/set", func(c *gin.Context) {
    c.SetCookie("token", "abc123", 3600, "/", "localhost", false, true)
    c.JSON(200, gin.H{"message": "cookie set"})
})

r.GET("/get", func(c *gin.Context) {
    token, err := c.Cookie("token")
    if err != nil {
        c.JSON(400, gin.H{"error": "no cookie"})
        return
    }
    c.JSON(200, gin.H{"token": token})
})

store := cookie.NewStore([]byte("secret"))
r.Use(sessions.Sessions("session", store))
r.GET("/session", func(c *gin.Context) {
    session := sessions.Default(c)
    session.Set("user", "Alice")
    session.Save()
})
```

## fa-triangle-exclamation 错误处理

```go
r.Use(func(c *gin.Context) {
    c.Next()
    if len(c.Errors) > 0 {
        c.JSON(-1, gin.H{"errors": c.Errors})
    }
})

r.GET("/error", func(c *gin.Context) {
    c.Error(fmt.Errorf("something went wrong"))
})

r.NoRoute(func(c *gin.Context) {
    c.JSON(404, gin.H{"error": "page not found"})
})

r.NoMethod(func(c *gin.Context) {
    c.JSON(405, gin.H{"error": "method not allowed"})
})
```
