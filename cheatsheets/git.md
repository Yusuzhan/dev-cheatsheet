---
title: Git
icon: fa-code-branch
primary: "#F05032"
lang: bash
---

## fa-gear Setup & Config

```bash
git config --global user.name "Your Name"     # set username
git config --global user.email "you@example.com"  # set email
git config --global core.editor vim            # set default editor
git config --global init.defaultBranch main    # set default branch name
git config --list                              # view all config
```

## fa-folder-plus Create & Clone

```bash
git init                                      # initialize new repo
git clone https://github.com/user/repo.git    # clone via HTTPS
git clone git@github.com:user/repo.git        # clone via SSH
git clone --depth 1 https://github.com/user/repo.git  # shallow clone
git clone -b develop https://github.com/user/repo.git # clone specific branch
```

## fa-plus Stage & Commit

```bash
git add file.txt              # stage a file
git add .                     # stage all changes
git add -p                    # interactive staging (hunk by hunk)
git commit -m "feat: add login page"   # commit with message
git commit -am "fix: correct typo"     # stage tracked files and commit
git commit --amend -m "new message"    # amend last commit message
```

## fa-code-branch Branching

```bash
git branch                    # list local branches
git branch feature/login      # create branch
git checkout feature/login    # switch to branch
git checkout -b feature/login # create and switch
git switch -c feature/login   # modern alternative to checkout -b
git branch -d feature/login   # delete merged branch
git branch -D feature/login   # force delete branch
git branch -m old-name new-name  # rename branch
```

## fa-code-merge Merge & Rebase

```bash
git merge feature/login       # merge branch into current
git merge --no-ff feature/login  # merge with merge commit (no fast-forward)
git merge --abort             # abort conflicted merge

git rebase main               # rebase current branch onto main
git rebase -i HEAD~3          # interactive rebase last 3 commits
git rebase --continue         # continue after resolving conflict
git rebase --abort            # abort rebase
```

## fa-cloud Remote

```bash
git remote -v                 # list remote URLs
git remote add origin https://github.com/user/repo.git  # add remote
git remote set-url origin git@github.com:user/repo.git  # change remote URL
git fetch origin              # fetch without merging
git pull origin main          # fetch and merge
git push origin main          # push to remote
git push -u origin feature/login  # push and set upstream
git push origin --delete feature/login  # delete remote branch
```

## fa-clock-rotate-left Log & Diff

```bash
git log --oneline --graph --all    # visual branch history
git log --oneline -10              # last 10 commits
git log --author="Alice"           # filter by author
git log --since="2 weeks ago"      # filter by date
git log -p file.txt                # commit history of a file

git diff                           # unstaged changes
git diff --staged                  # staged changes
git diff main..feature/login       # diff between branches
git diff HEAD~1                    # diff with last commit
```

## fa-rotate-left Undo & Reset

```bash
git restore file.txt           # discard working directory changes
git restore --staged file.txt  # unstage file (keep changes)
git reset HEAD~1              # undo last commit, keep changes unstaged
git reset --soft HEAD~1       # undo last commit, keep changes staged
git reset --hard HEAD~1       # undo last commit, discard all changes
git revert abc1234            # create new commit that undoes abc1234
git clean -fd                 # remove untracked files and directories
```

## fa-box-archive Stash

```bash
git stash                      # stash current changes
git stash -u                   # stash including untracked files
git stash list                 # list all stashes
git stash pop                  # apply and remove latest stash
git stash apply stash@{1}      # apply specific stash (keep in list)
git stash drop stash@{0}      # drop specific stash
git stash branch feature/from-stash  # create branch from stash
```

## fa-tags Tags

```bash
git tag v1.0.0                 # lightweight tag
git tag -a v1.0.0 -m "Release 1.0.0"  # annotated tag with message
git tag                        # list all tags
git tag -d v1.0.0             # delete local tag
git push origin v1.0.0        # push tag to remote
git push origin --tags        # push all tags to remote
```

## fa-magnifying-glass Search & Blame

```bash
git blame file.txt             # show who changed each line
git blame -L 10,20 file.txt   # blame specific line range
git log -S "function_name"    # find commits that added/removed string
git log -G "pattern"          # find commits matching regex in diff
git grep "TODO"               # search working directory
git bisect start              # binary search for bug-introducing commit
git bisect bad                # mark current as bad
git bisect good abc1234       # mark known-good commit
```

## fa-wand-magic-sparks Cherry-pick & Reflog

```bash
git cherry-pick abc1234        # apply specific commit to current branch
git cherry-pick abc1234 def5678  # cherry-pick multiple commits
git cherry-pick --abort       # abort cherry-pick

git reflog                    # show all HEAD movements
git reflog show HEAD@{5}      # show reflog entry
git checkout HEAD@{5}         # restore to previous state
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
git shortlog -sn              # contributor commit count
git log --pretty=format:"%h - %an : %s" --since="1 day ago"  # custom log format
git diff --stat               # show changed files summary
git stash clear               # remove all stashes
git rm file.txt               # remove file and stage deletion
git mv old.txt new.txt        # rename file and stage
git worktree add ../hotfix hotfix-branch  # create linked working tree
```
