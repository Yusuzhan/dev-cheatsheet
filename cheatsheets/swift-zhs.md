---
title: Swift
icon: fa-swift
primary: "#F05138"
lang: swift
locale: zhs
---

## fa-box 变量与类型

```swift
let x: Int = 42
var y: Double = 3.14
let s: String = "hello"
let b: Bool = true
let ch: Character = "A"

let inferred = 100
let hex = 0xFF
let binary = 0b1010

typealias Age = Int
let age: Age = 30
```

## fa-question 可选值

```swift
var name: String? = nil
name = "Alice"

if let n = name {
    print(n)
}

guard let value = optional else { return }
let val = optional ?? "default"
let unwrapped = try! mightThrow()
let result = try? mightThrow()
```

## fa-font 字符串

```swift
let s = "Hello, World!"
s.count
s.isEmpty
s.uppercased()
s.lowercased()
s.contains("World")
s.hasPrefix("Hello")
s.hasSuffix("!")

let parts = s.split(separator: " ")
let joined = ["a", "b"].joined(separator: ",")
let formatted = "value is \(42)"
let multi = """
  line one
  line two
  """
```

## fa-code-branch 控制流

```swift
if x > 0 {
    print("positive")
} else if x == 0 {
    print("zero")
} else {
    print("negative")
}

switch value {
case 0: print("zero")
case 1...10: print("small")
case let x where x > 100: print("big: \(x)")
default: break
}

for i in 0..<10 { print(i) }
for item in array { print(item) }
while condition { }
repeat { } while condition
```

## fa-pen-to-square 函数与闭包

```swift
func greet(name: String) -> String {
    return "Hello, \(name)"
}
greet(name: "Alice")

func add(_ a: Int, _ b: Int) -> Int { a + b }
add(3, 5)

func compute(values: [Int], transform: (Int) -> Int) -> [Int] {
    values.map(transform)
}

let double: (Int) -> Int = { $0 * 2 }
let add: (Int, Int) -> Int = { $0 + $1 }
let capturer = { [weak self] in /* ... */ }
```

## fa-layer-group 枚举与结构体

```swift
enum Direction {
    case north, south, east, west
}

enum Result<T> {
    case success(T)
    case failure(Error)
}

switch result {
case .success(let value): print(value)
case .failure(let err): print(err)
}

struct Point {
    var x: Double
    var y: Double
    mutating func moveBy(dx: Double, dy: Double) {
        x += dx; y += dy
    }
}
var p = Point(x: 1, y: 2)
```

## fa-cubes 类与继承

```swift
class Animal {
    var name: String
    init(name: String) { self.name = name }
    func speak() -> String { "..." }
    deinit { }
}

class Dog: Animal {
    override func speak() -> String { "Woof" }
}

let dog = Dog(name: "Rex")
dog.speak()

class View {
    weak var delegate: AnyObject?
    lazy var data: [Int] = { return [] }()
}
```

## fa-puzzle-piece 协议与扩展

```swift
protocol Describable {
    var description: String { get }
    func describe() -> String
}

extension Describable {
    func describe() -> String { description }
}

extension Int {
    var squared: Int { self * self }
    func clamp(_ min: Int, _ max: Int) -> Int {
        Swift.max(min, Swift.min(max, self))
    }
}

5.squared
3.clamp(0, 10)
```

## fa-layer-group 泛型

```swift
func first<T>(_ items: [T]) -> T? {
    items.isEmpty ? nil : items[0]
}

struct Stack<Element> {
    private var items: [Element] = []
    mutating func push(_ item: Element) { items.append(item) }
    mutating func pop() -> Element? { items.popLast() }
}

func item<T: Equatable & Comparable>(in array: [T], equalTo target: T) -> Int? {
    array.firstIndex(of: target)
}
```

## fa-triangle-exclamation 错误处理

```swift
enum AppError: Error {
    case notFound
    case unauthorized
    case custom(message: String)
}

func fetch(id: Int) throws -> String {
    guard id > 0 else { throw AppError.notFound }
    return "data"
}

do {
    let data = try fetch(id: 1)
} catch AppError.notFound {
    print("not found")
} catch {
    print(error)
}
```

## fa-list 集合

```swift
var arr = [1, 2, 3]
arr.append(4)
arr.insert(0, at: 0)
arr.removeLast()
arr.sort()
arr.reverse()
arr.filter { $0 > 2 }
arr.map { $0 * 2 }
arr.reduce(0, +)

var set: Set<Int> = [1, 2, 3]
set.insert(4)
set.contains(2)
set.intersection([2, 3, 4])

var dict: [String: Int] = ["a": 1, "b": 2]
dict["c"] = 3
dict.removeValue(forKey: "a")
```

## fa-arrows-spin 并发 (async/await)

```swift
func fetchData() async throws -> Data {
    let (data, _) = try await URLSession.shared.data(from: url)
    return data
}

Task {
    let result = try await fetchData()
}

func streamValues() -> AsyncStream<Int> {
    AsyncStream { continuation in
        for i in 0..<10 {
            continuation.yield(i)
        }
        continuation.finish()
    }
}

actor Counter {
    private var value = 0
    func increment() -> Int {
        value += 1
        return value
    }
}
```

## fa-brain 内存管理 (ARC)

```swift
class Parent {
    var child: Child?
    deinit { print("parent deallocated") }
}

class Child {
    weak var parent: Parent?
    deinit { print("child deallocated") }
}

let parent = Parent()
let child = Child()
parent.child = child
child.parent = parent

unowned let ref: SomeClass
lazy var prop = { [unowned self] in self.data }()
```

## fa-desktop SwiftUI 基础

```swift
struct ContentView: View {
    @State private var count = 0

    var body: some View {
        VStack {
            Text("Count: \(count)")
                .font(.title)
            Button("Increment") {
                count += 1
            }
            .buttonStyle(.borderedProminent)
        }
        .padding()
    }
}

struct UserView: View {
    let name: String
    var body: some View {
        Text(name)
    }
}
```
