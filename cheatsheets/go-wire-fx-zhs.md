---
title: Wire / fx
icon: fa-sitemap
primary: "#00ADD8"
lang: go
locale: zhs
---

## fa-sitemap Wire 概述

```go
// Wire：Go 的编译时依赖注入
// google.golang.org/wire
//
// 两个阶段：
//   1. 定义 Provider（如何创建每个类型）
//   2. 定义 Injector（如何将它们组装起来）
//
// 运行 `wire` 生成 wire_gen.go

//go:build wireinject

package main
```

## fa-cube Provider 函数

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

## fa-arrow-right Injector 函数

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

## fa-gears Wire 构建

```go
// wire.go（构建标签触发生成）

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

// 运行：wire ./...
// 生成 wire_gen.go，包含实际实现
```

## fa-link 绑定接口

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

## fa-layer-group 结构体 Provider

```go
type App struct {
    DB      *sql.DB
    Service *UserService
    Logger  *zap.Logger
}

func NewApp(db *sql.DB, svc *UserService, log *zap.Logger) *App {
    return &App{DB: db, Service: svc, Logger: log}
}

// wire.Struct 自动注入字段
var AppSet = wire.NewSet(
    wire.Struct(new(App), "*"),
)

// 选择性字段
var AppSet2 = wire.NewSet(
    wire.Struct(new(App), "DB", "Service"),
)
```

## fa-boxes Wire 集合

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

## fa-broom 清理函数

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

// Wire 自动链式组合清理函数
// 返回的 func() 按逆序调用所有清理函数
```

## fa-circle-play fx 应用生命周期

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

## fa-sliders fx 选项

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

// 模块模式
var Module = fx.Options(
    fx.Provide(NewUserService, NewUserRepo),
    fx.Invoke(RegisterUserRoutes),
)
```

## fa-terminal fx 日志

```go
fx.New(
    fx.WithLogger(func() fxevent.Logger {
        return &fxevent.ConsoleLogger{W: os.Stderr}
    }),
)

// 使用 Zap 的自定义日志
fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
    return &fxevent.ZapLogger{Logger: log}
})

// 测试中静默 fx 日志
fx.WithLogger(func() fxevent.Logger {
    return fxevent.NopLogger
})
```

## fa-power-off fx 关闭

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

// 基于信号的关闭（fx.New().Run() 默认行为）
// fx.New().Run() 自动处理 SIGINT 和 SIGTERM
```

## fa-scale-balanced Wire vs fx 对比

```
Wire:
  - 编译时 DI，构建时捕获错误
  - 生成 Go 代码（wire_gen.go）
  - 无反射，无运行时魔法
  - 需要 //go:build wireinject 标签
  - 适用于：性能关键、编译时安全

fx:
  - 基于反射的运行时 DI
  - 生命周期管理（OnStart/OnStop 钩子）
  - 无代码生成步骤
  - 丰富的日志和可视化
  - 适用于：长时间运行的服务、需要生命周期钩子的应用
```
