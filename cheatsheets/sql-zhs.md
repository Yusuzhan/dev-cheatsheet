---
title: SQL 速查表
icon: fa-database
primary: "#FF6B35"
lang: sql
locale: zhs
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

-- IS NULL / IS NOT NULL
SELECT * FROM users WHERE phone IS NULL;
```

## fa-sort 排序与分页

```sql
-- ORDER BY 排序
SELECT * FROM users ORDER BY age ASC;          -- 升序 (默认)
SELECT * FROM users ORDER BY age DESC;         -- 降序

-- LIMIT 和 OFFSET 分页
SELECT * FROM users LIMIT 10;                  -- 前 10 条
SELECT * FROM users LIMIT 10 OFFSET 20;        -- 第 3 页 (每页10条)
```

## fa-chart-bar 聚合函数

```sql
SELECT COUNT(*) FROM users;                    -- 总行数
SELECT AVG(age) FROM users;                    -- 平均值
SELECT SUM(salary) FROM employees;             -- 总和
SELECT MAX(age) FROM users;                    -- 最大值
SELECT MIN(age) FROM users;                    -- 最小值

-- GROUP BY 分组
SELECT city, COUNT(*) FROM users GROUP BY city;

-- HAVING 过滤分组
SELECT city, COUNT(*) as cnt FROM users
GROUP BY city HAVING cnt > 5;
```

## fa-link JOIN 连接

```sql
-- INNER JOIN (内连接)
SELECT u.name, o.total FROM users u
INNER JOIN orders o ON u.id = o.user_id;

-- LEFT JOIN (左连接)
SELECT u.name, o.total FROM users u
LEFT JOIN orders o ON u.id = o.user_id;

-- 自连接
SELECT e.name AS employee, m.name AS manager
FROM employees e
LEFT JOIN employees m ON e.manager_id = m.id;
```

## fa-code-branch 子查询

```sql
SELECT * FROM users
WHERE id IN (SELECT user_id FROM orders WHERE total > 1000);

SELECT city, avg_age
FROM (SELECT city, AVG(age) as avg_age FROM users GROUP BY city) t
WHERE avg_age > 25;

SELECT * FROM users u
WHERE EXISTS (SELECT 1 FROM orders o WHERE o.user_id = u.id);
```

## fa-pen-to-square 数据操作

```sql
-- INSERT 插入
INSERT INTO users (name, email, age) VALUES ('Alice', 'a@example.com', 25);

-- 批量插入
INSERT INTO users (name, email) VALUES
  ('Alice', 'a@example.com'), ('Bob', 'b@example.com');

-- UPDATE 更新
UPDATE users SET age = 26 WHERE name = 'Alice';

-- DELETE 删除
DELETE FROM users WHERE id = 100;

-- TRUNCATE 清空表
TRUNCATE TABLE temp_data;
```

## fa-table-columns 表操作

```sql
CREATE TABLE users (
    id         INT PRIMARY KEY AUTO_INCREMENT,
    name       VARCHAR(100) NOT NULL,
    email      VARCHAR(255) UNIQUE,
    age        INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE users ADD COLUMN phone VARCHAR(20);
ALTER TABLE users DROP COLUMN phone;
DROP TABLE IF EXISTS temp_data;
CREATE INDEX idx_city ON users(city);
```

## fa-window-restore 窗口函数

```sql
SELECT name, age,
    ROW_NUMBER() OVER (ORDER BY age DESC) AS rank
FROM users;

SELECT name, department, salary,
    RANK() OVER (PARTITION BY department ORDER BY salary DESC) AS dept_rank
FROM employees;

SELECT date, revenue,
    LAG(revenue, 1) OVER (ORDER BY date) AS prev_day
FROM daily_sales;
```

## fa-rotate 事务

```sql
BEGIN;
UPDATE accounts SET balance = balance - 100 WHERE id = 1;
UPDATE accounts SET balance = balance + 100 WHERE id = 2;
COMMIT;

BEGIN;
DELETE FROM users WHERE inactive = true;
ROLLBACK;
```

## fa-lightbulb 常用技巧

```sql
SELECT name,
    CASE WHEN age < 18 THEN '未成年'
         WHEN age < 65 THEN '成年'
         ELSE '老年'
    END AS age_group
FROM users;

SELECT name, COALESCE(phone, 'N/A') AS phone FROM users;
SELECT CAST(age AS CHAR);
SELECT CONCAT(first_name, ' ', last_name) AS full_name FROM users;
```
