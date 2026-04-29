---
title: Go 日志 (Logrus / Zap / Slog)
icon: fa-scroll
primary: "#00ADD8"
lang: go
locale: zhs
---

## fa-play Slog 基础

```go
import "log/slog"

slog.Info("用户登录", "user_id", 42, "ip", "192.168.1.1")
slog.Warn("磁盘空间不足", "percent_free", 5.2)
slog.Error("连接失败", "err", err, "host", "db.example.com")
slog.Debug("缓存命中", "key", "user:42")

slog.Info("请求完成",
    "method", "GET",
    "path", "/api/users",
    "status", 200,
    "duration", time.Since(start),
)
```

## fa-list Slog 结构化

```go
slog.Info("用户创建",
    slog.String("name", "Alice"),
    slog.Int("id", 42),
    slog.Float64("score", 98.5),
    slog.Bool("active", true),
    slog.Duration("elapsed", time.Since(start)),
    slog.Time("timestamp", time.Now()),
)

slog.LogAttrs(context.Background(), slog.LevelInfo, "下单成功",
    slog.String("order_id", "ORD-123"),
    slog.Int("items", 3),
    slog.Float64("total", 99.99),
)
```

## fa-sliders Slog Handler (Text/JSON)

```go
handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelDebug,
})
logger := slog.New(handler)
slog.SetDefault(logger)

handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
    Level:     slog.LevelInfo,
    AddSource: true,
})
logger := slog.New(handler)
slog.SetDefault(logger)
```

## fa-link Slog Context

```go
ctx := context.WithValue(context.Background(), "request_id", "req-abc123")

handler := slog.NewJSONHandler(os.Stdout, nil)
handler = middleware.ExtractRequestID(handler)
slog.SetDefault(slog.New(handler))

slog.InfoContext(ctx, "处理请求")

func WithRequestID(ctx context.Context, reqID string) context.Context {
    return context.WithValue(ctx, "request_id", reqID)
}
```

## fa-wrench Slog 自定义 Handler

```go
type FilterHandler struct {
    handler slog.Handler
    level   slog.Level
}

func (h *FilterHandler) Enabled(ctx context.Context, level slog.Level) bool {
    return level >= h.level
}

func (h *FilterHandler) Handle(ctx context.Context, r slog.Record) error {
    return h.handler.Handle(ctx, r)
}

func (h *FilterHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
    return &FilterHandler{handler: h.handler.WithAttrs(attrs), level: h.level}
}

func (h *FilterHandler) WithGroup(name string) slog.Handler {
    return &FilterHandler{handler: h.handler.WithGroup(name), level: h.level}
}
```

## fa-play Logrus 基础

```go
import log "github.com/sirupsen/logrus"

log.SetFormatter(&log.JSONFormatter{})
log.SetLevel(log.InfoLevel)

log.Info("服务启动")
log.Warn("配置缺失，使用默认值")
log.Error("连接失败")

log.WithFields(log.Fields{
    "user_id": 42,
    "action":  "login",
}).Info("用户操作")

log.SetFormatter(&log.TextFormatter{
    FullTimestamp: true,
    ForceColors:   true,
})
```

## fa-link Logrus 钩子

```go
type SlackHook struct{}

func (h *SlackHook) Levels() []log.Level {
    return []log.Level{log.ErrorLevel, log.FatalLevel, log.PanicLevel}
}

func (h *SlackHook) Fire(entry *log.Entry) error {
    msg, _ := entry.String()
    return sendToSlack(msg)
}

log.AddHook(&SlackHook{})

hook, _ := lfsHook.NewHook(
    lfsHook.WriterMap{
        log.ErrorLevel: os.Stderr,
        log.InfoLevel:  logFile,
    },
    &log.JSONFormatter{},
)
log.AddHook(hook)
```

## fa-list Logrus 字段

```go
entry := log.WithFields(log.Fields{
    "request_id": "req-123",
    "user_id":    42,
    "module":     "auth",
})

entry.Info("尝试登录")
entry.Warn("接近速率限制")
entry.Error("登录失败")

log.WithField("service", "api").Info("健康检查")
log.WithTime(time.Now()).Info("定时事件")

log.WithError(err).Error("操作失败")
log.WithError(err).WithField("query", sql).Fatal("数据库错误")
```

## fa-play Zap 基础

```go
import "go.uber.org/zap"

logger, _ := zap.NewProduction()
defer logger.Sync()

logger.Info("用户创建",
    zap.String("name", "Alice"),
    zap.Int("id", 42),
    zap.Error(err),
)

logger.Warn("慢查询",
    zap.String("query", "SELECT * FROM users"),
    zap.Duration("duration", 500*time.Millisecond),
)

logger.Error("连接丢失",
    zap.String("host", "db.example.com"),
    zap.Int("port", 5432),
)
```

## fa-list Zap 结构化

```go
logger, _ := zap.NewProduction()

logger.Info("请求",
    zap.String("method", "GET"),
    zap.String("path", "/api/users"),
    zap.Int("status", 200),
    zap.Duration("latency", time.Since(start)),
    zap.Any("headers", req.Header),
)

logger.Info("指标",
    zap.Int64("bytes_written", 2048),
    zap.Float64("cpu_usage", 0.75),
    zap.Bool("healthy", true),
    zap.Strings("tags", []string{"go", "web"}),
    zap.Ints("codes", []int{200, 201, 404}),
)
```

## fa-sliders Zap 编码器配置

```go
cfg := zap.Config{
    Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
    Development: false,
    Encoding:    "json",
    EncoderConfig: zapcore.EncoderConfig{
        TimeKey:        "ts",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,
        EncodeTime:     zapcore.ISO8601TimeEncoder,
        EncodeDuration: zapcore.StringDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
    },
    OutputPaths:      []string{"stdout", "/var/log/app.log"},
    ErrorOutputPaths: []string{"stderr"},
}
logger, _ := cfg.Build()
```

## fa-candy-cane Zap Sugar Logger

```go
logger, _ := zap.NewProduction()
sugar := logger.Sugar()
defer sugar.Sync()

sugar.Infow("收到请求",
    "method", "GET",
    "path", "/api/users",
    "latency", time.Since(start),
)

sugar.Infof("用户 %s 从 %s 登录", name, ip)
sugar.Errorw("数据库错误", "err", err, "query", query)

sugar.With("request_id", "abc123").Infow("处理中")
```

## fa-scale-balanced 如何选择

```
Slog（标准库）：
  - Go 1.21+ 内置
  - 类型化属性的结构化日志
  - 可插拔 Handler
  - 零依赖
  - 适用于：新项目、偏好标准库

Zap (uber-go)：
  - 最快的结构化日志库
  - 类型安全字段，无反射
  - Sugar logger 支持 printf 风格
  - 适用于：高性能服务

Logrus：
  - 成熟，广泛使用
  - 钩子系统可扩展
  - 比 Zap/Slog 慢
  - 适用于：现有代码库、需要钩子
```

## fa-code 常用模式

```go
// 请求日志中间件（Slog）
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        slog.Info("请求",
            "method", r.Method,
            "path", r.URL.Path,
            "duration", time.Since(start),
        )
    })
}

// 带堆栈跟踪的错误（Zap）
logger.Error("panic 恢复",
    zap.String("stack", string(debug.Stack())),
    zap.Any("request", reqInfo),
)

// 条件调试日志
if logger.Core().Enabled(zap.DebugLevel) {
    logger.Debug("大量数据", zap.Any("data", largeSlice))
}
```
