---
title: Vim
icon: fa-terminal
primary: "#019733"
lang: vim
---

## fa-arrow-right Mode Switching

```vim
i           " insert mode before cursor
a           " insert mode after cursor
I           " insert mode at line beginning
A           " insert mode at line end
o           " open new line below, enter insert
O           " open new line above, enter insert
v           " visual mode (character)
V           " visual mode (line)
Ctrl+v      " visual mode (block)
Esc         " back to normal mode
```

## fa-arrows-up-down-left-right Cursor Movement

```vim
h j k l             " left, down, up, right
w                   " next word start
b                   " previous word start
e                   " word end
0                   " line beginning
^                   " first non-blank character
$                   " line end
gg                  " first line
G                   " last line
5G                  " go to line 5
%                   " jump to matching bracket
fx                  " next occurrence of x on line
```

## fa-scissors Cut, Copy & Paste

```vim
dd                  " delete (cut) line
3dd                 " delete 3 lines
dw                  " delete word
d$                  " delete to end of line
D                   " delete to end of line (same as d$)
x                   " delete character under cursor
X                   " delete character before cursor

yy                  " yank (copy) line
3yy                 " yank 3 lines
yw                  " yank word
y$                  " yank to end of line

p                   " paste after cursor
P                   " paste before cursor
```

## fa-magnifying-glass Search & Replace

```vim
/pattern            " search forward
?pattern            " search backward
n                   " next match
N                   " previous match
*                   " search word under cursor forward
#                   " search word under cursor backward

:%s/old/new/g       " replace all in file
:%s/old/new/gc      " replace all with confirmation
:s/old/new/g        " replace in current line
:5,10s/old/new/g    " replace in lines 5-10
:noh                " clear search highlight
```

## fa-rotate Undo & Repeat

```vim
u                   " undo
Ctrl+r              " redo
.                   " repeat last change
```

## fa-font Text Objects

```vim
diw                 " delete inner word
daw                 " delete around word
di"                 " delete inside double quotes
da"                 " delete around double quotes
di(                 " delete inside parentheses
da{                 " delete around braces
ciw                 " change inner word
ci"                 " change inside quotes
yiw                 " yank inner word
vit                 " select inner HTML tag
vat                 " select around HTML tag
```

## fa-window-restore Windows & Tabs

```vim
:sp file.txt        " horizontal split
:vsp file.txt       " vertical split
Ctrl+w h/j/k/l      " navigate between splits
Ctrl+w =            " equal size splits
Ctrl+w +            " increase split height
Ctrl+w -            " decrease split height
:q                   " close split/window
:only                " close all other splits

:tabnew file.txt    " open file in new tab
:tabn               " next tab
:tabp               " previous tab
gt                  " next tab (normal mode)
gT                  " previous tab (normal mode)
:tabo               " close all other tabs
```

## fa-file File Operations

```vim
:w                  " save file
:w file.txt         " save as
:q                  " quit
:q!                 " quit without saving
:wq                 " save and quit
:x                  " save and quit (only if modified)
:e file.txt         " open file
:e!                 " reload file (discard changes)
:bn                 " next buffer
:bp                 " previous buffer
:ls                 " list buffers
:bd                 " close buffer
```

## fa-code-branch Marks & Jumps

```vim
ma                  " set mark 'a' in current line
'a                  " jump to mark 'a' (first col)
`a                  " jump to mark 'a' (exact position)
:marks              " list all marks

Ctrl+o              " jump to older position in jumplist
Ctrl+i              " jump to newer position in jumplist
:jumps              " show jumplist
```

## fa-wand-magic-sparks Macros & Commands

```vim
qa                  " start recording macro 'a'
q                   " stop recording macro
@a                  " run macro 'a'
5@a                 " run macro 'a' 5 times
@@                  " repeat last macro

:!ls                " run shell command
:r !date            " insert command output into file
:sort               " sort lines in visual selection
:%!jq .             " filter file through jq (pretty print JSON)
```

## fa-sliders Indentation & Formatting

```vim
>>                  " indent line right
<<                  " indent line left
5>>                 " indent 5 lines right
=                   " auto-indent (visual mode)
gg=G                " auto-indent entire file
:set tabstop=4      " tab display width
:set shiftwidth=4   " indent width
:set expandtab      " spaces instead of tabs
:retab              " convert tabs to spaces
```

## fa-lightbulb Useful Tips

```vim
:set number         " show line numbers
:set relativenumber " show relative line numbers
:set hlsearch       " highlight all search matches
:set incsearch      " incremental search
:set ignorecase     " case-insensitive search
:set smartcase      " smart case sensitivity
:syntax on          " enable syntax highlighting
:set wrap           " line wrapping
:set mouse=a        " enable mouse support

gf                  " open file under cursor
Ctrl+a              " increment number under cursor
Ctrl+x              " decrement number under cursor
J                   " join line below to current
~                   " toggle case of character
gUU                 " uppercase entire line
guu                 " lowercase entire line
```
