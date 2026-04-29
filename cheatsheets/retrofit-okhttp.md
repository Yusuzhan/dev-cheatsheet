---
title: Retrofit / OkHttp
icon: fa-plug
primary: "#3DDC84"
lang: kotlin
---

## fa-gears Retrofit Setup

```kotlin
interface Api {
    @GET("users")
    suspend fun getUsers(): List<User>

    @GET("users/{id}")
    suspend fun getUser(@Path("id") id: Long): User
}

val retrofit = Retrofit.Builder()
    .baseUrl("https://api.example.com/")
    .addConverterFactory(GsonConverterFactory.create())
    .build()

val api = retrofit.create(Api::class.java)
```

## fa-download GET Requests

```kotlin
interface Api {
    @GET("users")
    suspend fun getUsers(): List<User>

    @GET("users/{id}")
    suspend fun getUser(@Path("id") id: Long): User

    @GET("users")
    suspend fun searchUsers(@Query("q") query: String, @Query("page") page: Int = 1): List<User>

    @GET("users")
    suspend fun getUsers(@Header("Authorization") token: String): List<User>

    @GET("items")
    suspend fun getItems(@QueryMap filters: Map<String, String>): List<Item>
}
```

## fa-upload POST Requests

```kotlin
interface Api {
    @POST("users")
    suspend fun createUser(@Body user: CreateUserRequest): User

    @PUT("users/{id}")
    suspend fun updateUser(@Path("id") id: Long, @Body user: UpdateUserRequest): User

    @PATCH("users/{id}")
    suspend fun patchUser(@Path("id") id: Long, @Body fields: Map<String, @JvmSuppressWildcards Any>): User

    @DELETE("users/{id}")
    suspend fun deleteUser(@Path("id") id: Long)

    @FormUrlEncoded
    @POST("auth/login")
    suspend fun login(
        @Field("username") username: String,
        @Field("password") password: String
    ): TokenResponse
}
```

## fa-code Path & Query Params

```kotlin
interface Api {
    @GET("repos/{owner}/{repo}/issues")
    suspend fun getIssues(
        @Path("owner") owner: String,
        @Path("repo") repo: String,
        @Query("state") state: String = "open",
        @Query("per_page") perPage: Int = 30,
        @Query("page") page: Int = 1
    ): List<Issue>

    @GET("search")
    suspend fun search(
        @QueryMap(encoded = true) params: Map<String, String>
    ): SearchResult

    @Url
    @GET
    suspend fun getFromUrl(@Url url: String): ResponseBody
}
```

## fa-heading Headers

```kotlin
interface Api {
    @Headers(
        "Accept: application/json",
        "X-Custom-Header: value"
    )
    @GET("users")
    suspend fun getUsers(): List<User>

    @GET("users")
    suspend fun getUsers(@Header("Authorization") token: String): List<User>

    @GET("users")
    suspend fun getUsers(@HeaderMap headers: Map<String, String>): List<User>
}
```

```kotlin
val okhttp = OkHttpClient.Builder()
    .addInterceptor { chain ->
        val request = chain.request().newBuilder()
            .addHeader("Authorization", "Bearer $token")
            .addHeader("Accept", "application/json")
            .build()
        chain.proceed(request)
    }
    .build()
```

## fa-paperclip Multipart Upload

```kotlin
interface Api {
    @Multipart
    @POST("upload")
    suspend fun uploadFile(
        @Part file: MultipartBody.Part
    ): UploadResponse

    @Multipart
    @POST("upload")
    suspend fun uploadWithMeta(
        @Part file: MultipartBody.Part,
        @Part("description") description: RequestBody,
        @Part("tags") tags: RequestBody
    ): UploadResponse

    @Multipart
    @PUT("avatar")
    suspend fun uploadAvatar(@Part avatar: MultipartBody.Part): User
}
```

```kotlin
val file = File("/path/to/file.jpg")
val requestFile = file.asRequestBody("image/jpeg".toMediaType())
val multipart = MultipartBody.Part.createFormData("file", file.name, requestFile)
val description = "My photo".toRequestBody("text/plain".toMediaType())
api.uploadWithMeta(multipart, description)
```

## fa-server OkHttp Client

```kotlin
val client = OkHttpClient.Builder()
    .connectTimeout(30, TimeUnit.SECONDS)
    .readTimeout(30, TimeUnit.SECONDS)
    .writeTimeout(30, TimeUnit.SECONDS)
    .connectionPool(ConnectionPool(5, 5, TimeUnit.MINUTES))
    .cache(Cache(File(context.cacheDir, "http_cache"), 10L * 1024 * 1024))
    .cookieJar(JavaNetCookieJar(CookieManager()))
    .followRedirects(true)
    .followSslRedirects(true)
    .build()

val retrofit = Retrofit.Builder()
    .baseUrl("https://api.example.com/")
    .client(client)
    .addConverterFactory(GsonConverterFactory.create())
    .build()
```

## fa-filter Interceptors

```kotlin
class AuthInterceptor(private val tokenProvider: () -> String) : Interceptor {
    override fun intercept(chain: Interceptor.Chain): Response {
        val request = chain.request().newBuilder()
            .addHeader("Authorization", "Bearer ${tokenProvider()}")
            .build()
        return chain.proceed(request)
    }
}

class ErrorInterceptor : Interceptor {
    override fun intercept(chain: Interceptor.Chain): Response {
        val response = chain.proceed(chain.request())
        if (!response.isSuccessful) {
            throw HttpException(
                Response.error<Any>(
                    response.code,
                    response.body?.string()?.toResponseBody(response.body?.contentType())
                )
            )
        }
        return response
    }
}
```

```kotlin
val client = OkHttpClient.Builder()
    .addInterceptor(AuthInterceptor { token })
    .addNetworkInterceptor(CacheInterceptor())
    .build()
```

## fa-terminal Logging

```kotlin
val loggingInterceptor = HttpLoggingInterceptor().apply {
    level = if (BuildConfig.DEBUG)
        HttpLoggingInterceptor.Level.BODY
    else
        HttpLoggingInterceptor.Level.NONE
    redactHeader("Authorization")
    redactHeader("Cookie")
}

val client = OkHttpClient.Builder()
    .addInterceptor(loggingInterceptor)
    .build()
```

## fa-key Authenticator

```kotlin
class TokenAuthenticator(
    private val tokenApi: TokenApi,
    private val tokenStore: TokenStore
) : Authenticator {
    override fun authenticate(route: Route?, response: Response): Request? {
        if (responseCount(response) >= 3) return null

        val refreshToken = tokenStore.getRefreshToken() ?: return null
        val newToken = runBlocking {
            tokenApi.refreshToken("refresh_token", refreshToken)
        }

        tokenStore.saveTokens(newToken.accessToken, newToken.refreshToken)
        return response.request.newBuilder()
            .header("Authorization", "Bearer ${newToken.accessToken}")
            .build()
    }

    private fun responseCount(response: Response): Int {
        var count = 1
        var prior = response.priorResponse
        while (prior != null) { count++; prior = prior.priorResponse }
        return count
    }
}
```

## fa-triangle-exclamation Error Handling

```kotlin
sealed class ApiResult<out T> {
    data class Success<T>(val data: T) : ApiResult<T>()
    data class Error(val code: Int, val message: String) : ApiResult<Nothing>()
    data class Exception(val throwable: Throwable) : ApiResult<Nothing>()
}

suspend fun <T> safeApiCall(call: suspend () -> T): ApiResult<T> {
    return try {
        ApiResult.Success(call())
    } catch (e: HttpException) {
        val errorBody = e.response()?.errorBody()?.string()
        ApiResult.Error(e.code(), errorBody ?: e.message())
    } catch (e: IOException) {
        ApiResult.Exception(e)
    } catch (e: Exception) {
        ApiResult.Exception(e)
    }
}
```

```kotlin
val result = safeApiCall { api.getUsers() }
when (result) {
    is ApiResult.Success -> showUsers(result.data)
    is ApiResult.Error -> showError(result.code, result.message)
    is ApiResult.Exception -> showError(result.throwable.message ?: "Unknown error")
}
```

## fa-bolt Coroutine Support

```kotlin
class UserViewModel(private val api: Api) : ViewModel() {
    private val _users = MutableStateFlow<List<User>>(emptyList())
    val users: StateFlow<List<User>> = _users.asStateFlow()

    init {
        viewModelScope.launch {
            try {
                _users.value = api.getUsers()
            } catch (e: Exception) {
                Log.e("UserVM", "Failed to load users", e)
            }
        }
    }

    fun refresh() {
        viewModelScope.launch {
            _users.value = api.getUsers()
        }
    }
}
```

```kotlin
suspend fun fetchAll() = coroutineScope {
    val users = async { api.getUsers() }
    val posts = async { api.getPosts() }
    DashboardData(users.await(), posts.await())
}
```

## fa-file-arrow-down File Download

```kotlin
interface Api {
    @Streaming
    @GET("files/{id}")
    suspend fun downloadFile(@Path("id") id: String): ResponseBody
}
```

```kotlin
suspend fun downloadAndSave(api: Api, id: String, dest: File) {
    val body = api.downloadFile(id)
    body.byteStream().use { input ->
        dest.outputStream().use { output ->
            val buffer = ByteArray(8192)
            var bytesRead: Int
            while (input.read(buffer).also { bytesRead = it } != -1) {
                output.write(buffer, 0, bytesRead)
            }
        }
    }
}
```

## fa-box-archive Cache

```kotlin
val cache = Cache(File(context.cacheDir, "http_cache"), 10L * 1024 * 1024)

val client = OkHttpClient.Builder()
    .cache(cache)
    .addNetworkInterceptor { chain ->
        val response = chain.proceed(chain.request())
        response.newBuilder()
            .header("Cache-Control", "public, max-age=300")
            .build()
    }
    .addInterceptor { chain ->
        var request = chain.request()
        if (!isNetworkAvailable(context)) {
            request = request.newBuilder()
                .cacheControl(CacheControl.FORCE_CACHE)
                .build()
        }
        chain.proceed(request)
    }
    .build()
```

## fa-certificate Certificate Pinning

```kotlin
val certificatePinner = CertificatePinner.Builder()
    .add("api.example.com", "sha256/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=")
    .add("api.example.com", "sha256/BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB=")
    .build()

val client = OkHttpClient.Builder()
    .certificatePinner(certificatePinner)
    .build()
```

```kotlin
val pins = listOf(
    "sha256/AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=",
    "sha256/BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB="
)
val pinnedClient = OkHttpClient.Builder()
    .certificatePinner(
        CertificatePinner.Builder()
            .add("*.example.com", *pins.toTypedArray())
            .build()
    )
    .build()
```
