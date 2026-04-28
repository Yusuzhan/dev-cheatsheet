---
title: Make
icon: fa-gears
primary: "#6D8086"
lang: makefile
locale: zhs
---

## fa-rocket 基本规则

```makefile
# 目标: 依赖
# 	命令（必须使用 Tab 缩进，不能用空格）

build: main.o utils.o
	gcc -o app main.o utils.o

main.o: main.c
	gcc -c main.c

utils.o: utils.c
	gcc -c utils.c

clean:
	rm -f *.o app

.PHONY: clean                       # 声明伪目标（不对应文件）
```

## fa-circle-dot 变量

```makefile
CC      := gcc
CFLAGS  := -Wall -O2
TARGET  := app
SRCS    := main.c utils.c helper.c
OBJS    := $(SRCS:.c=.o)            # 变量替换：.c → .o

$(TARGET): $(OBJS)
	$(CC) $(CFLAGS) -o $@ $^

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

# 变量赋值方式
# := 即时赋值（定义时求值）
# =  递归赋值（使用时求值）
# ?= 仅在未定义时赋值
# += 追加到变量
```

## fa-code 自动变量

```makefile
# $@  目标名
# $<  第一个依赖
# $^  所有依赖（去重）
# $?  比目标新的依赖
# $*  模式规则匹配的词干

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@       # $< = %.c, $@ = %.o

build: main.o utils.o
	@echo "Building $@ from $^"        # $@ = build, $^ = main.o utils.o
```

## fa-layer-group 模式规则

```makefile
# 编译任意 .c 为 .o
%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

# 从 .o 链接生成可执行文件
%: %.o
	$(CC) $(LDFLAGS) -o $@ $^

# 从 .def 生成 .h
%.h: %.def
	./gen-header.pl $< > $@

# 构建所有 .c 文件
SRCS := $(wildcard src/*.c)
OBJS := $(patsubst src/%.c, build/%.o, $(SRCS))
```

## fa-diagram-project 实战示例

```makefile
CC       := gcc
CFLAGS   := -Wall -Wextra -O2 -Iinclude
LDFLAGS  := -lpthread
SRCS     := $(wildcard src/*.c)
OBJS     := $(SRCS:src/%.c=build/%.o)
TARGET   := bin/app

.PHONY: all clean run install

all: $(TARGET)

$(TARGET): $(OBJS)
	@mkdir -p bin
	$(CC) $(CFLAGS) -o $@ $^ $(LDFLAGS)

build/%.o: src/%.c
	@mkdir -p build
	$(CC) $(CFLAGS) -c $< -o $@

run: $(TARGET)
	./$(TARGET)

clean:
	rm -rf build/ bin/

install: $(TARGET)
	install -m 755 $(TARGET) /usr/local/bin/
```

## fa-wand-magic-sparks 内置函数

```makefile
# 字符串函数
$(subst from,to,text)             # 文本替换
$(patsubst %.c,%.o,x.c y.c)      # 模式替换：x.o y.o
$(strip text)                     # 去除多余空白
$(findstring find,text)           # 查找子串
$(filter %.c %.h,$(FILES))        # 保留匹配的模式
$(filter-out %.o,$(FILES))        # 移除匹配的模式

# 文件名函数
$(dir src/main.c)                 # src/
$(notdir src/main.c)              # main.c
$(basename src/main.c)            # src/main
$(suffix src/main.c)              # .c
$(addprefix build/,$(OBJS))       # build/main.o build/utils.o
$(addsuffix .o,$(BASES))          # main.o utils.o
$(wildcard src/*.c)               # 展开通配符

# 条件函数
$(if $(DEBUG),-g,-O2)             # if-then-else
$(or $(VAR1),$(VAR2),default)     # 第一个非空值
$(and $(VAR1),$(VAR2))            # 全部非空则返回最后一个
```

## fa-arrows-turn-to-dots 条件判断

```makefile
# ifeq / ifneq
ifeq ($(DEBUG),1)
	CFLAGS += -g -DDEBUG
else
	CFLAGS += -O2
endif

# ifdef / ifndef
ifndef CC
	CC := gcc
endif

ifdef VERBOSE
	Q :=
else
	Q := @                          # @ 静默执行命令
endif

# 在命令中使用
build:
	$(Q) echo "building..."
```

## fa-terminal 指令与包含

```makefile
# 包含其他 makefile（找不到会报错停止）
include config.mk

# 包含但找不到不报错
-include local.mk

# 条件包含
ifneq ($(wildcard .env),)
	include .env
	export $(shell sed 's/=.*//' .env)
endif

# 定义多行变量
define HELP_TEXT
用法:
  make build    构建项目
  make test     运行测试
  make clean    清理构建产物
endef
export HELP_TEXT
```

## fa-box VPATH 与目录

```makefile
# 依赖文件搜索路径
VPATH = src:lib:include

# vpath 指令（按模式指定搜索路径）
vpath %.c src
vpath %.h include
vpath %.a lib

# 自动创建目录
MKDIR_P := mkdir -p
BUILD_DIR := build
OBJ_DIR := $(BUILD_DIR)/obj
BIN_DIR := $(BUILD_DIR)/bin

$(OBJ_DIR)/%.o: src/%.c | $(OBJ_DIR)
	$(CC) $(CFLAGS) -c $< -o $@

$(OBJ_DIR):
	$(MKDIR_P) $@
```

## fa-message 调试技巧

```makefile
# 输出消息
@echo "Compiling $<"           # @ 抑制命令回显

# 警告（不停止）
$(warning DEBUG is $(DEBUG))

# 错误（停止执行）
$(error 缺少必要变量 TARGET)

# 打印变量用于调试
print-%:
	@echo '$* = $($*)'             # 用法：make print-CC

# 试运行（显示命令但不执行）
# make -n build

# 调试模式
# make -d build
```

## fa-lightbulb 实用技巧

```makefile
# 并行构建
# make -j$(nproc)

# 命令行传入变量
# make CC=clang CFLAGS="-O3"

# help 目标
help:                              # make help
	@echo "目标列表:"
	@echo "  all      - 构建所有"
	@echo "  test     - 运行测试"
	@echo "  clean    - 清理产物"

# 自文档化 makefile
SRCS := $(wildcard src/*.c)        ## 源文件
TARGET := app                      ## 输出二进制文件
```
