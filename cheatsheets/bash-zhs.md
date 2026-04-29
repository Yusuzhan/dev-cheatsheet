---
title: Bash
icon: fa-terminal
primary: "#FCC624"
lang: bash
locale: zhs
---

## fa-box 变量与展开

```bash
name="world"
echo "Hello, $name"
echo "Hello, ${name}"
greeting="Hello, ${name:-there}"
export PATH="$PATH:/usr/local/bin"
local var="inside function"
readonly CONST="value"
unset var
```

## fa-font 字符串

```bash
str="Hello World"
echo "${#str}"                     # 字符串长度
echo "${str^^}"                    # 转大写
echo "${str,,}"                    # 转小写
echo "${str:0:5}"                  # 子字符串
echo "${str/World/Bash}"           # 替换第一个
echo "${str//o/0}"                 # 替换所有
echo "${str#Hello }"               # 删除前缀
echo "${str%World}"                # 删除后缀
echo "${str^h}"                    # 首字母大写
```

## fa-list 数组

```bash
arr=(one two three)
echo "${arr[0]}"                   # 第一个元素
echo "${arr[@]}"                   # 所有元素
echo "${#arr[@]}"                  # 数组长度
echo "${arr[-1]}"                  # 最后一个元素
arr+=(four five)                   # 追加
echo "${arr[@]:1:2}"               # 切片
unset arr[1]                       # 删除元素
declare -A map                     # 关联数组
map=([name]="alice" [age]=30)
echo "${map[name]}"
```

## fa-code-compare 条件判断

```bash
if [ -f file.txt ]; then
  echo "文件存在"
elif [ -d dir ]; then
  echo "是目录"
else
  echo "未找到"
fi

[[ -n "$var" && "$var" == "yes" ]]
[[ "$str" =~ ^[0-9]+$ ]]
[[ -x script.sh ]]

case "$status" in
  running) echo "运行中" ;;
  stopped) echo "已停止" ;;
  *) echo "未知" ;;
esac
```

## fa-rotate 循环

```bash
for i in 1 2 3; do echo "$i"; done
for i in {1..10}; do echo "$i"; done
for i in {a..z}; do echo "$i"; done
for file in *.txt; do echo "$file"; done

while read -r line; do
  echo "$line"
done < file.txt

while [[ $count -lt 10 ]]; do
  ((count++))
done

until [[ -f /tmp/done ]]; do sleep 1; done
```

## fa-cube 函数

```bash
greet() {
  local name="$1"
  echo "Hello, $name"
}

add() {
  echo $(($1 + $2))
}

result=$(add 3 5)

return_status() {
  return 1
}

# 访问所有参数
args() {
  echo "全部: $@"
  echo "数量: $#"
  echo "第一个: $1"
  echo "脚本名: $0"
}
```

## fa-right-from-bracket 重定向与管道

```bash
command > out.txt 2>&1           # stdout+stderr 写入文件
command &> out.txt               # 同上 (简写)
command >> out.txt 2>> err.txt   # 分别追加
command 2>/dev/null              # 丢弃 stderr
command | tee output.txt         # 管道并保存
command | grep "pattern"         # 管道到 grep
command1 && command2             # 成功时执行
command1 || command2             # 失败时执行
```

## fa-file-lines Here Document

```bash
cat << EOF
多行
文本内容
EOF

cat <<- EOF
  缩进的 heredoc (去除前导制表符)
EOF

variable=$(cat <<EOF
存储的文本
EOF
)

ssh host <<EOF
  hostname
  uptime
EOF
```

## fa-sliders 参数展开

```bash
${var:-default}                   # 未设置/为空时使用默认值
${var:=default}                   # 未设置/为空时赋默认值
${var:+alternate}                 # 已设置时使用替代值
${var:?error message}             # 未设置/为空时报错
${#var}                           # 长度
${var:start:length}               # 子字符串
${var#pattern}                    # 删除最短前缀匹配
${var##pattern}                   # 删除最长前缀匹配
${var%pattern}                    # 删除最短后缀匹配
${var%%pattern}                   # 删除最长后缀匹配
${var/pattern/replacement}        # 替换第一个
${var//pattern/replacement}       # 替换所有
```

## fa-calculator 算术运算

```bash
echo $((1 + 2))
echo $((10 / 3))
echo $((2 ** 10))
echo $((x++))
echo $((x--))
echo $((x += 5))

((x = 3 * 4))
((x > 5)) && echo "greater"

# bc 浮点运算
echo "scale=2; 10/3" | bc
echo "sqrt(2)" | bc -l
```

## fa-shuffle 进程替换

```bash
diff <(sort file1.txt) <(sort file2.txt)
comm <(sort a.txt) <(sort b.txt)
while read -r left right; do
  echo "$left | $right"
done < <(paste <(cmd1) <(cmd2))
cat <(echo "header") file.txt
```

## fa-bell 信号与陷阱

```bash
trap 'echo "已中断"; exit 1' INT
trap 'rm -f /tmp/tempfile' EXIT
trap 'cleanup' TERM HUP

trap - INT                        # 重置陷阱

kill -l                           # 列出所有信号
kill -9 $PID                      # SIGKILL
kill -15 $PID                     # SIGTERM
kill -0 $PID 2>/dev/null && echo "运行中"
wait $PID                         # 等待进程
```

## fa-keyboard getopts

```bash
while getopts "a:b:c" opt; do
  case $opt in
    a) arg_a="$OPTARG" ;;
    b) arg_b="$OPTARG" ;;
    c) flag_c=1 ;;
    *) echo "用法: $0 [-a 参数] [-b 参数] [-c]" >&2; exit 1 ;;
  esac
done
shift $((OPTIND - 1))
```

## fa-lightbulb 实用示例

```bash
# 带时间戳的备份
cp file.txt "file_$(date +%Y%m%d_%H%M%S).bak"

# 批量重命名
for f in *.jpg; do mv "$f" "${f%.jpg}_thumb.jpg"; done

# 查找并终止进程
kill $(lsof -ti:8080)

# 目录大小汇总
du -sh */ | sort -rh | head -10

# 定时查看命令输出
watch -n 2 'df -h'

# 循环中的进度计数
total=$(ls *.jpg | wc -l)
for f in *.jpg; do
  ((count++))
  printf "\r[%d/%d] %s" "$count" "$total" "$f"
  convert "$f" -resize 50% "out/$f"
done

# 解析 key=value 配置
while IFS='=' read -r key value; do
  declare "$key=$value"
done < config.env

# 并行执行
for host in host1 host2 host3; do
  ssh "$host" "uptime" &
done
wait
```
