---
title: Wire / fx
icon: fa-sitemap
primary: "#00ADD8"
lang: go
---

## fa-sitemap Wire Overview

```go
// Wire: compile-time dependency injection for Go
// google.golang.org/wire
//
// Two phases:
//   1. Define providers (how to create each type)
//   2. Define injectors (how to wire them together)
//
// Run `wire` to generate wire_gen.go

//go:build wireinject

package main
```

## fa-cube Provider Functions

```go
func NewDB(cfg *Config) *sql.DB {
    db, _ := sql.Open("postgres", cfg.DSN)
    return db
}

func NewUserRepo(db *sql.DB) *UserRepo {
    return &UserRepo{db: db}
}

func NewUserService(repo *UserRepo) *UserService {
    return &UserService{repo: repo}
}

func NewLogger() *zap.Logger {
    logger, _ := zap.NewProduction()
    return logger
}
```

## fa-arrow-right Injector Functions

```go
//go:build wireinject

func InitializeApp(cfg *Config) (*App, error) {
    wire.Build(
        NewDB,
        NewUserRepo,
        NewUserService,
        NewLogger,
        NewApp,
    )
    return nil, nil
}
```

## fa-gears Wire Build

```go
// wire.go (build tag triggers generation)

//go:build wireinject

package main

import "github.com/google/wire"

func InitializeApp(cfg *Config) (*App, func(), error) {
    wire.Build(
        NewDB,
        NewUserRepo,
        NewUserService,
        NewLogger,
        NewApp,
    )
    return nil, nil, nil
}

// Run: wire ./...
// Produces wire_gen.go with actual implementation
```

## fa-link Bind Interfaces

```go
type UserRepository interface {
    GetByID(id int) (*User, error)
    Create(user *User) error
}

type postgresUserRepo struct {
    db *sql.DB
}

var UserRepoSet = wire.NewSet(
    NewPostgresUserRepo,
    wire.Bind(new(UserRepository), new(*postgresUserRepo)),
)
```

## fa-layer-group Struct Providers

```go
type App struct {
    DB      *sql.DB
    Service *UserService
    Logger  *zap.Logger
}

func NewApp(db *sql.DB, svc *UserService, log *zap.Logger) *App {
    return &App{DB: db, Service: svc, Logger: log}
}

// wire.Struct can inject fields automatically
var AppSet = wire.NewSet(
    wire.Struct(new(App), "*"),
)

// selective fields
var AppSet2 = wire.NewSet(
    wire.Struct(new(App), "DB", "Service"),
)
```

## fa-boxes Wire Sets

```go
var RepoSet = wire.NewSet(
    NewUserRepo,
    NewOrderRepo,
    NewProductRepo,
)

var ServiceSet = wire.NewSet(
    NewUserService,
    NewOrderService,
)

var InfraSet = wire.NewSet(
    NewDB,
    NewLogger,
    NewCache,
)

func InitializeApp(cfg *Config) (*App, error) {
    wire.Build(
        InfraSet,
        RepoSet,
        ServiceSet,
        NewApp,
    )
    return nil, nil
}
```

## fa-broom Cleanup Functions

```go
func NewDB(cfg *Config) (*sql.DB, func(), error) {
    db, err := sql.Open("postgres", cfg.DSN)
    if err != nil {
        return nil, nil, err
    }
    cleanup := func() {
        db.Close()
    }
    return db, cleanup, nil
}

func NewRedis(addr string) (*redis.Client, func(), error) {
    rdb := redis.NewClient(&redis.Options{Addr: addr})
    cleanup := func() { rdb.Close() }
    return rdb, cleanup, nil
}

// Wire chains cleanup functions automatically
// The returned func() calls all cleanup functions in reverse order
```

## fa-circle-play fx App Lifecycle

```go
func main() {
    fx.New(
        fx.Provide(
            NewDB,
            NewLogger,
            NewUserService,
        ),
        fx.Invoke(StartServer),
        fx.WithLogger(func() fxevent.Logger {
            return &fxevent.ConsoleLogger{W: os.Stderr}
        }),
    ).Run()
}
```

## fa-hand-holding-heart fx Provide / Invoke

```go
fx.New(
    fx.Provide(
        NewConfig,
        NewDB,
        NewUserRepo,
        NewUserService,
        NewHTTPServer,
    ),
    fx.Invoke(RegisterRoutes),
)

func NewHTTPServer(lc fx.Lifecycle, log *zap.Logger) *http.Server {
    srv := &http.Server{Addr: ":8080"}
    lc.Append(fx.Hook{
        OnStart: func(ctx context.Context) error {
            log.Info("starting server", zap.String("addr", srv.Addr))
            go srv.ListenAndServe()
            return nil
        },
        OnStop: func(ctx context.Context) error {
            log.Info("stopping server")
            return srv.Shutdown(ctx)
        },
    })
    return srv
}
```

## fa-sliders fx Options

```go
fx.New(
    fx.Options(
        fx.Provide(NewDB, NewLogger),
        fx.Invoke(StartServer),
    ),
    fx.Invoke(func(lc fx.Lifecycle) {
        lc.Append(fx.Hook{
            OnStart: func(ctx context.Context) error {
                fmt.Println("app started")
                return nil
            },
            OnStop: func(ctx context.Context) error {
                fmt.Println("app stopped")
                return nil
            },
        })
    }),
)

// Module pattern
var Module = fx.Options(
    fx.Provide(NewUserService, NewUserRepo),
    fx.Invoke(RegisterUserRoutes),
)
```

## fa-terminal fx Logging

```go
fx.New(
    fx.WithLogger(func() fxevent.Logger {
        return &fxevent.ConsoleLogger{W: os.Stderr}
    }),
)

// Custom logger using Zap
fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
    return &fxevent.ZapLogger{Logger: log}
})

// Suppress fx logs in tests
fx.WithLogger(func() fxevent.Logger {
    return fxevent.NopLogger
})
```

## fa-power-off fx Shutdown

```go
func main() {
    app := fx.New(
        fx.Provide(NewDB, NewHTTPServer),
        fx.Invoke(StartServer),
    )

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if err := app.Stop(ctx); err != nil {
        log.Fatal(err)
    }
}

// Signal-based shutdown (default with fx.New().Run())
// fx.New().Run() handles SIGINT and SIGTERM automatically
```

## fa-scale-balanced Comparison Wire vs fx

```
Wire:
  - Compile-time DI, errors caught at build time
  - Generates Go code (wire_gen.go)
  - No reflection, no runtime magic
  - Requires //go:build wireinject tag
  - Best for: performance-critical, compile-time safety

fx:
  - Runtime DI using reflection
  - Lifecycle management (OnStart/OnStop hooks)
  - No code generation step
  - Rich logging and visualization
  - Best for: long-running services, apps needing lifecycle hooks
```
