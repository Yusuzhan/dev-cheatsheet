---
title: Kotlin
icon: fa-code
primary: "#7F52FF"
lang: kotlin
locale: zhs
---

## fa-box 变量与类型

```kotlin
val name: String = "Kotlin"
var count: Int = 0
val inferred = 42
val pi: Double = 3.14
val hex = 0xFF
val binary = 0b1010
val longNum = 123L

val anyVal: Any = "可以是任意类型"
val unit: Unit = Unit
```

## fa-shield 空安全

```kotlin
var nullable: String? = null
var nonNull: String = "不可为空"

val len = nullable?.length               // 安全调用
val lenOrZero = nullable?.length ?: 0    // Elvis 运算符
val required = nullable ?: "default"     // 默认值
val crash = nullable!!                   // 非空断言，为空抛异常

val result = nullable?.let {
    it.uppercase()
}
```

## fa-font 字符串

```kotlin
val raw = """原始字符串
    |带边距""".trimMargin()

val template = "$name 有 ${name.length} 个字符"
val multi = """第一行
第二行
第三行"""

val sub = "hello".substring(1..3)
val upper = "hello".uppercase()
val parts = "a,b,c".split(",")
```

## fa-pen-to-square 函数

```kotlin
fun add(a: Int, b: Int): Int = a + b

fun greet(name: String, greeting: String = "Hello"): String {
    return "$greeting, $name"
}

fun varargExample(vararg items: String): List<String> {
    return items.toList()
}

greet("World")
greet(name = "Kotlin", greeting = "Hi")
```

## fa-arrow-right-arrow-left Lambda 与高阶函数

```kotlin
val square: (Int) -> Int = { x -> x * x }
val sum = { a: Int, b: Int -> a + b }

fun operate(x: Int, fn: (Int) -> Int): Int = fn(x)
operate(5) { it * 2 }

val names = listOf("Alice", "Bob")
names.map { it.uppercase() }
names.filter { it.length > 3 }
names.forEach { println(it) }
```

## fa-cube 类与对象

```kotlin
class Person(val name: String, var age: Int) {
    fun greet() = "你好，我是 $name"
}

val person = Person("Alice", 30)

object Database {
    val url = "localhost:5432"
    fun connect() { }
}

class Singleton private constructor() {
    companion object {
        val instance = Singleton()
    }
}
```

## fa-database 数据类

```kotlin
data class User(val name: String, val age: Int)

val user = User("Alice", 30)
val copy = user.copy(age = 31)
val (name, age) = user                         // 解构声明

val map = mapOf("key" to "value")
println(user.toString())
println(user == copy)
```

## fa-lock 密封类

```kotlin
sealed class Result<out T> {
    data class Success<T>(val value: T) : Result<T>()
    data class Error(val message: String) : Result<Nothing>()
    data object Loading : Result<Nothing>()
}

fun handle(result: Result<Int>) = when (result) {
    is Result.Success -> println(result.value)
    is Result.Error -> println(result.message)
    Result.Loading -> println("加载中")
}
```

## fa-list 枚举

```kotlin
enum class Direction {
    NORTH, SOUTH, EAST, WEST
}

enum class Status(val code: Int) {
    OK(200), NOT_FOUND(404), ERROR(500);

    fun isOk() = this == OK
}

val dir = Direction.valueOf("NORTH")
val entries = Direction.entries
```

## fa-puzzle-piece 扩展函数

```kotlin
fun String.addExclamation(): String = "$this!"

fun Int.isEven(): Boolean = this % 2 == 0

fun <T> MutableList<T>.swap(i: Int, j: Int) {
    val temp = this[i]
    this[i] = this[j]
    this[j] = temp
}

"hello".addExclamation()
4.isEven()
```

## fa-layer-group 集合

```kotlin
val list = listOf(1, 2, 3, 4, 5)
val mutableList = mutableListOf(1, 2, 3)
val set = setOf("a", "b", "c")
val map = mapOf(1 to "one", 2 to "two")

list.filter { it > 2 }
list.map { it * 2 }
list.sortedDescending()
list.groupBy { it % 2 }
list.associate { it to it * it }
list.fold(0) { acc, v -> acc + v }
list.chunked(2)
```

## fa-arrows-to-dot 作用域函数

```kotlin
val str: String? = "hello"
val length = str?.let {
    println(it)
    it.length
}

val config = StringBuilder().apply {
    append("host=")
    append("localhost")
}

val result = with(StringBuilder()) {
    append("a")
    append("b")
    toString()
}

val sq = "hello".also { println(it) }
```

## fa-handshake 委托

```kotlin
interface Printer {
    fun print(msg: String)
}

class ConsolePrinter : Printer {
    override fun print(msg: String) = println(msg)
}

class LogPrinter(private val printer: Printer) : Printer by printer

class Observable(var value: Int) {
    var observed by ::value
}

import kotlin.properties.Delegates
var checked: Boolean by Delegates.observable(false) { _, old, new ->
    println("$old -> $new")
}
```

## fa-code-branch 泛型

```kotlin
class Box<T>(val item: T)

fun <T> List<T>.second(): T = this[1]

fun <T : Comparable<T>> max(a: T, b: T): T = if (a > b) a else b

fun <T> clone(list: List<T>): List<T> {
    return list.toList()
}

val box: Box<out Number> = Box<Int>(42)      // 协变
fun add(list: MutableList<in Number>) {      // 逆变
    list.add(3.14)
}
```
