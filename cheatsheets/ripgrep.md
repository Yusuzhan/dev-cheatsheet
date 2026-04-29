---
title: Ripgrep
icon: fa-bolt
primary: "#FF2D20"
lang: bash
---

## fa-magnifying-glass Basic Search

```bash
rg "pattern"                      # search in current directory
rg "pattern" path/to/dir          # search in specific directory
rg -i "pattern"                   # case-insensitive
rg -w "word"                      # whole word match
rg "pattern" file.txt             # search in a single file
```

## fa-file File Filtering

```bash
rg -t py "pattern"                # search Python files
rg -T py "pattern"                # exclude Python files
rg -t py -t js "pattern"          # search Python and JS files
rg -g "*.toml" "pattern"          # glob include
rg -g "!*.min.js" "pattern"       # glob exclude
rg -g "*.{rs,toml}" "pattern"     # multiple extensions
```

## fa-folder-tree Directory Control

```bash
rg --max-depth 3 "pattern"        # limit search depth
rg -d 3 "pattern"                 # shorthand for max-depth
rg "pattern" --ignore-file .customignore
rg --no-ignore "pattern"          # ignore .gitignore
rg --no-ignore-vcs "pattern"      # ignore VCS ignore files only
```

## fa-code Regex Patterns

```bash
rg "start.*end"                   # basic regex
rg -e "pat1" -e "pat2"            # multiple patterns (OR)
rg -f patterns.txt                # patterns from file
rg -F "literal.*string"           # fixed string (no regex)
rg -U "line1\nline2"              # multi-line mode
rg "\bword\b"                     # word boundary
```

## fa-arrows-up-down Context & Display

```bash
rg -C 3 "pattern"                 # 3 lines before and after
rg -B 2 "pattern"                 # 2 lines before
rg -A 2 "pattern"                 # 2 lines after
rg --context-separator "---" "pattern"
rg -n "pattern"                   # show line numbers (default)
rg --no-line-number "pattern"     # hide line numbers
```

## fa-display Output Formats

```bash
rg -l "pattern"                   # list filenames only
rg --files-without-match "pattern"
rg -c "pattern"                   # count matches per file
rg --count-matches "pattern"      # count total matches
rg -o "pattern"                   # show only matching parts
rg --json "pattern"               # JSON output
rg --vimgrep "pattern"            # vim quickfix format
```

## fa-file-code File Types

```bash
rg --type-list                    # list all file types
rg -t css "color"                 # CSS files
rg -t html "href"                 # HTML files
rg -t rust "fn main"              # Rust files
rg -t markdown "TODO"             # Markdown files
rg -t sh "echo"                   # Shell scripts
```

## fa-eye-slash Binary & Hidden

```bash
rg --hidden "pattern"             # search hidden files
rg --no-hidden "pattern"          # exclude hidden files (default)
rg --binary "pattern"             # search binary files
rg --no-binary "pattern"          # skip binary files (default)
rg -a "pattern"                   # treat all files as text
```

## fa-arrows-rotate Replacing

```bash
rg -r "replacement" "pattern"     # preview replacement (no file change)
rg -r '$1' "(\w+)"                # use capture groups
rg -r '$1$new$3' "(old)(.*)($)" "pattern"
```

## fa-sort-amount-down Sort & Limit

```bash
rg --sort path "pattern"          # sort by file path
rg --sortr modified "pattern"     # sort by modification time (desc)
rg --sort accessed "pattern"      # sort by access time
rg -m 100 "pattern"               # max 100 matches
rg -m 0 "pattern"                 # unlimited matches
```

## fa-gear Configuration

```bash
rg --config-file ~/.ripgreprc     # use config file
rg --no-config "pattern"          # ignore config
echo '--smart-case' >> ~/.ripgreprc
echo '--max-columns=150' >> ~/.ripgreprc
echo '--colors="match:fg:red"' >> ~/.ripgreprc
```

## fa-gauge-high Performance Tips

```bash
rg -j 4 "pattern"                 # use 4 threads
rg --max-filesize 10M "pattern"   # skip files larger than 10MB
rg --mmap "pattern"               # use memory-mapped I/O
rg --no-mmap "pattern"            # disable memory-mapped I/O
rg --fast "pattern"               # reduce regex complexity
```

## fa-lightbulb Practical Examples

```bash
# Find all TODO/FIXME in code
rg -n "TODO|FIXME|HACK" -t py -t rs -t js

# Find function definitions
rg "fn \w+" -t rust
rg "(def|class) \w+" -t py
rg "function \w+" -t js

# Search for IP addresses
rg "\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}" access.log

# Find secret patterns
rg -i "password|secret|api_key|token" --type-add 'lock:*.lock' -t lock

# Extract URLs from files
rg -o "https?://[^\s)]+" -N --no-filename

# Count pattern per file, sorted
rg -c "error" --sort path

# Search in recently modified files
rg "pattern" --sortr modified --max-depth 5

# Preview find-and-replace
rg -r "newFunc" "oldFunc" -t go
```
