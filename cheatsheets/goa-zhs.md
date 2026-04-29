---
title: Goa
icon: fa-drafting-compass
primary: "#00ADD8"
lang: go
locale: zhs
---

## fa-drafting-compass 设计概述

```go
// goa.design/goa — API 优先设计框架
// 编写 DSL → 生成代码（服务端、客户端、文档）
//
// 项目结构：
//   design/   — API DSL 定义
//   gen/      — 生成代码（勿手动修改）
//   cmd/      — 服务入口

package design

import . "goa.design/goa/v3/dsl"
```

## fa-globe API DSL

```go
var _ = API("myapp", func() {
    Title("My Application")
    Description("A sample API")
    Version("1.0.0")

    Server("myapp", func() {
        Host("localhost", func() {
            URI("http://localhost:8080")
            URI("grpc://localhost:9090")
        })
    })
})
```

## fa-cubes 服务定义

```go
var _ = Service("users", func() {
    Description("用户管理服务")

    Error("not_found", NotFound, "User not found")
    Error("unauthorized", String, "Unauthorized")

    HTTP(func() {
        Path("/users")
    })

    GRPC(func() {
        // gRPC 传输映射
    })

    Method("list", func() {
        Description("List all users")
        Result(CollectionOf(UserResult))
        GRPC(func() {
            Response(CodeOK)
        })
        HTTP(func() {
            GET("/")
            Response(StatusOK)
        })
    })
})
```

## fa-code 方法定义

```go
Method("create", func() {
    Description("创建新用户")
    Payload(CreateUserPayload)
    Result(UserResult)

    HTTP(func() {
        POST("/")
        Response(StatusCreated)
    })

    GRPC(func() {
        Response(CodeOK)
    })
})

Method("show", func() {
    Payload(func() {
        Field(1, "id", UInt, "User ID")
        Required("id")
    })
    Result(UserResult)

    HTTP(func() {
        GET("/{id}")
        Response(StatusOK)
    })
})
```

## fa-database Payload 与 Result

```go
var CreateUserPayload = Type("CreateUserPayload", func() {
    Field(1, "name", String, "用户名", func() {
        MinLength(2)
        MaxLength(100)
    })
    Field(2, "email", String, "邮箱地址", func() {
        Format(FormatEmail)
    })
    Field(3, "age", Int, "年龄", func() {
        Minimum(0)
        Maximum(150)
    })
    Required("name", "email")
})

var UserResult = ResultType("application/vnd.user", func() {
    Field(1, "id", UInt, "用户 ID")
    Field(2, "name", String, "用户名")
    Field(3, "email", String, "邮箱")
    Required("id", "name")
})
```

## fa-triangle-exclamation 错误定义

```go
var NotFound = Type("NotFound", func() {
    Field(1, "message", String, "错误信息")
    Field(2, "id", UInt, "资源 ID")
    Required("message")
})

var _ = Service("users", func() {
    Error("not_found", NotFound, "User not found")

    Error("bad_request", func() {
        Field(1, "detail", String)
        Required("detail")
    })

    HTTP(func() {
        Response("not_found", StatusNotFound)
        Response("bad_request", StatusBadRequest)
    })

    GRPC(func() {
        Response("not_found", CodeNotFound)
        Response("bad_request", CodeInvalidArgument)
    })
})
```

## fa-network-wired HTTP 传输

```go
Method("update", func() {
    Payload(func() {
        Field(1, "id", UInt)
        Field(2, "name", String)
    })
    Result(UserResult)

    HTTP(func() {
        PUT("/{id}")
        Header("Authorization:String")
        Response(StatusOK)
        Response("not_found", StatusNotFound)
    })
})

Method("upload", func() {
    Payload(func() {
        Field(1, "file", Bytes)
    })
    HTTP(func() {
        POST("/upload")
        MultipartRequest()
    })
})

Method("search", func() {
    Payload(func() {
        Field(1, "q", String)
        Field(2, "page", Int, func() { Default(1) })
    })
    HTTP(func() {
        GET("/search")
        Param("q")
        Param("page")
    })
})
```

## fa-bolt gRPC 传输

```go
var _ = Service("users", func() {
    GRPC(func() {
        Path("/proto/users")
    })

    Method("create", func() {
        Payload(CreateUserPayload)
        Result(UserResult)

        GRPC(func() {
            Message(func() {
                Field(1, "name", String)
                Field(2, "email", String)
            })
            Response(CodeOK)
            Response("not_found", CodeNotFound)
        })
    })

    Method("stream", func() {
        Payload(func() {
            Field(1, "filter", String)
        })
        Result(StreamingResult(UserResult))

        GRPC(func() {
            Response(CodeOK)
        })
    })
})
```

## fa-wand-magic-sparkles Goa 生成命令

```bash
# 生成全部（HTTP + gRPC + 文档）
goa gen myapp/design

# 仅生成 HTTP
goa gen --http myapp/design

# 仅生成 gRPC
goa gen --grpc myapp/design

# 生成示例客户端/服务端
goa example myapp/design

# 常用 go generate 指令
# go:generate goa gen myapp/design
```

## fa-server 服务端搭建

```go
package main

import (
    "myapp/gen/users"
    "goa.design/goa/v3/http"
)

func main() {
    svc := &UserService{}
    endpoints := users.NewEndpoints(svc)

    mux := http.NewMuxer()
    usersServer := users.NewHTTPServer(endpoints, mux)
    usersServer.Mount(mux)

    httpsvc := http.NewServer(mux)
    httpsvc.Addr = ":8080"

    if err := httpsvc.ListenAndServe(); err != nil {
        log.Fatal(err)
    }
}
```

## fa-desktop 客户端生成

```go
// Goa 自动生成客户端代码

import "myapp/gen/client/users"

func main() {
    httpCl := http.NewClient("http://localhost:8080")
    cl := users.NewClient(httpCl)

    result, err := cl.List(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    user, err := cl.Show(context.Background(), &users.ShowPayload{ID: 1})
}
```

## fa-shield 中间件

```go
// 在设计 DSL 中
Method("admin", func() {
    Payload(func() {
        TokenField("token", String)
        Required("token")
    })

    HTTP(func() {
        GET("/admin")
        Header("token:Authorization")
    })
})

// 服务端中间件
mux := http.NewMuxer()
usersServer := users.NewHTTPServer(endpoints, mux)
usersServer.Mount(mux)

// 使用中间件包装
handler := http.RequestEncoder(
    middleware.Logging(mux),
)
```

## fa-check 数据验证

```go
var CreateUserPayload = Type("CreateUserPayload", func() {
    Field(1, "name", String, func() {
        MinLength(2)
        MaxLength(100)
        Pattern("^[a-zA-Z ]+$")
    })
    Field(2, "email", String, func() {
        Format(FormatEmail)
    })
    Field(3, "age", Int, func() {
        Minimum(0)
        Maximum(150)
    })
    Field(4, "role", String, func() {
        Enum("admin", "user", "guest")
    })
    Field(5, "tags", ArrayOf(String), func() {
        MinLength(1)
        MaxLength(10)
    })
    Required("name", "email")
})
```

## fa-file 文件流

```go
Method("download", func() {
    Result(func() {
        Field(1, "content", Bytes)
    })
    HTTP(func() {
        GET("/download/{id}")
        Response(StatusOK)
    })
})

Method("upload", func() {
    Payload(func() {
        Field(1, "file", Bytes)
        Field(2, "name", String)
    })
    HTTP(func() {
        POST("/upload")
        MultipartRequest()
        Response(StatusCreated)
    })
})

Method("events", func() {
    Result(StreamingResult(EventResult))
    HTTP(func() {
        GET("/events")
        Response(StatusOK)
    })
})
```
