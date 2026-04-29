---
title: PostgreSQL
icon: fa-database
primary: "#4169E1"
lang: sql
locale: zhs
---

## fa-list 数据类型

```sql
CREATE TABLE data_types (
  id          SERIAL PRIMARY KEY,
  name        VARCHAR(100) NOT NULL,
  content     TEXT,
  price       NUMERIC(10, 2),
  quantity    INTEGER DEFAULT 0,
  is_active   BOOLEAN DEFAULT true,
  created_at  TIMESTAMP DEFAULT NOW(),
  birth_date  DATE,
  data_json   JSONB,
  data_array  TEXT[],
  location    POINT,
  uuid_col    UUID DEFAULT gen_random_uuid(),
  money_col   MONEY
);
```

## fa-table 表操作

```sql
CREATE TABLE users (
  id         SERIAL PRIMARY KEY,
  name       VARCHAR(100) NOT NULL,
  email      VARCHAR(255) UNIQUE,
  created_at TIMESTAMPTZ DEFAULT NOW()
);

ALTER TABLE users ADD COLUMN age INTEGER;
ALTER TABLE users DROP COLUMN age;
ALTER TABLE users ALTER COLUMN name TYPE VARCHAR(200);
ALTER TABLE users RENAME COLUMN name TO full_name;
ALTER TABLE users ADD CONSTRAINT email_check CHECK (email ~* '@');

DROP TABLE users CASCADE;
TRUNCATE TABLE users RESTART IDENTITY;
```

## fa-lock 约束

```sql
ALTER TABLE orders ADD CONSTRAINT pk_orders PRIMARY KEY (id);
ALTER TABLE orders ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
ALTER TABLE orders ADD CONSTRAINT unique_order UNIQUE (user_id, product_id);
ALTER TABLE orders ADD CONSTRAINT chk_amount CHECK (amount > 0);
ALTER TABLE orders ADD CONSTRAINT chk_status CHECK (status IN ('pending', 'shipped', 'delivered'));

ALTER TABLE orders DROP CONSTRAINT fk_user;
SELECT conname FROM pg_constraint WHERE conrelid = 'orders'::regclass;
```

## fa-magnifying-glass 索引

```sql
CREATE INDEX idx_users_email ON users(email);
CREATE UNIQUE INDEX idx_users_name ON users(name);
CREATE INDEX idx_orders_date ON orders(created_at DESC NULLS LAST);
CREATE INDEX idx_users_lower ON users(LOWER(email));

CREATE INDEX idx_docs_gin ON documents USING GIN(to_tsvector('english', content));
CREATE INDEX idx_products_brin ON products USING BRIN(price);

CREATE INDEX CONCURRENTLY idx_users_name ON users(name);

DROP INDEX idx_users_email;
REINDEX TABLE users;
SELECT * FROM pg_indexes WHERE tablename = 'users';
```

## fa-code-branch 连接查询

```sql
SELECT u.name, o.total FROM users u INNER JOIN orders o ON u.id = o.user_id;
SELECT u.name, o.total FROM users u LEFT JOIN orders o ON u.id = o.user_id;
SELECT u.name, o.total FROM users u RIGHT JOIN orders o ON u.id = o.user_id;
SELECT u.name, o.total FROM users u FULL OUTER JOIN orders o ON u.id = o.user_id;

SELECT u1.name AS employee, u2.name AS manager
FROM users u1 LEFT JOIN users u2 ON u1.manager_id = u2.id;

SELECT c.name, p.title
FROM categories c CROSS JOIN products p;

SELECT u.name, o.total
FROM users u, LATERAL (SELECT * FROM orders WHERE user_id = u.id LIMIT 3) o;
```

## fa-window-maximize 窗口函数

```sql
SELECT name, salary,
  ROW_NUMBER() OVER (ORDER BY salary DESC) AS rank,
  RANK() OVER (ORDER BY salary DESC) AS rank_val,
  DENSE_RANK() OVER (ORDER BY salary DESC) AS dense_val,
  PERCENT_RANK() OVER (ORDER BY salary DESC) AS pct
FROM employees;

SELECT name, dept, salary,
  SUM(salary) OVER (PARTITION BY dept ORDER BY salary) AS running_total,
  AVG(salary) OVER (PARTITION BY dept) AS dept_avg,
  LAG(salary, 1) OVER (ORDER BY salary) AS prev,
  LEAD(salary, 1) OVER (ORDER BY salary) AS next,
  FIRST_VALUE(salary) OVER (PARTITION BY dept ORDER BY salary) AS dept_min,
  NTILE(4) OVER (ORDER BY salary) AS quartile
FROM employees;
```

## fa-sitemap CTE 与递归查询

```sql
WITH active_users AS (
  SELECT * FROM users WHERE status = 'active'
)
SELECT au.name, o.total
FROM active_users au JOIN orders o ON au.id = o.user_id;

WITH RECURSIVE org_chart AS (
  SELECT id, name, manager_id, 1 AS level
  FROM employees WHERE manager_id IS NULL
  UNION ALL
  SELECT e.id, e.name, e.manager_id, oc.level + 1
  FROM employees e JOIN org_chart oc ON e.manager_id = oc.id
)
SELECT * FROM org_chart ORDER BY level;

WITH RECURSIVE tree AS (
  SELECT id, ARRAY[id] AS path FROM categories WHERE parent_id IS NULL
  UNION ALL
  SELECT c.id, t.path || c.id FROM categories c JOIN tree t ON c.parent_id = t.id
)
SELECT * FROM tree;
```

## fa-brackets-curly JSONB

```sql
SELECT '{"name":"Alice","age":30}'::jsonb;
SELECT data->>'name' FROM users WHERE data->'age' > '25';
SELECT data->'address'->>'city' FROM users;

UPDATE users SET data = data || '{"active": true}'::jsonb;
UPDATE users SET data = data - 'temp';
UPDATE users SET data = jsonb_set(data, '{age}', '31');

SELECT * FROM users WHERE data @> '{"role": "admin"}'::jsonb;
SELECT * FROM users WHERE data ? 'email';
SELECT * FROM users WHERE data ?| array['email', 'phone'];

SELECT jsonb_pretty(data) FROM users;
SELECT jsonb_object_keys(data) FROM users;
SELECT jsonb_array_elements(data->'tags') FROM users;
```

## fa-search 全文搜索

```sql
ALTER TABLE documents ADD COLUMN tsv tsvector
  GENERATED ALWAYS AS (to_tsvector('english', title || ' ' || content)) STORED;
CREATE INDEX idx_docs_tsv ON documents USING GIN(tsv);

SELECT * FROM documents WHERE tsv @@ to_tsquery('postgresql & index');
SELECT * FROM documents WHERE tsv @@ plainto_tsquery('postgresql index');
SELECT * FROM documents WHERE tsv @@ phraseto_tsquery('full text search');
SELECT * FROM documents WHERE tsv @@ websearch_to_tsquery('"exact phrase" -exclude');

SELECT ts_headline('english', content, websearch_to_tsquery('postgresql')),
  ts_rank(tsv, websearch_to_tsquery('postgresql')) AS rank
FROM documents WHERE tsv @@ websearch_to_tsquery('postgresql')
ORDER BY rank DESC;
```

## fa-arrows-rotate 事务与隔离级别

```sql
BEGIN;
BEGIN ISOLATION LEVEL READ COMMITTED;
BEGIN ISOLATION LEVEL REPEATABLE READ;
BEGIN ISOLATION LEVEL SERIALIZABLE;
BEGIN ISOLATION LEVEL READ ONLY;

SAVEPOINT my_savepoint;
RELEASE SAVEPOINT my_savepoint;
ROLLBACK TO SAVEPOINT my_savepoint;

COMMIT;
ROLLBACK;

SET TRANSACTION ISOLATION LEVEL SERIALIZABLE READ ONLY DEFERRABLE;
SELECT * FROM pg_locks WHERE pid = pg_backend_pid();
```

## fa-user-shield 角色与权限

```sql
CREATE ROLE app_user WITH LOGIN PASSWORD 'secret';
CREATE ROLE readonly NOINHERIT;
GRANT readonly TO app_user;

GRANT SELECT ON ALL TABLES IN SCHEMA public TO readonly;
GRANT SELECT, INSERT, UPDATE ON users TO app_user;
GRANT USAGE, SELECT ON SEQUENCE users_id_seq TO app_user;
GRANT ALL PRIVILEGES ON DATABASE mydb TO app_user;

REVOKE DELETE ON users FROM app_user;

ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO readonly;

\du
SELECT * FROM pg_roles WHERE rolname = 'app_user';
SELECT grantee, table_name, privilege_type FROM information_schema.role_table_grants;
```

## fa-download 备份与恢复

```sql
pg_dump mydb > mydb_backup.sql
pg_dump -Fc mydb > mydb_backup.dump
pg_dump -j 4 -Fd mydb -f mydb_backup_dir/
pg_dumpall --globals-only > globals.sql

pg_restore -d mydb mydb_backup.dump
pg_restore -j 4 -d mydb mydb_backup_dir/
psql mydb < mydb_backup.sql

psql -c "SELECT pg_start_backup('backup', true);"
psql -c "SELECT pg_stop_backup();"

SELECT pg_switch_wal();
SELECT pg_current_wal_lsn();
SELECT pg_walfile_name(pg_current_wal_lsn());
```

## fa-gauge-high 性能调优

```sql
EXPLAIN ANALYZE SELECT * FROM users WHERE email = 'test@example.com';
EXPLAIN (ANALYZE, BUFFERS, FORMAT JSON) SELECT * FROM orders WHERE user_id = 1;

SELECT * FROM pg_stat_activity WHERE state = 'active';
SELECT * FROM pg_stat_user_indexes WHERE idx_scan = 0;
SELECT * FROM pg_statio_user_tables;

VACUUM ANALYZE users;
VACUUM FULL users;

ALTER SYSTEM SET shared_buffers = '256MB';
ALTER SYSTEM SET effective_cache_size = '1GB';
ALTER SYSTEM SET work_mem = '16MB';
SELECT pg_reload_conf();

SELECT pg_size_pretty(pg_database_size('mydb'));
SELECT relname, pg_size_pretty(pg_total_relation_size(relid))
  FROM pg_stat_user_tables ORDER BY pg_total_relation_size(relid) DESC;
```

## fa-puzzle-piece 扩展 (pg_stat_statements, hstore)

```sql
CREATE EXTENSION IF NOT EXISTS pg_stat_statements;
SELECT query, calls, total_exec_time, mean_exec_time, rows
  FROM pg_stat_statements ORDER BY total_exec_time DESC LIMIT 10;

CREATE EXTENSION IF NOT EXISTS hstore;
SELECT 'name => Alice, age => 30'::hstore;
SELECT data -> 'name' FROM users_hstore;
SELECT * FROM users_hstore WHERE data @> 'name => Alice';
SELECT * FROM users_hstore WHERE data ? 'email';
SELECT each(data) FROM users_hstore;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
SELECT uuid_generate_v4();
CREATE EXTENSION IF NOT EXISTS pg_trgm;
SELECT * FROM users WHERE name % 'Alic';
```
