---
title: GitHub Flow
icon: fa-code-branch
primary: "#F05032"
lang: bash
locale: zhs
---

## fa-stream 概述与原则

GitHub Flow 是一个轻量级分支模型，以 `main` 作为可部署分支。所有变更都从分支开始，经过 Pull Request 审查后合并回 `main`。

核心规则：
- `main` 始终保持可部署状态
- 每个变更都创建独立分支
- 尽早创建 Pull Request 进行讨论
- 审查通过且 CI 通过后方可合并
- 合并后立即从 `main` 部署

```bash
# 典型生命周期
git checkout main
git pull origin main
git checkout -b feature/add-search
# ... 编写代码 ...
git push -u origin feature/add-search
# 创建 PR → 审查 → 合并 → 部署
```

## fa-code-branch 创建分支

```bash
git checkout main
git pull origin main
git checkout -b feature/user-profile
git push -u origin feature/user-profile
```

始终从最新的 `main` 创建分支。保持分支短生命周期——几小时或几天内合并，不要拖几周。

```bash
# 创建 PR 前用最新 main 更新分支
git checkout feature/user-profile
git fetch origin
git rebase origin/main
git push --force-with-lease
```

## fa-upload 提交与推送

```bash
git add .
git commit -m "feat: 添加用户资料页"
git push origin feature/user-profile

# 修改提交后安全地强制推送
git commit --amend -m "feat: 添加用户资料页（含头像）"
git push --force-with-lease origin feature/user-profile
```

使用 Conventional Commits 规范：

```bash
git commit -m "feat: 添加搜索接口"
git commit -m "fix: 处理搜索空查询"
git commit -m "refactor: 提取搜索逻辑到 service 层"
```

## fa-code-pull-request 创建 Pull Request

```bash
# 推送分支后通过 GitHub CLI 创建 PR
gh pr create --title "feat: 添加用户资料页" --body "## 变更内容
- 添加资料页组件
- 添加头像上传功能
- 添加资料 API 接口

Closes #42"
```

```bash
# 创建草稿 PR 获取早期反馈
gh pr create --draft --title "WIP: 新搜索功能"

# 将草稿转为正式 PR
gh pr ready 123
```

```bash
# 列出和查看 PR
gh pr list
gh pr view 123
gh pr view 123 --web
```

## fa-users 代码审查

```bash
# 指定审查人
gh pr edit 123 --add-reviewer alice,bob

# 查看审查状态
gh pr view 123 --json reviews

# 通过 CLI 批准
gh pr review 123 --approve --body "LGTM!"

# 请求修改
gh pr review 123 --request-changes --body "请补充单元测试"
```

审查要点：
- 代码符合项目规范
- 测试覆盖新功能
- 无明显缺陷或安全问题
- 必要时更新文档

```bash
# 本地检出 PR 进行测试
gh pr checkout 123
# 本地运行测试
npm test
# 切回自己的分支
git checkout feature/user-profile
```

## fa-clipboard-check CI 检查

```bash
# 查看 PR 的 CI 状态
gh pr checks 123

# 实时监控检查
gh pr checks 123 --watch

# 查看详细的检查信息
gh api repos/OWNER/REPO/commits/SHA/check-runs
```

在仓库 Settings → Branches → Branch protection rules 中配置必须通过的状态检查。常见 CI 工作流：

```yaml
# .github/workflows/ci.yml
name: CI
on:
  pull_request:
    branches: [main]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: npm ci
      - run: npm test
      - run: npm run lint
```

## fa-code-merge 合并与部署

```bash
# 合并 PR（创建合并提交）
gh pr merge 123 --merge

# Squash 合并（推荐，保持历史整洁）
gh pr merge 123 --squash

# Rebase 合并（线性历史）
gh pr merge 123 --rebase

# 合并后删除远程分支
git push origin --delete feature/user-profile
```

合并后部署：

```bash
# 拉取最新 main 并部署
git checkout main
git pull origin main

# 打标签
git tag -a v1.2.0 -m "Release 1.2.0: 用户资料功能"
git push origin v1.2.0

# 触发部署
gh workflow run deploy.yml -f environment=production
```

## fa-tags 发布与标签

```bash
# 在 main 上创建附注标签
git tag -a v1.1.0 -m "Release 1.1.0"
git push origin v1.1.0

# 创建 GitHub Release 并附上说明
gh release create v1.1.0 --title "v1.1.0" --notes "## 新功能
- 用户资料页
- 搜索功能改进

## 缺陷修复
- 修复登录重定向问题"

# 列出所有 Release
gh release list

# 下载 Release 资源
gh release download v1.1.0
```

```bash
# 从提交记录生成变更日志
git log v1.0.0..v1.1.0 --oneline --no-merges
gh api repos/OWNER/REPO/compare/v1.0.0...v1.1.0 --jq '.commits[] | .commit.message'
```

## fa-label 分支命名规范

```bash
# 功能分支
feature/add-search
feature/user-profile

# 缺陷修复分支
fix/login-redirect
fix/null-pointer-on-save

# 其他前缀
hotfix/critical-security-patch
chore/upgrade-dependencies
docs/update-api-guide
refactor/extract-auth-middleware
test/add-user-service-tests
```

名称保持简短、小写、用连字符分隔。前缀表示类型，后面表示范围。

## fa-shield-halved 分支保护规则

```bash
# 通过 API 查看分支保护规则
gh api repos/OWNER/REPO/branches/main/protection

# 通过 GitHub Settings 设置：
# 1. 合并前必须创建 PR
# 2. 必须获得审批（1-2 人）
# 3. 必须通过状态检查
# 4. 分支必须是最新的
# 5. 必须签名提交
# 6. 禁止强制推送
```

```bash
# 通过 GitHub CLI 设置保护规则（企业/组织）
gh api repos/OWNER/REPO/branches/main/protection \
  --method PUT \
  --field required_status_checks='{"strict":true,"contexts":["ci/test"]}' \
  --field required_pull_request_reviews='{"required_approving_review_count":1}' \
  --field enforce_admins=true \
  --field restrictions=null
```

## fa-code-merge 冲突解决

```bash
# 变基到最新 main 以尽早发现冲突
git checkout feature/user-profile
git fetch origin
git rebase origin/main

# 出现冲突时，逐个解决
# 编辑冲突文件后：
git add <resolved-file>
git rebase --continue

# 推送解决后的分支
git push --force-with-lease origin feature/user-profile
```

```bash
# 或者将 main 合并到功能分支
git checkout feature/user-profile
git fetch origin
git merge origin/main
# 解决冲突后：
git add .
git commit -m "merge: 解决与 main 的冲突"
git push origin feature/user-profile
```

## fa-arrows-left-right Flow vs Git Flow

```bash
# GitHub Flow — 单个长期分支（main）
main ← feature/a
main ← feature/b
main ← hotfix/c

# Git Flow — 多个长期分支
main        ← 仅标签
develop     ← 集成分支
feature/*   ← 从 develop 创建
release/*   ← 从 develop 合入 main
hotfix/*    ← 从 main 修复合入 main 和 develop
```

| 方面 | GitHub Flow | Git Flow |
|------|------------|----------|
| 长期分支 | 仅 `main` | `main` + `develop` |
| 发布分支 | 无 | 有 |
| 复杂度 | 简单 | 复杂 |
| 适用于 | 持续部署 | 定期发布 |
| 热修路径 | 从 `main` 拉分支，PR 回 `main` | 专用 `hotfix/*` 分支 |

## fa-repeat 常用模式

```bash
# 开始新功能
git checkout main && git pull
git checkout -b feature/new-feature
# 开发、提交、推送、创建 PR

# 热修流程
git checkout main && git pull
git checkout -b hotfix/fix-payment-bug
# 修复、提交、推送、创建 PR、立即合并

# 长期功能分支同步 main
git checkout feature/long-feature
git fetch origin
git rebase origin/main
git push --force-with-lease

# 通过 squash 合并多个小改动
gh pr merge 123 --squash --subject "feat: 添加用户设置模块"
```
