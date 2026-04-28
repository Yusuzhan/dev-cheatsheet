---
title: Grep
icon: fa-magnifying-glass
primary: "#F05032"
lang: bash
locale: zhs
---

## fa-magnifying-glass 基础搜索

```bash
# 在文件中搜索
grep "pattern" file.txt

# 忽略大小写
grep -i "pattern" file.txt

# 在多个文件中搜索
grep "pattern" file1.txt file2.txt

# 通过标准输入搜索
cat file.txt | grep "pattern"
echo "hello world" | grep "hello"
```

## fa-folder-tree 递归搜索

```bash
# 递归搜索目录
grep -r "pattern" /path/to/dir

# 跟随符号链接
grep -R "pattern" /path/to/dir

# 递归搜索并限定文件类型
grep -r --include="*.py" "pattern" /path/to/dir

# 排除目录
grep -r --exclude-dir={node_modules,.git} "pattern" .

# 排除文件类型
grep -r --exclude="*.log" "pattern" .
```

## fa-code 正则表达式

```bash
# 基本正则
grep "^[A-Z]" file.txt           # 以大写字母开头的行
grep "[0-9]\{3\}" file.txt       # 包含连续3个数字的行
grep "end$" file.txt             # 以 "end" 结尾的行

# 扩展正则 (-E)
grep -E "cat|dog" file.txt       # 匹配 cat 或 dog
grep -E "go+d" file.txt          # 匹配 god, good, gooood...
grep -E "colou?r" file.txt       # 匹配 color 或 colour

# Perl 兼容正则 (-P)
grep -P "\d{3}-\d{4}" file.txt   # 匹配电话号码格式
grep -P "(?<=error: ).*" log.txt # 后顾断言
```

## fa-arrows-up-down 上下文行

```bash
# 显示匹配行后 N 行
grep -A 3 "pattern" file.txt

# 显示匹配行前 N 行
grep -B 3 "pattern" file.txt

# 显示匹配行前后各 N 行
grep -C 3 "pattern" file.txt

# 仅显示匹配部分
grep -o -E "[0-9]+" file.txt     # 提取所有数字
```

## fa-file 文件过滤

```bash
# 仅显示包含匹配的文件名
grep -l "pattern" *.txt

# 仅显示不包含匹配的文件名
grep -L "pattern" *.txt

# 搜索指定文件类型
grep "pattern" *.py
grep "pattern" *.md

# 递归搜索时包含/排除
grep -r --include="*.{js,ts}" "pattern" .
grep -r --exclude="*.min.js" "pattern" .
```

## fa-calculator 反转与计数

```bash
# 反转匹配 (显示不匹配的行)
grep -v "pattern" file.txt

# 统计匹配行数
grep -c "pattern" file.txt

# 显示行号
grep -n "pattern" file.txt

# 仅显示匹配部分
grep -o "pattern" file.txt

# 抑制错误信息
grep -s "pattern" /root/*
```

## fa-gear 搜索选项

```bash
# 全词匹配
grep -w "error" file.txt         # 匹配 "error" 不匹配 "errors"

# 整行精确匹配
grep -x "exact line" file.txt

# 固定字符串 (不解析正则)
grep -F ".*" file.txt            # 字面匹配 .*

# 将二进制文件当作文本处理
grep -a "pattern" binary_file

# 搜索压缩文件
zgrep "pattern" file.gz
```

## fa-display 输出控制

```bash
# 彩色输出
grep --color=auto "pattern" file.txt

# 显示文件名前缀 (多文件时有用)
grep -H "pattern" *.txt

# 不显示文件名前缀
grep -h "pattern" *.txt

# 显示字节偏移量
grep -b "pattern" file.txt

# 提取匹配内容并统计排序
grep -o "pattern" file.txt | sort | uniq -c | sort -rn
```

## fa-wand-magic-sparkles 高级模式

```bash
# 邮箱地址
grep -E "[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}" file.txt

# IP 地址
grep -E "[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}" file.txt

# 十六进制颜色值
grep -E "#[0-9a-fA-F]{6}" file.txt

# URL 模式
grep -E "https?://[^[:space:]]+" file.txt

# 空行
grep -E "^$" file.txt
```

## fa-lightbulb 实用示例

```bash
# 查找项目中的 TODO 注释
grep -rn "TODO\|FIXME\|HACK" --include="*.{py,js,ts,go}" .

# 搜索进程
ps aux | grep nginx

# 搜索命令历史
history | grep "docker"

# 提取访问日志中的 HTTP 状态码并统计
grep -oE " [0-9]{3} " access.log | sort | uniq -c | sort -rn

# 查找不包含 license 的文件
grep -rL "license" --include="*.py" .

# 统计 JS 文件中 function 出现次数
grep -ro "function" --include="*.js" . | wc -l
```
