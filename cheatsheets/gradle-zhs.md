---
title: Gradle
icon: fa-gears
primary: "#02303A"
lang: groovy
locale: zhs
---

## fa-folder-tree 项目结构

```
MyApp/
├── build.gradle(.kts)        # 项目级构建文件
├── settings.gradle(.kts)      # 模块列表
├── gradle.properties          # JVM/Gradle 属性
├── gradlew / gradlew.bat      # Wrapper 脚本
├── gradle/
│   └── wrapper/
│       ├── gradle-wrapper.jar
│       └── gradle-wrapper.properties
└── app/
    └── build.gradle(.kts)     # 模块级构建文件
```

## fa-mobile-screen build.gradle (app)

```groovy
plugins {
    id 'com.android.application'
    id 'org.jetbrains.kotlin.android'
    id 'kotlin-kapt'
}

android {
    namespace 'com.example.app'
    compileSdk 34

    defaultConfig {
        applicationId "com.example.app"
        minSdk 24
        targetSdk 34
        versionCode 1
        versionName "1.0.0"
        testInstrumentationRunner "androidx.test.runner.AndroidJUnitRunner"
    }
}

dependencies {
    implementation "androidx.core:core-ktx:1.12.0"
    implementation "androidx.appcompat:appcompat:1.6.1"
    testImplementation "junit:junit:4.13.2"
}
```

## fa-link 依赖管理

```groovy
dependencies {
    implementation "androidx.core:core-ktx:1.12.0"
    api "androidx.appcompat:appcompat:1.6.1"
    compileOnly "javax.annotation:javax.annotation-api:1.3.2"
    annotationProcessor "com.github.bumptech.glide:compiler:4.16.0"
    kapt "com.google.dagger:hilt-compiler:2.50"
    debugImplementation "com.squareup.leakcanary:leakcanary-android:2.12"
    testImplementation "junit:junit:4.13.2"
    androidTestImplementation "androidx.test.ext:junit:1.1.5"
}
```

依赖配置说明：`implementation`、`api`、`compileOnly`、`runtimeOnly`、`annotationProcessor`、`kapt`、`ksp`、`testImplementation`、`androidTestImplementation`、`debugImplementation`。

## fa-sliders 构建类型与变体

```groovy
android {
    buildTypes {
        debug {
            minifyEnabled false
            debuggable true
            applicationIdSuffix ".debug"
            buildConfigField "String", "BASE_URL", '"https://api.dev.example.com"'
        }
        release {
            minifyEnabled true
            shrinkResources true
            proguardFiles getDefaultProguardFile('proguard-android-optimize.txt'), 'proguard-rules.pro'
            signingConfig signingConfigs.release
        }
    }

    flavorDimensions "environment"
    productFlavors {
        dev {
            dimension "environment"
            applicationIdSuffix ".dev"
            versionNameSuffix "-dev"
        }
        prod {
            dimension "environment"
        }
    }
}
```

## fa-key 签名配置

```groovy
android {
    signingConfigs {
        debug {
            storeFile file("debug.keystore")
            storePassword "android"
            keyAlias "androiddebugkey"
            keyPassword "android"
        }
        release {
            storeFile file(System.getenv("KEYSTORE_FILE") ?: "release.keystore")
            storePassword System.getenv("KEYSTORE_PASSWORD") ?: ""
            keyAlias System.getenv("KEY_ALIAS") ?: ""
            keyPassword System.getenv("KEY_PASSWORD") ?: ""
        }
    }
}
```

## fa-shield-halved ProGuard / R8

```groovy
android {
    buildTypes {
        release {
            minifyEnabled true
            shrinkResources true
            proguardFiles getDefaultProguardFile('proguard-android-optimize.txt'), 'proguard-rules.pro'
        }
    }
}
```

```proguard
-keepattributes *Annotation*
-keepattributes SourceFile,LineNumberTable
-keep public class * extends java.lang.Exception
-keep class com.example.model.** { *; }
-dontwarn retrofit2.**
-keepclassmembers,allowobfuscation class * {
    @com.google.gson.annotations.SerializedName <fields>;
}
```

## fa-wrench 自定义任务

```groovy
tasks.register('copyAssets', Copy) {
    from "src/main/assets"
    into "${buildDir}/output/assets"
}

tasks.register('generateVersionFile') {
    def outputDir = file("${buildDir}/generated/version")
    outputs.dir outputDir
    doLast {
        outputDir.mkdirs()
        new File(outputDir, "version.txt").text = """
            versionName=${android.defaultConfig.versionName}
            versionCode=${android.defaultConfig.versionCode}
        """.stripIndent()
    }
}

tasks.named('preBuild') {
    dependsOn 'generateVersionFile'
}
```

## fa-building build.gradle (project)

```groovy
plugins {
    id 'com.android.application' version '8.2.0' apply false
    id 'com.android.library' version '8.2.0' apply false
    id 'org.jetbrains.kotlin.android' version '1.9.22' apply false
}

ext {
    kotlinVersion = '1.9.22'
    coroutinesVersion = '1.7.3'
    lifecycleVersion = '2.7.0'
    roomVersion = '2.6.1'
}

tasks.register('clean', Delete) {
    delete rootProject.layout.buildDirectory
}
```

## fa-gear settings.gradle

```groovy
pluginManagement {
    repositories {
        google()
        mavenCentral()
        gradlePluginPortal()
    }
}

dependencyResolutionManagement {
    repositoriesMode.set(RepositoriesMode.FAIL_ON_PROJECT_REPOS)
    repositories {
        google()
        mavenCentral()
        maven { url 'https://jitpack.io' }
    }
}

rootProject.name = "MyApp"
include ':app'
include ':core'
include ':feature:home'
include ':feature:profile'
```

## fa-file-lines gradle.properties

```properties
org.gradle.jvmargs=-Xmx4096m -XX:+HeapDumpOnOutOfMemoryError -Dfile.encoding=UTF-8
org.gradle.parallel=true
org.gradle.caching=true
org.gradle.configuration-cache=true
android.useAndroidX=true
android.nonTransitiveRClass=true
kotlin.code.style=official
```

## fa-terminal 构建命令

```bash
./gradlew assembleDebug                     # 构建 Debug APK
./gradlew assembleRelease                   # 构建 Release APK
./gradlew bundleRelease                     # 构建 AAB (Play Store)
./gradlew installDebug                      # 构建并安装到设备
./gradlew clean                             # 清理构建输出
./gradlew :app:dependencies                 # 查看依赖树
./gradlew :app:dependencies --configuration releaseRuntimeClasspath
./gradlew tasks --all                       # 列出所有任务
./gradlew --refresh-dependencies            # 强制刷新依赖
```

## fa-object-group 复合构建

```groovy
includeBuild('shared-utils') {
    dependencySubstitution {
        substitute module('com.example:shared-utils') using project(':')
    }
}

includeBuild('build-logic') {
    dependencySubstitution {
        substitute module('com.example.build-logic:convention') using project(':convention')
    }
}
```

## fa-tags 版本目录

```groovy
enableFeaturePreview("VERSION_CATALOG")
```

```toml
# gradle/libs.versions.toml
[versions]
kotlin = "1.9.22"
coroutines = "1.7.3"
room = "2.6.1"

[libraries]
kotlinx-coroutines-android = { module = "org.jetbrains.kotlinx:kotlinx-coroutines-android", version.ref = "coroutines" }
room-runtime = { module = "androidx.room:room-runtime", version.ref = "room" }
room-ktx = { module = "androidx.room:room-ktx", version.ref = "room" }

[plugins]
android-application = { id = "com.android.application", version = "8.2.0" }
kotlin-android = { id = "org.jetbrains.kotlin.android", version.ref = "kotlin" }
```

```groovy
dependencies {
    implementation libs.room.runtime
    implementation libs.room.ktx
    implementation libs.kotlinx.coroutines.android
}
```
