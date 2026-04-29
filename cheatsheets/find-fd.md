---
title: Find / Fd
icon: fa-magnifying-glass
primary: "#5A4FCF"
lang: bash
---

## fa-magnifying-glass Basic Find

```bash
find /path/to/search
find .                          # current directory
find /home/user -maxdepth 2    # limit depth
find / -name "*.conf" 2>/dev/null  # suppress permission errors
```

## fa-font Find by Name

```bash
find . -name "file.txt"
find . -iname "file.txt"       # case-insensitive
find . -name "*.jpg"
find . -name "*.jpg" -o -name "*.png"   # OR matching
find . -not -name "*.log"      # NOT matching
```

## fa-folder Find by Type

```bash
find . -type f                 # regular files
find . -type d                 # directories
find . -type l                 # symbolic links
find . -type b                 # block devices
find . -type s                 # sockets
```

## fa-weight-hanging Find by Size

```bash
find . -size 0                 # empty files
find . -size +100M             # larger than 100MB
find . -size -1k               # smaller than 1KB
find . -size +10M -size -100M  # between 10MB and 100MB
find . -empty                  # empty files or directories
```

## fa-clock Find by Time

```bash
find . -mtime -7               # modified within 7 days
find . -mtime +30              # modified more than 30 days ago
find . -atime -1               # accessed within 1 day
find . -ctime -1               # status changed within 1 day
find . -mmin -60               # modified within 60 minutes
find . -newer reference.txt    # newer than reference file
```

## fa-lock Find by Permissions

```bash
find . -perm 755               # exact permission
find . -perm -u+x              # executable by owner
find . -perm /u+s              # has setuid bit
find . -readable               # readable by current user
find . -writable               # writable by current user
```

## fa-user Find by Owner

```bash
find . -user root
find . -group developers
find . -user alice -group staff
find . -nouser                 # no matching user
find . -nogroup                # no matching group
```

## fa-play Execute Actions

```bash
find . -name "*.log" -print                    # print path (default)
find . -name "*.log" -ls                       # ls -dils style
find . -name "*.tmp" -delete                   # delete matched files
find . -type f -exec chmod 644 {} \;           # run command per file
find . -type f -exec chmod 644 {} +            # batch files into one command
find . -name "*.py" -exec grep "TODO" {} +
```

## fa-bolt fd Basic Usage

```bash
fd pattern                     # search by name in current dir
fd "\.jpg$"                    # regex pattern
fd -e txt pattern              # by extension
fd -i readme                   # case-insensitive
fd -g '*.py'                   # glob pattern
```

## fa-sliders fd Advanced

```bash
fd -t f pattern                # files only
fd -t d pattern                # directories only
fd -t l pattern                # symlinks only
fd -d 3 pattern                # max depth 3
fd -S +1G pattern              # larger than 1GB
fd --changed-within 1week      # modified recently
fd -x wc -l                    # execute command per result
fd -X rm                       # execute command with all results
```

## fa-link xargs + find

```bash
find . -name "*.log" -print0 | xargs -0 rm
find . -type f -name "*.js" -print0 | xargs -0 wc -l
find . -name "*.bak" -print0 | xargs -0 -I {} mv {} /tmp/backup/
find . -type f -size +0 -print0 | xargs -0 grep -l "pattern"
```

## fa-ban Prune & Exclude

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

## fa-lightbulb Practical Examples

```bash
# Find and archive old logs
find /var/log -name "*.log" -mtime +30 | tar -czf old-logs.tar.gz -T -

# List top 20 largest files
find . -type f -printf '%s %p\n' | sort -rn | head -20

# Find duplicate filenames across directories
fd -t f -x basename | sort | uniq -d

# Find files with specific content
fd -e py -x grep -l "import django"

# Count files by extension
fd -t f | sed 's/.*\.//' | sort | uniq -c | sort -rn

# Find broken symlinks
find . -type l ! -exec test -e {} \; -print

# Change ownership recursively
find /var/www -type f -exec chown www-data:www-data {} +
```
