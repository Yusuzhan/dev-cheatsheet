---
title: Go Logging (Logrus / Zap / Slog)
icon: fa-scroll
primary: "#00ADD8"
lang: go
---

## fa-play Slog Basic

```go
import "log/slog"

slog.Info("user logged in", "user_id", 42, "ip", "192.168.1.1")
slog.Warn("disk space low", "percent_free", 5.2)
slog.Error("connection failed", "err", err, "host", "db.example.com")
slog.Debug("cache hit", "key", "user:42")

slog.Info("request completed",
    "method", "GET",
    "path", "/api/users",
    "status", 200,
    "duration", time.Since(start),
)
```

## fa-list Slog Structured

```go
slog.Info("user created",
    slog.String("name", "Alice"),
    slog.Int("id", 42),
    slog.Float64("score", 98.5),
    slog.Bool("active", true),
    slog.Duration("elapsed", time.Since(start)),
    slog.Time("timestamp", time.Now()),
)

slog.LogAttrs(context.Background(), slog.LevelInfo, "order placed",
    slog.String("order_id", "ORD-123"),
    slog.Int("items", 3),
    slog.Float64("total", 99.99),
)
```

## fa-sliders Slog Handlers (Text/JSON)

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

slog.InfoContext(ctx, "processing request")

func WithRequestID(ctx context.Context, reqID string) context.Context {
    return context.WithValue(ctx, "request_id", reqID)
}
```

## fa-wrench Slog Custom Handler

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

## fa-play Logrus Basic

```go
import log "github.com/sirupsen/logrus"

log.SetFormatter(&log.JSONFormatter{})
log.SetLevel(log.InfoLevel)

log.Info("server started")
log.Warn("config missing, using defaults")
log.Error("failed to connect")

log.WithFields(log.Fields{
    "user_id": 42,
    "action":  "login",
}).Info("user action")

log.SetFormatter(&log.TextFormatter{
    FullTimestamp: true,
    ForceColors:   true,
})
```

## fa-link Logrus Hooks

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

## fa-list Logrus Fields

```go
entry := log.WithFields(log.Fields{
    "request_id": "req-123",
    "user_id":    42,
    "module":     "auth",
})

entry.Info("attempting login")
entry.Warn("rate limit approaching")
entry.Error("login failed")

log.WithField("service", "api").Info("health check")
log.WithTime(time.Now()).Info("timed event")

log.WithError(err).Error("operation failed")
log.WithError(err).WithField("query", sql).Fatal("db error")
```

## fa-play Zap Basic

```go
import "go.uber.org/zap"

logger, _ := zap.NewProduction()
defer logger.Sync()

logger.Info("user created",
    zap.String("name", "Alice"),
    zap.Int("id", 42),
    zap.Error(err),
)

logger.Warn("slow query",
    zap.String("query", "SELECT * FROM users"),
    zap.Duration("duration", 500*time.Millisecond),
)

logger.Error("connection lost",
    zap.String("host", "db.example.com"),
    zap.Int("port", 5432),
)
```

## fa-list Zap Structured

```go
logger, _ := zap.NewProduction()

logger.Info("request",
    zap.String("method", "GET"),
    zap.String("path", "/api/users"),
    zap.Int("status", 200),
    zap.Duration("latency", time.Since(start)),
    zap.Any("headers", req.Header),
)

logger.Info("metrics",
    zap.Int64("bytes_written", 2048),
    zap.Float64("cpu_usage", 0.75),
    zap.Bool("healthy", true),
    zap.Strings("tags", []string{"go", "web"}),
    zap.Ints("codes", []int{200, 201, 404}),
)
```

## fa-sliders Zap Encoder Config

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

sugar.Infow("request received",
    "method", "GET",
    "path", "/api/users",
    "latency", time.Since(start),
)

sugar.Infof("user %s logged in from %s", name, ip)
sugar.Errorw("db error", "err", err, "query", query)

sugar.With("request_id", "abc123").Infow("processing")
```

## fa-scale-balanced Choosing a Library

```
Slog (stdlib):
  - Built into Go 1.21+
  - Structured logging with typed attributes
  - Pluggable handlers
  - Zero dependencies
  - Best for: new projects, stdlib preference

Zap (uber-go):
  - Fastest structured logger
  - Type-safe fields, no reflection
  - Sugar logger for printf-style
  - Best for: high-performance services

Logrus:
  - Mature, widely adopted
  - Hooks system for extensibility
  - Slower than Zap/Slog
  - Best for: existing codebases, hook needs
```

## fa-code Common Patterns

```go
// Request logger middleware (Slog)
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        slog.Info("request",
            "method", r.Method,
            "path", r.URL.Path,
            "duration", time.Since(start),
        )
    })
}

// Error with stack trace (Zap)
logger.Error("panic recovered",
    zap.String("stack", string(debug.Stack())),
    zap.Any("request", reqInfo),
)

// Conditional debug logging
if logger.Core().Enabled(zap.DebugLevel) {
    logger.Debug("large payload", zap.Any("data", largeSlice))
}
```
