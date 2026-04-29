---
title: Android Notification
icon: fa-bell
primary: "#3DDC84"
lang: kotlin
---

## fa-bell Basic Notification

```kotlin
val notification = NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setContentTitle("Title")
    .setContentText("Message text")
    .setPriority(NotificationCompat.PRIORITY_DEFAULT)
    .setContentIntent(pendingIntent)
    .setAutoCancel(true)
    .build()

NotificationManagerCompat.from(context).notify(notificationId, notification)
```

## fa-layer-group Channels

```kotlin
fun createChannel(context: Context) {
    if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
        val channel = NotificationChannel(
            CHANNEL_ID,
            "Messages",
            NotificationManager.IMPORTANCE_DEFAULT
        ).apply {
            description = "General messages"
            enableLights(true)
            lightColor = Color.GREEN
            enableVibration(true)
            vibrationPattern = longArrayOf(0, 200, 100, 200)
            setShowBadge(true)
        }
        val manager = context.getSystemService(Context.NOTIFICATION_SERVICE) as NotificationManager
        manager.createNotificationChannel(channel)
    }
}
```

## fa-hand-pointer Actions

```kotlin
val replyIntent = Intent(context, ReplyReceiver::class.java)
val replyPendingIntent = PendingIntent.getBroadcast(
    context, 0, replyIntent,
    PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_MUTABLE
)

val replyAction = RemoteInput.Builder("key_reply")
    .setLabel("Reply")
    .build()
    .let { remoteInput ->
        NotificationCompat.Action.Builder(
            R.drawable.ic_reply, "Reply", replyPendingIntent
        )
            .addRemoteInput(remoteInput)
            .build()
    }

val dismissAction = NotificationCompat.Action.Builder(
    R.drawable.ic_dismiss, "Dismiss",
    dismissPendingIntent
).build()

val notification = NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setContentTitle("New message")
    .setContentText("You have a new message")
    .addAction(replyAction)
    .addAction(dismissAction)
    .build()
```

## fa-align-left Large Text & Inbox Style

```kotlin
val bigTextStyle = NotificationCompat.BigTextStyle()
    .bigText("This is a long message that will be expanded when the user expands the notification. It supports multiline text display.")
    .setBigContentTitle("Expanded Title")
    .setSummaryText("3 new messages")

NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setStyle(bigTextStyle)
    .build()
```

```kotlin
val inboxStyle = NotificationCompat.InboxStyle()
    .setBigContentTitle("3 new messages")
    .setSummaryText("user@example.com")
    .addLine("Alice: Hey, are you free?")
    .addLine("Bob: Check out this link")
    .addLine("Charlie: See you tomorrow")

NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setContentTitle("3 new messages")
    .setContentText("user@example.com")
    .setStyle(inboxStyle)
    .setNumber(3)
    .build()
```

## fa-spinner Progress

```kotlin
val builder = NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setContentTitle("Downloading")
    .setContentText("Download in progress")
    .setOngoing(true)

Thread {
    for (progress in 0..100 step 5) {
        builder.setProgress(100, progress, false)
        NotificationManagerCompat.from(context).notify(PROGRESS_ID, builder.build())
        Thread.sleep(500)
    }
    builder.setContentText("Download complete")
        .setProgress(0, 0, false)
        .setOngoing(false)
    NotificationManagerCompat.from(context).notify(PROGRESS_ID, builder.build())
}.start()
```

## fa-palette Custom Layout

```kotlin
val remoteViews = RemoteViews(context.packageName, R.layout.notification_custom).apply {
    setTextViewText(R.id.title, "Custom Title")
    setTextViewText(R.id.message, "Custom message")
    setImageViewResource(R.id.icon, R.drawable.ic_notification)
    setOnClickPendingIntent(R.id.action_button, actionPendingIntent)
}

val notification = NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setCustomContentView(remoteViews)
    .setCustomBigContentView(remoteViews)
    .setStyle(NotificationCompat.DecoratedCustomViewStyle())
    .build()
```

## fa-satellite-dish Foreground Service

```kotlin
class LocationService : Service() {
    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        val notification = NotificationCompat.Builder(this, CHANNEL_ID)
            .setSmallIcon(R.drawable.ic_location)
            .setContentTitle("Location tracking")
            .setContentText("Tracking your location")
            .setOngoing(true)
            .build()

        startForeground(FOREGROUND_ID, notification)
        return START_STICKY
    }

    override fun onBind(intent: Intent?): IBinder? = null
}
```

```kotlin
if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
    val channel = NotificationChannel(
        "location_channel",
        "Location",
        NotificationManager.IMPORTANCE_LOW
    )
    NotificationManagerCompat.from(context).createNotificationChannel(channel)
}

ContextCompat.startForegroundService(context, Intent(context, LocationService::class.java))
```

## fa-shield Permission

```kotlin
if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.TIRAMISU) {
    if (ContextCompat.checkSelfPermission(context, Manifest.permission.POST_NOTIFICATIONS)
        != PackageManager.PERMISSION_GRANTED
    ) {
        requestPermissionLauncher.launch(Manifest.permission.POST_NOTIFICATIONS)
    }
}

val requestPermissionLauncher = registerForActivityResult(
    ActivityResultContracts.RequestPermission()
) { isGranted: Boolean ->
    if (isGranted) {
        showNotification()
    }
}
```

## fa-link Deep Link

```kotlin
val deepLinkIntent = TaskStackBuilder.create(context).run {
    addNextIntentWithParentStack(
        Intent(Intent.ACTION_VIEW, Uri.parse("myapp://detail/42"))
    )
    getPendingIntent(0, PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_IMMUTABLE)
}

val notification = NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setContentTitle("New article")
    .setContentText("Tap to read")
    .setContentIntent(deepLinkIntent)
    .setAutoCancel(true)
    .build()
```

## fa-object-group Grouped Notifications

```kotlin
val groupKey = "group_messages"

val summary = NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setContentTitle("2 new messages")
    .setContentText("Alice, Bob")
    .setStyle(
        NotificationCompat.InboxStyle()
            .addLine("Alice: Hello!")
            .addLine("Bob: What's up?")
            .setBigContentTitle("2 new messages")
    )
    .setGroup(groupKey)
    .setGroupSummary(true)
    .build()

val msg1 = NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setContentTitle("Alice")
    .setContentText("Hello!")
    .setGroup(groupKey)
    .build()

val msg2 = NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setContentTitle("Bob")
    .setContentText("What's up?")
    .setGroup(groupKey)
    .build()

NotificationManagerCompat.from(context).apply {
    notify(1, msg1)
    notify(2, msg2)
    notify(SUMMARY_ID, summary)
}
```

## fa-circle-dot Badge

```kotlin
if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
    val channel = NotificationChannel(CHANNEL_ID, "Messages", NotificationManager.IMPORTANCE_DEFAULT).apply {
        setShowBadge(true)
        setSound(null, null)
    }
    val manager = context.getSystemService(NotificationManager::class.java)
    manager.createNotificationChannel(channel)
}

fun clearBadge(context: Context) {
    NotificationManagerCompat.from(context).cancelAll()
}
```

## fa-xmark Cancel & Update

```kotlin
fun updateNotification(context: Context, id: Int, text: String) {
    val notification = NotificationCompat.Builder(context, CHANNEL_ID)
        .setSmallIcon(R.drawable.ic_notification)
        .setContentTitle("Updated")
        .setContentText(text)
        .build()
    NotificationManagerCompat.from(context).notify(id, notification)
}

fun cancelNotification(context: Context, id: Int) {
    NotificationManagerCompat.from(context).cancel(id)
}

fun cancelAll(context: Context) {
    NotificationManagerCompat.from(context).cancelAll()
}
```
