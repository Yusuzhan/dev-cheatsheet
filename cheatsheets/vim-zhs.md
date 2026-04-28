---
title: Vim
icon: fa-terminal
primary: "#019733"
lang: vim
locale: zhs
---

## fa-arrow-right 模式切换

```vim
i           " 在光标前进入插入模式
a           " 在光标后进入插入模式
I           " 在行首进入插入模式
A           " 在行尾进入插入模式
o           " 在下方新开一行并进入插入模式
O           " 在上方新开一行并进入插入模式
v           " 可视模式（字符选择）
V           " 可视模式（行选择）
Ctrl+v      " 可视模式（块选择）
Esc         " 返回普通模式
```

## fa-arrows-up-down-left-right 光标移动

```vim
h j k l             " 左、下、上、右
w                   " 下一个单词开头
b                   " 上一个单词开头
e                   " 单词结尾
0                   " 行首
^                   " 第一个非空白字符
$                   " 行尾
gg                  " 跳到第一行
G                   " 跳到最后一行
5G                  " 跳到第 5 行
%                   " 跳到匹配的括号
fx                  " 当前行下一个字符 x
```

## fa-scissors 剪切、复制与粘贴

```vim
dd                  " 删除（剪切）当前行
3dd                 " 删除 3 行
dw                  " 删除单词
d$                  " 删除到行尾
D                   " 删除到行尾（同 d$）
x                   " 删除光标处字符
X                   " 删除光标前字符

yy                  " 复制当前行
3yy                 " 复制 3 行
yw                  " 复制单词
y$                  " 复制到行尾

p                   " 在光标后粘贴
P                   " 在光标前粘贴
```

## fa-magnifying-glass 搜索与替换

```vim
/pattern            " 向下搜索
?pattern            " 向上搜索
n                   " 下一个匹配
N                   " 上一个匹配
*                   " 向下搜索光标下单词
#                   " 向上搜索光标下单词

:%s/old/new/g       " 全文替换
:%s/old/new/gc      " 全文替换（逐个确认）
:s/old/new/g        " 当前行替换
:5,10s/old/new/g    " 第 5-10 行替换
:noh                " 清除搜索高亮
```

## fa-rotate 撤销与重复

```vim
u                   " 撤销
Ctrl+r              " 重做
.                   " 重复上次修改
```

## fa-font 文本对象

```vim
diw                 " 删除光标所在单词
daw                 " 删除光标所在单词及周围空格
di"                 " 删除双引号内的内容
da"                 " 删除双引号及其内容
di(                 " 删除括号内的内容
da{                 " 删除花括号及其内容
ciw                 " 修改光标所在单词
ci"                 " 修改引号内的内容
yiw                 " 复制光标所在单词
vit                 " 选中 HTML 标签内的内容
vat                 " 选中 HTML 标签及其内容
```

## fa-window-restore 窗口与标签页

```vim
:sp file.txt        " 水平分屏
:vsp file.txt       " 垂直分屏
Ctrl+w h/j/k/l      " 在分屏间移动
Ctrl+w =            " 等比分屏大小
Ctrl+w +            " 增加分屏高度
Ctrl+w -            " 减小分屏高度
:q                   " 关闭分屏/窗口
:only                " 关闭其他所有分屏

:tabnew file.txt    " 在新标签页打开文件
:tabn               " 下一个标签页
:tabp               " 上一个标签页
gt                  " 下一个标签页（普通模式）
gT                  " 上一个标签页（普通模式）
:tabo               " 关闭其他所有标签页
```

## fa-file 文件操作

```vim
:w                  " 保存文件
:w file.txt         " 另存为
:q                  " 退出
:q!                 " 不保存退出
:wq                 " 保存并退出
:x                  " 保存并退出（仅修改时保存）
:e file.txt         " 打开文件
:e!                 " 重新加载（丢弃修改）
:bn                 " 下一个缓冲区
:bp                 " 上一个缓冲区
:ls                 " 列出所有缓冲区
:bd                 " 关闭缓冲区
```

## fa-code-branch 标记与跳转

```vim
ma                  " 在当前行设置标记 'a'
'a                  " 跳到标记 'a'（行首）
`a                  " 跳到标记 'a'（精确位置）
:marks              " 列出所有标记

Ctrl+o              " 跳转列表中后退
Ctrl+i              " 跳转列表中前进
:jumps              " 显示跳转列表
```

## fa-wand-magic-sparks 宏与命令

```vim
qa                  " 开始录制宏 'a'
q                   " 停止录制宏
@a                  " 运行宏 'a'
5@a                 " 运行宏 'a' 5 次
@@                  " 重复上次宏

:!ls                " 执行 shell 命令
:r !date            " 将命令输出插入文件
:sort               " 对可视选区排序
:%!jq .             " 通过 jq 格式化 JSON
```

## fa-sliders 缩进与格式化

```vim
>>                  " 当前行右缩进
<<                  " 当前行左缩进
5>>                 " 5 行右缩进
=                   " 自动缩进（可视模式）
gg=G                " 整个文件自动缩进
:set tabstop=4      " Tab 显示宽度
:set shiftwidth=4   " 缩进宽度
:set expandtab      " 用空格替代 Tab
:retab              " 将 Tab 转为空格
```

## fa-lightbulb 实用技巧

```vim
:set number         " 显示行号
:set relativenumber " 显示相对行号
:set hlsearch       " 高亮所有搜索匹配
:set incsearch      " 增量搜索
:set ignorecase     " 搜索忽略大小写
:set smartcase      " 智能大小写
:syntax on          " 启用语法高亮
:set wrap           " 自动换行
:set mouse=a        " 启用鼠标支持

gf                  " 打开光标下的文件
Ctrl+a              " 递增光标下数字
Ctrl+x              " 递减光标下数字
J                   " 将下一行合并到当前行
~                   " 切换光标处字符大小写
gUU                 " 当前行转大写
guu                 " 当前行转小写
```
