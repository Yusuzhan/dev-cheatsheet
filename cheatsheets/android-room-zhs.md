---
title: Android Room
icon: fa-database
primary: "#3DDC84"
lang: kotlin
locale: zhs
---

## fa-table 实体定义

```kotlin
@Entity(tableName = "users")
data class User(
    @PrimaryKey(autoGenerate = true) val id: Long = 0,
    @ColumnInfo(name = "user_name") val userName: String,
    @ColumnInfo(name = "email") val email: String,
    @ColumnInfo(name = "age") val age: Int,
    @ColumnInfo(name = "is_active") val isActive: Boolean = true,
    @ColumnInfo(name = "created_at") val createdAt: Long = System.currentTimeMillis()
)
```

```kotlin
@Entity(
    tableName = "articles",
    indices = [Index(value = ["slug"], unique = true)],
    foreignKeys = [ForeignKey(
        entity = User::class,
        parentColumns = ["id"],
        childColumns = ["author_id"],
        onDelete = ForeignKey.CASCADE
    )]
)
data class Article(
    @PrimaryKey val id: Long,
    val title: String,
    val slug: String,
    @ColumnInfo(name = "author_id") val authorId: Long
)
```

## fa-magnifying-glass DAO（查询）

```kotlin
@Dao
interface UserDao {
    @Query("SELECT * FROM users")
    fun getAll(): List<User>

    @Query("SELECT * FROM users WHERE id = :id")
    fun getById(id: Long): User?

    @Query("SELECT * FROM users WHERE user_name LIKE '%' || :query || '%'")
    fun searchByName(query: String): List<User>

    @Query("SELECT * FROM users WHERE age BETWEEN :minAge AND :maxAge")
    fun getByAgeRange(minAge: Int, maxAge: Int): List<User>

    @Query("SELECT COUNT(*) FROM users WHERE is_active = 1")
    fun getActiveUserCount(): Int
}
```

## fa-pen-to-square DAO（增删改）

```kotlin
@Dao
interface UserDao {
    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insert(user: User): Long

    @Insert(onConflict = OnConflictStrategy.IGNORE)
    suspend fun insertAll(users: List<User>): List<Long>

    @Update
    suspend fun update(user: User)

    @Delete
    suspend fun delete(user: User)

    @Query("DELETE FROM users WHERE is_active = 0")
    suspend fun deleteInactiveUsers(): Int
}
```

## fa-database Room 数据库配置

```kotlin
@Database(
    entities = [User::class, Article::class],
    version = 2,
    exportSchema = true
)
abstract class AppDatabase : RoomDatabase() {
    abstract fun userDao(): UserDao
    abstract fun articleDao(): ArticleDao

    companion object {
        @Volatile
        private var INSTANCE: AppDatabase? = null

        fun getInstance(context: Context): AppDatabase =
            INSTANCE ?: synchronized(this) {
                INSTANCE ?: Room.databaseBuilder(
                    context.applicationContext,
                    AppDatabase::class.java,
                    "app_database"
                )
                    .addMigrations(MIGRATION_1_2)
                    .fallbackToDestructiveMigration()
                    .build()
                    .also { INSTANCE = it }
            }
    }
}
```

## fa-link 关系（一对多）

```kotlin
data class UserWithArticles(
    @Embedded val user: User,
    @Relation(
        parentColumn = "id",
        entityColumn = "author_id"
    )
    val articles: List<Article>
)
```

```kotlin
@Dao
interface ArticleDao {
    @Transaction
    @Query("SELECT * FROM users WHERE id = :userId")
    suspend fun getUserWithArticles(userId: Long): UserWithArticles

    @Transaction
    @Query("SELECT * FROM users")
    suspend fun getAllUsersWithArticles(): List<UserWithArticles>
}
```

## fa-arrows-left-right 关系（多对多）

```kotlin
@Entity(primaryKeys = ["articleId", "tagId"])
data class ArticleTagCrossRef(
    val articleId: Long,
    val tagId: Long
)

@Entity
data class Tag(
    @PrimaryKey(autoGenerate = true) val id: Long = 0,
    val name: String
)

data class ArticleWithTags(
    @Embedded val article: Article,
    @Relation(
        parentColumn = "id",
        entityColumn = "id",
        associateBy = Junction(
            ArticleTagCrossRef::class,
            parentColumn = "articleId",
            entityColumn = "tagId"
        )
    )
    val tags: List<Tag>
)
```

## fa-exchange-alt 类型转换器

```kotlin
class Converters {
    @TypeConverter
    fun fromTimestamp(value: Long?): Date? = value?.let { Date(it) }

    @TypeConverter
    fun dateToTimestamp(date: Date?): Long? = date?.time

    @TypeConverter
    fun fromStringList(value: String?): List<String> =
        value?.split(",")?.map { it.trim() } ?: emptyList()

    @TypeConverter
    fun toStringList(list: List<String>?): String? =
        list?.joinToString(",")
}
```

```kotlin
@TypeConverters(Converters::class)
abstract class AppDatabase : RoomDatabase() { ... }
```

## fa-arrow-up 数据库迁移

```kotlin
val MIGRATION_1_2 = object : Migration(1, 2) {
    override fun migrate(db: SupportSQLiteDatabase) {
        db.execSQL("ALTER TABLE users ADD COLUMN avatar_url TEXT DEFAULT NULL")
        db.execSQL("CREATE INDEX index_users_email ON users(email)")
    }
}

val MIGRATION_2_3 = object : Migration(2, 3) {
    override fun migrate(db: SupportSQLiteDatabase) {
        db.execSQL("CREATE TABLE tags (id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, name TEXT NOT NULL)")
    }
}
```

## fa-wave-square Flow/LiveData 查询

```kotlin
@Dao
interface UserDao {
    @Query("SELECT * FROM users WHERE is_active = 1")
    fun observeActiveUsers(): Flow<List<User>>

    @Query("SELECT * FROM users WHERE id = :id")
    fun observeUser(id: Long): Flow<User?>

    @Query("SELECT * FROM users")
    fun getAllLiveData(): LiveData<List<User>>

    @Query("SELECT COUNT(*) FROM users")
    fun getUserCountLiveData(): LiveData<Int>
}
```

```kotlin
@ViewModelScoped
class UserViewModel @Inject constructor(
    private val userDao: UserDao
) : ViewModel() {
    val activeUsers: StateFlow<List<User>> = userDao
        .observeActiveUsers()
        .stateIn(viewModelScope, SharingStarted.WhileSubscribed(5000), emptyList())
}
```

## fa-lock 事务

```kotlin
@Dao
interface ArticleDao {
    @Transaction
    @Query("SELECT * FROM articles WHERE id = :id")
    suspend fun getArticleWithAuthor(id: Long): ArticleWithAuthor

    @Transaction
    suspend fun replaceAllArticles(articles: List<Article>) {
        deleteAll()
        insertAll(articles)
    }
}

@Dao
abstract class MixedDao {
    @Transaction
    open suspend fun updateUserAndArticle(user: User, article: Article) {
        updateUser(user)
        insertArticle(article)
    }

    abstract suspend fun updateUser(user: User)
    abstract suspend fun insertArticle(article: Article)
}
```

## fa-layer-group 嵌入对象

```kotlin
data class Address(
    val street: String,
    val city: String,
    val zipCode: String
)

@Entity(tableName = "users")
data class User(
    @PrimaryKey(autoGenerate = true) val id: Long = 0,
    val userName: String,
    @Embedded(prefix = "home_") val homeAddress: Address,
    @Embedded(prefix = "work_") val workAddress: Address?
)
```

## fa-search 全文搜索 (FTS)

```kotlin
@Fts4(contentEntity = Article::class)
@Entity(tableName = "articles_fts")
data class ArticleFts(
    @PrimaryKey @ColumnInfo(name = "rowid") val rowId: Long,
    val title: String,
    val slug: String
)
```

```kotlin
@Dao
interface ArticleDao {
    @Query("SELECT articles.* FROM articles JOIN articles_fts ON articles.rowid = articles_fts.rowid WHERE articles_fts MATCH :query")
    suspend fun searchArticles(query: String): List<Article>
}
```

## fa-vial 测试

```kotlin
@RunWith(AndroidJUnit4::class)
class UserDaoTest {
    private lateinit var database: AppDatabase
    private lateinit var userDao: UserDao

    @Before
    fun setup() {
        database = Room.inMemoryDatabaseBuilder(
            ApplicationProvider.getApplicationContext(),
            AppDatabase::class.java
        ).allowMainThreadQueries().build()
        userDao = database.userDao()
    }

    @After
    fun teardown() = database.close()

    @Test
    fun insertAndRetrieve() = runTest {
        val user = User(userName = "alice", email = "alice@test.com", age = 25)
        val id = userDao.insert(user)
        val loaded = userDao.getById(id)
        assertEquals("alice", loaded?.userName)
    }

    @Test
    fun searchByName() = runTest {
        userDao.insertAll(listOf(
            User(userName = "alice", email = "a@test.com", age = 25),
            User(userName = "bob", email = "b@test.com", age = 30)
        ))
        val results = userDao.searchByName("ali")
        assertEquals(1, results.size)
    }
}
```
