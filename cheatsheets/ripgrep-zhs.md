---
title: Ripgrep
icon: fa-bolt
primary: "#FF2D20"
lang: bash
locale: zhs
---

## fa-magnifying-glass 基础搜索

```bash
rg "pattern"                      # 在当前目录搜索
rg "pattern" path/to/dir          # 在指定目录搜索
rg -i "pattern"                   # 忽略大小写
rg -w "word"                      # 全词匹配
rg "pattern" file.txt             # 在单个文件中搜索
```

## fa-file 文件过滤

```bash
rg -t py "pattern"                # 搜索 Python 文件
rg -T py "pattern"                # 排除 Python 文件
rg -t py -t js "pattern"          # 搜索 Python 和 JS 文件
rg -g "*.toml" "pattern"          # glob 包含
rg -g "!*.min.js" "pattern"       # glob 排除
rg -g "*.{rs,toml}" "pattern"     # 多个扩展名
```

## fa-folder-tree 目录控制

```bash
rg --max-depth 3 "pattern"        # 限制搜索深度
rg -d 3 "pattern"                 # 限制深度的简写
rg "pattern" --ignore-file .customignore
rg --no-ignore "pattern"          # 忽略 .gitignore
rg --no-ignore-vcs "pattern"      # 仅忽略 VCS 忽略文件
```

## fa-code 正则模式

```bash
rg "start.*end"                   # 基本正则
rg -e "pat1" -e "pat2"            # 多个模式 (或)
rg -f patterns.txt                # 从文件读取模式
rg -F "literal.*string"           # 固定字符串 (不解析正则)
rg -U "line1\nline2"              # 多行模式
rg "\bword\b"                     # 词边界
```

## fa-arrows-up-down 上下文与显示

```bash
rg -C 3 "pattern"                 # 前后各 3 行
rg -B 2 "pattern"                 # 前 2 行
rg -A 2 "pattern"                 # 后 2 行
rg --context-separator "---" "pattern"
rg -n "pattern"                   # 显示行号 (默认)
rg --no-line-number "pattern"     # 隐藏行号
```

## fa-display 输出格式

```bash
rg -l "pattern"                   # 仅列出文件名
rg --files-without-match "pattern"
rg -c "pattern"                   # 每个文件的匹配数
rg --count-matches "pattern"      # 总匹配数
rg -o "pattern"                   # 仅显示匹配部分
rg --json "pattern"               # JSON 输出
rg --vimgrep "pattern"            # vim quickfix 格式
```

## fa-file-code 文件类型

```bash
rg --type-list                    # 列出所有文件类型
rg -t css "color"                 # CSS 文件
rg -t html "href"                 # HTML 文件
rg -t rust "fn main"              # Rust 文件
rg -t markdown "TODO"             # Markdown 文件
rg -t sh "echo"                   # Shell 脚本
```

## fa-eye-slash 二进制与隐藏文件

```bash
rg --hidden "pattern"             # 搜索隐藏文件
rg --no-hidden "pattern"          # 排除隐藏文件 (默认)
rg --binary "pattern"             # 搜索二进制文件
rg --no-binary "pattern"          # 跳过二进制文件 (默认)
rg -a "pattern"                   # 将所有文件当作文本
```

## fa-arrows-rotate 替换预览

```bash
rg -r "replacement" "pattern"     # 预览替换 (不修改文件)
rg -r '$1' "(\w+)"                # 使用捕获组
rg -r '$1$new$3' "(old)(.*)($)" "pattern"
```

## fa-sort-amount-down 排序与限制

```bash
rg --sort path "pattern"          # 按路径排序
rg --sortr modified "pattern"     # 按修改时间降序
rg --sort accessed "pattern"      # 按访问时间排序
rg -m 100 "pattern"               # 最多 100 个匹配
rg -m 0 "pattern"                 # 不限匹配数
```

## fa-gear 配置

```bash
rg --config-file ~/.ripgreprc     # 使用配置文件
rg --no-config "pattern"          # 忽略配置
echo '--smart-case' >> ~/.ripgreprc
echo '--max-columns=150' >> ~/.ripgreprc
echo '--colors="match:fg:red"' >> ~/.ripgreprc
```

## fa-gauge-high 性能技巧

```bash
rg -j 4 "pattern"                 # 使用 4 个线程
rg --max-filesize 10M "pattern"   # 跳过大于 10MB 的文件
rg --mmap "pattern"               # 使用内存映射 I/O
rg --no-mmap "pattern"            # 禁用内存映射 I/O
rg --fast "pattern"               # 降低正则复杂度
```

## fa-lightbulb 实用示例

```bash
# 查找代码中的 TODO/FIXME
rg -n "TODO|FIXME|HACK" -t py -t rs -t js

# 查找函数定义
rg "fn \w+" -t rust
rg "(def|class) \w+" -t py
rg "function \w+" -t js

# 搜索 IP 地址
rg "\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}" access.log

# 查找敏感信息模式
rg -i "password|secret|api_key|token" --type-add 'lock:*.lock' -t lock

# 提取 URL
rg -o "https?://[^\s)]+" -N --no-filename

# 按文件统计模式出现次数并排序
rg -c "error" --sort path

# 搜索最近修改的文件
rg "pattern" --sortr modified --max-depth 5

# 预览查找替换
rg -r "newFunc" "oldFunc" -t go
```
