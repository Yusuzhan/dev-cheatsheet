---
title: Bash
icon: fa-terminal
primary: "#FCC624"
lang: bash
---

## fa-box Variables & Expansion

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

## fa-font Strings

```bash
str="Hello World"
echo "${#str}"                     # string length
echo "${str^^}"                    # to uppercase
echo "${str,,}"                    # to lowercase
echo "${str:0:5}"                  # substring
echo "${str/World/Bash}"           # replace first
echo "${str//o/0}"                 # replace all
echo "${str#Hello }"               # remove prefix
echo "${str%World}"                # remove suffix
echo "${str^h}"                    # capitalize first char
```

## fa-list Arrays

```bash
arr=(one two three)
echo "${arr[0]}"                   # first element
echo "${arr[@]}"                   # all elements
echo "${#arr[@]}"                  # array length
echo "${arr[-1]}"                  # last element
arr+=(four five)                   # append
echo "${arr[@]:1:2}"               # slice
unset arr[1]                       # delete element
declare -A map                     # associative array
map=([name]="alice" [age]=30)
echo "${map[name]}"
```

## fa-code-compare Conditionals

```bash
if [ -f file.txt ]; then
  echo "file exists"
elif [ -d dir ]; then
  echo "is directory"
else
  echo "not found"
fi

[[ -n "$var" && "$var" == "yes" ]]
[[ "$str" =~ ^[0-9]+$ ]]
[[ -x script.sh ]]

case "$status" in
  running) echo "ok" ;;
  stopped) echo "down" ;;
  *) echo "unknown" ;;
esac
```

## fa-rotate Loops

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

## fa-cube Functions

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

# access all arguments
args() {
  echo "all: $@"
  echo "count: $#"
  echo "first: $1"
  echo "script: $0"
}
```

## fa-right-from-bracket Redirections & Pipes

```bash
command > out.txt 2>&1           # stdout+stderr to file
command &> out.txt               # same (shorter)
command >> out.txt 2>> err.txt   # append separately
command 2>/dev/null              # discard stderr
command | tee output.txt         # pipe and save
command | grep "pattern"         # pipe to grep
command1 && command2             # run if success
command1 || command2             # run if failure
```

## fa-file-lines Here Documents

```bash
cat << EOF
multi line
text here
EOF

cat <<- EOF
  indented heredoc (tabs stripped)
EOF

variable=$(cat <<EOF
stored text
EOF
)

ssh host <<EOF
  hostname
  uptime
EOF
```

## fa-sliders Parameter Expansion

```bash
${var:-default}                   # use default if unset/empty
${var:=default}                   # assign default if unset/empty
${var:+alternate}                 # use alternate if set
${var:?error message}             # error if unset/empty
${#var}                           # length
${var:start:length}               # substring
${var#pattern}                    # remove shortest prefix
${var##pattern}                   # remove longest prefix
${var%pattern}                    # remove shortest suffix
${var%%pattern}                   # remove longest suffix
${var/pattern/replacement}        # replace first
${var//pattern/replacement}       # replace all
```

## fa-calculator Arithmetic

```bash
echo $((1 + 2))
echo $((10 / 3))
echo $((2 ** 10))
echo $((x++))
echo $((x--))
echo $((x += 5))

((x = 3 * 4))
((x > 5)) && echo "greater"

# bc for floating point
echo "scale=2; 10/3" | bc
echo "sqrt(2)" | bc -l
```

## fa-shuffle Process Substitution

```bash
diff <(sort file1.txt) <(sort file2.txt)
comm <(sort a.txt) <(sort b.txt)
while read -r left right; do
  echo "$left | $right"
done < <(paste <(cmd1) <(cmd2))
cat <(echo "header") file.txt
```

## fa-bell Traps & Signals

```bash
trap 'echo "Interrupted"; exit 1' INT
trap 'rm -f /tmp/tempfile' EXIT
trap 'cleanup' TERM HUP

trap - INT                        # reset trap

kill -l                           # list all signals
kill -9 $PID                      # SIGKILL
kill -15 $PID                     # SIGTERM
kill -0 $PID 2>/dev/null && echo "running"
wait $PID                         # wait for process
```

## fa-keyboard getopts

```bash
while getopts "a:b:c" opt; do
  case $opt in
    a) arg_a="$OPTARG" ;;
    b) arg_b="$OPTARG" ;;
    c) flag_c=1 ;;
    *) echo "Usage: $0 [-a arg] [-b arg] [-c]" >&2; exit 1 ;;
  esac
done
shift $((OPTIND - 1))
```

## fa-lightbulb Practical Examples

```bash
# Backup with timestamp
cp file.txt "file_$(date +%Y%m%d_%H%M%S).bak"

# Batch rename
for f in *.jpg; do mv "$f" "${f%.jpg}_thumb.jpg"; done

# Find and kill process
kill $(lsof -ti:8080)

# Directory size summary
du -sh */ | sort -rh | head -10

# Watch command output
watch -n 2 'df -h'

# Progress counter in loop
total=$(ls *.jpg | wc -l)
for f in *.jpg; do
  ((count++))
  printf "\r[%d/%d] %s" "$count" "$total" "$f"
  convert "$f" -resize 50% "out/$f"
done

# Parse key=value from config
while IFS='=' read -r key value; do
  declare "$key=$value"
done < config.env

# Parallel execution
for host in host1 host2 host3; do
  ssh "$host" "uptime" &
done
wait
```
