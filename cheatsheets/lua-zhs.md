---
title: Lua
icon: fa-moon
primary: "#000080"
lang: lua
locale: zhs
---

## fa-box 变量与类型

```lua
local name = "Lua"
local version = 5.4
local active = true
local nothing = nil

print(type(name))
print(type(version))
print(type(active))
print(type(nothing))
```

## fa-font 字符串

```lua
local s = "hello world"
print(#s)
print(string.upper(s))
print(string.sub(s, 1, 5))
print(string.format("Hi %s, v%.1f", "Lua", 5.4))
print(string.rep("ab", 3))
print(string.find(s, "world"))
```

## fa-calculator 运算符

```lua
local a, b = 10, 3
print(a + b)
print(a - b)
print(a * b)
print(a / b)
print(a % b)
print(a ^ b)
print(a // b)

print(a == b)
print(a ~= b)
print(a > b)

local x = true
print(x and false)
print(x or true)
print(not x)
```

## fa-code-branch 控制流

```lua
local x = 10
if x > 20 then
  print("big")
elseif x > 5 then
  print("medium")
else
  print("small")
end

for i = 1, 10, 2 do
  print(i)
end

for i = 1, 5 do
  print(i)
end

local arr = { "a", "b", "c" }
for i, v in ipairs(arr) do
  print(i, v)
end

local t = { name = "Lua", ver = 5.4 }
for k, v in pairs(t) do
  print(k, v)
end

local n = 0
while n < 5 do
  n = n + 1
end

repeat
  n = n + 1
until n > 10
```

## fa-cube 函数

```lua
local function greet(name)
  return "Hello, " .. name
end

local function add(a, b)
  return a + b
end

local function swap(a, b)
  return b, a
end

local x, y = swap(1, 2)

local function variadic(...)
  local args = { ... }
  return #args
end

print(variadic(1, 2, 3))
```

## fa-table 表

```lua
local fruits = { "apple", "banana", "cherry" }
table.insert(fruits, "date")
table.remove(fruits, 1)
table.sort(fruits)
print(#fruits)

local person = {
  name = "Alice",
  age = 30,
}
person.email = "alice@example.com"

for k, v in pairs(person) do
  print(k, v)
end

local copy = {}
for k, v in pairs(person) do
  copy[k] = v
end
```

## fa-layer-group 元表

```lua
local mt = {
  __add = function(a, b)
    return setmetatable({ x = a.x + b.x, y = a.y + b.y }, getmetatable(a))
  end,
  __tostring = function(v)
    return "(" .. v.x .. "," .. v.y .. ")"
  end,
}

local function Vec(x, y)
  return setmetatable({ x = x, y = y }, mt)
end

local v1 = Vec(1, 2)
local v2 = Vec(3, 4)
print(v1 + v2)
```

## fa-puzzle-piece 模块

```lua
-- mymodule.lua
local M = {}

function M.hello(name)
  return "Hello, " .. name
end

return M

-- main.lua
local mod = require("mymodule")
print(mod.hello("world"))
```

## fa-rotate 协程

```lua
local co = coroutine.create(function(a, b)
  print("received:", a, b)
  local c = coroutine.yield(a + b)
  print("resumed with:", c)
  return "done"
end)

local ok, val = coroutine.resume(co, 1, 2)
print("yielded:", val)

local ok2, val2 = coroutine.resume(co, 99)
print("result:", val2)
```

## fa-file 文件 I/O

```lua
local f = io.open("test.txt", "w")
f:write("hello\n")
f:close()

local f = io.open("test.txt", "r")
local content = f:read("*a")
f:close()
print(content)

for line in io.lines("test.txt") do
  print(line)
end
```

## fa-magnifying-glass 字符串模式匹配

```lua
local s = "Hello, World! 123"
print(string.match(s, "%d+"))
print(string.match(s, "(%a+)"))

for w in string.gmatch(s, "%a+") do
  print(w)
end

print(string.gsub(s, "%d", "X"))
print(string.gsub("2024-01-15", "(%d+)-(%d+)-(%d+)", "%3/%2/%1"))
```

## fa-shield 错误处理

```lua
local ok, err = pcall(function()
  error("something went wrong")
end)
if not ok then
  print("caught:", err)
end

local function safe_div(a, b)
  if b == 0 then return nil, "division by zero" end
  return a / b
end

local result, err = safe_div(10, 0)
if not result then
  print(err)
end
```

## fa-sitemap 面向对象（元表实现）

```lua
local Dog = {}
Dog.__index = Dog

function Dog.new(name, breed)
  return setmetatable({ name = name, breed = breed }, Dog)
end

function Dog:bark()
  print(self.name .. " says: Woof!")
end

function Dog:info()
  print(self.name .. " is a " .. self.breed)
end

local d = Dog.new("Rex", "German Shepherd")
d:bark()
d:info()
```

## fa-wrench 常用模式

```lua
local function deepcopy(orig)
  local copy = {}
  for k, v in pairs(orig) do
    if type(v) == "table" then
      copy[k] = deepcopy(v)
    else
      copy[k] = v
    end
  end
  return copy
end

local function keys(t)
  local ks = {}
  for k in pairs(t) do
    ks[#ks + 1] = k
  end
  return ks
end

local function values(t)
  local vs = {}
  for _, v in pairs(t) do
    vs[#vs + 1] = v
  end
  return vs
end

local function merge(a, b)
  local r = {}
  for k, v in pairs(a) do r[k] = v end
  for k, v in pairs(b) do r[k] = v end
  return r
end
```
