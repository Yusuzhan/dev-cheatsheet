---
title: Pipes & Redirection
icon: fa-right-left
primary: "#FCC624"
lang: bash
---

## fa-arrow-right Output Redirection

```bash
ls > files.txt             # overwrite
ls >> files.txt            # append
echo "hello" > output.txt
cat /etc/hosts > hosts.bak

# discard output
command > /dev/null
```

## fa-arrow-left Input Redirection

```bash
mail -s "Report" user@host < report.txt
mysql -u root -p < backup.sql
sort < names.txt
wc -l < file.txt           # only count, no filename shown
```

## fa-right-left Pipes

```bash
cat access.log | grep "404"
ps aux | grep nginx
cat file.txt | sort | uniq
dmesg | tail -20
history | grep "git"
```

## fa-triangle-exclamation stderr Redirection

```bash
command 2> errors.log          # redirect stderr to file
command 2>> errors.log         # append stderr
command 2>/dev/null            # discard stderr

# separate stdout and stderr
command > output.log 2> error.log
```

## fa-code-merge Combined Redirection

```bash
command &> all.log              # stdout + stderr to file (bash)
command &>> all.log             # append both
command > all.log 2>&1          # POSIX compatible
command >> all.log 2>&1         # append, POSIX

# redirect stderr to stdout (for piping)
command 2>&1 | grep "error"
```

## fa-grip-lines Here Document

```bash
cat << EOF
Line one
Line two
EOF

cat << 'EOF'       # no variable expansion
$HOME stays literal
EOF

cat <<- EOF        # strip leading tabs
	indented content
	EOF
```

## fa-arrow-right-arrow-left Here String

```bash
grep "root" <<< "/etc/passwd content here"
read -r a b c <<< "1 2 3"
bc <<< "2^10"
awk '{print $1}' <<< "hello world"
```

## fa-t tee

```bash
ls | tee files.txt              # output to screen and file
ls | tee -a files.txt           # append mode
echo "log entry" | tee -a log.txt

# tee to multiple files
ls | tee a.txt b.txt c.txt

# pipe through while saving
cat bigfile.log | tee >(grep ERROR > errors.txt) | wc -l
```

## fa-list xargs

```bash
find . -name "*.log" | xargs rm
find . -type f | xargs grep "TODO"
cat urls.txt | xargs -n1 curl -s

# handle filenames with spaces
find . -name "*.txt" -print0 | xargs -0 rm
ls | xargs -I {} cp {} /backup/{}
ls | xargs -P 4 -n1 process.sh   # run 4 in parallel
```

## fa-shuffle Process Substitution

```bash
diff <(sort file1.txt) <(sort file2.txt)
comm <(sort a.txt) <(sort b.txt)

while read line; do
    echo "$line"
done < <(grep "error" log.txt)

# feed output as input file
cat <(echo "generated content")
```

## fa-layer-group File Descriptors

```bash
exec 3> custom.log          # open fd 3 for writing
echo "info" >&3
exec 3>&-                   # close fd 3

exec 4< input.txt           # open fd 4 for reading
read -r line <&4
exec 4<&-                   # close fd 4

exec 5>&1                   # save stdout to fd 5
exec > log.txt              # redirect stdout
exec 1>&5 5>&-              # restore stdout
```

## fa-lightbulb Practical Examples

```bash
# count unique IPs from access log
awk '{print $1}' access.log | sort | uniq -c | sort -rn | head

# backup with progress
tar cf - /data | pv > backup.tar

# search and replace across files
grep -rl "oldtext" . | xargs sed -i 's/oldtext/newtext/g'

# merge sorted files
sort -m <(sort file1) <(sort file2) > merged.txt

# monitor log in real time
tail -f /var/log/syslog | grep --line-buffered "error"
```
