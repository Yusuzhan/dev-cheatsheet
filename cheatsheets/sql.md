---
title: SQL
icon: fa-database
primary: "#FF6B35"
lang: sql
---

## fa-table 基础查询

```sql
-- 查询所有数据
SELECT * FROM users;

-- 查询指定列
SELECT name, email FROM users;

-- 限制结果数量
SELECT * FROM users LIMIT 10;

-- 去重
SELECT DISTINCT city FROM users;

-- 别名
SELECT name AS user_name, age AS user_age FROM users;
```

## fa-filter 条件过滤

```sql
-- WHERE 条件
SELECT * FROM users WHERE age > 18;
SELECT * FROM users WHERE name = 'Alice';
SELECT * FROM users WHERE age BETWEEN 20 AND 30;

-- AND / OR / NOT
SELECT * FROM users WHERE age > 18 AND city = 'Beijing';
SELECT * FROM users WHERE city = 'Shanghai' OR city = 'Beijing';

-- IN 和 NOT IN
SELECT * FROM users WHERE city IN ('Beijing', 'Shanghai', 'Guangzhou');

-- LIKE 模糊匹配
SELECT * FROM users WHERE name LIKE 'A%';    -- 以 A 开头
SELECT * FROM users WHERE name LIKE '%son';   -- 以 son 结尾
SELECT * FROM users WHERE email LIKE '%@gmail.com';

-- IS NULL / IS NOT NULL
SELECT * FROM users WHERE phone IS NULL;
```

## fa-sort 排序与分页

```sql
-- ORDER BY 排序
SELECT * FROM users ORDER BY age ASC;          -- 升序 (默认)
SELECT * FROM users ORDER BY age DESC;         -- 降序
SELECT * FROM users ORDER BY city, age DESC;   -- 多列排序

-- LIMIT 和 OFFSET 分页
SELECT * FROM users LIMIT 10;                  -- 前 10 条
SELECT * FROM users LIMIT 10 OFFSET 20;        -- 第 3 页 (每页10条)
```

## fa-chart-bar 聚合函数

```sql
-- 常用聚合函数
SELECT COUNT(*) FROM users;                    -- 总行数
SELECT AVG(age) FROM users;                    -- 平均值
SELECT SUM(salary) FROM employees;             -- 总和
SELECT MAX(age) FROM users;                    -- 最大值
SELECT MIN(age) FROM users;                    -- 最小值

-- GROUP BY 分组
SELECT city, COUNT(*) FROM users GROUP BY city;
SELECT department, AVG(salary) FROM employees
GROUP BY department;

-- HAVING 过滤分组
SELECT city, COUNT(*) as cnt FROM users
GROUP BY city
HAVING cnt > 5;
```

## fa-link JOIN 连接

```sql
-- INNER JOIN (内连接)
SELECT u.name, o.total
FROM users u
INNER JOIN orders o ON u.id = o.user_id;

-- LEFT JOIN (左连接)
SELECT u.name, o.total
FROM users u
LEFT JOIN orders o ON u.id = o.user_id;

-- RIGHT JOIN (右连接)
SELECT u.name, o.total
FROM users u
RIGHT JOIN orders o ON u.id = o.user_id;

-- FULL OUTER JOIN (全连接)
SELECT u.name, o.total
FROM users u
FULL OUTER JOIN orders o ON u.id = o.user_id;

-- CROSS JOIN (交叉连接)
SELECT u.name, d.name
FROM users u CROSS JOIN departments d;

-- 自连接
SELECT e.name AS employee, m.name AS manager
FROM employees e
LEFT JOIN employees m ON e.manager_id = m.id;
```

## fa-code-branch 子查询

```sql
-- WHERE 中的子查询
SELECT * FROM users
WHERE id IN (SELECT user_id FROM orders WHERE total > 1000);

-- FROM 中的子查询
SELECT city, avg_age
FROM (SELECT city, AVG(age) as avg_age FROM users GROUP BY city) t
WHERE avg_age > 25;

-- EXISTS
SELECT * FROM users u
WHERE EXISTS (SELECT 1 FROM orders o WHERE o.user_id = u.id);
```

## fa-pen-to-square 数据操作 (DML)

```sql
-- INSERT 插入
INSERT INTO users (name, email, age) VALUES ('Alice', 'a@example.com', 25);
INSERT INTO users (name, email) VALUES ('Bob', 'b@example.com');

-- 批量插入
INSERT INTO users (name, email) VALUES
  ('Alice', 'a@example.com'),
  ('Bob', 'b@example.com'),
  ('Charlie', 'c@example.com');

-- UPDATE 更新
UPDATE users SET age = 26 WHERE name = 'Alice';
UPDATE users SET age = age + 1 WHERE city = 'Beijing';

-- DELETE 删除
DELETE FROM users WHERE id = 100;
DELETE FROM users WHERE last_login < '2023-01-01';

-- TRUNCATE 清空表
TRUNCATE TABLE temp_data;
```

## fa-table-columns 表操作 (DDL)

```sql
-- CREATE TABLE
CREATE TABLE users (
    id         INT PRIMARY KEY AUTO_INCREMENT,
    name       VARCHAR(100) NOT NULL,
    email      VARCHAR(255) UNIQUE,
    age        INT DEFAULT 0,
    city       VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ALTER TABLE
ALTER TABLE users ADD COLUMN phone VARCHAR(20);
ALTER TABLE users DROP COLUMN phone;
ALTER TABLE users MODIFY COLUMN name VARCHAR(200);
ALTER TABLE users RENAME COLUMN name TO full_name;

-- DROP TABLE
DROP TABLE IF EXISTS temp_data;

-- CREATE INDEX
CREATE INDEX idx_city ON users(city);
CREATE UNIQUE INDEX idx_email ON users(email);
```

## fa-window-restore 窗口函数

```sql
-- ROW_NUMBER
SELECT name, age,
    ROW_NUMBER() OVER (ORDER BY age DESC) AS rank
FROM users;

-- RANK / DENSE_RANK
SELECT name, department, salary,
    RANK() OVER (PARTITION BY department ORDER BY salary DESC) AS dept_rank
FROM employees;

-- LAG / LEAD
SELECT date, revenue,
    LAG(revenue, 1) OVER (ORDER BY date) AS prev_day,
    LEAD(revenue, 1) OVER (ORDER BY date) AS next_day
FROM daily_sales;
```

## fa-rotate 事务

```sql
-- 基本事务
BEGIN;
UPDATE accounts SET balance = balance - 100 WHERE id = 1;
UPDATE accounts SET balance = balance + 100 WHERE id = 2;
COMMIT;

-- 回滚
BEGIN;
DELETE FROM users WHERE inactive = true;
ROLLBACK;

-- SAVEPOINT
BEGIN;
UPDATE orders SET status = 'processing' WHERE id = 1;
SAVEPOINT sp1;
UPDATE orders SET status = 'shipped' WHERE id = 1;
ROLLBACK TO sp1;
COMMIT;
```

## fa-lightbulb 常用技巧

```sql
-- CASE 表达式
SELECT name,
    CASE WHEN age < 18 THEN 'minor'
         WHEN age < 65 THEN 'adult'
         ELSE 'senior'
    END AS age_group
FROM users;

-- COALESCE 空值处理
SELECT name, COALESCE(phone, 'N/A') AS phone FROM users;

-- 类型转换
SELECT CAST(age AS CHAR);
SELECT created_at::DATE FROM users;  -- PostgreSQL

-- CONCAT 字符串拼接
SELECT CONCAT(first_name, ' ', last_name) AS full_name FROM users;

-- IFNULL / NULLIF
SELECT IFNULL(phone, 'unknown') FROM users;   -- MySQL
SELECT NULLIF(age, 0) FROM users;              -- age 为 0 时返回 NULL
```
