---
title: GitHub Actions
icon: fa-gear
primary: "#2088FF"
lang: yaml
locale: zhs
---

## fa-file-code Workflow 基础

```yaml
# .github/workflows/ci.yml
name: CI
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: 20
      - run: npm ci
      - run: npm test
```

工作流文件位于 `.github/workflows/`。每个文件定义一个工作流，包含触发器、任务和步骤。

## fa-bolt 触发器 (on)

```yaml
# 推送和拉取请求
on:
  push:
    branches: [main, develop]
    tags: ["v*"]
  pull_request:
    branches: [main]

# 定时任务 (cron)
on:
  schedule:
    - cron: "0 2 * * *"   # 每天 02:00 UTC

# 手动触发
on:
  workflow_dispatch:
    inputs:
      environment:
        description: "部署目标"
        default: staging
        type: string
```

```yaml
# Issue 或 Discussion 触发
on:
  issues:
    types: [opened, labeled]
  discussion:
    types: [created]

# Release 触发
on:
  release:
    types: [published]

# 被其他工作流调用
on:
  workflow_call:
    inputs:
      ref:
        type: string
        default: main
```

## fa-list-check Jobs 与 Steps

```yaml
jobs:
  test:
    runs-on: ubuntu-latest
    timeout-minutes: 30
    defaults:
      run:
        working-directory: ./app
    steps:
      - uses: actions/checkout@v4
      - name: 安装依赖
        run: npm ci
      - name: 运行测试
        run: npm test
      - name: 代码检查
        run: npm run lint

  build:
    needs: test
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: npm run build
```

`needs` 定义任务依赖。没有 `needs` 的任务并行执行。

## fa-server Runners 与 Matrix

```yaml
jobs:
  test:
    strategy:
      matrix:
        node-version: [18, 20, 22]
        os: [ubuntu-latest, macos-latest]
      fail-fast: false
    runs-on: ${{ matrix.os }}
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: ${{ matrix.node-version }}
      - run: npm ci && npm test
```

```yaml
# 自托管 Runner
runs-on: [self-hosted, linux, x64]

# 包含/排除特定组合
strategy:
  matrix:
    python: ["3.10", "3.11", "3.12"]
    include:
      - python: "3.12"
        experimental: true
    exclude:
      - python: "3.10"
```

## fa-key 环境变量与 Secrets

```yaml
env:
  NODE_ENV: ci
  API_URL: https://api.example.com

jobs:
  deploy:
    runs-on: ubuntu-latest
    environment: production
    env:
      DEPLOY_TARGET: production
    steps:
      - name: 部署
        env:
          API_KEY: ${{ secrets.API_KEY }}
          DB_URL: ${{ secrets.DATABASE_URL }}
        run: |
          echo "Deploying to $DEPLOY_TARGET"
          ./deploy.sh
```

```yaml
# 在仓库 Settings → Secrets and variables → Actions 中定义
# 在工作流中访问：
${{ secrets.MY_SECRET }}
${{ vars.MY_VARIABLE }}

# GitHub 内置环境变量（无需 secrets 前缀）
# $GITHUB_SHA, $GITHUB_REF, $GITHUB_REPOSITORY
# $GITHUB_RUN_ID, $GITHUB_RUN_NUMBER
# $GITHUB_EVENT_NAME, $GITHUB_WORKSPACE
```

## fa-puzzle-piece Actions (uses)

```yaml
steps:
  # 官方 Actions
  - uses: actions/checkout@v4
  - uses: actions/setup-node@v4
    with:
      node-version: 20
      cache: npm

  # 社区 Actions
  - uses: docker/login-action@v3
    with:
      registry: ghcr.io
      username: ${{ github.actor }}
      password: ${{ secrets.GITHUB_TOKEN }}

  # 本地 Action
  - uses: ./.github/actions/my-action
    with:
      param: value
```

```yaml
# 使用 commit SHA 引用（最安全）
- uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11

# 使用标签引用
- uses: actions/checkout@v4

# 使用分支引用
- uses: actions/checkout@main
```

## fa-download Checkout 与环境配置

```yaml
steps:
  - uses: actions/checkout@v4
    with:
      fetch-depth: 0          # 完整历史（用于版本管理）
      ref: develop            # 检出指定分支/标签
      submodules: true        # 包含子模块
      token: ${{ secrets.PAT }}

  - uses: actions/setup-go@v5
    with:
      go-version: "1.22"

  - uses: actions/setup-python@v5
    with:
      python-version: "3.12"
      cache: pip

  - uses: actions/setup-java@v4
    with:
      distribution: temurin
      java-version: "21"
```

```yaml
# 使用服务容器（类似 docker-compose）
services:
  postgres:
    image: postgres:16
    env:
      POSTGRES_PASSWORD: postgres
    ports:
      - 5432:5432
    options: >-
      --health-cmd pg_isready
      --health-interval 10s
      --health-timeout 5s
      --health-retries 5
```

## fa-hammer 构建与测试

```yaml
jobs:
  build-and-test:
    runs-on: ubuntu-latest
    services:
      redis:
        image: redis:7-alpine
        ports:
          - 6379:6379
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: 20
          cache: npm
      - run: npm ci
      - run: npm run lint
      - run: npm run typecheck
      - run: npm test -- --coverage
      - uses: codecov/codecov-action@v4
        with:
          token: ${{ secrets.CODECOV_TOKEN }}
```

## fa-docker Docker 操作

```yaml
jobs:
  docker:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: docker/setup-buildx-action@v3

      - uses: docker/login-action@v3
        with:
          registry: ghcr.io
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

      - uses: docker/build-push-action@v6
        with:
          context: .
          push: true
          tags: |
            ghcr.io/${{ github.repository }}:latest
            ghcr.io/${{ github.repository }}:${{ github.sha }}
          cache-from: type=gha
          cache-to: type=gha,mode=max
```

## fa-box-archive 缓存

```yaml
# 内置缓存 Action
- uses: actions/cache@v4
  with:
    path: |
      ~/.npm
      ${{ github.workspace }}/.next/cache
    key: ${{ runner.os }}-node-${{ hashFiles('**/package-lock.json') }}
    restore-keys: |
      ${{ runner.os }}-node-

# 通过 setup Actions 启用工具专属缓存
- uses: actions/setup-go@v5
  with:
    go-version: "1.22"
    cache: true
```

```yaml
# 缓存 Docker 层
- uses: actions/cache@v4
  with:
    path: /tmp/.buildx-cache
    key: ${{ runner.os }}-buildx-${{ github.sha }}
    restore-keys: |
      ${{ runner.os }}-buildx-
```

## fa-box Artifacts

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: npm ci && npm run build

      - uses: actions/upload-artifact@v4
        with:
          name: dist
          path: dist/
          retention-days: 5

  deploy:
    needs: build
    runs-on: ubuntu-latest
    steps:
      - uses: actions/download-artifact@v4
        with:
          name: dist
          path: dist/
      - run: ./deploy.sh
```

## fa-rocket 部署

```yaml
jobs:
  deploy:
    runs-on: ubuntu-latest
    needs: build
    if: github.ref == 'refs/heads/main'
    environment:
      name: production
      url: https://app.example.com
    steps:
      - uses: actions/checkout@v4
      - uses: actions/download-artifact@v4
        with:
          name: dist
      - name: 部署到服务器
        env:
          SSH_KEY: ${{ secrets.SSH_PRIVATE_KEY }}
        run: |
          mkdir -p ~/.ssh
          echo "$SSH_KEY" > ~/.ssh/deploy_key
          chmod 600 ~/.ssh/deploy_key
          rsync -avz -e "ssh -i ~/.ssh/deploy_key -o StrictHostKeyChecking=no" \
            dist/ deploy@server:/var/www/app/
```

## fa-code-branch 条件步骤

```yaml
steps:
  - name: 仅在 main 分支运行
    if: github.ref == 'refs/heads/main'
    run: echo "on main branch"

  - name: 仅在 PR 中运行
    if: github.event_name == 'pull_request'
    run: echo "this is a PR"

  - name: 前序步骤成功时运行
    if: success()
    run: echo "previous steps passed"

  - name: 前序步骤失败时运行
    if: failure()
    run: echo "something failed"

  - name: 始终运行
    if: always()
    run: echo "runs no matter what"
```

```yaml
# 任务级别条件
jobs:
  nightly:
    if: github.event_name == 'schedule'
    runs-on: ubuntu-latest
    steps:
      - run: echo "scheduled job"

  deploy:
    needs: test
    if: ${{ needs.test.result == 'success' }}
    runs-on: ubuntu-latest
    steps:
      - run: echo "tests passed, deploying"
```

## fa-recycle 可复用工作流

```yaml
# .github/workflows/reusable-test.yml
on:
  workflow_call:
    inputs:
      node-version:
        required: true
        type: string
    secrets:
      api-key:
        required: true

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with:
          node-version: ${{ inputs.node-version }}
      - run: npm ci && npm test
        env:
          API_KEY: ${{ secrets.api-key }}
```

```yaml
# .github/workflows/ci.yml — 调用方
jobs:
  test:
    uses: ./.github/workflows/reusable-test.yml
    with:
      node-version: "20"
    secrets:
      api-key: ${{ secrets.API_KEY }}
```
