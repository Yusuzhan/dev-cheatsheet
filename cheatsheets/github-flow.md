---
title: GitHub Flow
icon: fa-code-branch
primary: "#F05032"
lang: bash
---

## fa-stream Overview & Principles

GitHub Flow is a lightweight branching model centered on `main` as the deployable branch. Every change starts as a branch, goes through a pull request, gets reviewed, and merges back to `main`.

Core rules:
- `main` is always deployable
- Create a descriptive branch for every change
- Open a pull request early for discussion
- Merge only after review and passing CI
- Deploy from `main` immediately after merge

```bash
# typical lifecycle
git checkout main
git pull origin main
git checkout -b feature/add-search
# ... make changes ...
git push -u origin feature/add-search
# open PR → review → merge → deploy
```

## fa-code-branch Create Branch

```bash
git checkout main
git pull origin main
git checkout -b feature/user-profile
git push -u origin feature/user-profile
```

Branch always from the latest `main`. Keep branches short-lived—merge within hours or a few days, not weeks.

```bash
# update your branch with latest main before PR
git checkout feature/user-profile
git fetch origin
git rebase origin/main
git push --force-with-lease
```

## fa-upload Commit & Push

```bash
git add .
git commit -m "feat: add user profile page"
git push origin feature/user-profile

# amend and force-push safely
git commit --amend -m "feat: add user profile page with avatar"
git push --force-with-lease origin feature/user-profile
```

Use Conventional Commits for traceability:

```bash
git commit -m "feat: add search endpoint"
git commit -m "fix: handle empty query in search"
git commit -m "refactor: extract search logic to service"
```

## fa-code-pull-request Open Pull Request

```bash
# push branch then create PR via GitHub CLI
gh pr create --title "feat: add user profile" --body "## Changes
- Add profile page component
- Add avatar upload
- Add profile API endpoint

Closes #42"
```

```bash
# draft PR for early feedback
gh pr create --draft --title "WIP: new search feature"

# convert draft to ready
gh pr ready 123
```

```bash
# list and view PRs
gh pr list
gh pr view 123
gh pr view 123 --web
```

## fa-users Code Review

```bash
# request reviewers
gh pr edit 123 --add-reviewer alice,bob

# check review status
gh pr view 123 --json reviews

# approve from CLI
gh pr review 123 --approve --body "LGTM!"

# request changes
gh pr review 123 --request-changes --body "Please add unit tests"
```

Review checklist:
- Code follows project conventions
- Tests cover new behavior
- No obvious bugs or security issues
- Documentation updated if needed

```bash
# checkout a PR locally to test it
gh pr checkout 123
# run tests locally
npm test
# go back to your branch
git checkout feature/user-profile
```

## fa-clipboard-check CI Checks

```bash
# view CI status for a PR
gh pr checks 123

# watch checks in real-time
gh pr checks 123 --watch

# view detailed check info
gh api repos/OWNER/REPO/commits/SHA/check-runs
```

Configure required status checks in repo Settings → Branches → Branch protection rules. Common CI workflows:

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

## fa-code-merge Merge & Deploy

```bash
# merge PR (creates merge commit)
gh pr merge 123 --merge

# squash merge (preferred for clean history)
gh pr merge 123 --squash

# rebase merge (linear history)
gh pr merge 123 --rebase

# delete remote branch after merge
git push origin --delete feature/user-profile
```

Deploy after merge:

```bash
# pull latest main and deploy
git checkout main
git pull origin main

# tag the release
git tag -a v1.2.0 -m "Release 1.2.0: user profile feature"
git push origin v1.2.0

# trigger deployment
gh workflow run deploy.yml -f environment=production
```

## fa-tags Release & Tag

```bash
# create annotated tag on main
git tag -a v1.1.0 -m "Release 1.1.0"
git push origin v1.1.0

# create GitHub Release with notes
gh release create v1.1.0 --title "v1.1.0" --notes "## What's New
- User profile page
- Search improvements

## Bug Fixes
- Fixed login redirect"

# list releases
gh release list

# download release assets
gh release download v1.1.0
```

```bash
# generate changelog from commits
git log v1.0.0..v1.1.0 --oneline --no-merges
gh api repos/OWNER/REPO/compare/v1.0.0...v1.1.0 --jq '.commits[] | .commit.message'
```

## fa-label Branch Naming Conventions

```bash
# feature branches
feature/add-search
feature/user-profile

# bug fix branches
fix/login-redirect
fix/null-pointer-on-save

# other prefixes
hotfix/critical-security-patch
chore/upgrade-dependencies
docs/update-api-guide
refactor/extract-auth-middleware
test/add-user-service-tests
```

Keep names short, lowercase, hyphen-separated. The prefix tells the type, the rest tells the scope.

## fa-shield-halved Protect Branch Rules

```bash
# view branch protection via API
gh api repos/OWNER/REPO/branches/main/protection

# set up protection via GitHub Settings:
# 1. Require PR before merging
# 2. Require approvals (1-2 reviewers)
# 3. Require status checks to pass
# 4. Require branches to be up to date
# 5. Require signed commits
# 6. Do not allow force pushes
```

```bash
# enforce with GitHub CLI (enterprise/org)
gh api repos/OWNER/REPO/branches/main/protection \
  --method PUT \
  --field required_status_checks='{"strict":true,"contexts":["ci/test"]}' \
  --field required_pull_request_reviews='{"required_approving_review_count":1}' \
  --field enforce_admins=true \
  --field restrictions=null
```

## fa-code-merge Conflict Resolution

```bash
# rebase on latest main to find conflicts early
git checkout feature/user-profile
git fetch origin
git rebase origin/main

# if conflicts occur, resolve each file
# edit conflicted files, then:
git add <resolved-file>
git rebase --continue

# push the resolved branch
git push --force-with-lease origin feature/user-profile
```

```bash
# alternatively, merge main into your branch
git checkout feature/user-profile
git fetch origin
git merge origin/main
# resolve conflicts, then:
git add .
git commit -m "merge: resolve conflicts with main"
git push origin feature/user-profile
```

## fa-arrows-left-right Flow vs Git Flow

```bash
# GitHub Flow — single long-lived branch (main)
main ← feature/a
main ← feature/b
main ← hotfix/c

# Git Flow — multiple long-lived branches
main        ← tags only
develop     ← integration branch
feature/*   ← from develop
release/*   ← from develop into main
hotfix/*    ← from main back to main and develop
```

| Aspect | GitHub Flow | Git Flow |
|--------|------------|----------|
| Long-lived branches | `main` only | `main` + `develop` |
| Release branches | No | Yes |
| Complexity | Simple | Complex |
| Best for | Continuous deployment | Scheduled releases |
| Hotfix path | Branch from `main`, PR to `main` | Dedicated `hotfix/*` branch |

## fa-repeat Common Patterns

```bash
# start a new feature
git checkout main && git pull
git checkout -b feature/new-feature
# develop, commit, push, open PR

# hotfix workflow
git checkout main && git pull
git checkout -b hotfix/fix-payment-bug
# fix, commit, push, open PR, merge immediately

# sync long-running feature with main
git checkout feature/long-feature
git fetch origin
git rebase origin/main
git push --force-with-lease

# batch multiple small changes via squash
gh pr merge 123 --squash --subject "feat: add user settings module"
```
