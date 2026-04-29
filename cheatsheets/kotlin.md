---
title: Kotlin
icon: fa-code
primary: "#7F52FF"
lang: kotlin
---

## fa-box Variables & Types

```kotlin
val name: String = "Kotlin"
var count: Int = 0
val inferred = 42
val pi: Double = 3.14
val hex = 0xFF
val binary = 0b1010
val longNum = 123L

val anyVal: Any = "can be anything"
val unit: Unit = Unit
```

## fa-shield Null Safety

```kotlin
var nullable: String? = null
var nonNull: String = "required"

val len = nullable?.length
val lenOrZero = nullable?.length ?: 0
val required = nullable ?: "default"
val crash = nullable!!

val result = nullable?.let {
    it.uppercase()
}
```

## fa-font Strings

```kotlin
val raw = """raw string
    |with margin""".trimMargin()

val template = "$name has ${name.length} chars"
val multi = """line1
line2
line3"""

val sub = "hello".substring(1..3)
val upper = "hello".uppercase()
val parts = "a,b,c".split(",")
```

## fa-pen-to-square Functions

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

## fa-arrow-right-arrow-left Lambdas & Higher-Order

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

## fa-cube Classes & Objects

```kotlin
class Person(val name: String, var age: Int) {
    fun greet() = "Hi, I'm $name"
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

## fa-database Data Classes

```kotlin
data class User(val name: String, val age: Int)

val user = User("Alice", 30)
val copy = user.copy(age = 31)
val (name, age) = user

val map = mapOf("key" to "value")
println(user.toString())
println(user == copy)
```

## fa-lock Sealed Classes

```kotlin
sealed class Result<out T> {
    data class Success<T>(val value: T) : Result<T>()
    data class Error(val message: String) : Result<Nothing>()
    data object Loading : Result<Nothing>()
}

fun handle(result: Result<Int>) = when (result) {
    is Result.Success -> println(result.value)
    is Result.Error -> println(result.message)
    Result.Loading -> println("loading")
}
```

## fa-list Enums

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

## fa-puzzle-piece Extension Functions

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

## fa-layer-group Collections

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

## fa-arrows-to-dot Scope Functions

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

## fa-handshake Delegation

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

## fa-code-branch Generics

```kotlin
class Box<T>(val item: T)

fun <T> List<T>.second(): T = this[1]

fun <T : Comparable<T>> max(a: T, b: T): T = if (a > b) a else b

fun <T> clone(list: List<T>): List<T> {
    return list.toList()
}

val box: Box<out Number> = Box<Int>(42)
fun add(list: MutableList<in Number>) {
    list.add(3.14)
}
```
