---
title: SQLite
icon: fa-feather
primary: "#003B57"
lang: sql
locale: zhs
---

## fa-terminal CLI 命令

```sql
.open mydb.db
.databases
.tables
.schema users
.dump users
.headers on
.mode column
.mode csv
.output result.csv
.width 10 20 15
.timer on
.import data.csv users
.read script.sql
.quit
```

## fa-list 数据类型（类型亲和性）

```sql
CREATE TABLE type_demo (
  a TEXT,
  b NUMERIC,
  c INTEGER,
  d REAL,
  e BLOB
);

INSERT INTO type_demo VALUES ('hello', 3.14, 42, 2.718, x'DEADBEEF');

SELECT typeof(a), typeof(b), typeof(c), typeof(d), typeof(e) FROM type_demo;

SELECT CAST('42' AS INTEGER);
SELECT CAST(3.14 AS TEXT);
```

## fa-table 表操作

```sql
CREATE TABLE users (
  id       INTEGER PRIMARY KEY AUTOINCREMENT,
  name     TEXT NOT NULL,
  email    TEXT UNIQUE,
  age      INTEGER DEFAULT 0,
  created  TEXT DEFAULT (datetime('now'))
);

ALTER TABLE users ADD COLUMN city TEXT;
ALTER TABLE users RENAME COLUMN name TO full_name;
ALTER TABLE users RENAME TO customers;
DROP TABLE IF EXISTS temp_data;

CREATE TABLE IF NOT EXISTS logs AS SELECT * FROM users WHERE 0;
```

## fa-lock 约束

```sql
CREATE TABLE orders (
  id          INTEGER PRIMARY KEY,
  user_id     INTEGER NOT NULL,
  product_id  INTEGER NOT NULL,
  quantity    INTEGER CHECK (quantity > 0),
  price       REAL CHECK (price >= 0),
  total       REAL GENERATED ALWAYS AS (quantity * price) STORED,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (product_id) REFERENCES products(id) ON UPDATE CASCADE,
  UNIQUE (user_id, product_id)
);

PRAGMA foreign_keys = ON;
PRAGMA foreign_key_list(orders);
```

## fa-magnifying-glass 索引

```sql
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_name_lower ON users(LOWER(name));
CREATE UNIQUE INDEX idx_users_email_uq ON users(email);

CREATE INDEX idx_orders_composite ON orders(user_id, created_at);

DROP INDEX IF EXISTS idx_users_email;

REINDEX idx_users_email;
REINDEX;

SELECT * FROM pragma_index_list('users');
SELECT * FROM pragma_index_info('idx_users_email');

EXPLAIN QUERY PLAN SELECT * FROM users WHERE email = 'test@example.com';
```

## fa-code-branch 连接查询

```sql
SELECT u.name, o.total
FROM users u
INNER JOIN orders o ON u.id = o.user_id;

SELECT u.name, o.total
FROM users u
LEFT JOIN orders o ON u.id = o.user_id;

SELECT u.name, o.total
FROM users u
LEFT JOIN orders o ON u.id = o.user_id
WHERE o.id IS NULL;

SELECT a.name, b.name AS colleague
FROM users a, users b
WHERE a.dept = b.dept AND a.id < b.id;

SELECT * FROM users
WHERE id IN (SELECT user_id FROM orders WHERE total > 100);
```

## fa-layer-group 子查询

```sql
SELECT name, salary FROM employees
WHERE salary > (SELECT AVG(salary) FROM employees);

SELECT dept, MAX(salary) AS max_sal FROM employees
GROUP BY dept
HAVING MAX(salary) > (SELECT AVG(salary) FROM employees);

SELECT e.name, e.salary, e.dept
FROM employees e
INNER JOIN (
  SELECT dept, MAX(salary) AS max_sal FROM employees GROUP BY dept
) m ON e.dept = m.dept AND e.salary = m.max_sal;

UPDATE users SET status = 'premium'
WHERE id IN (SELECT user_id FROM orders GROUP BY user_id HAVING SUM(total) > 1000);
```

## fa-arrow-up-up-across Upsert

```sql
INSERT INTO users (id, name, email) VALUES (1, 'Alice', 'alice@example.com')
  ON CONFLICT (id) DO UPDATE SET name = excluded.name, email = excluded.email;

INSERT INTO counters (key, value) VALUES ('hits', 1)
  ON CONFLICT (key) DO UPDATE SET value = value + 1;

INSERT INTO users (id, name, email) VALUES (1, 'Alice', 'new@example.com')
  ON CONFLICT (email) DO NOTHING;

INSERT OR REPLACE INTO config (key, value) VALUES ('theme', 'dark');
INSERT OR IGNORE INTO users (email) VALUES ('alice@example.com');
```

## fa-window-maximize 窗口函数

```sql
SELECT name, salary,
  ROW_NUMBER() OVER (ORDER BY salary DESC) AS row_num,
  RANK() OVER (ORDER BY salary DESC) AS rank_val,
  DENSE_RANK() OVER (ORDER BY salary DESC) AS dense_rank_val
FROM employees;

SELECT name, dept, salary,
  SUM(salary) OVER (PARTITION BY dept ORDER BY salary) AS running_total,
  AVG(salary) OVER (PARTITION BY dept) AS dept_avg,
  LAG(salary, 1) OVER (ORDER BY salary) AS prev_salary,
  LEAD(salary, 1) OVER (ORDER BY salary) AS next_salary,
  FIRST_VALUE(salary) OVER (PARTITION BY dept ORDER BY salary) AS dept_min
FROM employees;

SELECT name, salary,
  NTILE(4) OVER (ORDER BY salary) AS quartile,
  PERCENT_RANK() OVER (ORDER BY salary) AS pct_rank
FROM employees;
```

## fa-sitemap CTE

```sql
WITH high_spenders AS (
  SELECT user_id, SUM(total) AS total_spent
  FROM orders GROUP BY user_id HAVING SUM(total) > 500
)
SELECT u.name, hs.total_spent
FROM users u JOIN high_spenders hs ON u.id = hs.user_id;

WITH RECURSIVE hierarchy AS (
  SELECT id, name, manager_id, 1 AS level FROM employees WHERE manager_id IS NULL
  UNION ALL
  SELECT e.id, e.name, e.manager_id, h.level + 1
  FROM employees e JOIN hierarchy h ON e.manager_id = h.id
)
SELECT id, name, level FROM hierarchy ORDER BY level;

WITH RECURSIVE nums AS (
  SELECT 1 AS n UNION ALL SELECT n + 1 FROM nums WHERE n < 100
)
SELECT n FROM nums;
```

## fa-brackets-curly JSON1 扩展

```sql
SELECT json('{"name":"Alice","age":30}');
SELECT json_extract('{"a":1,"b":2}', '$.a');
SELECT json_extract(data, '$.address.city') FROM users;

INSERT INTO logs (data) VALUES (json('{"level":"info","msg":"started"}'));

SELECT * FROM users WHERE json_extract(data, '$.age') > 25;
SELECT json_array_length('[1,2,3]');
SELECT json_type('{"a":1}', '$.a');

SELECT json_group_array(name) FROM users;
SELECT json_group_object(key, value) FROM config;

UPDATE users SET data = json_set(data, '$.age', 31);
UPDATE users SET data = json_remove(data, '$.temp');
```

## fa-search FTS5 全文搜索

```sql
CREATE VIRTUAL TABLE docs USING fts5(title, content);

INSERT INTO docs VALUES ('SQLite Guide', 'SQLite is a lightweight database engine');
INSERT INTO docs VALUES ('PostgreSQL Guide', 'PostgreSQL is an advanced database');

SELECT * FROM docs WHERE docs MATCH 'sqlite';
SELECT * FROM docs WHERE docs MATCH 'database AND engine';
SELECT * FROM docs WHERE docs MATCH '"database engine"';

SELECT *, rank FROM docs WHERE docs MATCH 'sqlite' ORDER BY rank;

SELECT *, snippet(docs, 0, '>>>', '<<<', '...', 10) FROM docs WHERE docs MATCH 'database';

CREATE VIRTUAL TABLE docs USING fts5(title, content, tokenize='porter unicode61');
```

## fa-bolt 触发器

```sql
CREATE TRIGGER update_timestamp
AFTER UPDATE ON users
FOR EACH ROW
BEGIN
  UPDATE users SET updated_at = datetime('now') WHERE id = NEW.id;
END;

CREATE TRIGGER log_delete
AFTER DELETE ON users
BEGIN
  INSERT INTO users_audit (user_id, name, deleted_at)
  VALUES (OLD.id, OLD.name, datetime('now'));
END;

CREATE TRIGGER validate_email
BEFORE INSERT ON users
FOR EACH ROW
BEGIN
  SELECT CASE
    WHEN NEW.email NOT LIKE '%@%' THEN
      RAISE(ABORT, 'Invalid email address')
  END;
END;

DROP TRIGGER IF EXISTS update_timestamp;
SELECT * FROM sqlite_master WHERE type = 'trigger';
```

## fa-sliders Pragma 配置

```sql
PRAGMA journal_mode = WAL;
PRAGMA journal_mode;
PRAGMA synchronous = NORMAL;
PRAGMA busy_timeout = 5000;

PRAGMA foreign_keys = ON;
PRAGMA foreign_key_check;

PRAGMA cache_size = -64000;
PRAGMA temp_store = MEMORY;
PRAGMA mmap_size = 268435456;

PRAGMA table_info(users);
PRAGMA table_xinfo(users);
PRAGMA index_list(users);
PRAGMA database_list;

PRAGMA integrity_check;
PRAGMA wal_checkpoint(TRUNCATE);
PRAGMA page_count;
PRAGMA page_size;
```
