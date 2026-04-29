---
title: Go Modules & Toolchain
icon: fa-toolbox
primary: "#00ADD8"
lang: bash
---

## fa-box init & basic

```bash
go mod init github.com/user/project
go mod tidy
go mod vendor
go get github.com/gin-gonic/gin@latest
go get github.com/gin-gonic/gin@v1.9.0
go install golang.org/x/tools/cmd/stringer@latest
```

## fa-pen-to-square go.mod

```
module github.com/user/project

go 1.22

require (
    github.com/gin-gonic/gin v1.9.0
    golang.org/x/text v0.14.0
)

require (
    github.com/json-iterator/go v1.1.12 // indirect
)

replace github.com/broken/pkg => ../fixed-pkg

exclude github.com/broken/pkg v1.0.0

retract [v1.0.0, v1.0.5]
```

## fa-arrows-left-right Dependencies

```bash
go get ./...                          # install all deps
go get -u ./...                       # update all deps to latest
go get -u=patch ./...                 # update to latest patch
go get -t ./...                       # include test deps
go mod tidy                           # remove unused, add missing
go mod verify                         # verify downloaded checksums
go list -m all                        # list all deps
go list -m -versions github.com/gin-gonic/gin
go list -m -json github.com/gin-gonic/gin
```

## fa-code-branch Workspaces

```bash
go work init ./app ./lib
go work use ./app ./lib
go work edit -dropuse=./old
```

```
go.work

go 1.22

use (
    ./app
    ./lib
)

replace github.com/user/lib => ./lib
```

## fa-wrench Build & Run

```bash
go build -o bin/app ./cmd/app
go build -race ./...
go build -ldflags="-s -w -X main.version=1.0.0"
go build -trimpath ./...
go run ./cmd/app
go run -race main.go
GOOS=linux GOARCH=amd64 go build -o app ./cmd/app
```

## fa-vial Test

```bash
go test ./...
go test -v ./...
go test -run TestFunc ./pkg
go test -run TestSuite/TestCase ./pkg
go test -race ./...
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
go test -bench=. -benchmem ./...
go test -fuzz=FuzzFunc ./pkg
```

## fa-broom Vet & Lint

```bash
go vet ./...
go vet -structtag ./...
staticcheck ./...
golangci-lint run
golangci-lint run --enable-all
golangci-lint run ./pkg/...
```

## fa-file-code Generate & Format

```bash
go generate ./...
stringer -type=Pill
goimports -w .
gofmt -w .
gofmt -d .                        # show diff
goimports -local github.com/user/project -w .
```

## fa-magnifying-glass Doc & Inspect

```bash
go doc fmt.Println
go doc -all fmt
go doc -src sync.Map
go doc github.com/user/project/pkg.Func
go fix ./...
go tool dist list                    # list all GOOS/GOARCH
go version -m ./bin/app              # print build info
```

## fa-globe Cross Compilation

```bash
GOOS=linux   GOARCH=amd64  go build -o app-linux-amd64  ./cmd/app
GOOS=darwin  GOARCH=arm64  go build -o app-darwin-arm64  ./cmd/app
GOOS=windows GOARCH=amd64  go build -o app-windows.exe   ./cmd/app

# common targets
# GOOS: linux, darwin, windows, freebsd, js
# GOARCH: amd64, arm64, arm, 386, wasm

CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/app
```

## fa-code version directives

```go
// go:build linux && amd64
//go:generate stringer -type=Status

package main

import _ "github.com/lib/pq"

//go:embed static/*
var staticFiles embed.FS

//go:build !js
```

## fa-layer-group Private Modules

```bash
# ~/.gitconfig or env
git config --global url."git@github.com:".insteadOf "https://github.com/"

# GOPRIVATE
export GOPRIVATE=github.com/myorg/*
export GONOSUMCHECK=github.com/myorg/*
export GOFLAGS=-insecure

# netrc for private proxy
# ~/.netrc
# machine github.com login user password ghp_xxx
```

## fa-puzzle-piece Module Release

```bash
git tag v1.0.0
git push origin v1.0.0

# pre-release
git tag v1.0.0-rc.1
git push origin v1.0.0-rc.1

# retract broken version in go.mod
# retract v1.0.0

# major version v2+ requires /v2 suffix
go mod init github.com/user/project/v2

# list available versions
go list -m -versions github.com/user/project
```

## fa-gauge Build Info

```bash
# embed version at build time
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
