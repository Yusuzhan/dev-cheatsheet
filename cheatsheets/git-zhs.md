---
title: Git
icon: fa-code-branch
primary: "#F05032"
lang: bash
locale: zhs
---

## fa-gear 初始配置

```bash
git config --global user.name "Your Name"
git config --global user.email "you@example.com"
git config --global core.editor vim
git config --global init.defaultBranch main
git config --list
```

## fa-folder-plus 创建与克隆

```bash
git init
git clone https://github.com/user/repo.git
git clone git@github.com:user/repo.git
git clone --depth 1 https://github.com/user/repo.git    # 浅克隆
git clone -b develop https://github.com/user/repo.git    # 克隆指定分支
```

## fa-plus 暂存与提交

```bash
git add file.txt
git add .
git add -p                  # 交互式暂存
git commit -m "feat: 添加登录页"
git commit -am "fix: 修正拼写错误"
git commit --amend -m "新提交信息"
```

## fa-code-branch 分支操作

```bash
git branch                  # 列出分支
git branch feature/login    # 创建分支
git checkout feature/login  # 切换分支
git checkout -b feature/login  # 创建并切换
git switch -c feature/login    # 新版切换命令
git branch -d feature/login   # 删除已合并分支
git branch -D feature/login   # 强制删除分支
git branch -m old-name new-name  # 重命名分支
```

## fa-code-merge 合并与变基

```bash
git merge feature/login
git merge --no-ff feature/login   # 禁用快进合并
git merge --abort

git rebase main                   # 变基到 main
git rebase -i HEAD~3              # 交互式变基最近3条
git rebase --continue
git rebase --abort
```

## fa-cloud 远程操作

```bash
git remote -v
git remote add origin https://github.com/user/repo.git
git remote set-url origin git@github.com:user/repo.git
git fetch origin
git pull origin main
git push origin main
git push -u origin feature/login   # 推送并设置上游
git push origin --delete feature/login  # 删除远程分支
```

## fa-clock-rotate-left 日志与差异

```bash
git log --oneline --graph --all
git log --oneline -10
git log --author="Alice"
git log --since="2 weeks ago"
git log -p file.txt

git diff                    # 工作区与暂存区差异
git diff --staged           # 暂存区与最新提交差异
git diff main..feature/login  # 两个分支差异
git diff HEAD~1
```

## fa-rotate-left 撤销与重置

```bash
git restore file.txt            # 恢复工作区文件
git restore --staged file.txt   # 取消暂存
git reset HEAD~1               # 软重置 (保留修改)
git reset --soft HEAD~1        # 保留暂存区
git reset --hard HEAD~1        # 彻底重置 (丢弃修改)
git revert abc1234             # 创建反向提交
git clean -fd                  # 清理未跟踪文件
```

## fa-box-archive 暂存工作区

```bash
git stash
git stash -u                   # 包含未跟踪文件
git stash list
git stash pop                  # 恢复并删除
git stash apply stash@{1}      # 恢复但保留记录
git stash drop stash@{0}
git stash branch feature/from-stash
```

## fa-tags 标签管理

```bash
git tag v1.0.0                  # 轻量标签
git tag -a v1.0.0 -m "发布 1.0.0"  # 附注标签
git tag                         # 列出标签
git tag -d v1.0.0              # 删除本地标签
git push origin v1.0.0         # 推送标签
git push origin --tags         # 推送所有标签
```

## fa-magnifying-glass 搜索与追溯

```bash
git blame file.txt              # 查看每行最后修改者
git blame -L 10,20 file.txt    # 指定行范围
git log -S "function_name"     # 搜索包含该字符串的提交
git log -G "pattern"           # 用正则搜索差异
git grep "TODO"                # 在工作区搜索
git bisect start               # 二分查找引入bug的提交
git bisect bad
git bisect good abc1234
```

## fa-wand-magic-sparks 精选与引用日志

```bash
git cherry-pick abc1234         # 将指定提交应用到当前分支
git cherry-pick abc1234 def5678
git cherry-pick --abort

git reflog                     # 查看所有操作记录
git reflog show HEAD@{5}
git checkout HEAD@{5}          # 恢复到历史状态
```

## fa-lightbulb 实用技巧

```bash
git shortlog -sn               # 按提交数统计贡献者
git log --pretty=format:"%h - %an : %s" --since="1 day ago"
git diff --stat                # 显示变更文件统计
git stash clear                # 清空所有 stash
git rm file.txt                # 删除并暂存
git mv old.txt new.txt         # 重命名并暂存
git worktree add ../hotfix hotfix-branch  # 创建工作树
```
