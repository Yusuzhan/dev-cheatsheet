# Cheatsheet MD Format

When generating cheatsheet markdown files for this project, follow these rules exactly.

## File Location

Place files in `cheatsheets/` directory:
- Default language: `cheatsheets/<name>.md`
- Chinese variant: `cheatsheets/<name>-zhs.md`

## File Structure

Every file MUST start with YAML frontmatter, then `##` sections with code blocks.

### Frontmatter (required)

```yaml
---
title: SQL
icon: fa-database
primary: "#FF6B35"
lang: sql
locale: zhs          # omit for English (default)
---
```

| Field | Required | Description |
|-------|----------|-------------|
| `title` | Yes | Page title |
| `icon` | Yes | Font Awesome icon name with `fa-` prefix. Browse at https://fontawesome.com/search?o=r&m=free |
| `primary` | Yes | Theme color as HEX string in quotes |
| `lang` | Yes | Code language for Prism.js syntax highlighting (sql, go, bash, python, rust, javascript, typescript, vim, etc.) |
| `locale` | No | Language variant code. Omit for English. Use `zhs` for Simplified Chinese, `zht` for Traditional Chinese, `ja` for Japanese, etc. |

### Sections

Each section is a `##` heading with an FA icon prefix and title:

```markdown
## fa-table Basic Queries

​```sql
SELECT * FROM users;
​```
```

Rules:
- Heading MUST be exactly `##` (not `###` or `#`)
- First word starting with `fa-` is the icon, rest is the title
- Use FA Solid icons only: https://fontawesome.com/search?o=r&m=free
- Code blocks use triple backticks with language identifier
- Optionally add plain text between code blocks for descriptions
- Use `-- ` for SQL comments, `// ` for Go/Rust, `# ` for bash
- Keep code examples concise — real snippets, not explanations
- Aim for 8-15 sections per cheatsheet

### Complete Example

```markdown
---
title: SQL
icon: fa-database
primary: "#FF6B35"
lang: sql
---

## fa-table Basic Queries

​```sql
SELECT * FROM users;
SELECT name, email FROM users;
SELECT DISTINCT city FROM users;
​```

## fa-filter Filtering

​```sql
SELECT * FROM users WHERE age > 18;
SELECT * FROM users WHERE name LIKE 'A%';
​```

Optional description text between code blocks.

​```sql
SELECT * FROM users WHERE city IN ('Beijing', 'Shanghai');
​```
```

## Multi-language Pairs

When creating a cheatsheet with multiple languages, create two files:
- `cheatsheets/<name>.md` — English (no `locale` field)
- `cheatsheets/<name>-zhs.md` — Chinese (`locale: zhs`)

Both files MUST have the same sections in the same order, with matching FA icons.

## Suggested Primary Colors

| Topic | Color |
|-------|-------|
| SQL | `#FF6B35` |
| Go | `#00ADD8` |
| Python | `#3776AB` |
| Rust | `#CE422B` |
| Vim | `#019733` |
| Linux/Bash | `#FCC624` |
| JavaScript | `#F7DF1E` |
| TypeScript | `#3178C6` |
| Git | `#F05032` |
| Docker | `#2496ED` |
| Kubernetes | `#326CE5` |
| fd | `#5A4FCF` |
