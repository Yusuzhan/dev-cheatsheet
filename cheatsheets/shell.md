---
title: Shell (POSIX sh)
icon: fa-terminal
primary: "#4EAA25"
lang: sh
---

## fa-font Variables & Quoting

```sh
name="world"
echo "Hello, $name"
echo "Hello, ${name}"
echo 'literal $name'

greeting="${name:-there}"
export PATH="$PATH:/usr/local/bin"
readonly CONST="value"
unset var
```

## fa-text-width Strings

```sh
str="Hello World"
echo "${#str}"
echo "${str%World}"
echo "${str%%World}"
echo "${str#Hello}"
echo "${str##Hello}"
echo "${str/World/Bash}"
```

## fa-code-compare Conditionals (test/[ ])

```sh
if [ -f file.txt ]; then
    echo "regular file"
elif [ -d dir ]; then
    echo "directory"
else
    echo "not found"
fi

[ -n "$var" ]
[ -z "$var" ]
[ "$a" = "$b" ]
[ "$a" != "$b" ]
[ "$x" -gt "$y" ]
[ "$x" -lt "$y" ]
[ "$x" -eq "$y" ]
[ -r file ]
[ -w file ]
[ -x file ]
[ -s file ]
```

## fa-rotate Loops

```sh
for item in one two three; do
    echo "$item"
done

for f in *.txt; do
    echo "$f"
done

while [ "$count" -lt 10 ]; do
    count=$((count + 1))
done

until [ -f /tmp/done ]; do
    sleep 1
done
```

## fa-cube Functions

```sh
greet() {
    echo "Hello, $1"
}

add() {
    echo $(($1 + $2))
}

result=$(add 3 5)

return_status() {
    return 1
}
```

## fa-right-from-bracket Redirections

```sh
command > out.txt 2>&1
command >> out.txt
command 2>/dev/null
command >/dev/null 2>&1
command | tee output.txt
command1 | command2
cmd > out.txt 2> err.txt
```

## fa-door-open Exit Codes

```sh
command
if [ $? -eq 0 ]; then
    echo "success"
else
    echo "failed with code $?"
fi

exit 0
exit 1

command && echo "ok" || echo "fail"
```

## fa-sliders Positional Parameters

```sh
echo "script: $0"
echo "first: $1"
echo "second: $2"
echo "all: $@"
echo "count: $#"
echo "pid: $$"
echo "last exit: $?"
echo "background: $!"

shift
shift 2
```

## fa-file-lines Here Documents

```sh
cat <<EOF
multi line
text here
EOF

cat <<'EOF'
$not_expanded
EOF

variable=$(cat <<EOF
stored text
EOF
)
```

## fa-calculator Arithmetic

```sh
echo $((1 + 2))
echo $((10 / 3))
echo $((2 ** 10))
echo $((x += 5))
echo $((x++))
echo $((x--))
echo $((x % 3))
echo $((x & 0xFF))

x=$((x + 1))
```

## fa-stream Pattern Matching (case)

```sh
case "$status" in
    running) echo "ok" ;;
    stopped) echo "down" ;;
    start*) echo "starts with start" ;;
    [0-9]*) echo "starts with digit" ;;
    *) echo "unknown" ;;
esac

case "$extension" in
    jpg|png|gif) echo "image" ;;
    mp4|avi) echo "video" ;;
    *) echo "other" ;;
esac
```

## fa-bell Signal Handling

```sh
trap 'echo "Interrupted"; exit 1' INT
trap 'rm -f /tmp/tempfile' EXIT
trap 'cleanup' TERM HUP

trap - INT

kill -l
kill -15 $PID
kill -9 $PID
wait $PID
```

## fa-terminal Command Substitution

```sh
today=$(date +%Y-%m-%d)
files=$(ls *.txt)
lines=$(wc -l < file.txt)

result=$(echo "scale=2; 10/3" | bc)

dir=$(dirname "$0")
base=$(basename "$0")
```

## fa-shield Portability Tips

```sh
#!/bin/sh

[ "$var" = "yes" ]
[ -n "$var" ]
[ -z "$var" ]

echo "hello"
printf "name: %s\n" "$name"

[ x"$var" = x"yes" ]

command -v grep >/dev/null 2>&1 || { echo "grep required"; exit 1; }

: "${VAR:=default}"
```
