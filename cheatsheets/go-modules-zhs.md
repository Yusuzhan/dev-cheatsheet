---
title: Go 模块与工具链
icon: fa-toolbox
primary: "#00ADD8"
lang: bash
locale: zhs
---

## fa-box 初始化与基础命令

```bash
go mod init github.com/user/project    # 初始化模块
go mod tidy                             # 整理依赖
go mod vendor                           # 复制依赖到 vendor
go get github.com/gin-gonic/gin@latest  # 添加最新依赖
go get github.com/gin-gonic/gin@v1.9.0  # 指定版本
go install golang.org/x/tools/cmd/stringer@latest  # 安装工具
```

## fa-pen-to-square go.mod 文件

```
module github.com/user/project

go 1.22

require (
    github.com/gin-gonic/gin v1.9.0
    golang.org/x/text v0.14.0
)

require (
    github.com/json-iterator/go v1.1.12 // indirect  # 间接依赖
)

replace github.com/broken/pkg => ../fixed-pkg         # 替换依赖

exclude github.com/broken/pkg v1.0.0                  # 排除版本

retract [v1.0.0, v1.0.5]                              # 撤回版本
```

## fa-arrows-left-right 依赖管理

```bash
go get ./...                          # 安装所有依赖
go get -u ./...                       # 更新到最新版本
go get -u=patch ./...                 # 更新到最新补丁版本
go get -t ./...                       # 包含测试依赖
go mod tidy                           # 移除未使用、添加缺失依赖
go mod verify                         # 校验下载校验和
go list -m all                        # 列出所有依赖
go list -m -versions github.com/gin-gonic/gin  # 列出可用版本
go list -m -json github.com/gin-gonic/gin      # JSON 格式输出
```

## fa-code-branch Workspace 工作区

```bash
go work init ./app ./lib               # 初始化工作区
go work use ./app ./lib                # 添加模块
go work edit -dropuse=./old            # 移除模块
```

```
go.work

go 1.22

use (
    ./app
    ./lib
)

replace github.com/user/lib => ./lib   # 本地替换
```

## fa-wrench 构建与运行

```bash
go build -o bin/app ./cmd/app          # 指定输出路径
go build -race ./...                   # 启用竞态检测
go build -ldflags="-s -w -X main.version=1.0.0"  # 去除调试信息、注入变量
go build -trimpath ./...               # 去除本地路径信息
go run ./cmd/app                       # 编译并运行
go run -race main.go                   # 带竞态检测运行
GOOS=linux GOARCH=amd64 go build -o app ./cmd/app  # 交叉编译
```

## fa-vial 测试

```bash
go test ./...                          # 运行所有测试
go test -v ./...                       # 详细输出
go test -run TestFunc ./pkg            # 运行匹配的测试
go test -run TestSuite/TestCase ./pkg  # 运行子测试
go test -race ./...                    # 竞态检测
go test -cover ./...                   # 覆盖率
go test -coverprofile=coverage.out ./...  # 生成覆盖率报告
go tool cover -html=coverage.out       # 浏览器查看覆盖率
go test -bench=. -benchmem ./...       # 基准测试
go test -fuzz=FuzzFunc ./pkg           # 模糊测试
```

## fa-broom 代码检查

```bash
go vet ./...                           # 内置静态分析
go vet -structtag ./...                # 检查结构体标签
staticcheck ./...                      # staticcheck 检查
golangci-lint run                      # golangci-lint
golangci-lint run --enable-all         # 启用所有 linter
golangci-lint run ./pkg/...            # 指定包
```

## fa-file-code 生成与格式化

```bash
go generate ./...                      # 执行所有 go:generate
stringer -type=Pill                    # 生成 String 方法
goimports -w .                         # 管理导入并格式化
gofmt -w .                             # 格式化代码
gofmt -d .                             # 显示格式差异
goimports -local github.com/user/project -w .  # 分组导入
```

## fa-magnifying-glass 文档与检查

```bash
go doc fmt.Println                     # 查看文档
go doc -all fmt                        # 显示所有导出
go doc -src sync.Map                   # 显示源码
go doc github.com/user/project/pkg.Func  # 第三方包文档
go fix ./...                           # 自动修复旧语法
go tool dist list                      # 列出所有 GOOS/GOARCH
go version -m ./bin/app                # 查看构建信息
```

## fa-globe 交叉编译

```bash
GOOS=linux   GOARCH=amd64  go build -o app-linux-amd64  ./cmd/app
GOOS=darwin  GOARCH=arm64  go build -o app-darwin-arm64  ./cmd/app
GOOS=windows GOARCH=amd64  go build -o app-windows.exe   ./cmd/app

# 常用目标平台
# GOOS: linux, darwin, windows, freebsd, js
# GOARCH: amd64, arm64, arm, 386, wasm

CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app  # 禁用 CGO 静态编译
```

## fa-code 构建标签与指令

```go
// go:build linux && amd64           // 构建约束
//go:generate stringer -type=Status  // 代码生成
//go:embed static/*                  // 嵌入文件
//go:build !js                       // 排除平台

package main

import _ "github.com/lib/pq"         // 匿名导入

//go:embed static/*
var staticFiles embed.FS
```

## fa-layer-group 私有模块

```bash
# ~/.gitconfig 或环境变量
git config --global url."git@github.com:".insteadOf "https://github.com/"

# 设置私有模块
export GOPRIVATE=github.com/myorg/*
export GONOSUMCHECK=github.com/myorg/*
export GOFLAGS=-insecure

# netrc 配置代理认证
# ~/.netrc
# machine github.com login user password ghp_xxx
```

## fa-puzzle-piece 模块发布

```bash
git tag v1.0.0                        # 打标签
git push origin v1.0.0                # 推送标签

# 预发布版本
git tag v1.0.0-rc.1
git push origin v1.0.0-rc.1

# 在 go.mod 中撤回有问题的版本
# retract v1.0.0

# v2+ 主版本需要 /v2 后缀
go mod init github.com/user/project/v2

# 查看可用版本
go list -m -versions github.com/user/project
```

## fa-gauge 构建信息注入

```bash
# 构建时注入版本信息
go build -ldflags="-X main.version=$(git describe --tags) \
-X main.commit=$(git rev-parse --short HEAD) \
-X main.date=$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
-o app ./cmd/app
```

```go
var (
    version = "dev"
    commit  = "none"
    date    = "unknown"
)

func main() {
    fmt.Printf("version: %s, commit: %s, built: %s\n", version, commit, date)
}
```
