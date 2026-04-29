---
title: Jetpack ViewModel
icon: fa-layer-group
primary: "#7F52FF"
lang: kotlin
---

## fa-cube Basic ViewModel

```kotlin
class MyViewModel : ViewModel() {
    private val _name = mutableStateOf("")
    val name: String get() = _name.value

    fun setName(value: String) {
        _name.value = value
    }
}

class MyActivity : AppCompatActivity() {
    private val viewModel: MyViewModel by viewModels()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        viewModel.setName("Hello")
    }
}
```

## fa-database ViewModel with State

```kotlin
data class UiState(
    val items: List<String> = emptyList(),
    val isLoading: Boolean = false,
    val error: String? = null
)

class ListViewModel : ViewModel() {
    private val _state = MutableStateFlow(UiState())
    val state: StateFlow<UiState> = _state.asStateFlow()

    fun loadItems() {
        _state.update { it.copy(isLoading = true) }
        viewModelScope.launch {
            try {
                val items = repository.getItems()
                _state.update { it.copy(items = items, isLoading = false) }
            } catch (e: Exception) {
                _state.update { it.copy(error = e.message, isLoading = false) }
            }
        }
    }
}
```

## fa-bolt ViewModel + LiveData

```kotlin
class UserViewModel(private val repository: UserRepository) : ViewModel() {
    private val _user = MutableLiveData<User>()
    val user: LiveData<User> = _user

    private val _loading = MutableLiveData(false)
    val loading: LiveData<Boolean> = _loading

    fun loadUser(id: Int) {
        _loading.value = true
        viewModelScope.launch {
            _user.value = repository.getUser(id)
            _loading.value = false
        }
    }
}

val user by viewModel.user.observeAsState()
```

## fa-wave-square ViewModel + StateFlow

```kotlin
class CounterViewModel : ViewModel() {
    private val _count = MutableStateFlow(0)
    val count: StateFlow<Int> = _count.asStateFlow()

    fun increment() { _state.update { it + 1 } }
    fun decrement() { _state.update { it - 1 } }
    fun reset() { _state.value = 0 }
}

class SearchViewModel : ViewModel() {
    private val _query = MutableStateFlow("")
    val query: StateFlow<String> = _query.asStateFlow()

    val results: StateFlow<List<Item>> = _query
        .debounce(300)
        .mapLatest { query -> repository.search(query) }
        .stateIn(viewModelScope, SharingStarted.WhileSubscribed(5000), emptyList())
}
```

## fa-floppy-disk SavedStateHandle

```kotlin
class EditViewModel(
    private val savedStateHandle: SavedStateHandle
) : ViewModel() {
    var title: String
        get() = savedStateHandle["title"] ?: ""
        set(value) { savedStateHandle["title"] = value }

    var draft: String
        get() = savedStateHandle.get<String>("draft") ?: ""
        set(value) { savedStateHandle["draft"] = value }

    companion object {
        const val KEY_USER_ID = "userId"
    }

    val userId: String = savedStateHandle[KEY_USER_ID] ?: ""
}
```

## fa-industry ViewModelFactory

```kotlin
class MyViewModelFactory(
    private val repository: DataRepository
) : ViewModelProvider.Factory {
    override fun <T : ViewModel> create(modelClass: Class<T>): T {
        return MyViewModel(repository) as T
    }
}

class MyActivity : AppCompatActivity() {
    private val repository by lazy { DataRepository() }
    private val viewModel: MyViewModel by viewModels {
        MyViewModelFactory(repository)
    }
}

class MyFragment : Fragment() {
    private val viewModel: MyViewModel by activityViewModels {
        MyViewModelFactory(DataRepository())
    }
}
```

## fa-share-nodes Sharing with Activity

```kotlin
class SharedViewModel : ViewModel() {
    private val _selected = MutableStateFlow<Item?>(null)
    val selected: StateFlow<Item?> = _selected.asStateFlow()

    fun select(item: Item) { _selected.value = item }
}

class MasterFragment : Fragment() {
    private val viewModel: SharedViewModel by activityViewModels()
}

class DetailFragment : Fragment() {
    private val viewModel: SharedViewModel by activityViewModels()

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        viewLifecycleOwner.lifecycleScope.launch {
            viewLifecycleOwner.repeatOnLifecycle(Lifecycle.State.STARTED) {
                viewModel.selected.collect { item -> bind(item) }
            }
        }
    }
}
```

## fa-palette Compose + ViewModel

```kotlin
@Composable
fun MyScreen(viewModel: MyViewModel = viewModel()) {
    val state by viewModel.state.collectAsStateWithLifecycle()

    Column {
        if (state.isLoading) {
            CircularProgressIndicator()
        }
        LazyColumn {
            items(state.items) { item ->
                Text(item, modifier = Modifier.clickable {
                    viewModel.select(item)
                })
            }
        }
    }
}

@Composable
fun NavHostScreen() {
    val navController = rememberNavController()
    val sharedViewModel: SharedViewModel = viewModel()
    NavHost(navController, "home") {
        composable("home") { HomeScreen(sharedViewModel) }
        composable("detail") { DetailScreen(sharedViewModel) }
    }
}
```

## fa-clock ViewModelScope

```kotlin
class TimerViewModel : ViewModel() {
    private val _elapsed = MutableStateFlow(0L)
    val elapsed: StateFlow<Long> = _elapsed.asStateFlow()

    private var timerJob: Job? = null

    fun start() {
        timerJob = viewModelScope.launch {
            while (isActive) {
                delay(1000)
                _elapsed.update { it + 1 }
            }
        }
    }

    fun stop() {
        timerJob?.cancel()
    }
}

class PaginatedViewModel(private val repo: Repository) : ViewModel() {
    private var pageJob: Job? = null

    fun loadPage(page: Int) {
        pageJob?.cancel()
        pageJob = viewModelScope.launch {
            val data = repo.getPage(page)
        }
    }
}
```

## fa-bell One-time Events

```kotlin
sealed class UiEvent {
    data class ShowToast(val message: String) : UiEvent()
    data class Navigate(val route: String) : UiEvent()
    data object ScrollToTop : UiEvent()
}

class EventViewModel : ViewModel() {
    private val _events = Channel<UiEvent>()
    val events = _events.receiveAsFlow()

    fun saveAndNavigate() {
        viewModelScope.launch {
            repository.save(data)
            _events.send(UiEvent.Navigate("detail"))
        }
    }
}

@Composable
fun MyScreen(viewModel: EventViewModel = viewModel()) {
    LaunchedEffect(Unit) {
        viewModel.events.collect { event ->
            when (event) {
                is UiEvent.ShowToast -> println(event.message)
                is UiEvent.Navigate -> navController.navigate(event.route)
                UiEvent.ScrollToTop -> listState.scrollToItem(0)
            }
        }
    }
}
```

## fa-vial Testing ViewModel

```kotlin
class CounterViewModelTest {
    @get:Rule
    val mainRule = MainCoroutineRule()

    private lateinit var viewModel: CounterViewModel

    @Before
    fun setup() {
        viewModel = CounterViewModel(FakeRepository())
    }

    @Test
    fun `increment updates count`() = runTest {
        viewModel.increment()
        assertEquals(1, viewModel.count.value)
    }

    @Test
    fun `load items sets state`() = runTest {
        viewModel.loadItems()
        val state = viewModel.state.value
        assertFalse(state.isLoading)
        assertTrue(state.items.isNotEmpty())
    }
}

class MainCoroutineRule : TestRule {
    val dispatcher = StandardTestDispatcher()
    override fun apply(base: Statement, description: Description): Statement {
        Dispatchers.setMain(dispatcher)
        return object : Statement() {
            override fun evaluate() {
                try { base.evaluate() } finally {
                    Dispatchers.resetMain()
                }
            }
        }
    }
}
```

## fa-syringe Hilt Integration

```kotlin
@HiltViewModel
class UserViewModel @Inject constructor(
    private val repository: UserRepository,
    savedStateHandle: SavedStateHandle
) : ViewModel() {

    private val userId: String = savedStateHandle["userId"] ?: ""

    private val _user = MutableStateFlow<User?>(null)
    val user: StateFlow<User?> = _user.asStateFlow()

    init {
        viewModelScope.launch {
            _user.value = repository.getUser(userId)
        }
    }
}

@AndroidEntryPoint
class UserActivity : AppCompatActivity() {
    private val viewModel: UserViewModel by viewModels()
}

@Composable
fun UserScreen(viewModel: UserViewModel = hiltViewModel()) {
    val user by viewModel.user.collectAsStateWithLifecycle()
}
```
