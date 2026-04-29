---
title: Android Notification
icon: fa-bell
primary: "#3DDC84"
lang: kotlin
locale: zhs
---

## fa-bell 基本通知

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

## fa-layer-group 通知渠道

```kotlin
fun createChannel(context: Context) {
    if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
        val channel = NotificationChannel(
            CHANNEL_ID,
            "消息",
            NotificationManager.IMPORTANCE_DEFAULT
        ).apply {
            description = "通用消息"
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

## fa-hand-pointer 操作按钮

```kotlin
val replyIntent = Intent(context, ReplyReceiver::class.java)
val replyPendingIntent = PendingIntent.getBroadcast(
    context, 0, replyIntent,
    PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_MUTABLE
)

val replyAction = RemoteInput.Builder("key_reply")
    .setLabel("回复")
    .build()
    .let { remoteInput ->
        NotificationCompat.Action.Builder(
            R.drawable.ic_reply, "回复", replyPendingIntent
        )
            .addRemoteInput(remoteInput)
            .build()
    }

val dismissAction = NotificationCompat.Action.Builder(
    R.drawable.ic_dismiss, "忽略",
    dismissPendingIntent
).build()

val notification = NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setContentTitle("新消息")
    .setContentText("您有一条新消息")
    .addAction(replyAction)
    .addAction(dismissAction)
    .build()
```

## fa-align-left 大文本与收件箱样式

```kotlin
val bigTextStyle = NotificationCompat.BigTextStyle()
    .bigText("这是一条很长的消息，用户展开通知后会完整显示。支持多行文本。")
    .setBigContentTitle("展开标题")
    .setSummaryText("3 条新消息")

NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setStyle(bigTextStyle)
    .build()
```

```kotlin
val inboxStyle = NotificationCompat.InboxStyle()
    .setBigContentTitle("3 条新消息")
    .setSummaryText("user@example.com")
    .addLine("Alice: 你有空吗？")
    .addLine("Bob: 看看这个链接")
    .addLine("Charlie: 明天见")

NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setContentTitle("3 条新消息")
    .setContentText("user@example.com")
    .setStyle(inboxStyle)
    .setNumber(3)
    .build()
```

## fa-spinner 进度通知

```kotlin
val builder = NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setContentTitle("下载中")
    .setContentText("正在下载")
    .setOngoing(true)

Thread {
    for (progress in 0..100 step 5) {
        builder.setProgress(100, progress, false)
        NotificationManagerCompat.from(context).notify(PROGRESS_ID, builder.build())
        Thread.sleep(500)
    }
    builder.setContentText("下载完成")
        .setProgress(0, 0, false)
        .setOngoing(false)
    NotificationManagerCompat.from(context).notify(PROGRESS_ID, builder.build())
}.start()
```

## fa-palette 自定义布局

```kotlin
val remoteViews = RemoteViews(context.packageName, R.layout.notification_custom).apply {
    setTextViewText(R.id.title, "自定义标题")
    setTextViewText(R.id.message, "自定义消息")
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

## fa-satellite-dish 前台服务

```kotlin
class LocationService : Service() {
    override fun onStartCommand(intent: Intent?, flags: Int, startId: Int): Int {
        val notification = NotificationCompat.Builder(this, CHANNEL_ID)
            .setSmallIcon(R.drawable.ic_location)
            .setContentTitle("定位追踪")
            .setContentText("正在追踪您的位置")
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
        "定位",
        NotificationManager.IMPORTANCE_LOW
    )
    NotificationManagerCompat.from(context).createNotificationChannel(channel)
}

ContextCompat.startForegroundService(context, Intent(context, LocationService::class.java))
```

## fa-shield 通知权限

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

## fa-link 深度链接

```kotlin
val deepLinkIntent = TaskStackBuilder.create(context).run {
    addNextIntentWithParentStack(
        Intent(Intent.ACTION_VIEW, Uri.parse("myapp://detail/42"))
    )
    getPendingIntent(0, PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_IMMUTABLE)
}

val notification = NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setContentTitle("新文章")
    .setContentText("点击查看")
    .setContentIntent(deepLinkIntent)
    .setAutoCancel(true)
    .build()
```

## fa-object-group 分组通知

```kotlin
val groupKey = "group_messages"

val summary = NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setContentTitle("2 条新消息")
    .setContentText("Alice, Bob")
    .setStyle(
        NotificationCompat.InboxStyle()
            .addLine("Alice: 你好！")
            .addLine("Bob: 最近怎么样？")
            .setBigContentTitle("2 条新消息")
    )
    .setGroup(groupKey)
    .setGroupSummary(true)
    .build()

val msg1 = NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setContentTitle("Alice")
    .setContentText("你好！")
    .setGroup(groupKey)
    .build()

val msg2 = NotificationCompat.Builder(context, CHANNEL_ID)
    .setSmallIcon(R.drawable.ic_notification)
    .setContentTitle("Bob")
    .setContentText("最近怎么样？")
    .setGroup(groupKey)
    .build()

NotificationManagerCompat.from(context).apply {
    notify(1, msg1)
    notify(2, msg2)
    notify(SUMMARY_ID, summary)
}
```

## fa-circle-dot 角标

```kotlin
if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
    val channel = NotificationChannel(CHANNEL_ID, "消息", NotificationManager.IMPORTANCE_DEFAULT).apply {
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

## fa-xmark 取消与更新

```kotlin
fun updateNotification(context: Context, id: Int, text: String) {
    val notification = NotificationCompat.Builder(context, CHANNEL_ID)
        .setSmallIcon(R.drawable.ic_notification)
        .setContentTitle("已更新")
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
