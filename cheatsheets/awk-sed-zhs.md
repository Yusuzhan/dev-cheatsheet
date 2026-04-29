---
title: Awk / Sed
icon: fa-file-lines
primary: "#C74634"
lang: bash
locale: zhs
---

## fa-pen sed 基础替换

```bash
sed 's/old/new/' file.txt              # 每行替换第一个匹配
sed 's/old/new/g' file.txt             # 替换所有匹配
sed 's/old/new/2' file.txt             # 每行替换第 2 个匹配
sed -i 's/old/new/g' file.txt          # 就地编辑
echo "hello123" | sed 's/[0-9]//g'     # 删除所有数字
```

## fa-eraser sed 删除与插入

```bash
sed '/pattern/d' file.txt              # 删除匹配行
sed '3d' file.txt                      # 删除第 3 行
sed '1,5d' file.txt                    # 删除第 1-5 行
sed '/^$/d' file.txt                   # 删除空行
sed '2i\inserted line' file.txt        # 在第 2 行前插入
sed '2a\appended line' file.txt        # 在第 2 行后追加
sed '2c\replaced line' file.txt        # 替换第 2 行
```

## fa-layer-group sed 多行操作

```bash
sed 'N;s/\n/ /' file.txt              # 合并相邻两行
sed ':a;N;$!ba;s/\n/ /g' file.txt     # 合并所有行为一行
sed '/start/,/end/d' file.txt         # 删除两个模式之间的内容
sed '/start/,/end/s/old/new/g' file.txt
```

## fa-floppy-disk sed 就地编辑

```bash
sed -i.bak 's/old/new/g' file.txt     # 就地编辑并备份
sed -i '' 's/old/new/g' file.txt      # macOS 就地编辑 (无备份)
sed -i -e 's/old/new/g' file.txt
```

## fa-map-pin sed 地址与范围

```bash
sed -n '5p' file.txt                   # 仅打印第 5 行
sed -n '5,10p' file.txt                # 打印第 5-10 行
sed -n '$p' file.txt                   # 打印最后一行
sed -n '/pattern/p' file.txt           # 打印匹配行
sed -n '/start/,/end/p' file.txt       # 打印两个模式之间的行
sed '1~2d' file.txt                    # 从第 1 行起隔行删除
```

## fa-terminal awk 基础用法

```bash
awk '{print}' file.txt                 # 打印所有行
awk '{print $1}' file.txt              # 打印第一个字段
awk '{print $1, $3}' file.txt          # 打印第 1 和第 3 个字段
awk '/pattern/{print}' file.txt        # 打印匹配行
awk '{print NR, $0}' file.txt          # 带行号打印
```

## fa-table-columns awk 字段与分隔符

```bash
awk -F',' '{print $1, $2}' file.csv   # 逗号分隔
awk -F'\t' '{print $1}' file.tsv      # 制表符分隔
awk -F':' '{print $1}' /etc/passwd
awk 'BEGIN{FS=","; OFS="|"} {print $1, $2}' file.csv
awk '{print NF}' file.txt              # 每行字段数
awk '{print $NF}' file.txt             # 最后一个字段
awk '{print $(NF-1)}' file.txt         # 倒数第二个字段
```

## fa-filter awk 模式与条件

```bash
awk '$3 > 100' file.txt                # 第 3 列大于 100
awk '$1 == "root"' /etc/passwd
awk '/start/,/end/' file.txt           # 范围模式
awk 'NR > 1' file.txt                  # 跳过表头
awk 'NR >= 5 && NR <= 10' file.txt     # 第 5 到 10 行
awk '$3 ~ /^[0-9]+$/' file.txt         # 正则匹配第 3 列
awk '$3 !~ /pattern/' file.txt         # 正则不匹配
```

## fa-boxes-stacked awk 变量

```bash
awk '{sum += $1} END {print sum}' file.txt
awk '{sum += $1; count++} END {print sum/count}' file.txt
awk 'BEGIN {pi=3.14159; print pi}'
awk '{arr[NR] = $0} END {for(i=NR;i>0;i--) print arr[i]}' file.txt
awk '{len = length($0); if(len > max) max = len} END {print max}' file.txt
```

## fa-code-branch awk 控制流

```bash
awk '{if ($3 > 50) print $1, "high"; else print $1, "low"}' file.txt
awk '{
  for (i = 1; i <= NF; i++) {
    if ($i > 100) print $i
  }
}' file.txt
awk '{
  i = 1
  while (i <= NF) {
    print $i; i++
  }
}' file.txt
```

## fa-database awk 数组

```bash
awk '{count[$1]++} END {for(k in count) print k, count[k]}' file.txt
awk '{sum[$1] += $2} END {for(k in sum) print k, sum[k]}' file.txt
awk '!seen[$1]++' file.txt             # 按第 1 列去重
awk '{a[$1] = a[$1] ? a[$1]","$2 : $2} END {for(k in a) print k, a[k]}' file.txt
```

## fa-wand-magic-sparkles awk 函数

```bash
awk '{print length($0)}' file.txt      # 字符串长度
awk '{print substr($0, 1, 10)}' file.txt
awk '{print toupper($1)}' file.txt
awk '{print tolower($1)}' file.txt
awk '{print index($0, "pattern")}' file.txt
awk '{print split($0, arr, ",")}' file.txt
awk '{gsub(/old/, "new"); print}' file.txt
```

## fa-bolt awk 单行命令

```bash
awk 'END{print NR}' file.txt           # 统计行数
awk '{s+=$1} END{print s}' file.txt    # 求和第 1 列
awk 'NR%2==0' file.txt                 # 打印偶数行
awk '{gsub(/\r/,""); print}' file.txt  # 去除回车符
awk '{$1=$1; print}' file.txt          # 规范化空白
awk '{print length, $0}' file.txt | sort -rn | head -5  # 最长的行
```

## fa-lightbulb 实用示例

```bash
# 从 passwd 提取用户名
awk -F':' '{print $1}' /etc/passwd

# 统计目录中文件总大小
ls -l | awk '{sum += $5} END {print sum " bytes"}'

# CSV 列提取 (跳过表头)
awk -F',' 'NR > 1 {print $2, $5}' data.csv

# 统计访问日志中 Top 10 IP
awk '{print $1}' access.log | sort | uniq -c | sort -rn | head -10

# 交换两列
awk '{print $2, $1, $3}' file.txt

# 统计第 1 列唯一值出现次数
awk -F',' '{count[$1]++} END {for(k in count) print k, count[k]}' data.csv

# 去除 HTML 标签
sed 's/<[^>]*>//g' file.html

# 给文件加行号
sed '=' file.txt | sed 'N;s/\n/\t/'
```
