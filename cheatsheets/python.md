---
title: Python
icon: fa-python
primary: "#3776AB"
lang: python
---

## fa-box Variables & Types

```python
x = 10
y = 3.14
name = "Alice"
flag = True
z = None
a, b = 1, 2

print(type(x))
```

## fa-font Strings & f-strings

```python
s = "hello world"
s.upper()
s.split()
s.replace("hello", "hi")
s.startswith("hello")
", ".join(["a", "b", "c"])

name = "Alice"
age = 30
f"{name} is {age} years old"
f"{'centered':^20}"
f"{3.14159:.2f}"
```

## fa-list Lists & Tuples

```python
nums = [1, 2, 3]
nums.append(4)
nums.extend([5, 6])
nums.pop()
nums.sort()
nums.reverse()
nums[1:3]
nums[-1]

point = (10, 20)
x, y = point
len(point)
```

## fa-database Dicts & Sets

```python
d = {"name": "Alice", "age": 30}
d["name"]
d.get("email", "N/A")
d.keys()
d.values()
d.items()
d.update({"age": 31})

s = {1, 2, 3}
s.add(4)
s.discard(2)
a | b
a & b
a - b
```

## fa-code-branch Control Flow

```python
if x > 0:
    print("positive")
elif x == 0:
    print("zero")
else:
    print("negative")

for i in range(10):
    print(i)

for item in iterable:
    if skip_condition:
        continue
    if stop_condition:
        break

while x > 0:
    x -= 1

match value:
    case 1:
        print("one")
    case _:
        print("other")
```

## fa-pen-to-square Functions

```python
def greet(name, greeting="Hello"):
    return f"{greeting}, {name}"

def add(*args, **kwargs):
    return sum(args)

add(1, 2, 3, extra=4)

square = lambda x: x ** 2

def outer(x):
    def inner(y):
        return x + y
    return inner
```

## fa-layer-group List/Dict Comprehensions

```python
squares = [x**2 for x in range(10)]
evens = [x for x in range(20) if x % 2 == 0]

matrix = [[1, 2], [3, 4]]
flat = [x for row in matrix for x in row]

word_len = {w: len(w) for w in ["hello", "world"]}
unique = {x for x in [1, 2, 2, 3]}
```

## fa-cubes Classes

```python
class Animal:
    count = 0

    def __init__(self, name):
        self.name = name
        Animal.count += 1

    def speak(self):
        raise NotImplementedError

    def __repr__(self):
        return f"Animal({self.name})"

class Dog(Animal):
    def speak(self):
        return f"{self.name} says Woof!"

d = Dog("Rex")
isinstance(d, Animal)
issubclass(Dog, Animal)
```

## fa-triangle-exclamation Error Handling

```python
try:
    result = 1 / 0
except ZeroDivisionError as e:
    print(f"Error: {e}")
except (TypeError, ValueError):
    print("Type or value error")
else:
    print("No error")
finally:
    print("Always runs")

raise ValueError("invalid value")
```

## fa-file File I/O

```python
with open("file.txt", "r") as f:
    content = f.read()
    lines = f.readlines()

with open("out.txt", "w") as f:
    f.write("hello\n")

import json
data = json.load(open("data.json"))
json.dumps(data, indent=2)

import csv
with open("data.csv") as f:
    reader = csv.DictReader(f)
    for row in reader:
        print(row)
```

## fa-puzzle-piece Modules & Packages

```python
import os
from os.path import join, exists
from collections import Counter, defaultdict

dir(module)
help(module.func)
__name__ == "__main__"

import sys
sys.path.append("/custom/path")
```

## fa-wand-magic-sparkles Decorators

```python
import functools

def timer(func):
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        start = time.time()
        result = func(*args, **kwargs)
        print(f"{func.__name__}: {time.time() - start:.4f}s")
        return result
    return wrapper

@timer
def slow_func():
    time.sleep(1)

@staticmethod
def utility():
    pass

@classmethod
def create(cls, value):
    return cls(value)

@property
def name(self):
    return self._name
```

## fa-arrows-spin Generators

```python
def count_up(n):
    i = 0
    while i < n:
        yield i
        i += 1

gen = (x**2 for x in range(10))

def fibonacci():
    a, b = 0, 1
    while True:
        yield a
        a, b = b, a + b

from itertools import count, chain, islice
```

## fa-shield-halved Context Managers

```python
class ManagedResource:
    def __enter__(self):
        print("acquired")
        return self
    def __exit__(self, *exc):
        print("released")

with ManagedResource() as r:
    pass

from contextlib import contextmanager

@contextmanager
def temp_dir():
    d = tempfile.mkdtemp()
    yield d
    shutil.rmtree(d)
```

## fa-toolbox Common Standard Library

```python
from collections import Counter, defaultdict, deque
from datetime import datetime, timedelta
from pathlib import Path
from dataclasses import dataclass
from typing import Optional, List, Dict

p = Path("file.txt")
p.exists()
p.read_text()
p.stem

now = datetime.now()
now.strftime("%Y-%m-%d %H:%M")

Counter("abracadabra").most_common(3)
dd = defaultdict(list)
```
