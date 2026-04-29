---
title: Gradle
icon: fa-gears
primary: "#02303A"
lang: groovy
---

## fa-folder-tree Project Structure

```
MyApp/
├── build.gradle(.kts)        # project-level
├── settings.gradle(.kts)      # module list
├── gradle.properties          # JVM/Gradle props
├── gradlew / gradlew.bat      # wrapper scripts
├── gradle/
│   └── wrapper/
│       ├── gradle-wrapper.jar
│       └── gradle-wrapper.properties
└── app/
    └── build.gradle(.kts)     # module-level
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

## fa-link Dependencies

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

Dependency configurations: `implementation`, `api`, `compileOnly`, `runtimeOnly`, `annotationProcessor`, `kapt`, `ksp`, `testImplementation`, `androidTestImplementation`, `debugImplementation`.

## fa-sliders Build Types & Flavors

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

## fa-key Signing Config

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

## fa-wrench Custom Tasks

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

## fa-terminal Build Commands

```bash
./gradlew assembleDebug                     # debug APK
./gradlew assembleRelease                   # release APK
./gradlew bundleRelease                     # AAB for Play Store
./gradlew installDebug                      # build & install on device
./gradlew clean                             # clean build outputs
./gradlew :app:dependencies                 # dependency tree
./gradlew :app:dependencies --configuration releaseRuntimeClasspath
./gradlew tasks --all                       # list all tasks
./gradlew --refresh-dependencies            # force refresh deps
```

## fa-object-group Composite Builds

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

## fa-tags Version Catalog

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
