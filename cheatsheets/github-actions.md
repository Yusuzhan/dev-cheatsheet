---
title: GitHub Actions
icon: fa-gear
primary: "#2088FF"
lang: yaml
---

## fa-file-code Workflow Basics

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

Workflow files live in `.github/workflows/`. Each file defines one workflow with triggers, jobs, and steps.

## fa-bolt Triggers (on)

```yaml
# push and pull request
on:
  push:
    branches: [main, develop]
    tags: ["v*"]
  pull_request:
    branches: [main]

# schedule (cron)
on:
  schedule:
    - cron: "0 2 * * *"   # daily at 02:00 UTC

# manual trigger
on:
  workflow_dispatch:
    inputs:
      environment:
        description: "Deploy target"
        default: staging
        type: string
```

```yaml
# trigger on issue or discussion
on:
  issues:
    types: [opened, labeled]
  discussion:
    types: [created]

# trigger on release
on:
  release:
    types: [published]

# trigger on other workflows
on:
  workflow_call:
    inputs:
      ref:
        type: string
        default: main
```

## fa-list-check Jobs & Steps

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
      - name: Install deps
        run: npm ci
      - name: Run tests
        run: npm test
      - name: Run lint
        run: npm run lint

  build:
    needs: test
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: npm run build
```

`needs` defines job dependencies. Jobs without `needs` run in parallel.

## fa-server Runners & Matrix

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
# self-hosted runner
runs-on: [self-hosted, linux, x64]

# include/exclude specific combinations
strategy:
  matrix:
    python: ["3.10", "3.11", "3.12"]
    include:
      - python: "3.12"
        experimental: true
    exclude:
      - python: "3.10"
```

## fa-key Environment Variables & Secrets

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
      - name: Deploy
        env:
          API_KEY: ${{ secrets.API_KEY }}
          DB_URL: ${{ secrets.DATABASE_URL }}
        run: |
          echo "Deploying to $DEPLOY_TARGET"
          ./deploy.sh
```

```yaml
# define secrets in repo Settings → Secrets and variables → Actions
# access in workflow:
${{ secrets.MY_SECRET }}
${{ vars.MY_VARIABLE }}

# GitHub-provided env vars (no secrets prefix needed)
# $GITHUB_SHA, $GITHUB_REF, $GITHUB_REPOSITORY
# $GITHUB_RUN_ID, $GITHUB_RUN_NUMBER
# $GITHUB_EVENT_NAME, $GITHUB_WORKSPACE
```

## fa-puzzle-piece Actions (uses)

```yaml
steps:
  # official actions
  - uses: actions/checkout@v4
  - uses: actions/setup-node@v4
    with:
      node-version: 20
      cache: npm

  # community actions
  - uses: docker/login-action@v3
    with:
      registry: ghcr.io
      username: ${{ github.actor }}
      password: ${{ secrets.GITHUB_TOKEN }}

  # local action
  - uses: ./.github/actions/my-action
    with:
      param: value
```

```yaml
# reference by commit SHA (most secure)
- uses: actions/checkout@b4ffde65f46336ab88eb53be808477a3936bae11

# reference by tag
- uses: actions/checkout@v4

# reference by branch
- uses: actions/checkout@main
```

## fa-download Checkout & Setup

```yaml
steps:
  - uses: actions/checkout@v4
    with:
      fetch-depth: 0          # full history for versioning
      ref: develop            # checkout specific branch/tag
      submodules: true        # include submodules
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
# cache multiple dependency managers
- uses: actions/setup-node@v4
  with:
    node-version: 20
    cache: npm

# use service containers (like docker-compose)
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

## fa-hammer Build & Test

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

## fa-docker Docker Actions

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

## fa-box-archive Caching

```yaml
# built-in cache action
- uses: actions/cache@v4
  with:
    path: |
      ~/.npm
      ${{ github.workspace }}/.next/cache
    key: ${{ runner.os }}-node-${{ hashFiles('**/package-lock.json') }}
    restore-keys: |
      ${{ runner.os }}-node-

# tool-specific cache via setup actions
- uses: actions/setup-go@v5
  with:
    go-version: "1.22"
    cache: true
```

```yaml
# cache Docker layers
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

## fa-rocket Deploy

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
      - name: Deploy to server
        env:
          SSH_KEY: ${{ secrets.SSH_PRIVATE_KEY }}
        run: |
          mkdir -p ~/.ssh
          echo "$SSH_KEY" > ~/.ssh/deploy_key
          chmod 600 ~/.ssh/deploy_key
          rsync -avz -e "ssh -i ~/.ssh/deploy_key -o StrictHostKeyChecking=no" \
            dist/ deploy@server:/var/www/app/
```

## fa-code-branch Conditional Steps

```yaml
steps:
  - name: Run only on main
    if: github.ref == 'refs/heads/main'
    run: echo "on main branch"

  - name: Run only on PRs
    if: github.event_name == 'pull_request'
    run: echo "this is a PR"

  - name: Run on success
    if: success()
    run: echo "previous steps passed"

  - name: Run on failure
    if: failure()
    run: echo "something failed"

  - name: Always run
    if: always()
    run: echo "runs no matter what"
```

```yaml
# job-level conditions
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

## fa-recycle Reusable Workflows

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
# .github/workflows/ci.yml — caller
jobs:
  test:
    uses: ./.github/workflows/reusable-test.yml
    with:
      node-version: "20"
    secrets:
      api-key: ${{ secrets.API_KEY }}
```
