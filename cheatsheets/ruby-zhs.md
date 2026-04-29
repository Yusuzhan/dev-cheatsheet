---
title: Ruby
icon: fa-gem
primary: "#CC342D"
lang: ruby
locale: zhs
---

## fa-box 变量与类型

```ruby
name = "Alice"
age = 30
score = 95.5
active = true
nothing = nil

CONST = 100

x, y = 1, 2
a, *rest = [1, 2, 3, 4]

42.class
"hello".class
nil.class
```

## fa-font 字符串与符号

```ruby
s = "hello world"
s.length
s.upcase
s.downcase
s.strip
s.include?("world")
s.start_with?("hello")
s.gsub(/o/, "0")
s.split(", ")
"hello " + "world"
"#{name} is #{age}"

sym = :status
sym.to_s
"status".to_sym
sym == :status
```

## fa-list 数组

```ruby
arr = [1, 2, 3, 4, 5]
arr << 6
arr.push(7)
arr.pop
arr.shift
arr.unshift(0)
arr[1..3]
arr.first
arr.last
arr.size
arr.empty?
arr.include?(3)
arr.sort
arr.reverse
arr.uniq
arr.flatten
arr.join(", ")
[1, 2, 3] + [4, 5]
```

## fa-table 哈希

```ruby
h = { name: "Alice", age: 30 }
h[:name]
h[:email] = "a@b.com"
h.key?(:name)
h.delete(:age)
h.keys
h.values
h.each { |k, v| puts "#{k}: #{v}" }
h.merge(other: 1)
h.transform_values { |v| v.to_s }
h.fetch(:missing, "default")
```

## fa-code-branch 控制流

```ruby
if x > 0
  puts "positive"
elsif x == 0
  puts "zero"
else
  puts "negative"
end

result = x > 0 ? "pos" : "neg"

unless done
  process
end

case status
when "running" then puts "ok"
when "stopped" then puts "down"
else puts "unknown"
end
```

## fa-pen-to-square 方法与块

```ruby
def greet(name)
  "Hello, #{name}"
end

def add(a, b) = a + b

def wrap(text, before: "<", after: ">")
  "#{before}#{text}#{after}"
end

def sum(*args)
  args.sum
end

def with_block
  yield if block_given?
end

with_block { puts "hi" }
```

## fa-arrows-spin 迭代器

```ruby
[1, 2, 3].each { |n| puts n }
[1, 2, 3].map { |n| n * 2 }
[1, 2, 3].select { |n| n > 1 }
[1, 2, 3].reject { |n| n < 2 }
[1, 2, 3].find { |n| n == 2 }
[1, 2, 3].reduce(0) { |sum, n| sum + n }
[1, 2, 3].any? { |n| n > 2 }
[1, 2, 3].all? { |n| n > 0 }
[1, 2, 3].none? { |n| n < 0 }
[1, 2, 3].count { |n| n > 1 }
[1, 2, 3].sort_by { |n| -n }
[1, 2, 3].flat_map { |n| [n, n * 2] }
(1..10).each { |i| puts i }
3.times { puts "hi" }
```

## fa-cubes 类与模块

```ruby
class Animal
  attr_accessor :name
  attr_reader :age

  def initialize(name, age)
    @name = name
    @age = age
  end

  def to_s
    "#{@name} (#{@age})"
  end
end

class Dog < Animal
  def speak
    "Woof! I'm #{@name}"
  end
end

module Printable
  def print_info
    puts to_s
  end
end

class Cat < Animal
  include Printable
end
```

## fa-triangle-exclamation 异常处理

```ruby
begin
  risky_operation
rescue StandardError => e
  puts "Error: #{e.message}"
  puts e.backtrace.first(5)
rescue ArgumentError => e
  retry
ensure
  cleanup
end

raise "something went wrong"
raise ArgumentError, "invalid input"
```

## fa-file 文件 I/O

```ruby
File.write("out.txt", "hello")
content = File.read("in.txt")
lines = File.readlines("in.txt")

File.open("out.txt", "w") do |f|
  f.puts "line 1"
  f.puts "line 2"
end

File.exist?("file.txt")
File.size("file.txt")
File.delete("file.txt")
Dir.glob("*.rb")
Dir.mkdir("new_dir")
```

## fa-filter 正则表达式

```ruby
/text/.match?("some text")
"hello world" =~ /world/
"2024-01-15".match(/(\d{4})-(\d{2})-(\d{2})/)
"hello".sub(/l/, "L")
"hello".gsub(/l/, "L")
"a1b2c3".scan(/\d/)
"hello world".split(/\s+/)
```

## fa-layer-group Enumerable

```ruby
[1, 2, 3, 4, 5].group_by { |n| n.even? }
[1, 2, 3, 4, 5].partition { |n| n.even? }
[1, 2, 3].zip([4, 5, 6])
[1, 2, 3].each_with_index { |n, i| puts "#{i}: #{n}" }
(1..10).first(3)
(1..10).take_while { |n| n < 5 }
(1..10).drop_while { |n| n < 5 }
[1, 2, 3].min
[1, 2, 3].max
[1, 2, 3].minmax
[1, 2, 3].tally
[1, 2, 3].one? { |n| n > 2 }
```

## fa-boxes-stacked Gems 与 Bundler

```sh
gem install rails
gem list
gem update
gem uninstall rails

bundle init
bundle install
bundle add rspec
bundle exec rspec
bundle update
bundle outdated
```

## fa-lightbulb 常用单行代码

```ruby
10.times.map { rand(100) }
(1..100).inject(:+)
File.readlines("f.txt").map(&:strip)
array.uniq.sort
hash.invert
string.chars.tally
array.each_slice(3).to_a
array.max_by(&:length)
string.delete("aeiou")
array.shuffle.sample(3)
```
