---
title: Jetpack Compose UI
icon: fa-palette
primary: "#7F52FF"
lang: kotlin
---

## fa-play Composable Basics

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

## fa-grip Layout (Row/Column/Box)

```kotlin
Column(
    verticalArrangement = Arrangement.spacedBy(8.dp),
    horizontalAlignment = Alignment.CenterHorizontally
) {
    Text("First")
    Text("Second")
}

Row(
    horizontalArrangement = Arrangement.SpaceBetween,
    modifier = Modifier.fillMaxWidth()
) {
    Text("Left")
    Text("Right")
}

Box(modifier = Modifier.size(100.dp)) {
    Text("Center", modifier = Modifier.align(Alignment.Center))
    Text("Bottom", modifier = Modifier.align(Alignment.BottomCenter))
}
```

## fa-sliders Modifiers

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

## fa-heading Text & Styling

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
    append("Normal ")
    withStyle(SpanStyle(fontWeight = FontWeight.Bold, color = Color.Blue)) {
        append("Bold")
    }
})
```

## fa-keyboard TextField

```kotlin
var text by remember { mutableStateOf("") }

OutlinedTextField(
    value = text,
    onValueChange = { text = it },
    label = { Text("Email") },
    placeholder = { Text("Enter email") },
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

## fa-hand-pointer Button & Click

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
    Text("Add Item")
}

IconButton(onClick = { }) {
    Icon(Icons.Default.Favorite, contentDescription = "Like")
}

Text("Clickable", modifier = Modifier.clickable { })
```

## fa-list List (LazyColumn)

```kotlin
LazyColumn(
    verticalArrangement = Arrangement.spacedBy(4.dp),
    contentPadding = PaddingValues(16.dp)
) {
    items(100) { index ->
        Text("Item $index", modifier = Modifier.padding(8.dp))
    }

    items(listOf("A", "B", "C"), key = { it }) { item ->
        Text(item)
    }

    item {
        Text("Footer")
    }

    stickyHeader {
        Text("Header", style = MaterialTheme.typography.titleMedium)
    }
}
```

## fa-image Image

```kotlin
Image(
    painter = painterResource(R.drawable.photo),
    contentDescription = "Profile photo",
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

## fa-paintbrush Theming

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
    content = { /* your composable */ }
)

val color = MaterialTheme.colorScheme.primary
val typography = MaterialTheme.typography.bodyLarge
```

## fa-toggle-on State Management

```kotlin
var count by remember { mutableStateOf(0) }
var checked by remember { mutableStateOf(false) }
val list = remember { mutableStateListOf(1, 2, 3) }

Column {
    Text("Count: $count")
    Button(onClick = { count++ }) { Text("Increment") }
}

val saved by rememberSaveable { mutableStateOf("") }

val state by remember {
    derivedStateOf { count > 10 }
}
```

## fa-bolt Side Effects

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

## fa-route Navigation

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

## fa-table-columns Custom Layout

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

## fa-wand-magic-sparkles Animations

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
    Text("Now visible")
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
            title = { Text("My App") },
            colors = TopAppBarDefaults.topAppBarColors(
                containerColor = MaterialTheme.colorScheme.primaryContainer
            )
        )
    },
    floatingActionButton = {
        FloatingActionButton(onClick = { }) {
            Icon(Icons.Default.Add, "Add")
        }
    },
    snackbarHost = { SnackbarHost(snackbarHostState) },
    bottomBar = {
        NavigationBar {
            NavigationBarItem(
                selected = true,
                onClick = { },
                icon = { Icon(Icons.Default.Home, "Home") },
                label = { Text("Home") }
            )
        }
    }
) { padding ->
    Content(modifier = Modifier.padding(padding))
}
```
