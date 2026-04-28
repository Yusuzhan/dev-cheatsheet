# Dev Cheatsheets

Convert Markdown files into beautiful single-page HTML cheatsheets, auto-deployed to GitHub Pages.

## Adding a Cheatsheet

Drop a `.md` file into `cheatsheets/` and push. Each file needs a YAML frontmatter:

```markdown
---
title: SQL
icon: fa-database
primary: "#FF6B35"
lang: sql
---

## fa-table Basic Queries

```sql
SELECT * FROM users;
```
```

### Frontmatter Fields

| Field | Required | Description |
|-------|----------|-------------|
| `title` | Yes | Page title |
| `icon` | Yes | Font Awesome icon name (e.g. `fa-database`) |
| `primary` | Yes | Theme color as HEX (e.g. `"#FF6B35"`) |
| `lang` | Yes | Code language for Prism.js syntax highlighting |

### Section Format

`## fa-icon-name Title` becomes a card.

```markdown
## fa-filter Filtering

```sql
SELECT * FROM users WHERE age > 18;
```

Optionally add description text between code blocks.
```

## Local Development

```bash
# Build
go build -o cheatsheetgen ./cmd/cheatsheetgen/

# Generate single
./cheatsheetgen cheatsheets/sql.md -o sql.html

# Generate all + landing page
./cheatsheetgen --all --output site

# Test
go test ./...
```

## Icons

Uses [Font Awesome 6 Free](https://fontawesome.com/search?o=r&m=free) Solid icons. Just use the `fa-` prefixed name, e.g. `fa-database`, `fa-code`, `fa-terminal`.

## Tech Stack

- Go + [goldmark](https://github.com/yuin/goldmark) (Markdown parsing)
- Font Awesome 6 (icons)
- Prism.js (syntax highlighting)
- GitHub Actions → GitHub Pages (auto deployment)
