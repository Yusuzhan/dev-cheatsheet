---
title: Android Four Components
icon: fa-puzzle-piece
primary: "#3DDC84"
lang: kotlin
---

## fa-rotate Activity Lifecycle

```kotlin
class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
    }
    override fun onStart() { super.onStart() }
    override fun onResume() { super.onResume() }
    override fun onPause() { super.onPause() }
    override fun onStop() { super.onStop() }
    override fun onDestroy() { super.onDestroy() }
    override fun onRestart() { super.onRestart() }

    override fun onSaveInstanceState(outState: Bundle) {
        super.onSaveInstanceState(outState)
        outState.putString("key_text", "saved data")
    }

    override fun onRestoreInstanceState(savedInstanceState: Bundle) {
        super.onRestoreInstanceState(savedInstanceState)
        val text = savedInstanceState.getString("key_text")
    }
}
```

## fa-arrow-right Activity Intent & Navigation

```kotlin
val intent = Intent(this, DetailActivity::class.java).apply {
    putExtra("item_id", 42L)
    putExtra("item_name", "example")
}
startActivity(intent)

startActivity(Intent(Intent.ACTION_VIEW, Uri.parse("https://example.com")))

val launcher = registerForActivityResult(ActivityResultContracts.StartActivityForResult()) { result ->
    if (result.resultCode == RESULT_OK) {
        val data = result.data?.getStringExtra("result_key")
    }
}
launcher.launch(Intent(this, EditActivity::class.java))
```

## fa-window-maximize Fragment

```kotlin
class HomeFragment : Fragment(R.layout.fragment_home) {
    private var _binding: FragmentHomeBinding? = null
    private val binding get() = _binding!!

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        _binding = FragmentHomeBinding.bind(view)
        binding.textView.text = "Hello"
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }
}
```

```kotlin
parentFragmentManager.beginTransaction()
    .replace(R.id.container, HomeFragment::class.java, bundleOf("key" to "value"))
    .addToBackStack("home")
    .commit()

val fragment = parentFragmentManager.findFragmentById(R.id.container)
```

## fa-play Service (Started)

```kotlin
class DownloadService : Service() {
    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        val url = intent?.getStringExtra("url")
        Thread { downloadFile(url) }.start()
        return START_NOT_STICKY
    }

    private fun downloadFile(url: String?) { /* ... */ }
    override fun onBind(intent: Intent): IBinder? = null
}
```

```kotlin
val intent = Intent(context, DownloadService::class.java).apply {
    putExtra("url", "https://example.com/file.zip")
}
ContextCompat.startForegroundService(context, intent)
context.stopService(intent)
```

## fa-link Service (Bound)

```kotlin
class MusicService : Service() {
    private val binder = LocalBinder()

    inner class LocalBinder : Binder() {
        fun getService(): MusicService = this@MusicService
    }

    fun play() { /* ... */ }
    fun pause() { /* ... */ }
    override fun onBind(intent: Intent): IBinder = binder
}
```

```kotlin
private lateinit var musicService: MusicService
private var bound = false

private val connection = object : ServiceConnection {
    override fun onServiceConnected(name: ComponentName, service: IBinder) {
        musicService = (service as MusicService.LocalBinder).getService()
        bound = true
    }
    override fun onServiceDisconnected(name: ComponentName) { bound = false }
}

override fun onStart() {
    super.onStart()
    bindService(Intent(this, MusicService::class.java), connection, Context.BIND_AUTO_CREATE)
}

override fun onStop() {
    super.onStop()
    if (bound) { unbindService(connection); bound = false }
}
```

## fa-tower-broadcast BroadcastReceiver

```kotlin
class BatteryReceiver : BroadcastReceiver() {
    override fun onReceive(context: Context, intent: Intent) {
        val level = intent.getIntExtra(BatteryManager.EXTRA_LEVEL, -1)
        val scale = intent.getIntExtra(BatteryManager.EXTRA_SCALE, -1)
        val percent = (level * 100) / scale
    }
}
```

```kotlin
val receiver = BatteryReceiver()
registerReceiver(receiver, IntentFilter(Intent.ACTION_BATTERY_CHANGED))
unregisterReceiver(receiver)
```

```xml
<receiver android:name=".BatteryReceiver" android:exported="false">
    <intent-filter>
        <action android:name="android.intent.action.BOOT_COMPLETED" />
    </intent-filter>
</receiver>
```

## fa-database ContentProvider

```kotlin
class UserProvider : ContentProvider() {
    private lateinit var dbHelper: SQLiteOpenHelper

    override fun onCreate(): Boolean {
        dbHelper = DbHelper(context!!)
        return true
    }

    override fun query(
        uri: Uri, projection: Array<String>?, selection: String?,
        selectionArgs: Array<String>?, sortOrder: String?
    ): Cursor? {
        val db = dbHelper.readableDatabase
        return db.query("users", projection, selection, selectionArgs, null, null, sortOrder)
    }

    override fun getType(uri: Uri): String = "vnd.android.cursor.dir/vnd.example.users"
    override fun insert(uri: Uri, values: ContentValues?): Uri? = null
    override fun delete(uri: Uri, selection: String?, selectionArgs: Array<String>?): Int = 0
    override fun update(uri: Uri, values: ContentValues?, selection: String?, selectionArgs: Array<String>?): Int = 0
}
```

```kotlin
val cursor = contentResolver.query(
    Uri.parse("content://com.example.provider/users"),
    null, null, null, null
)
cursor?.use {
    while (it.moveToNext()) {
        val name = it.getString(it.getColumnIndexOrThrow("name"))
    }
}
```

## fa-file-code Manifest

```xml
<manifest xmlns:android="http://schemas.android.com/apk/res/android"
    package="com.example.app">

    <uses-permission android:name="android.permission.INTERNET" />
    <uses-permission android:name="android.permission.ACCESS_FINE_LOCATION" />

    <application
        android:name=".App"
        android:allowBackup="true"
        android:icon="@mipmap/ic_launcher"
        android:label="@string/app_name"
        android:theme="@style/Theme.AppCompat">

        <activity android:name=".MainActivity" android:exported="true">
            <intent-filter>
                <action android:name="android.intent.action.MAIN" />
                <category android:name="android.intent.category.LAUNCHER" />
            </intent-filter>
        </activity>

        <service android:name=".DownloadService" android:exported="false" />
        <receiver android:name=".BatteryReceiver" android:exported="false" />
        <provider
            android:name=".UserProvider"
            android:authorities="com.example.provider"
            android:exported="false" />
    </application>
</manifest>
```

## fa-filter Intent Filters

```xml
<activity android:name=".DetailActivity" android:exported="true">
    <intent-filter>
        <action android:name="android.intent.action.VIEW" />
        <category android:name="android.intent.category.DEFAULT" />
        <category android:name="android.intent.category.BROWSABLE" />
        <data
            android:scheme="myapp"
            android:host="detail"
            android:pathPrefix="/item" />
    </intent-filter>
</activity>
```

```kotlin
val intent = Intent(Intent.ACTION_SEND).apply {
    type = "text/plain"
    putExtra(Intent.EXTRA_TEXT, "Share this content")
}
startActivity(Intent.createChooser(intent, "Share via"))
```

## fa-lock Permissions

```kotlin
if (ContextCompat.checkSelfPermission(this, Manifest.permission.ACCESS_FINE_LOCATION)
    != PackageManager.PERMISSION_GRANTED
) {
    requestPermissionLauncher.launch(Manifest.permission.ACCESS_FINE_LOCATION)
}

val requestPermissionLauncher = registerForActivityResult(
    ActivityResultContracts.RequestPermission()
) { granted ->
    if (granted) startLocationUpdates()
}
```

```kotlin
val multipleLauncher = registerForActivityResult(
    ActivityResultContracts.RequestMultiplePermissions()
) { permissions ->
    val allGranted = permissions.all { it.value }
    if (allGranted) startCamera()
}
multipleLauncher.launch(
    arrayOf(
        Manifest.permission.CAMERA,
        Manifest.permission.RECORD_AUDIO
    )
)
```

## fa-file-contract AIDL

```java
// IRemoteService.aidl
interface IRemoteService {
    int getPid();
    String getMessage();
}
```

```kotlin
class RemoteService : Service() {
    private val binder = object : IRemoteService.Stub() {
        override fun getPid(): Int = Process.myPid()
        override fun getMessage(): String = "Hello from remote"
    }
    override fun onBind(intent: Intent): IBinder = binder
}
```

```kotlin
private var remoteService: IRemoteService? = null

private val connection = object : ServiceConnection {
    override fun onServiceConnected(name: ComponentName, service: IBinder) {
        remoteService = IRemoteService.Stub.asInterface(service)
        val pid = remoteService?.pid
        val msg = remoteService?.message
    }
    override fun onServiceDisconnected(name: ComponentName) { remoteService = null }
}

bindService(Intent(this, RemoteService::class.java), connection, Context.BIND_AUTO_CREATE)
```

## fa-rocket App Startup

```kotlin
class TimberInitializer : Initializer<Unit> {
    override fun create(context: Context) {
        Timber.plant(Timber.DebugTree())
    }
    override fun dependencies(): List<Class<out Initializer<*>>> = emptyList()
}

class WorkManagerInitializer : Initializer<WorkManager> {
    override fun create(context: Context): WorkManager {
        val config = Configuration.Builder().build()
        WorkManager.initialize(context, config)
        return WorkManager.getInstance(context)
    }
    override fun dependencies(): List<Class<out Initializer<*>>> = listOf(TimberInitializer::class.java)
}
```

```xml
<provider
    android:name="androidx.startup.InitializationProvider"
    android:authorities="${applicationId}.androidx-startup"
    android:exported="false"
    tools:node="merge">
    <meta-data
        android:name="com.example.TimberInitializer"
        android:value="androidx.startup" />
</provider>
```
