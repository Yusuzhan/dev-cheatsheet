---
title: sqlx
icon: fa-database
primary: "#00ADD8"
lang: go
---

## fa-plug Connect

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

## fa-table Query Rows

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
    rows.StructScan(&u)
}
```

## fa-crosshairs Query Row

```go
var u User
err := db.Get(&u, "SELECT id, name, email FROM users WHERE id = $1", 1)

var name string
err = db.QueryRowx("SELECT name FROM users WHERE id = $1", 1).Scan(&name)
```

## fa-pen Exec

```go
result, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", "Alice", "a@b.com")
id, _ := result.LastInsertId()
affected, _ := result.RowsAffected()

result, err = db.Exec("UPDATE users SET name = $1 WHERE id = $2", "Bob", 1)
result, err = db.Exec("DELETE FROM users WHERE id = $1", 1)
```

## fa-tags Named Queries

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
result, err = db.NamedExec("INSERT INTO users (name, email) VALUES (:name, :email)", users)

rows, err := db.NamedQuery("SELECT * FROM users WHERE name = :name", map[string]interface{}{"name": "Alice"})
```

## fa-layer-group Struct Scanning

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

## fa-rotate Transactions

```go
tx, err := db.Beginx()
if err != nil {
    log.Fatal(err)
}
defer tx.Rollback()

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

## fa-file-lines Prepared Statements

```go
stmt, err := db.Preparex("SELECT id, name FROM users WHERE id = $1")
defer stmt.Close()

var u User
stmt.Get(&u, 1)

stmt, err = db.Preparex("INSERT INTO users (name, email) VALUES ($1, $2)")
stmt.Exec("Alice", "a@b.com")
```

## fa-circle-xmark NULL Handling

```go
type User struct {
    ID    int            `db:"id"`
    Name  string         `db:"name"`
    Email sql.NullString `db:"email"`
    Age   sql.NullInt64  `db:"age"`
}

var u User
db.Get(&u, "SELECT * FROM users WHERE id = $1", 1)
if u.Email.Valid {
    fmt.Println(u.Email.String)
}
```

## fa-network-wired Connection Pool

```go
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(10)
db.SetConnMaxLifetime(5 * time.Minute)
db.SetConnMaxIdleTime(1 * time.Minute)

err := db.Ping()

stats := db.Stats()
fmt.Printf("open: %d, idle: %d, inuse: %d\n",
    stats.OpenConnections, stats.Idle, stats.InUse)
```

## fa-list IN Queries

```go
ids := []int{1, 2, 3, 4, 5}
query, args, err := sqlx.In("SELECT * FROM users WHERE id IN (?)", ids)
query = db.Rebind(query)
var users []User
db.Select(&users, query, args...)

names := []string{"Alice", "Bob", "Charlie"}
query, args, err = sqlx.In("SELECT * FROM users WHERE name IN (?)", names)
```

## fa-truck-fast Migrations

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
