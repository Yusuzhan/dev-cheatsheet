---
title: C
icon: fa-c
primary: "#A8B9CC"
lang: c
locale: zhs
---

## fa-box 变量与类型

```c
int x = 10;
long big = 1000000L;
unsigned int age = 30;
float f = 3.14f;
double d = 3.1415926535;
char ch = 'A';
char name[] = "Alice";

sizeof(int);
sizeof(double);
```

## fa-calculator 运算符

```c
a + b; a - b; a * b; a / b; a % b;
a++; b--; ++a; --b;

a == b; a != b; a < b; a > b; a <= b; a >= b;
a && b; a || b; !a;

x = (cond) ? val1 : val2;

a & b; a | b; a ^ b; ~a;
a << 2; a >> 2;
```

## fa-code-branch 控制流

```c
if (x > 0) {
    printf("positive\n");
} else if (x == 0) {
    printf("zero\n");
} else {
    printf("negative\n");
}

for (int i = 0; i < 10; i++) {
    printf("%d\n", i);
}

while (condition) {
    /* ... */
}

do {
    /* ... */
} while (condition);

switch (x) {
    case 1: printf("one\n"); break;
    case 2: printf("two\n"); break;
    default: printf("other\n"); break;
}
```

## fa-pen-to-square 函数

```c
int add(int a, int b) {
    return a + b;
}

void greet(const char *name) {
    printf("Hello, %s!\n", name);
}

static int counter = 0;

int sum(int count, ...) {
    va_list args;
    va_start(args, count);
    int total = 0;
    for (int i = 0; i < count; i++)
        total += va_arg(args, int);
    va_end(args);
    return total;
}
```

## fa-arrow-pointer 指针

```c
int x = 42;
int *p = &x;
printf("%d\n", *p);
*p = 100;

int **pp = &p;

void swap(int *a, int *b) {
    int tmp = *a;
    *a = *b;
    *b = tmp;
}

const char *str = "hello";
char *const ptr = str;
const char *const cptr = str;

int arr[5] = {1, 2, 3, 4, 5};
int *p = arr;
*(p + 2);
```

## fa-list 数组与字符串

```c
int nums[5] = {1, 2, 3, 4, 5};
int zeros[10] = {0};
int matrix[3][3] = {{1,2,3},{4,5,6},{7,8,9}};

char str[] = "hello";
char buf[256];
snprintf(buf, sizeof(buf), "value: %d", x);

strlen(str);
strcmp(a, b);
strncmp(a, b, n);
strcpy(dst, src);
strncpy(dst, src, n);
strcat(dst, src);
strstr(haystack, needle);
strchr(str, 'o');
```

## fa-cubes 结构体与联合体

```c
struct Point {
    double x;
    double y;
};

struct Point p = {1.0, 2.0};
struct Point *pp = &p;
pp->x = 3.0;

typedef struct {
    char name[50];
    int age;
} Person;

Person people[10];

union Data {
    int i;
    float f;
    char c[4];
};

union Data d;
d.i = 42;
printf("%f\n", d.f);
```

## fa-memory-stick 动态内存

```c
int *arr = malloc(10 * sizeof(int));
if (!arr) { perror("malloc"); exit(1); }

arr = realloc(arr, 20 * sizeof(int));

int *zeros = calloc(10, sizeof(int));

free(arr);
arr = NULL;

int **matrix = malloc(rows * sizeof(int *));
for (int i = 0; i < rows; i++)
    matrix[i] = malloc(cols * sizeof(int));

for (int i = 0; i < rows; i++)
    free(matrix[i]);
free(matrix);
```

## fa-file 文件读写

```c
FILE *f = fopen("file.txt", "r");
if (!f) { perror("fopen"); exit(1); }

char line[256];
while (fgets(line, sizeof(line), f)) {
    printf("%s", line);
}

fprintf(f, "Hello %s\n", name);
fscanf(f, "%d %s", &age, name);

fclose(f);

size_t n = fwrite(buf, 1, size, f);
size_t n = fread(buf, 1, size, f);

fseek(f, 0, SEEK_SET);
ftell(f);
```

## fa-hashtag 预处理器

```c
#include <stdio.h>
#include "myheader.h"

#define PI 3.14159
#define MAX(a, b) ((a) > (b) ? (a) : (b))
#define STR(x) #x
#define CAT(a, b) a##b

#ifdef DEBUG
    #define LOG(msg) printf("DEBUG: %s\n", msg)
#else
    #define LOG(msg)
#endif

#pragma once
```

## fa-tag typedef 与 enum

```c
typedef unsigned long usize;
typedef int (*Comparator)(const void *, const void *);
typedef struct Node Node;

enum Color { RED, GREEN, BLUE };
typedef enum { OK = 0, ERR = -1 } Status;

enum Color c = GREEN;
```

## fa-arrow-right-arrow-left 函数指针

```c
int (*op)(int, int);

int add(int a, int b) { return a + b; }
int mul(int a, int b) { return a * b; }

op = add;
printf("%d\n", op(3, 4));

void sort(int *arr, int n, int (*cmp)(int, int)) {
    /* ... */
}

typedef void (*Callback)(int event);
void register_cb(Callback cb);
```

## fa-gears 位操作

```c
unsigned int set_bit(unsigned int x, int n)    { return x | (1U << n); }
unsigned int clear_bit(unsigned int x, int n)  { return x & ~(1U << n); }
unsigned int toggle_bit(unsigned int x, int n) { return x ^ (1U << n); }
int get_bit(unsigned int x, int n)             { return (x >> n) & 1; }

unsigned int mask = 0xFF00;
x = (x & mask) >> 8;

int is_power_of_two(unsigned int x) { return x && !(x & (x - 1)); }
```

## fa-wrench 常用模式

```c
int *arr = malloc(n * sizeof(int));
if (!arr) { perror("malloc"); return -1; }

#define ARRAY_LEN(arr) (sizeof(arr) / sizeof((arr)[0]))

qsort(arr, n, sizeof(int), cmp_func);

void *found = bsearch(&key, arr, n, sizeof(int), cmp_func);

assert(ptr != NULL);

srand(time(NULL));
int r = rand() % 100;

snprintf(buf, sizeof(buf), "%s/%s", dir, file);
```
