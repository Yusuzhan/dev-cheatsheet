---
title: Goa
icon: fa-drafting-compass
primary: "#00ADD8"
lang: go
---

## fa-drafting-compass Design Overview

```go
// goa.design/goa — API-first design framework
// Write DSL → Generate code (server, client, docs)
//
// Project layout:
//   design/   — API DSL definitions
//   gen/      — generated code (do not edit)
//   cmd/      — service entry points

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

## fa-cubes Service Definition

```go
var _ = Service("users", func() {
    Description("User management service")

    Error("not_found", NotFound, "User not found")
    Error("unauthorized", String, "Unauthorized")

    HTTP(func() {
        Path("/users")
    })

    GRPC(func() {
        // gRPC transport mapping
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

## fa-code Method Definition

```go
Method("create", func() {
    Description("Create a new user")
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

## fa-database Payload & Result

```go
var CreateUserPayload = Type("CreateUserPayload", func() {
    Field(1, "name", String, "User name", func() {
        MinLength(2)
        MaxLength(100)
    })
    Field(2, "email", String, "Email address", func() {
        Format(FormatEmail)
    })
    Field(3, "age", Int, "Age", func() {
        Minimum(0)
        Maximum(150)
    })
    Required("name", "email")
})

var UserResult = ResultType("application/vnd.user", func() {
    Field(1, "id", UInt, "User ID")
    Field(2, "name", String, "User name")
    Field(3, "email", String, "Email")
    Required("id", "name")
})
```

## fa-triangle-exclamation Error Definitions

```go
var NotFound = Type("NotFound", func() {
    Field(1, "message", String, "Error message")
    Field(2, "id", UInt, "Resource ID")
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

## fa-network-wired HTTP Transport

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

## fa-bolt gRPC Transport

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

## fa-wand-magic-sparkles Goa gen Command

```bash
# Generate all (HTTP + gRPC + docs)
goa gen myapp/design

# Generate only HTTP
goa gen --http myapp/design

# Generate only gRPC
goa gen --grpc myapp/design

# Generate example client/server
goa example myapp/design

# Typical go generate directive
# go:generate goa gen myapp/design
```

## fa-server Server Setup

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

## fa-desktop Client Generation

```go
// Goa generates client code automatically

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

## fa-shield Middleware

```go
// In the design DSL
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

// Server-side middleware
mux := http.NewMuxer()
usersServer := users.NewHTTPServer(endpoints, mux)
usersServer.Mount(mux)

// Wrap with middleware
handler := http.RequestEncoder(
    middleware.Logging(mux),
)
```

## fa-check Validation

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

## fa-file File Streaming

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
