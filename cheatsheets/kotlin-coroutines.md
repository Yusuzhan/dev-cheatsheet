---
title: Kotlin Coroutines & Flow
icon: fa-shuffle
primary: "#7F52FF"
lang: kotlin
---

## fa-rocket Launch Coroutines

```kotlin
fun main() = runBlocking {
    launch {
        delay(1000)
        println("World!")
    }
    println("Hello")
}

GlobalScope.launch {
    delay(500)
    println("Global scope")
}

suspend fun doWork() {
    coroutineScope {
        launch { delay(100); println("A") }
        launch { delay(200); println("B") }
    }
}
```

## fa-circle-dot CoroutineScope

```kotlin
class MyService {
    private val scope = CoroutineScope(Dispatchers.Default + SupervisorJob())

    fun start() {
        scope.launch { doWork() }
    }

    fun destroy() {
        scope.cancel()
    }
}

suspend fun parallelTasks() = coroutineScope {
    val deferred1 = async { fetchUser() }
    val deferred2 = async { fetchOrders() }
    awaitAll(deferred1, deferred2)
}
```

## fa-gears Dispatchers

```kotlin
withContext(Dispatchers.IO) {
    val data = networkCall()
}

withContext(Dispatchers.Default) {
    heavyComputation()
}

withContext(Dispatchers.Main) {
    updateUI()
}

launch(Dispatchers.IO) {
    val result = readFile()
    withContext(Dispatchers.Main) {
        textView.text = result
    }
}
```

## fa-pause suspend Functions

```kotlin
suspend fun fetchUser(id: Int): User {
    delay(1000)
    return User(id, "Alice")
}

suspend fun fetchData(): List<Item> = coroutineScope {
    val items = async { api.getItems() }
    items.await()
}

suspend fun retryFetch(maxRetries: Int = 3): Data {
    repeat(maxRetries) { attempt ->
        try {
            return api.fetch()
        } catch (e: Exception) {
            if (attempt == maxRetries - 1) throw e
            delay(1000L * (attempt + 1))
        }
    }
    error("unreachable")
}
```

## fa-code-branch async & await

```kotlin
suspend fun fetchAll() = coroutineScope {
    val user = async { api.getUser() }
    val posts = async { api.getPosts() }
    val comments = async { api.getComments() }
    Result(user.await(), posts.await(), comments.await())
}

suspend fun race(): String = coroutineScope {
    select {
        async { fastService() }.onAwait { "fast: $it" }
        async { slowService() }.onAwait { "slow: $it" }
    }
}
```

## fa-triangle-exclamation Exception Handling

```kotlin
try {
    coroutineScope {
        launch { throw RuntimeException("fail") }
    }
} catch (e: Exception) {
    println("Caught: ${e.message}")
}

runBlocking {
    val handler = CoroutineExceptionHandler { _, throwable ->
        println("Error: ${throwable.message}")
    }
    launch(handler) {
        throw RuntimeException("fail")
    }
}

val result: Result<Data> = runCatching { api.fetch() }
```

## fa-ban Cancellation

```kotlin
val job = launch {
    try {
        repeat(1000) {
            ensureActive()
            println("Working $it")
            delay(100)
        }
    } finally {
        println("Cleanup")
    }
}

delay(500)
job.cancelAndJoin()

suspend fun cancellableWork() {
    for (item in items) {
        yield()
        process(item)
    }
}
```

## fa-water Flow Basics

```kotlin
fun numbers(): Flow<Int> = flow {
    for (i in 1..10) {
        emit(i)
        delay(100)
    }
}

flowOf(1, 2, 3)
listOf(1, 2, 3).asFlow()

suspend fun collectExample() {
    numbers().collect { value ->
        println(value)
    }
}
```

## fa-filter Flow Operators

```kotlin
flowOf(1, 2, 3, 4, 5)
    .map { it * 2 }
    .filter { it > 4 }
    .take(3)
    .distinctUntilChanged()
    .collect { println(it) }

flow {
    emit(1)
    emit(2)
}
    .onStart { emit(0) }
    .onCompletion { println("Done") }
    .onEach { println("Processing $it") }
    .catch { e -> emit(-1) }
    .collect()
```

## fa-tower-broadcast StateFlow & SharedFlow

```kotlin
class MyViewModel : ViewModel() {
    private val _state = MutableStateFlow("initial")
    val state: StateFlow<String> = _state.asStateFlow()

    fun update(value: String) { _state.value = value }
}

val shared = MutableSharedFlow<String>(
    extraBufferCapacity = 10,
    onBufferOverflow = BufferOverflow.DROP_OLDEST
)

launch {
    shared.collect { println("Received: $it") }
}

shared.tryEmit("event")
```

## fa-right-left Channel

```kotlin
val channel = Channel<Int>()

launch {
    repeat(5) { channel.send(it) }
    channel.close()
}

launch {
    for (value in channel) {
        println(value)
    }
}

val produceFlow = produce {
    repeat(5) { send(it) }
}

Channel.Factory.UNLIMITED
Channel.Factory.CONFLATED
```

## fa-arrows-rotate withContext

```kotlin
suspend fun loadAndShow() {
    val data = withContext(Dispatchers.IO) {
        database.query()
    }
    withContext(Dispatchers.Main) {
        show(data)
    }
}

suspend fun <T> safeApiCall(block: suspend () -> T): Result<T> {
    return try {
        Result.success(withContext(Dispatchers.IO) { block() })
    } catch (e: Exception) {
        Result.failure(e)
    }
}
```

## fa-shield SupervisorJob

```kotlin
val scope = CoroutineScope(SupervisorJob() + Dispatchers.Default)

scope.launch {
    launch { throw Error("Child 1 fails") }
    launch { delay(100); println("Child 2 still runs") }
}

val supervisor = SupervisorJob()
with(CoroutineScope(supervisor)) {
    val child1 = launch { throw Error("fail") }
    val child2 = launch { delay(100); println("OK") }
}
```

## fa-code-compare Select

```kotlin
suspend fun selectFirst(): String {
    val channel1 = Channel<String>()
    val channel2 = Channel<String>()

    return select<String> {
        channel1.onReceive { it }
        channel2.onReceive { it }
    }
}

suspend fun raceDeferred(): Int = coroutineScope {
    select {
        async { serviceA() }.onAwait { it }
        async { serviceB() }.onAwait { it }
    }
}
```

## fa-vial Flow Testing

```kotlin
class FlowTest {
    @Test
    fun testFlow() = runTest {
        val flow = flowOf(1, 2, 3)

        val results = flow.toList()
        assertEquals(listOf(1, 2, 3), results)
    }

    @Test
    fun testStateFlow() = runTest {
        val state = MutableStateFlow(0)
        launch { state.value = 42 }

        val first = state.first()
        assertEquals(0, first)
    }
}

@Test
fun testWithVirtualTime() = runTest {
    val flow = flow {
        emit(1)
        delay(1000)
        emit(2)
    }
    val values = mutableListOf<Int>()
    flow.toList(values)
    assertEquals(listOf(1, 2), values)
}
```
