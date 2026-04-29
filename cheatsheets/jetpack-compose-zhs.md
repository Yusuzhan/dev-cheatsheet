---
title: Jetpack Compose UI
icon: fa-palette
primary: "#7F52FF"
lang: kotlin
locale: zhs
---

## fa-play Composable 基础

```kotlin
@Composable
fun Greeting(name: String) {
    Text(text = "Hello, $name!")
}

@Preview(showBackground = true)
@Composable
fun GreetingPreview() {
    MyTheme {
        Greeting("Android")
    }
}
```

## fa-grip 布局（Row/Column/Box）

```kotlin
Column(
    verticalArrangement = Arrangement.spacedBy(8.dp),
    horizontalAlignment = Alignment.CenterHorizontally
) {
    Text("第一项")
    Text("第二项")
}

Row(
    horizontalArrangement = Arrangement.SpaceBetween,
    modifier = Modifier.fillMaxWidth()
) {
    Text("左")
    Text("右")
}

Box(modifier = Modifier.size(100.dp)) {
    Text("居中", modifier = Modifier.align(Alignment.Center))
    Text("底部", modifier = Modifier.align(Alignment.BottomCenter))
}
```

## fa-sliders 修饰符

```kotlin
Modifier
    .fillMaxWidth()
    .padding(16.dp)
    .background(Color.White, RoundedCornerShape(8.dp))
    .clickable { }
    .size(48.dp)
    .offset(x = 10.dp, y = 5.dp)
    .border(1.dp, Color.Gray, CircleShape)
    .clip(RoundedCornerShape(12.dp))
    .alpha(0.8f)
    .weight(1f)
```

## fa-heading 文本与样式

```kotlin
Text(
    text = "Hello",
    fontSize = 24.sp,
    fontWeight = FontWeight.Bold,
    color = Color.Red,
    textAlign = TextAlign.Center,
    maxLines = 2,
    overflow = TextOverflow.Ellipsis,
    style = MaterialTheme.typography.headlineMedium,
    modifier = Modifier.fillMaxWidth()
)

Text(buildAnnotatedString {
    append("普通 ")
    withStyle(SpanStyle(fontWeight = FontWeight.Bold, color = Color.Blue)) {
        append("加粗")
    }
})
```

## fa-keyboard 输入框

```kotlin
var text by remember { mutableStateOf("") }

OutlinedTextField(
    value = text,
    onValueChange = { text = it },
    label = { Text("邮箱") },
    placeholder = { Text("请输入邮箱") },
    singleLine = true,
    keyboardOptions = KeyboardOptions(keyboardType = KeyboardType.Email),
    modifier = Modifier.fillMaxWidth()
)

TextField(
    value = text,
    onValueChange = { text = it },
    colors = TextFieldDefaults.colors(
        focusedContainerColor = Color.Transparent
    )
)
```

## fa-hand-pointer 按钮与点击

```kotlin
Button(
    onClick = { },
    enabled = true,
    colors = ButtonDefaults.buttonColors(
        containerColor = Color.Blue,
        contentColor = Color.White
    ),
    shape = RoundedCornerShape(8.dp)
) {
    Icon(Icons.Default.Add, contentDescription = null)
    Spacer(modifier = Modifier.width(8.dp))
    Text("添加")
}

IconButton(onClick = { }) {
    Icon(Icons.Default.Favorite, contentDescription = "收藏")
}

Text("可点击", modifier = Modifier.clickable { })
```

## fa-list 列表（LazyColumn）

```kotlin
LazyColumn(
    verticalArrangement = Arrangement.spacedBy(4.dp),
    contentPadding = PaddingValues(16.dp)
) {
    items(100) { index ->
        Text("第 $index 项", modifier = Modifier.padding(8.dp))
    }

    items(listOf("A", "B", "C"), key = { it }) { item ->
        Text(item)
    }

    item {
        Text("底部")
    }

    stickyHeader {
        Text("标题", style = MaterialTheme.typography.titleMedium)
    }
}
```

## fa-image 图片

```kotlin
Image(
    painter = painterResource(R.drawable.photo),
    contentDescription = "头像",
    contentScale = ContentScale.Crop,
    modifier = Modifier
        .size(120.dp)
        .clip(CircleShape)
        .border(2.dp, Color.Gray, CircleShape)
)

AsyncImage(
    model = "https://example.com/photo.jpg",
    contentDescription = null,
    placeholder = painterResource(R.drawable.placeholder),
    error = painterResource(R.drawable.error),
    contentScale = ContentScale.Crop,
    modifier = Modifier.fillMaxWidth()
)
```

## fa-paintbrush 主题

```kotlin
val LightColors = lightColorScheme(
    primary = Color(0xFF6200EE),
    secondary = Color(0xFF03DAC6),
    surface = Color.White,
    onSurface = Color.Black
)

MaterialTheme(
    colorScheme = LightColors,
    typography = Typography,
    content = { /* 你的可组合函数 */ }
)

val color = MaterialTheme.colorScheme.primary
val typography = MaterialTheme.typography.bodyLarge
```

## fa-toggle-on 状态管理

```kotlin
var count by remember { mutableStateOf(0) }
var checked by remember { mutableStateOf(false) }
val list = remember { mutableStateListOf(1, 2, 3) }

Column {
    Text("计数: $count")
    Button(onClick = { count++ }) { Text("递增") }
}

val saved by rememberSaveable { mutableStateOf("") }

val state by remember {
    derivedStateOf { count > 10 }
}
```

## fa-bolt 副作用

```kotlin
LaunchedEffect(userId) {
    val data = repository.fetchUser(userId)
    result = data
}

DisposableEffect(lifecycle) {
    val observer = LifecycleEventObserver { _, event -> }
    lifecycle.addObserver(observer)
    onDispose { lifecycle.removeObserver(observer) }
}

var key by remember { mutableStateOf(0) }
SideEffect { key++ }

val snackbarHostState = remember { SnackbarHostState() }
LaunchedEffect(errorMessage) {
    errorMessage?.let {
        snackbarHostState.showSnackbar(it)
    }
}
```

## fa-route 导航

```kotlin
val navController = rememberNavController()

NavHost(navController, startDestination = "home") {
    composable("home") {
        HomeScreen(onClick = { navController.navigate("detail/$id") })
    }
    composable(
        "detail/{id}",
        arguments = listOf(navArgument("id") { type = NavType.IntType })
    ) { backStack ->
        val id = backStack.arguments?.getInt("id")
        DetailScreen(id)
    }
}

navController.navigate("home") {
    popUpTo("home") { inclusive = true }
}

navController.popBackStack()
```

## fa-table-columns 自定义布局

```kotlin
@Composable
fun StaggeredGrid(
    modifier: Modifier = Modifier,
    rows: Int = 3,
    content: @Composable () -> Unit
) {
    Layout(content, modifier) { measurables, constraints ->
        val rowHeights = IntArray(rows) { 0 }
        val placeables = measurables.mapIndexed { index, measurable ->
            val placeable = measurable.measure(constraints)
            val row = index % rows
            rowHeights[row] += placeable.height
            placeable
        }
        val height = rowHeights.maxOrNull() ?: 0
        layout(constraints.maxWidth, height) {
            val y = IntArray(rows) { 0 }
            placeables.forEachIndexed { index, placeable ->
                val row = index % rows
                placeable.placeRelative(0, y[row])
                y[row] += placeable.height
            }
        }
    }
}
```

## fa-wand-magic-sparkles 动画

```kotlin
var expanded by remember { mutableStateOf(false) }
val size by animateDpAsState(
    targetValue = if (expanded) 200.dp else 80.dp,
    animationSpec = spring(dampingRatio = Spring.DampingRatioMediumBouncy)
)

val infiniteTransition = rememberInfiniteTransition()
val alpha by infiniteTransition.animateFloat(
    initialValue = 0f,
    targetValue = 1f,
    animationSpec = infiniteRepeatable(
        animation = tween(1000),
        repeatMode = RepeatMode.Reverse
    )
)

AnimatedVisibility(visible = expanded) {
    Text("现在可见了")
}

val color by animateColorAsState(
    if (checked) Color.Green else Color.Gray
)
```

## fa-shapes Material 3

```kotlin
Scaffold(
    topBar = {
        TopAppBar(
            title = { Text("我的应用") },
            colors = TopAppBarDefaults.topAppBarColors(
                containerColor = MaterialTheme.colorScheme.primaryContainer
            )
        )
    },
    floatingActionButton = {
        FloatingActionButton(onClick = { }) {
            Icon(Icons.Default.Add, "添加")
        }
    },
    snackbarHost = { SnackbarHost(snackbarHostState) },
    bottomBar = {
        NavigationBar {
            NavigationBarItem(
                selected = true,
                onClick = { },
                icon = { Icon(Icons.Default.Home, "首页") },
                label = { Text("首页") }
            )
        }
    }
) { padding ->
    Content(modifier = Modifier.padding(padding))
}
```
