---
title: Make
icon: fa-gears
primary: "#6D8086"
lang: makefile
---

## fa-rocket Basic Rules

```makefile
# target: prerequisites
# 	recipe (must use TAB, not spaces)

build: main.o utils.o
	gcc -o app main.o utils.o

main.o: main.c
	gcc -c main.c

utils.o: utils.c
	gcc -c utils.c

clean:
	rm -f *.o app

.PHONY: clean
```

## fa-circle-dot Variables

```makefile
CC      := gcc
CFLAGS  := -Wall -O2
TARGET  := app
SRCS    := main.c utils.c helper.c
OBJS    := $(SRCS:.c=.o)

$(TARGET): $(OBJS)
	$(CC) $(CFLAGS) -o $@ $^

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

# variable flavors
# := immediate assignment (evaluated once)
# =  recursive assignment (evaluated when used)
# ?= set only if not already defined
# += append to variable
```

## fa-code Automatic Variables

```makefile
# $@  target name
# $<  first prerequisite
# $^  all prerequisites (no duplicates)
# $?  prerequisites newer than target
# $*  stem of pattern rule match

%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@       # $< = %.c, $@ = %.o

build: main.o utils.o
	@echo "Building $@ from $^"        # $@ = build, $^ = main.o utils.o
```

## fa-layer-group Pattern Rules

```makefile
# compile any .c to .o
%.o: %.c
	$(CC) $(CFLAGS) -c $< -o $@

# link from .o files
%: %.o
	$(CC) $(LDFLAGS) -o $@ $^

# generate .h from .def
%.h: %.def
	./gen-header.pl $< > $@

# build all .c files
SRCS := $(wildcard src/*.c)
OBJS := $(patsubst src/%.c, build/%.o, $(SRCS))
```

## fa-diagram-project Real-World Example

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

## fa-wand-magic-sparks Built-in Functions

```makefile
# string functions
$(subst from,to,text)             # replace text
$(patsubst %.c,%.o,x.c y.c)      # pattern substitution: x.o y.o
$(strip text)                     # remove extra whitespace
$(findstring find,text)           # search substring
$(filter %.c %.h,$(FILES))        # keep matching patterns
$(filter-out %.o,$(FILES))        # remove matching patterns

# filename functions
$(dir src/main.c)                 # src/
$(notdir src/main.c)              # main.c
$(basename src/main.c)            # src/main
$(suffix src/main.c)              # .c
$(addprefix build/,$(OBJS))       # build/main.o build/utils.o
$(addsuffix .o,$(BASES))          # main.o utils.o
$(wildcard src/*.c)               # expand glob

# conditional
$(if $(DEBUG),-g,-O2)             # if-then-else
$(or $(VAR1),$(VAR2),default)     # first non-empty
$(and $(VAR1),$(VAR2))            # last if all non-empty
```

## fa-arrows-turn-to-dots Conditionals

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
	Q := @
endif

# use in recipe
build:
	$(Q) echo "building..."
```

## fa-terminal Directives & Includes

```makefile
# include other makefiles (stops if not found)
include config.mk

# include but don't error if missing
-include local.mk

# conditional include
ifneq ($(wildcard .env),)
	include .env
	export $(shell sed 's/=.*//' .env)
endif

# define multi-line variable
define HELP_TEXT
Usage:
  make build    Build the project
  make test     Run tests
  make clean    Remove build artifacts
endef
export HELP_TEXT
```

## fa-box VPATH & Directories

```makefile
# search paths for prerequisites
VPATH = src:lib:include

# vpath directive (pattern-specific)
vpath %.c src
vpath %.h include
vpath %.a lib

# create directories automatically
MKDIR_P := mkdir -p
BUILD_DIR := build
OBJ_DIR := $(BUILD_DIR)/obj
BIN_DIR := $(BUILD_DIR)/bin

$(OBJ_DIR)/%.o: src/%.c | $(OBJ_DIR)
	$(CC) $(CFLAGS) -c $< -o $@

$(OBJ_DIR):
	$(MKDIR_P) $@
```

## fa-message Debugging

```makefile
# echo messages
@echo "Compiling $<"           # @ suppresses command echo

# warning (does not stop)
$(warning DEBUG is $(DEBUG))

# error (stops make)
$(error Missing required variable TARGET)

# print variable for debugging
print-%:
	@echo '$* = $($*)'             # usage: make print-CC

# dry run (show commands without executing)
# make -n build

# debug mode
# make -d build
```

## fa-lightbulb Useful Tips

```makefile
# parallel builds
# make -j$(nproc)

# pass variables from command line
# make CC=clang CFLAGS="-O3"

# help target
help:                              # make help
	@echo "Targets:"
	@echo "  all      - build everything"
	@echo "  test     - run tests"
	@echo "  clean    - remove artifacts"

# self-documenting makefile
SRCS := $(wildcard src/*.c)        ## source files
TARGET := app                      ## output binary
```
