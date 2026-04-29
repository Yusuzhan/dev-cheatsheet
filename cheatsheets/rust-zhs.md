---
title: Rust
icon: fa-rust
primary: "#CE422B"
lang: rust
locale: zhs
---

## fa-box 变量与类型

```rust
let x: i32 = 42;
let y = 3.14_f64;
let b: bool = true;
let ch: char = 'A';
let s: &str = "hello";
let tuple: (i32, f64) = (1, 3.14);
let (a, b) = tuple;

let mut count = 0;
count += 1;

const MAX: i32 = 100;
static GLOBAL: &str = "shared";
```

## fa-pen-to-square 函数

```rust
fn add(a: i32, b: i32) -> i32 {
    a + b
}

fn greet(name: &str) -> String {
    format!("Hello, {}!", name)
}

fn safe_divide(a: f64, b: f64) -> Option<f64> {
    if b == 0.0 { None } else { Some(a / b) }
}

fn first_or<'a>(slice: &'a [i32], default: &'a i32) -> &'a i32 {
    slice.first().unwrap_or(default)
}

let double = |x: i32| x * 2;
let capture = |y| x + y;
```

## fa-code-branch 控制流

```rust
if x > 0 {
    println!("positive");
} else if x == 0 {
    println!("zero");
} else {
    println!("negative");
}

let abs = if x < 0 { -x } else { x };

for i in 0..10 {
    println!("{}", i);
}

for item in &collection {
    println!("{:?}", item);
}

while let Some(val) = iter.next() {
    println!("{}", val);
}

loop {
    if done { break; }
}

'outer: for x in 0..10 {
    for y in 0..10 {
        if x + y == 15 { break 'outer; }
    }
}
```

## fa-cubes 结构体与枚举

```rust
struct User {
    name: String,
    age: u32,
}

let user = User { name: String::from("Alice"), age: 30 };
let User { name, .. } = user;

struct Point(i32, i32);
let p = Point(1, 2);

#[derive(Debug)]
enum Shape {
    Circle(f64),
    Rectangle { w: f64, h: f64 },
    Triangle,
}

let area = match &shape {
    Shape::Circle(r) => std::f64::consts::PI * r * r,
    Shape::Rectangle { w, h } => w * h,
    Shape::Triangle => 0.0,
};

enum Option<T> { Some(T), None }
enum Result<T, E> { Ok(T), Err(E) }
```

## fa-layer-group 模式匹配

```rust
match value {
    0 => println!("zero"),
    1..=10 => println!("small"),
    _ => println!("other"),
}

if let Some(x) = option {
    println!("{}", x);
}

while let Some(item) = iter.next() {
    println!("{:?}", item);
}

let (x, y, _) = tuple;

let [first, .., last] = array;

match slice {
    [first, second] => println!("{} {}", first, second),
    [first, ..] => println!("{}", first),
    [] => println!("empty"),
}
```

## fa-key 所有权与借用

```rust
let s1 = String::from("hello");
let s2 = s1;

let s3 = s2.clone();

fn takes_ownership(s: String) { }
fn borrows(s: &String) { }
fn mutates(s: &mut String) { s.push_str("!"); }

let mut msg = String::from("hello");
borrows(&msg);
mutates(&mut msg);

let data = vec![1, 2, 3];
let r1 = &data;
let r2 = &data;
// let r3 = &mut data;

fn first_word(s: &str) -> &str {
    s
}
```

## fa-font 字符串

```rust
let s: String = String::from("hello world");
let slice: &str = &s[0..5];

s.push_str("!");
s.push('!');
s.replace("hello", "hi");
s.contains("world");

let parts: Vec<&str> = "a,b,c".split(',').collect();
s.len();
s.is_empty();
s.to_uppercase();
s.trim();

let s = format!("{} is {}", "answer", 42);
```

## fa-list 向量

```rust
let mut v: Vec<i32> = Vec::new();
let v = vec![1, 2, 3, 4, 5];

v.push(6);
v.pop();
v[0];
v.get(2);

for val in &v {
    println!("{}", val);
}

for val in &mut v {
    *val += 1;
}

v.iter().map(|x| x * 2).collect::<Vec<_>>();
v.iter().filter(|&&x| x > 2).collect::<Vec<_>>();
```

## fa-triangle-exclamation 错误处理

```rust
use std::fs;

let content = fs::read_to_string("file.txt")
    .expect("failed to read");

match fs::read_to_string("file.txt") {
    Ok(s) => println!("{}", s),
    Err(e) => eprintln!("Error: {}", e),
}

let val = option.ok_or("was None")?;
let val = result.map_err(|e| format!("wrap: {}", e))?;

let result: Result<i32, &str> = Ok(42);
let doubled = result.map(|x| x * 2);
let flattened: Result<i32, &str> = Ok(Ok(42)).unwrap();

thiserror, anyhow
```

## fa-puzzle-piece 特征

```rust
trait Summary {
    fn summarize(&self) -> String;

    fn default_summary(&self) -> String {
        String::from("...")
    }
}

impl Summary for User {
    fn summarize(&self) -> String {
        format!("{} ({})", self.name, self.age)
    }
}

fn print_summary(item: &impl Summary) {
    println!("{}", item.summarize());
}

fn print_summary<T: Summary>(item: &T) {
    println!("{}", item.summarize());
}

fn compare<T: PartialOrd>(a: T, b: T) -> std::cmp::Ordering {
    a.partial_cmp(&b).unwrap()
}

impl Display for Point {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "({}, {})", self.0, self.1)
    }
}
```

## fa-layer-group 泛型

```rust
fn largest<T: PartialOrd>(list: &[T]) -> &T {
    let mut max = &list[0];
    for item in &list[1..] {
        if item > max { max = item; }
    }
    max
}

struct Container<T> {
    value: T,
}

impl<T> Container<T> {
    fn new(value: T) -> Self {
        Container { value }
    }
}

enum Result<T, E> {
    Ok(T),
    Err(E),
}

fn identity<T>(x: T) -> T { x }
```

## fa-arrows-spin 迭代器

```rust
let v = vec![1, 2, 3, 4, 5];

v.iter().map(|x| x * 2);
v.iter().filter(|x| *x % 2 == 0);
v.iter().fold(0, |acc, x| acc + x);
v.iter().take(3);
v.iter().skip(2);
v.iter().enumerate();
v.iter().zip(v.iter().skip(1));
v.iter().find(|x| **x == 3);
v.iter().any(|x| *x > 3);
v.iter().all(|x| *x > 0);
v.iter().min();
v.iter().max();
v.iter().count();

(1..=10).for_each(|x| println!("{}", x));
```

## fa-boxes-stacked 模块与 Crate

```rust
mod network {
    pub fn connect() {}
    pub mod server {
        pub fn listen() {}
    }
}

use network::server::listen;
use std::collections::HashMap;

mod my_module;

pub fn api() {}
pub(crate) fn internal() {}

// lib.rs or main.rs
mod utils;
```

## fa-brain 智能指针

```rust
use std::rc::Rc;
use std::sync::Arc;
use std::cell::RefCell;
use std::sync::Mutex;

let a = Rc::new(vec![1, 2, 3]);
let b = Rc::clone(&a);

let a = Arc::new(vec![1, 2, 3]);
let b = Arc::clone(&a);

let val = RefCell::new(5);
*val.borrow_mut() += 1;

let data = Mutex::new(0);
let mut locked = data.lock().unwrap();
*locked += 1;

use std::borrow::Cow;
let cow: Cow<str> = Cow::Borrowed("hello");
```

## fa-bolt 并发

```rust
use std::thread;
use std::sync::{mpsc, Arc, Mutex};

let handle = thread::spawn(|| {
    42
});
let result = handle.join().unwrap();

let data = Arc::new(Mutex::new(vec![]));
let data_clone = Arc::clone(&data);
thread::spawn(move || {
    data_clone.lock().unwrap().push(1);
});

let (tx, rx) = mpsc::channel();
tx.send(42).unwrap();
let val = rx.recv().unwrap();

use std::sync::mpsc;
let (tx, rx) = mpsc::channel();
for i in 0..5 {
    let tx = tx.clone();
    thread::spawn(move || { tx.send(i).unwrap(); });
}
drop(tx);
for val in rx { println!("{}", val); }
```

## fa-truck Cargo 命令

```rust
cargo new my_project
cargo build
cargo build --release
cargo run
cargo test
cargo test test_name
cargo check
cargo doc --open
cargo update
cargo add serde
cargo publish
cargo fmt
cargo clippy
```
