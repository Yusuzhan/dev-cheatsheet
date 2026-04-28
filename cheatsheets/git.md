---
title: Git
icon: fa-code-branch
primary: "#F05032"
lang: bash
---

## fa-gear Setup & Config

```bash
git config --global user.name "Your Name"
git config --global user.email "you@example.com"
git config --global core.editor vim
git config --global init.defaultBranch main
git config --list
```

## fa-folder-plus Create & Clone

```bash
git init
git clone https://github.com/user/repo.git
git clone git@github.com:user/repo.git
git clone --depth 1 https://github.com/user/repo.git
git clone -b develop https://github.com/user/repo.git
```

## fa-plus Stage & Commit

```bash
git add file.txt
git add .
git add -p
git commit -m "feat: add login page"
git commit -am "fix: correct typo"
git commit --amend -m "new message"
```

## fa-code-branch Branching

```bash
git branch
git branch feature/login
git checkout feature/login
git checkout -b feature/login
git switch -c feature/login
git branch -d feature/login
git branch -D feature/login
git branch -m old-name new-name
```

## fa-code-merge Merge & Rebase

```bash
git merge feature/login
git merge --no-ff feature/login
git merge --abort

git rebase main
git rebase -i HEAD~3
git rebase --continue
git rebase --abort
```

## fa-cloud Remote

```bash
git remote -v
git remote add origin https://github.com/user/repo.git
git remote set-url origin git@github.com:user/repo.git
git fetch origin
git pull origin main
git push origin main
git push -u origin feature/login
git push origin --delete feature/login
```

## fa-clock-rotate-left Log & Diff

```bash
git log --oneline --graph --all
git log --oneline -10
git log --author="Alice"
git log --since="2 weeks ago"
git log -p file.txt

git diff
git diff --staged
git diff main..feature/login
git diff HEAD~1
```

## fa-rotate-left Undo & Reset

```bash
git restore file.txt
git restore --staged file.txt
git reset HEAD~1
git reset --soft HEAD~1
git reset --hard HEAD~1
git revert abc1234
git clean -fd
```

## fa-box-archive Stash

```bash
git stash
git stash -u
git stash list
git stash pop
git stash apply stash@{1}
git stash drop stash@{0}
git stash branch feature/from-stash
```

## fa-tags Tags

```bash
git tag v1.0.0
git tag -a v1.0.0 -m "Release 1.0.0"
git tag
git tag -d v1.0.0
git push origin v1.0.0
git push origin --tags
```

## fa-magnifying-glass Search & Blame

```bash
git blame file.txt
git blame -L 10,20 file.txt
git log -S "function_name"
git log -G "pattern"
git grep "TODO"
git bisect start
git bisect bad
git bisect good abc1234
```

## fa-wand-magic-sparks Cherry-pick & Reflog

```bash
git cherry-pick abc1234
git cherry-pick abc1234 def5678
git cherry-pick --abort

git reflog
git reflog show HEAD@{5}
git checkout HEAD@{5}
```

## fa-scroll Conventional Commits

```bash
# format: <type>(<scope>): <description>
git commit -m "feat: add user login page"
git commit -m "fix: correct password validation"
git commit -m "docs: update API documentation"
git commit -m "refactor: extract auth middleware"
git commit -m "test: add unit tests for user service"
git commit -m "chore: upgrade dependencies"
git commit -m "feat(auth): support OAuth2 login"
git commit -m "fix(api): handle null response"
```

Commit types:
- `feat` new feature
- `fix` bug fix
- `docs` documentation
- `style` formatting (no code change)
- `refactor` code restructuring
- `test` adding tests
- `chore` build/tooling changes

```
feat: add login page

Implement OAuth2 login with GitHub provider.
Includes redirect callback and session handling.

Closes #123
```

## fa-lightbulb Useful Tips

```bash
git shortlog -sn
git log --pretty=format:"%h - %an : %s" --since="1 day ago"
git diff --stat
git stash clear
git rm file.txt
git mv old.txt new.txt
git worktree add ../hotfix hotfix-branch
```
