---
title: Regex
icon: fa-asterisk
primary: "#E44D26"
lang: bash
locale: zhs
---

## fa-bullseye 字符类

```regex
.         除换行符外的任意字符
\d        数字 [0-9]
\D        非数字 [^0-9]
\w        单词字符 [a-zA-Z0-9_]
\W        非单词字符
\s        空白字符（空格、制表、换行）
\S        非空白字符
[abc]     a、b 或 c
[^abc]    非 a、b、c
[a-z]     小写字母 a 到 z
[A-Z0-9]  大写字母或数字
```

## fa-repeat 量词

```regex
a*        零次或多次
a+        一次或多次
a?        零次或一次（可选）
a{3}      恰好 3 次
a{3,5}    3 到 5 次
a{3,}     3 次及以上
a*?       零次或多次（懒惰模式）
a+?       一次或多次（懒惰模式）
a{3,5}?   3 到 5 次（懒惰模式）
```

## fa-map-pin 锚点与边界

```regex
^abc      字符串/行首
abc$      字符串/行尾
\A        仅字符串开头
\z        仅字符串结尾
\b        单词边界
\B        非单词边界

# 示例
^\d+      以数字开头
\.$       以点号结尾
\bcat\b   精确匹配 "cat" 整个单词
```

## fa-code-branch 分组与引用

```regex
(abc)       捕获分组
(?:abc)     非捕获分组
(?<name>abc)  命名捕获分组
\1          反向引用第 1 组
\2          反向引用第 2 组
$1          替换中引用第 1 组

# 示例
(\w+)\s+\1           # 匹配重复单词
(\d{4})-(\d{2})-(\d{2})   # 日期: $1-$2-$3
(?<year>\d{4})       # 命名分组 "year"
```

## fa-code 或与前后断言

```regex
cat|dog        cat 或 dog
(?=foo)        正向先行断言：后面跟着 foo
(?!bar)        负向先行断言：后面不跟 bar
(?<=foo)       正向后行断言：前面是 foo
(?<!bar)       负向后行断言：前面不是 bar

# 示例
\d+(?=px)           # 后面跟着 "px" 的数字
\d+(?!px)           # 后面不跟 "px" 的数字
(?<=\$)\d+          # 前面是 "$" 的数字
(?<!un)\w+          # 前面不是 "un" 的单词
```

## fa-flag 修饰符

```regex
/pattern/gi     g = 全局匹配（匹配所有）
                i = 忽略大小写
                m = 多行模式（^/$ 匹配行首行尾）
                s = 点号通配（. 匹配换行符）
                x = 宽松模式（忽略空白）
                u = Unicode 模式

# 常见用法
/foo/g          # 查找所有 "foo"
/^bar/m         # 匹配每行开头的 "bar"
/./s            # 点号匹配包括 \n 的所有字符
```

## fa-shield 转义字符

```regex
\.       匹配字面点号
\\       匹配字面反斜杠
\*       匹配字面星号
\+       匹配字面加号
\?       匹配字面问号
\^       匹配字面脱字符
\$       匹配字面美元符
\[       匹配字面左方括号
\]       匹配字面右方括号
\(       匹配字面左括号
\)       匹配字面右括号
\|       匹配字面管道符
```

## fa-table 常用模式

```regex
# 邮箱（基础）
[\w.+-]+@[\w-]+\.[\w.]+

# IPv4 地址
\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}

# 日期 (YYYY-MM-DD)
\d{4}-(?:0[1-9]|1[0-2])-(?:0[1-9]|[12]\d|3[01])

# URL
https?:\/\/[\w\-]+(?:\.[\w\-]+)+(?:\/[\w\-._~:?#@!$&'()*+,;=]*)?

# 十六进制颜色
#[0-9a-fA-F]{3,8}

# 手机号（中国）
1[3-9]\d{9}
```

## fa-terminal 编程语言用法

```python
import re
re.search(r'\d+', 'abc123')          # 第一个匹配
re.findall(r'\d+', 'a1b22c333')      # 所有匹配
re.sub(r'\d+', 'X', 'a1b22')         # 替换
re.match(r'^\w+', 'hello world')     # 从开头匹配
```

```javascript
'abc123'.match(/\d+/)                // 第一个匹配
'abc123'.match(/\d+/g)               // 所有匹配
'abc123'.replace(/\d+/, 'X')         // 替换
/^hello/.test('hello world')         // 测试是否匹配
```

```go
import "regexp"
re := regexp.MustCompile(`\d+`)
re.FindString("abc123")              // "123"
re.FindAllString("a1b22c333", -1)    // ["1", "22", "333"]
re.ReplaceAllString("a1b22", "X")    // "aXbX"
re.MatchString("abc123")             // true
```

## fa-lightbulb 实用技巧

```regex
# 匹配引号内的内容
"[^"]*"

# 匹配 HTML/XML 标签
<(\w+)[^>]*>.*?<\/\1>

# 密码（至少8位，含大小写和数字）
^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$

# 去除首尾空白
^\s+|\s+$

# 提取文件名（不含扩展名）
^(.+?)(?:\.[^.]+)?$

# 数字（含可选小数和负号）
-?\d+(?:\.\d+)?
```
