---
title: GORM
icon: fa-database
primary: "#00ADD8"
lang: go
---

## fa-plug Connect & Setup

```go
import "gorm.io/gorm"
import "gorm.io/driver/postgres"

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
    log.Fatal(err)
}

sqlDB, _ := db.DB()
sqlDB.SetMaxOpenConns(100)
sqlDB.SetMaxIdleConns(10)
sqlDB.SetConnMaxLifetime(time.Hour)
```

## fa-sitemap Model Definition

```go
type User struct {
    ID        uint           `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
    Name      string         `gorm:"size:100;not null"`
    Email     string         `gorm:"size:255;uniqueIndex"`
    Age       int
    Orders    []Order
}

type Order struct {
    ID     uint  `gorm:"primarykey"`
    UserID uint
    User   User
    Amount float64
}
```

## fa-wand-magic-sparkles Auto Migrate

```go
db.AutoMigrate(&User{}, &Order{})

db.Migrator().CreateTable(&User{})
db.Migrator().DropTable(&User{})
db.Migrator().HasTable(&User{})
db.Migrator().AddColumn(&User{}, "Age")
db.Migrator().DropColumn(&User{}, "Age")
db.Migrator().AlterColumn(&User{}, "Name")
```

## fa-plus Create

```go
user := User{Name: "Alice", Email: "a@b.com", Age: 25}
result := db.Create(&user)

result.RowsAffected
result.Error

users := []User{
    {Name: "Alice", Email: "a@b.com"},
    {Name: "Bob", Email: "b@b.com"},
}
db.Create(&users)

db.Select("Name", "Email").Create(&user)
```

## fa-magnifying-glass Read

```go
var user User
db.First(&user, 1)
db.First(&user, "name = ?", "Alice")

var users []User
db.Find(&users)
db.Where("age > ?", 18).Find(&users)
db.Where("name IN ?", []string{"Alice", "Bob"}).Find(&users)

db.Take(&user)  // without ordering
db.Last(&user)
```

## fa-pen-to-square Update

```go
db.Model(&user).Update("name", "Bob")

db.Model(&user).Updates(User{Name: "Bob", Age: 30})
db.Model(&user).Updates(map[string]interface{}{"name": "Bob", "age": 30})

db.Model(&User{}).Where("age < ?", 18).Update("status", "minor")

db.Model(&user).UpdateColumn("name", "Bob")
```

## fa-trash Delete

```go
db.Delete(&user)
db.Delete(&User{}, 1)
db.Where("name = ?", "Alice").Delete(&User{})

db.Unscoped().Where("id = ?", 1).Delete(&User{})
db.Unscoped().Find(&users)

db.Where("1 = 1").Delete(&User{})
```

## fa-filter Where Conditions

```go
db.Where("name = ?", "Alice").First(&user)
db.Where("name <> ?", "Alice").Find(&users)
db.Where("name IN ?", []string{"Alice", "Bob"}).Find(&users)
db.Where("name LIKE ?", "%ali%").Find(&users)
db.Where("age BETWEEN ? AND ?", 18, 30).Find(&users)
db.Where("age > ? AND name = ?", 18, "Alice").Find(&users)

db.Not("name = ?", "Alice").Find(&users)
db.Or("name = ?", "Bob").Find(&users)
```

## fa-link Associations

```go
user := User{Name: "Alice", Orders: []Order{{Amount: 99.99}}}
db.Create(&user)

db.Model(&user).Association("Orders").Append([]Order{{Amount: 49.99}})
db.Model(&user).Association("Orders").Delete(&order)
db.Model(&user).Association("Orders").Replace([]Order{{Amount: 199.99}})
db.Model(&user).Association("Orders").Clear()
db.Model(&user).Association("Orders").Count()
```

## fa-download Preloading

```go
db.Preload("Orders").Find(&users)

db.Preload("Orders", func(db *gorm.DB) *gorm.DB {
    return db.Order("amount DESC")
}).Find(&users)

db.Joins("Orders").Find(&users)

db.Preload("Orders").Preload("Orders.Items").Find(&users)
```

## fa-rotate Transactions

```go
db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&User{Name: "Alice"}).Error; err != nil {
        return err
    }
    if err := tx.Create(&Order{UserID: 1, Amount: 99.99}).Error; err != nil {
        return err
    }
    return nil
})

tx := db.Begin()
tx.Create(&User{Name: "Alice"})
tx.Commit()
tx.Rollback()
```

## fa-filter-circle-xmark Scopes

```go
func AgeGreaterThan(age int) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        return db.Where("age > ?", age)
    }
}

func Paginate(page, size int) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        offset := (page - 1) * size
        return db.Offset(offset).Limit(size)
    }
}

db.Scopes(AgeGreaterThan(18), Paginate(1, 10)).Find(&users)
```

## fa-bolt Hooks

```go
func (u *User) BeforeCreate(tx *gorm.DB) error {
    if u.Name == "" {
        return errors.New("name is required")
    }
    return nil
}

func (u *User) AfterCreate(tx *gorm.DB) error {
    log.Printf("created user: %s", u.Name)
    return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error { return nil }
func (u *User) AfterUpdate(tx *gorm.DB) error   { return nil }
func (u *User) BeforeDelete(tx *gorm.DB) error  { return nil }
func (u *User) AfterDelete(tx *gorm.DB) error   { return nil }
```

## fa-terminal Raw SQL

```go
type Result struct {
    Name  string
    Total int
}

var results []Result
db.Raw("SELECT name, count(*) as total FROM users GROUP BY name").Scan(&results)

db.Exec("UPDATE users SET name = ? WHERE id = ?", "Bob", 1)

db.Raw("SELECT * FROM users WHERE id = ?", 1).Scan(&user)

rows, _ := db.Raw("SELECT id, name FROM users").Rows()
defer rows.Close()
for rows.Next() {
    var id int
    var name string
    rows.Scan(&id, &name)
}
```
