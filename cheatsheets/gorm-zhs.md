---
title: GORM
icon: fa-database
primary: "#00ADD8"
lang: go
locale: zhs
---

## fa-plug 连接与配置

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

## fa-sitemap 模型定义

```go
type User struct {
    ID        uint           `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`  // 软删除
    Name      string         `gorm:"size:100;not null"`
    Email     string         `gorm:"size:255;uniqueIndex"`
    Age       int
    Orders    []Order        // 一对多关联
}

type Order struct {
    ID     uint  `gorm:"primarykey"`
    UserID uint
    User   User
    Amount float64
}
```

## fa-wand-magic-sparkles 自动迁移

```go
db.AutoMigrate(&User{}, &Order{})

db.Migrator().CreateTable(&User{})
db.Migrator().DropTable(&User{})
db.Migrator().HasTable(&User{})
db.Migrator().AddColumn(&User{}, "Age")
db.Migrator().DropColumn(&User{}, "Age")
db.Migrator().AlterColumn(&User{}, "Name")
```

## fa-plus 创建记录

```go
user := User{Name: "Alice", Email: "a@b.com", Age: 25}
result := db.Create(&user)

result.RowsAffected  // 影响行数
result.Error         // 错误信息

users := []User{
    {Name: "Alice", Email: "a@b.com"},
    {Name: "Bob", Email: "b@b.com"},
}
db.Create(&users)  // 批量创建

db.Select("Name", "Email").Create(&user)  // 指定字段创建
```

## fa-magnifying-glass 查询记录

```go
var user User
db.First(&user, 1)                       // 按主键查询
db.First(&user, "name = ?", "Alice")     // 条件查询

var users []User
db.Find(&users)                          // 查询全部
db.Where("age > ?", 18).Find(&users)     // 条件查询
db.Where("name IN ?", []string{"Alice", "Bob"}).Find(&users)

db.Take(&user)   // 不排序取一条
db.Last(&user)   // 最后一条
```

## fa-pen-to-square 更新记录

```go
db.Model(&user).Update("name", "Bob")    // 单字段更新

db.Model(&user).Updates(User{Name: "Bob", Age: 30})         // 结构体更新（零值会被忽略）
db.Model(&user).Updates(map[string]interface{}{"name": "Bob", "age": 30})  // map 更新（零值不忽略）

db.Model(&User{}).Where("age < ?", 18).Update("status", "minor")  // 批量更新

db.Model(&user).UpdateColumn("name", "Bob")  // 跳过钩子
```

## fa-trash 删除记录

```go
db.Delete(&user)                                    // 软删除（有 DeletedAt 字段时）
db.Delete(&User{}, 1)                               // 按主键删除
db.Where("name = ?", "Alice").Delete(&User{})       // 条件删除

db.Unscoped().Where("id = ?", 1).Delete(&User{})    // 永久删除
db.Unscoped().Find(&users)                          // 包含已软删除的记录

db.Where("1 = 1").Delete(&User{})  // 批量删除
```

## fa-filter 条件查询

```go
db.Where("name = ?", "Alice").First(&user)
db.Where("name <> ?", "Alice").Find(&users)
db.Where("name IN ?", []string{"Alice", "Bob"}).Find(&users)
db.Where("name LIKE ?", "%ali%").Find(&users)
db.Where("age BETWEEN ? AND ?", 18, 30).Find(&users)
db.Where("age > ? AND name = ?", 18, "Alice").Find(&users)

db.Not("name = ?", "Alice").Find(&users)   // NOT 条件
db.Or("name = ?", "Bob").Find(&users)      // OR 条件
```

## fa-link 关联操作

```go
user := User{Name: "Alice", Orders: []Order{{Amount: 99.99}}}
db.Create(&user)

db.Model(&user).Association("Orders").Append([]Order{{Amount: 49.99}})  // 追加关联
db.Model(&user).Association("Orders").Delete(&order)                    // 删除关联
db.Model(&user).Association("Orders").Replace([]Order{{Amount: 199.99}})  // 替换关联
db.Model(&user).Association("Orders").Clear()     // 清空关联
db.Model(&user).Association("Orders").Count()     // 关联数量
```

## fa-download 预加载

```go
db.Preload("Orders").Find(&users)  // 预加载关联

db.Preload("Orders", func(db *gorm.DB) *gorm.DB {  // 带条件的预加载
    return db.Order("amount DESC")
}).Find(&users)

db.Joins("Orders").Find(&users)  // JOIN 预加载

db.Preload("Orders").Preload("Orders.Items").Find(&users)  // 嵌套预加载
```

## fa-rotate 事务

```go
db.Transaction(func(tx *gorm.DB) error {  // 自动提交/回滚
    if err := tx.Create(&User{Name: "Alice"}).Error; err != nil {
        return err
    }
    if err := tx.Create(&Order{UserID: 1, Amount: 99.99}).Error; err != nil {
        return err
    }
    return nil
})

tx := db.Begin()       // 手动事务
tx.Create(&User{Name: "Alice"})
tx.Commit()
tx.Rollback()
```

## fa-filter-circle-xmatter Scopes 作用域

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

db.Scopes(AgeGreaterThan(18), Paginate(1, 10)).Find(&users)  // 组合使用
```

## fa-bolt 钩子函数

```go
func (u *User) BeforeCreate(tx *gorm.DB) error {  // 创建前
    if u.Name == "" {
        return errors.New("name is required")
    }
    return nil
}

func (u *User) AfterCreate(tx *gorm.DB) error {   // 创建后
    log.Printf("created user: %s", u.Name)
    return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error { return nil }
func (u *User) AfterUpdate(tx *gorm.DB) error   { return nil }
func (u *User) BeforeDelete(tx *gorm.DB) error  { return nil }
func (u *User) AfterDelete(tx *gorm.DB) error   { return nil }
```

## fa-terminal 原生 SQL

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
