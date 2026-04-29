---
title: ADB
icon: fa-mobile-screen
primary: "#3DDC84"
lang: bash
---

## fa-link Devices & Connection

```bash
adb devices
adb devices -l
adb connect 192.168.1.100:5555
adb disconnect
adb usb
adb tcpip 5555
adb kill-server
adb start-server
adb -s <serial> shell
```

## fa-terminal Shell Commands

```bash
adb shell
adb shell ls /sdcard/
adb shell cat /proc/cpuinfo
adb shell uptime
adb shell date
adb shell id
adb shell whoami
adb shell ps
adb shell top -n 5
```

## fa-box Package Management

```bash
adb shell pm list packages
adb shell pm list packages -3
adb shell pm list packages -s
adb shell pm path com.example.app
adb shell pm clear com.example.app
adb shell pm dump com.example.app
adb shell pm enable com.example.app
adb shell pm disable com.example.app
```

## fa-download Install & Uninstall APK

```bash
adb install app.apk
adb install -r app.apk
adb install -g app.apk
adb install -t app.apk
adb install-multiple base.apk split.apk
adb uninstall com.example.app
adb uninstall -k com.example.app
```

## fa-arrows-up-down File Transfer

```bash
adb push local.txt /sdcard/remote.txt
adb push ./folder /sdcard/backup/
adb pull /sdcard/file.txt ./
adb pull /sdcard/folder/ ./backup/
adb shell run-as com.example.app cat databases/app.db > app.db
```

## fa-scroll Logcat

```bash
adb logcat
adb logcat *:E
adb logcat *:W
adb logcat -s TAG:V
adb logcat --pid=<pid>
adb logcat -d
adb logcat -c
adb logcat -f /sdcard/log.txt
adb logcat -v time
adb logcat -v threadtime
adb logcat | grep "keyword"
adb logcat ActivityManager:I *:S
```

## fa-camera Screenshot & Screen Record

```bash
adb shell screencap /sdcard/screen.png
adb pull /sdcard/screen.png
adb shell screencap -p /sdcard/screen.png
adb shell screenrecord /sdcard/video.mp4
adb shell screenrecord --time-limit 30 /sdcard/video.mp4
adb shell screenrecord --bit-rate 8000000 /sdcard/video.mp4
adb shell screenrecord --size 1280x720 /sdcard/video.mp4
```

## fa-keyboard Input Simulation

```bash
adb shell input text "hello"
adb shell input tap 500 500
adb shell input swipe 500 1000 500 200
adb shell input swipe 500 500 500 500 1000
adb shell input keyevent KEYCODE_HOME
adb shell input keyevent KEYCODE_BACK
adb shell input keyevent KEYCODE_ENTER
adb shell input keyevent KEYCODE_VOLUME_UP
adb shell input keyevent 26
```

## fa-layer-group Activity & Intent

```bash
adb shell am start -n com.example.app/.MainActivity
adb shell am start -a android.intent.action.VIEW -d https://example.com
adb shell am start -a android.intent.action.CALL -d tel:123456
adb shell am broadcast -a android.intent.action.BOOT_COMPLETED
adb shell am force-stop com.example.app
adb shell am start -n com.example.app/.Activity -e key value
adb shell dumpsys activity top
```

## fa-magnifying-glass Dumpsys & Profiling

```bash
adb shell dumpsys battery
adb shell dumpsys meminfo com.example.app
adb shell dumpsys cpuinfo
adb shell dumpsys window displays
adb shell dumpsys activity activities
adb shell dumpsys notification
adb shell dumpsys gfxinfo com.example.app
adb shell dumpsys dbinfo
adb shell top -m 10 -t
```

## fa-network-wired Network

```bash
adb shell netstat
adb shell ifconfig
adb shell ip addr
adb shell ping -c 4 google.com
adb shell curl -I https://example.com
adb forward tcp:8080 tcp:8080
adb forward tcp:9222 localabstract:chrome_devtools_remote
adb forward --list
adb forward --remove tcp:8080
adb reverse tcp:3000 tcp:3000
```

## fa-circle-info Device Info

```bash
adb shell getprop ro.build.version.release
adb shell getprop ro.build.version.sdk
adb shell getprop ro.product.model
adb shell getprop ro.product.brand
adb shell getprop ro.hardware
adb shell settings get secure android_id
adb shell wm size
adb shell wm density
adb shell df -h
```

## fa-desktop Emulator

```bash
adb shell wm size 1080x1920
adb shell wm size reset
adb shell wm density 480
adb shell wm density reset
adb shell settings put global screen_brightness 128
adb shell settings put system screen_off_timeout 60000
adb shell svc power stayon true
adb shell svc wifi enable
adb shell svc wifi disable
```

## fa-bolt Useful One-liners

```bash
adb shell pm list packages -3 | cut -d: -f2
adb shell "dumpsys window | grep -i mCurrentFocus"
adb shell getprop | grep build
adb shell am start -a android.settings.APPLICATION_DETAILS_SETTINGS -d package:com.example.app
adb shell monkey -p com.example.app -c android.intent.category.LAUNCHER 1
adb shell cmd package compile -m speed -f com.example.app
for pkg in $(adb shell pm list packages -3 -u | tr -d '\r' | cut -d: -f2); do adb uninstall "$pkg"; done
```
