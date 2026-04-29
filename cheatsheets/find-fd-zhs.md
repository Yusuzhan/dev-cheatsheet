---
title: Find / Fd
icon: fa-magnifying-glass
primary: "#5A4FCF"
lang: bash
locale: zhs
---

## fa-magnifying-glass 基础查找

```bash
find /path/to/search
find .                          # 当前目录
find /home/user -maxdepth 2    # 限制深度
find / -name "*.conf" 2>/dev/null  # 抑制权限错误
```

## fa-font 按名称查找

```bash
find . -name "file.txt"
find . -iname "file.txt"       # 忽略大小写
find . -name "*.jpg"
find . -name "*.jpg" -o -name "*.png"   # 或匹配
find . -not -name "*.log"      # 取反匹配
```

## fa-folder 按类型查找

```bash
find . -type f                 # 普通文件
find . -type d                 # 目录
find . -type l                 # 符号链接
find . -type b                 # 块设备
find . -type s                 # 套接字
```

## fa-weight-hanging 按大小查找

```bash
find . -size 0                 # 空文件
find . -size +100M             # 大于 100MB
find . -size -1k               # 小于 1KB
find . -size +10M -size -100M  # 10MB 到 100MB 之间
find . -empty                  # 空文件或空目录
```

## fa-clock 按时间查找

```bash
find . -mtime -7               # 7 天内修改过
find . -mtime +30              # 超过 30 天前修改
find . -atime -1               # 1 天内访问过
find . -ctime -1               # 1 天内状态变更
find . -mmin -60               # 60 分钟内修改过
find . -newer reference.txt    # 比参考文件新
```

## fa-lock 按权限查找

```bash
find . -perm 755               # 精确权限匹配
find . -perm -u+x              # 所有者可执行
find . -perm /u+s              # 具有 setuid 位
find . -readable               # 当前用户可读
find . -writable               # 当前用户可写
```

## fa-user 按属主查找

```bash
find . -user root
find . -group developers
find . -user alice -group staff
find . -nouser                 # 无对应用户
find . -nogroup                # 无对应组
```

## fa-play 执行操作

```bash
find . -name "*.log" -print                    # 打印路径 (默认)
find . -name "*.log" -ls                       # ls -dils 格式
find . -name "*.tmp" -delete                   # 删除匹配文件
find . -type f -exec chmod 644 {} \;           # 每个文件执行一次
find . -type f -exec chmod 644 {} +            # 批量传参执行
find . -name "*.py" -exec grep "TODO" {} +
```

## fa-bolt fd 基础用法

```bash
fd pattern                     # 在当前目录按名称搜索
fd "\.jpg$"                    # 正则模式
fd -e txt pattern              # 按扩展名
fd -i readme                   # 忽略大小写
fd -g '*.py'                   # glob 模式
```

## fa-sliders fd 高级用法

```bash
fd -t f pattern                # 仅文件
fd -t d pattern                # 仅目录
fd -t l pattern                # 仅符号链接
fd -d 3 pattern                # 最大深度 3
fd -S +1G pattern              # 大于 1GB
fd --changed-within 1week      # 近期修改
fd -x wc -l                    # 对每个结果执行命令
fd -X rm                       # 用所有结果执行命令
```

## fa-link xargs + find

```bash
find . -name "*.log" -print0 | xargs -0 rm
find . -type f -name "*.js" -print0 | xargs -0 wc -l
find . -name "*.bak" -print0 | xargs -0 -I {} mv {} /tmp/backup/
find . -type f -size +0 -print0 | xargs -0 grep -l "pattern"
```

## fa-ban 排除与过滤

```bash
find . -path ./node_modules -prune -o -name "*.js" -print
find . -path ./.git -prune -o -type f -print
find . \( -path ./vendor -o -path ./dist \) -prune -o -print
fd --exclude node_modules --exclude .git pattern
fd -E '*.min.js' pattern
```

## fa-trash-can find + delete

```bash
find . -name "*.tmp" -delete
find . -name "*.log" -mtime +30 -delete
find . -type d -empty -delete
find . -name "core" -type f -delete
```

## fa-lightbulb 实用示例

```bash
# 查找并归档旧日志
find /var/log -name "*.log" -mtime +30 | tar -czf old-logs.tar.gz -T -

# 列出最大的 20 个文件
find . -type f -printf '%s %p\n' | sort -rn | head -20

# 查找跨目录的重复文件名
fd -t f -x basename | sort | uniq -d

# 查找包含特定内容的文件
fd -e py -x grep -l "import django"

# 按扩展名统计文件数
fd -t f | sed 's/.*\.//' | sort | uniq -c | sort -rn

# 查找失效的符号链接
find . -type l ! -exec test -e {} \; -print

# 递归修改文件属主
find /var/www -type f -exec chown www-data:www-data {} +
```
