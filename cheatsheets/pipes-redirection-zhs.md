---
title: 管道与重定向
icon: fa-right-left
primary: "#FCC624"
lang: bash
locale: zhs
---

## fa-arrow-right 输出重定向

```bash
ls > files.txt             # 覆盖写入
ls >> files.txt            # 追加写入
echo "hello" > output.txt
cat /etc/hosts > hosts.bak

# 丢弃输出
command > /dev/null
```

## fa-arrow-left 输入重定向

```bash
mail -s "报告" user@host < report.txt
mysql -u root -p < backup.sql
sort < names.txt
wc -l < file.txt           # 仅输出行数，不显示文件名
```

## fa-right-left 管道

```bash
cat access.log | grep "404"
ps aux | grep nginx
cat file.txt | sort | uniq
dmesg | tail -20
history | grep "git"
```

## fa-triangle-exclamation 标准错误重定向

```bash
command 2> errors.log          # stderr 重定向到文件
command 2>> errors.log         # 追加 stderr
command 2>/dev/null            # 丢弃 stderr

# 分离 stdout 和 stderr
command > output.log 2> error.log
```

## fa-code-merge 组合重定向

```bash
command &> all.log              # stdout + stderr 写入文件 (bash)
command &>> all.log             # 追加两者
command > all.log 2>&1          # POSIX 兼容写法
command >> all.log 2>&1         # 追加，POSIX 写法

# 将 stderr 转为 stdout 以便管道处理
command 2>&1 | grep "error"
```

## fa-grip-lines Here Document

```bash
cat << EOF
第一行内容
第二行内容
EOF

cat << 'EOF'       # 不展开变量
$HOME 保持原样
EOF

cat <<- EOF        # 去除前导制表符
	缩进内容
	EOF
```

## fa-arrow-right-arrow-left Here String

```bash
grep "root" <<< "内容在这里"
read -r a b c <<< "1 2 3"
bc <<< "2^10"
awk '{print $1}' <<< "hello world"
```

## fa-t tee 命令

```bash
ls | tee files.txt              # 同时输出到屏幕和文件
ls | tee -a files.txt           # 追加模式
echo "日志记录" | tee -a log.txt

# 同时写入多个文件
ls | tee a.txt b.txt c.txt

# 管道传输的同时保存中间结果
cat bigfile.log | tee >(grep ERROR > errors.txt) | wc -l
```

## fa-list xargs

```bash
find . -name "*.log" | xargs rm
find . -type f | xargs grep "TODO"
cat urls.txt | xargs -n1 curl -s

# 处理含空格的文件名
find . -name "*.txt" -print0 | xargs -0 rm
ls | xargs -I {} cp {} /backup/{}
ls | xargs -P 4 -n1 process.sh   # 并行执行4个
```

## fa-shuffle 进程替换

```bash
diff <(sort file1.txt) <(sort file2.txt)
comm <(sort a.txt) <(sort b.txt)

while read line; do
    echo "$line"
done < <(grep "error" log.txt)

cat <(echo "生成的内容")
```

## fa-layer-group 文件描述符

```bash
exec 3> custom.log          # 打开 fd 3 用于写入
echo "info" >&3
exec 3>&-                   # 关闭 fd 3

exec 4< input.txt           # 打开 fd 4 用于读取
read -r line <&4
exec 4<&-                   # 关闭 fd 4

exec 5>&1                   # 保存 stdout 到 fd 5
exec > log.txt              # 重定向 stdout
exec 1>&5 5>&-              # 恢复 stdout
```

## fa-lightbulb 实用示例

```bash
# 统计访问日志中的独立 IP
awk '{print $1}' access.log | sort | uniq -c | sort -rn | head

# 带进度的备份
tar cf - /data | pv > backup.tar

# 批量搜索替换
grep -rl "旧文本" . | xargs sed -i 's/旧文本/新文本/g'

# 合并已排序的文件
sort -m <(sort file1) <(sort file2) > merged.txt

# 实时监控日志
tail -f /var/log/syslog | grep --line-buffered "error"
```
