---
title: Grep
icon: fa-magnifying-glass
primary: "#F05032"
lang: bash
---

## fa-magnifying-glass Basic Search

```bash
# Search in a file
grep "pattern" file.txt

# Case-insensitive search
grep -i "pattern" file.txt

# Search in multiple files
grep "pattern" file1.txt file2.txt

# Search using stdin
cat file.txt | grep "pattern"
echo "hello world" | grep "hello"
```

## fa-folder-tree Recursive Search

```bash
# Search recursively in directory
grep -r "pattern" /path/to/dir

# Follow symlinks
grep -R "pattern" /path/to/dir

# Recursive with file pattern
grep -r --include="*.py" "pattern" /path/to/dir

# Exclude directories
grep -r --exclude-dir={node_modules,.git} "pattern" .

# Exclude file patterns
grep -r --exclude="*.log" "pattern" .
```

## fa-code Regular Expressions

```bash
# Basic regex
grep "^[A-Z]" file.txt           # lines starting with uppercase
grep "[0-9]\{3\}" file.txt       # lines with 3 consecutive digits
grep "end$" file.txt             # lines ending with "end"

# Extended regex (-E)
grep -E "cat|dog" file.txt       # match cat OR dog
grep -E "go+d" file.txt          # match god, good, gooood...
grep -E "colou?r" file.txt       # match color or colour
grep -E "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b" file.txt

# Perl-compatible regex (-P)
grep -P "\d{3}-\d{4}" file.txt   # match phone-like patterns
grep -P "(?<=error: ).*" log.txt # lookbehind assertion
```

## fa-arrows-up-down Context Lines

```bash
# Show N lines after match
grep -A 3 "pattern" file.txt

# Show N lines before match
grep -B 3 "pattern" file.txt

# Show N lines before and after
grep -C 3 "pattern" file.txt

# Show only matching part (-o) with context
grep -o -E "[0-9]+" file.txt     # extract all numbers
```

## fa-file File Filtering

```bash
# Only show filenames with matches
grep -l "pattern" *.txt

# Only show filenames WITHOUT matches
grep -L "pattern" *.txt

# Search only specific file types
grep "pattern" *.py
grep "pattern" *.md

# Recursive with include/exclude
grep -r --include="*.{js,ts}" "pattern" .
grep -r --exclude="*.min.js" "pattern" .
```

## fa-calculator Invert & Count

```bash
# Invert match (non-matching lines)
grep -v "pattern" file.txt

# Count matching lines
grep -c "pattern" file.txt

# Show line numbers
grep -n "pattern" file.txt

# Show only the matched parts
grep -o "pattern" file.txt

# Suppress error messages
grep -s "pattern" /root/*
```

## fa-gear Search Options

```bash
# Whole word match
grep -w "error" file.txt         # matches "error" not "errors"

# Whole line match
grep -x "exact line" file.txt

# Match fixed strings (no regex)
grep -F ".*" file.txt            # literally matches .*

# Binary files as text
grep -a "pattern" binary_file

# Read compressed files
zgrep "pattern" file.gz
```

## fa-display Output Control

```bash
# Colorized output
grep --color=auto "pattern" file.txt

# Prefix with filename (useful with multiple files)
grep -H "pattern" *.txt

# No filename prefix
grep -h "pattern" *.txt

# Line number with byte offset
grep -b "pattern" file.txt

# Suppress prefix entirely (for piping)
grep -o "pattern" file.txt | sort | uniq -c | sort -rn
```

## fa-wand-magic-sparkles Advanced Patterns

```bash
# Email addresses
grep -E "[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}" file.txt

# IP addresses
grep -E "[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}" file.txt

# Hex color codes
grep -E "#[0-9a-fA-F]{6}" file.txt

# URL patterns
grep -E "https?://[^[:space:]]+" file.txt

# Empty lines
grep -E "^$" file.txt

# Duplicate consecutive lines
uniq -d file.txt
```

## fa-lightbulb Practical Examples

```bash
# Find all TODO comments in a project
grep -rn "TODO\|FIXME\|HACK" --include="*.{py,js,ts,go}" .

# Search process output
ps aux | grep nginx

# Search command history
history | grep "docker"

# Find large files and filter
find / -type f -size +100M 2>/dev/null | grep -i log

# Extract HTTP status codes from access log
grep -oE " [0-9]{3} " access.log | sort | uniq -c | sort -rn

# Find files not matching a pattern
grep -rL "license" --include="*.py" .

# Count occurrences across files
grep -ro "function" --include="*.js" . | wc -l
```
