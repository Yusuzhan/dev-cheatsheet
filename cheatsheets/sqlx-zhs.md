---
title: sqlx
icon: fa-database
primary: "#00ADD8"
lang: go
locale: zhs
---

## fa-plug 连接数据库

```go
import "github.com/jmoiron/sqlx"
import _ "github.com/lib/pq"

db, err := sqlx.Connect("postgres", "user=postgres dbname=mydb sslmode=disable")
if err != nil {
    log.Fatal(err)
}
defer db.Close()

db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)
db.SetConnMaxLifetime(5 * time.Minute)
```

## fa-table 查询多行

```go
type User struct {
    ID    int    `db:"id"`
    Name  string `db:"name"`
    Email string `db:"email"`
}

users := []User{}
err := db.Select(&users, "SELECT id, name, email FROM users ORDER BY id")

rows, err := db.Queryx("SELECT id, name FROM users WHERE active = $1", true)
for rows.Next() {
    var u User
    rows.StructScan(&u)  // 扫描到结构体
}
```

## fa-crosshairs 查询单行

```go
var u User
err := db.Get(&u, "SELECT id, name, email FROM users WHERE id = $1", 1)

var name string
err = db.QueryRowx("SELECT name FROM users WHERE id = $1", 1).Scan(&name)
```

## fa-pen 执行写操作

```go
result, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", "Alice", "a@b.com")
id, _ := result.LastInsertId()
affected, _ := result.RowsAffected()

result, err = db.Exec("UPDATE users SET name = $1 WHERE id = $2", "Bob", 1)
result, err = db.Exec("DELETE FROM users WHERE id = $1", 1)
```

## fa-tags 命名查询

```go
type User struct {
    Name  string `db:"name"`
    Email string `db:"email"`
}

u := User{Name: "Alice", Email: "a@b.com"}
result, err := db.NamedExec("INSERT INTO users (name, email) VALUES (:name, :email)", u)

users := []User{
    {Name: "Alice", Email: "a@b.com"},
    {Name: "Bob", Email: "b@b.com"},
}
result, err = db.NamedExec("INSERT INTO users (name, email) VALUES (:name, :email)", users)  // 批量插入

rows, err := db.NamedQuery("SELECT * FROM users WHERE name = :name", map[string]interface{}{"name": "Alice"})
```

## fa-layer-group 结构体扫描

```go
type User struct {
    ID       int    `db:"id"`
    Name     string `db:"name"`
    Email    string `db:"email"`
    Age      int    `db:"age"`
}

var u User
db.Get(&u, "SELECT * FROM users WHERE id = $1", 1)

var users []User
db.Select(&users, "SELECT * FROM users WHERE age > $1", 18)

row := db.QueryRowx("SELECT * FROM users WHERE id = $1", 1)
row.StructScan(&u)
```

## fa-rotate 事务

```go
tx, err := db.Beginx()
if err != nil {
    log.Fatal(err)
}
defer tx.Rollback()  // 失败时自动回滚

tx.Exec("INSERT INTO orders (user_id, amount) VALUES ($1, $2)", 1, 99.99)
tx.Exec("UPDATE users SET orders = orders + 1 WHERE id = $1", 1)

if err := tx.Commit(); err != nil {
    log.Fatal(err)
}

err = sqlx.InTx(context.Background(), db, func(tx *sqlx.Tx) error {
    tx.Exec("INSERT INTO orders (user_id, amount) VALUES ($1, $2)", 1, 99.99)
    tx.Exec("UPDATE users SET orders = orders + 1 WHERE id = $1", 1)
    return nil
})
```

## fa-file-lines 预处理语句

```go
stmt, err := db.Preparex("SELECT id, name FROM users WHERE id = $1")
defer stmt.Close()

var u User
stmt.Get(&u, 1)

stmt, err = db.Preparex("INSERT INTO users (name, email) VALUES ($1, $2)")
stmt.Exec("Alice", "a@b.com")
```

## fa-circle-xmark NULL 处理

```go
type User struct {
    ID    int            `db:"id"`
    Name  string         `db:"name"`
    Email sql.NullString `db:"email"`  // 可能为 NULL 的字段
    Age   sql.NullInt64  `db:"age"`
}

var u User
db.Get(&u, "SELECT * FROM users WHERE id = $1", 1)
if u.Email.Valid {
    fmt.Println(u.Email.String)
}
```

## fa-network-wired 连接池

```go
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(10)
db.SetConnMaxLifetime(5 * time.Minute)
db.SetConnMaxIdleTime(1 * time.Minute)

err := db.Ping()  // 检测连接

stats := db.Stats()
fmt.Printf("open: %d, idle: %d, inuse: %d\n",
    stats.OpenConnections, stats.Idle, stats.InUse)
```

## fa-list IN 查询

```go
ids := []int{1, 2, 3, 4, 5}
query, args, err := sqlx.In("SELECT * FROM users WHERE id IN (?)", ids)
query = db.Rebind(query)  // 将 ? 转为 $1, $2... 等占位符
var users []User
db.Select(&users, query, args...)

names := []string{"Alice", "Bob", "Charlie"}
query, args, err = sqlx.In("SELECT * FROM users WHERE name IN (?)", names)
```

## fa-truck-fast 数据库迁移

```go
schema := `
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);
`
db.MustExec(schema)

db.MustExec("ALTER TABLE users ADD COLUMN IF NOT EXISTS age INTEGER")
db.MustExec("DROP TABLE IF EXISTS temp_data")
```
