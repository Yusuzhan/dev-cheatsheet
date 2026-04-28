---
title: SQL
icon: fa-database
primary: "#FF6B35"
lang: sql
---

## fa-table Basic Queries

```sql
-- Query all data
SELECT * FROM users;

-- Select specific columns
SELECT name, email FROM users;

-- Limit results
SELECT * FROM users LIMIT 10;

-- Distinct values
SELECT DISTINCT city FROM users;

-- Aliases
SELECT name AS user_name, age AS user_age FROM users;
```

## fa-filter Filtering

```sql
-- WHERE clause
SELECT * FROM users WHERE age > 18;
SELECT * FROM users WHERE name = 'Alice';
SELECT * FROM users WHERE age BETWEEN 20 AND 30;

-- AND / OR / NOT
SELECT * FROM users WHERE age > 18 AND city = 'Beijing';
SELECT * FROM users WHERE city = 'Shanghai' OR city = 'Beijing';

-- IN and NOT IN
SELECT * FROM users WHERE city IN ('Beijing', 'Shanghai', 'Guangzhou');

-- LIKE pattern matching
SELECT * FROM users WHERE name LIKE 'A%';    -- starts with A
SELECT * FROM users WHERE name LIKE '%son';   -- ends with son
SELECT * FROM users WHERE email LIKE '%@gmail.com';

-- IS NULL / IS NOT NULL
SELECT * FROM users WHERE phone IS NULL;
```

## fa-sort Sorting & Pagination

```sql
-- ORDER BY
SELECT * FROM users ORDER BY age ASC;          -- ascending (default)
SELECT * FROM users ORDER BY age DESC;         -- descending
SELECT * FROM users ORDER BY city, age DESC;   -- multi-column

-- LIMIT and OFFSET pagination
SELECT * FROM users LIMIT 10;                  -- first 10 rows
SELECT * FROM users LIMIT 10 OFFSET 20;        -- page 3 (10 per page)
```

## fa-chart-bar Aggregate Functions

```sql
-- Common aggregate functions
SELECT COUNT(*) FROM users;                    -- total rows
SELECT AVG(age) FROM users;                    -- average
SELECT SUM(salary) FROM employees;             -- sum
SELECT MAX(age) FROM users;                    -- maximum
SELECT MIN(age) FROM users;                    -- minimum

-- GROUP BY
SELECT city, COUNT(*) FROM users GROUP BY city;
SELECT department, AVG(salary) FROM employees
GROUP BY department;

-- HAVING (filter groups)
SELECT city, COUNT(*) as cnt FROM users
GROUP BY city
HAVING cnt > 5;
```

## fa-link JOINs

```sql
-- INNER JOIN
SELECT u.name, o.total
FROM users u
INNER JOIN orders o ON u.id = o.user_id;

-- LEFT JOIN
SELECT u.name, o.total
FROM users u
LEFT JOIN orders o ON u.id = o.user_id;

-- RIGHT JOIN
SELECT u.name, o.total
FROM users u
RIGHT JOIN orders o ON u.id = o.user_id;

-- FULL OUTER JOIN
SELECT u.name, o.total
FROM users u
FULL OUTER JOIN orders o ON u.id = o.user_id;

-- CROSS JOIN
SELECT u.name, d.name
FROM users u CROSS JOIN departments d;

-- Self join
SELECT e.name AS employee, m.name AS manager
FROM employees e
LEFT JOIN employees m ON e.manager_id = m.id;
```

## fa-code-branch Subqueries

```sql
-- Subquery in WHERE
SELECT * FROM users
WHERE id IN (SELECT user_id FROM orders WHERE total > 1000);

-- Subquery in FROM
SELECT city, avg_age
FROM (SELECT city, AVG(age) as avg_age FROM users GROUP BY city) t
WHERE avg_age > 25;

-- EXISTS
SELECT * FROM users u
WHERE EXISTS (SELECT 1 FROM orders o WHERE o.user_id = u.id);
```

## fa-pen-to-square Data Manipulation (DML)

```sql
-- INSERT
INSERT INTO users (name, email, age) VALUES ('Alice', 'a@example.com', 25);
INSERT INTO users (name, email) VALUES ('Bob', 'b@example.com');

-- Batch insert
INSERT INTO users (name, email) VALUES
  ('Alice', 'a@example.com'),
  ('Bob', 'b@example.com'),
  ('Charlie', 'c@example.com');

-- UPDATE
UPDATE users SET age = 26 WHERE name = 'Alice';
UPDATE users SET age = age + 1 WHERE city = 'Beijing';

-- DELETE
DELETE FROM users WHERE id = 100;
DELETE FROM users WHERE last_login < '2023-01-01';

-- TRUNCATE
TRUNCATE TABLE temp_data;
```

## fa-table-columns Table Operations (DDL)

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

## fa-window-restore Window Functions

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

## fa-rotate Transactions

```sql
-- Basic transaction
BEGIN;
UPDATE accounts SET balance = balance - 100 WHERE id = 1;
UPDATE accounts SET balance = balance + 100 WHERE id = 2;
COMMIT;

-- Rollback
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

## fa-lightbulb Useful Tips

```sql
-- CASE expression
SELECT name,
    CASE WHEN age < 18 THEN 'minor'
         WHEN age < 65 THEN 'adult'
         ELSE 'senior'
    END AS age_group
FROM users;

-- COALESCE (handle nulls)
SELECT name, COALESCE(phone, 'N/A') AS phone FROM users;

-- Type casting
SELECT CAST(age AS CHAR);
SELECT created_at::DATE FROM users;  -- PostgreSQL

-- String concatenation
SELECT CONCAT(first_name, ' ', last_name) AS full_name FROM users;

-- IFNULL / NULLIF
SELECT IFNULL(phone, 'unknown') FROM users;   -- MySQL
SELECT NULLIF(age, 0) FROM users;              -- returns NULL when age = 0
```
