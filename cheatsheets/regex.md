---
title: Regex
icon: fa-asterisk
primary: "#E44D26"
lang: bash
---

## fa-bullseye Character Classes

```regex
.         any character except newline
\d        digit [0-9]
\D        non-digit [^0-9]
\w        word character [a-zA-Z0-9_]
\W        non-word character
\s        whitespace (space, tab, newline)
\S        non-whitespace
[abc]     a, b, or c
[^abc]    not a, b, or c
[a-z]     lowercase letters a to z
[A-Z0-9]  uppercase letters or digits
```

## fa-repeat Quantifiers

```regex
a*        zero or more
a+        one or more
a?        zero or one (optional)
a{3}      exactly 3
a{3,5}    between 3 and 5
a{3,}     3 or more
a*?       zero or more (lazy)
a+?       one or more (lazy)
a{3,5}?   between 3 and 5 (lazy)
```

## fa-map-pin Anchors & Boundaries

```regex
^abc      start of string/line
abc$      end of string/line
\A        start of string only
\z        end of string only
\b        word boundary
\B        non-word boundary

# examples
^\d+      string starts with digits
\.$       string ends with a dot
\bcat\b   match "cat" as whole word
```

## fa-code-branch Groups & References

```regex
(abc)       capturing group
(?:abc)     non-capturing group
(?<name>abc)  named capturing group
\1          backreference to group 1
\2          backreference to group 2
$1          group 1 in replacement

# examples
(\w+)\s+\1           # match repeated word
(\d{4})-(\d{2})-(\d{2})   # date: $1-$2-$3
(?<year>\d{4})       # named group "year"
```

## fa-code Alternation & Lookaround

```regex
cat|dog        cat or dog
(?=foo)        lookahead: followed by foo
(?!bar)        negative lookahead: not followed by bar
(?<=foo)       lookbehind: preceded by foo
(?<!bar)       negative lookbehind: not preceded by bar

# examples
\d+(?=px)           # digits followed by "px"
\d+(?!px)           # digits NOT followed by "px"
(?<=\$)\d+          # digits preceded by "$"
(?<!un)\w+          # word NOT preceded by "un"
```

## fa-flag Flags

```regex
/pattern/gi     g = global (all matches)
                i = case-insensitive
                m = multiline (^/$ match line start/end)
                s = dotall (. matches newline)
                x = verbose (ignore whitespace)
                u = unicode mode

# common usage
/foo/g          # find all "foo" in string
/^bar/m         # match "bar" at start of any line
/./s            # dot matches everything including \n
```

## fa-shield Escaping

```regex
\.       literal dot
\\       literal backslash
\*       literal asterisk
\+       literal plus
\?       literal question mark
\^       literal caret
\$       literal dollar sign
\[       literal bracket
\]       literal bracket
\(       literal parenthesis
\)       literal parenthesis
\|       literal pipe
```

## fa-table Common Patterns

```regex
# email (basic)
[\w.+-]+@[\w-]+\.[\w.]+

# IPv4 address
\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}

# date (YYYY-MM-DD)
\d{4}-(?:0[1-9]|1[0-2])-(?:0[1-9]|[12]\d|3[01])

# URL
https?:\/\/[\w\-]+(?:\.[\w\-]+)+(?:\/[\w\-._~:?#@!$&'()*+,;=]*)?

# hex color
#[0-9a-fA-F]{3,8}

# phone (US)
\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4}
```

## fa-terminal Usage in Code

```python
import re
re.search(r'\d+', 'abc123')          # first match
re.findall(r'\d+', 'a1b22c333')      # all matches
re.sub(r'\d+', 'X', 'a1b22')         # replace
re.match(r'^\w+', 'hello world')     # match at start
```

```javascript
'abc123'.match(/\d+/)                // first match
'abc123'.match(/\d+/g)               // all matches
'abc123'.replace(/\d+/, 'X')         // replace
/^hello/.test('hello world')         // test match
```

```go
import "regexp"
re := regexp.MustCompile(`\d+`)
re.FindString("abc123")              // "123"
re.FindAllString("a1b22c333", -1)    // ["1", "22", "333"]
re.ReplaceAllString("a1b22", "X")    // "aXbX"
re.MatchString("abc123")             // true
```

## fa-lightbulb Tips & Tricks

```regex
# match everything between quotes
"[^"]*"

# match HTML/XML tag
<(\w+)[^>]*>.*?<\/\1>

# password (min 8, uppercase, lowercase, digit)
^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$

# trim whitespace
^\s+|\s+$

# capture filename without extension
^(.+?)(?:\.[^.]+)?$

# number with optional decimals
-?\d+(?:\.\d+)?
```
