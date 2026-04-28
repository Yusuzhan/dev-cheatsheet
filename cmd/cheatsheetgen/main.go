package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Yusuzhan/dev-cheatsheet/internal/parser"
	"github.com/Yusuzhan/dev-cheatsheet/internal/renderer"
)

func main() {
	output := flag.String("output", "index.html", "Output HTML file path")
	primary := flag.String("primary", "#00ADD8", "Primary theme color (hex)")
	lang := flag.String("lang", "", "Code language for syntax highlighting (auto-detected if empty)")
	all := flag.Bool("all", false, "Generate all cheatsheets from directory")
	cheatsheetsDir := flag.String("cheatsheets-dir", "cheatsheets", "Directory containing cheatsheet subdirectories (used with --all)")
	flag.Parse()

	if *all {
		if err := generateAll(*cheatsheetsDir, *output); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: cheatsheetgen [flags] <input.md>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	inputPath := args[0]
	if err := generateSingle(inputPath, *output, *lang, *primary); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func generateSingle(inputPath, outputPath, langOverride, primaryOverride string) error {
	md, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("reading %s: %w", inputPath, err)
	}

	if langOverride == "" {
		langOverride = detectLang(inputPath, md)
	}

	cs, err := parser.Parse(md, langOverride, primaryOverride)
	if err != nil {
		return fmt.Errorf("parsing markdown: %w", err)
	}

	html, err := renderer.Render(cs)
	if err != nil {
		return fmt.Errorf("rendering HTML: %w", err)
	}

	if err := os.WriteFile(outputPath, []byte(html), 0644); err != nil {
		return fmt.Errorf("writing %s: %w", outputPath, err)
	}

	fmt.Printf("Generated %s from %s (%d sections)\n", outputPath, inputPath, len(cs.Sections))
	return nil
}

type mdEntry struct {
	slug string
	path string
	lang string
}

func generateAll(cheatsheetsDir, outputDir string) error {
	entries, err := os.ReadDir(cheatsheetsDir)
	if err != nil {
		return fmt.Errorf("reading cheatsheets dir: %w", err)
	}

	var mdEntries []mdEntry
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		name := strings.TrimSuffix(entry.Name(), ".md")
		mdEntries = append(mdEntries, mdEntry{
			slug: name,
			path: filepath.Join(cheatsheetsDir, entry.Name()),
			lang: "",
		})
	}

	var landingEntries []landingEntry
	for _, e := range mdEntries {
		md, err := os.ReadFile(e.path)
		if err != nil {
			return fmt.Errorf("reading %s: %w", e.path, err)
		}

		cs, err := parser.Parse(md, e.lang, "#00ADD8")
		if err != nil {
			return fmt.Errorf("parsing %s: %w", e.path, err)
		}

		html, err := renderer.Render(cs)
		if err != nil {
			return fmt.Errorf("rendering %s: %w", e.path, err)
		}

		outPath := filepath.Join(outputDir, e.slug, "index.html")
		if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(outPath, []byte(html), 0644); err != nil {
			return fmt.Errorf("writing %s: %w", outPath, err)
		}
		fmt.Printf("Generated %s (%d sections)\n", outPath, len(cs.Sections))

		landingEntries = append(landingEntries, landingEntry{
			Slug:    e.slug,
			Title:   cs.Title,
			Icon:    cs.Icon,
			Primary: cs.Primary,
			Lang:    cs.Lang,
		})
	}

	return generateLanding(outputDir, landingEntries)
}

type landingEntry struct {
	Slug    string
	Title   string
	Icon    string
	Primary string
	Lang    string
}

func generateLanding(outputDir string, entries []landingEntry) error {
	landing := buildLandingHTML(entries)
	outPath := filepath.Join(outputDir, "index.html")
	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return err
	}
	return os.WriteFile(outPath, []byte(landing), 0644)
}

func buildLandingHTML(entries []landingEntry) string {
	var sb strings.Builder
	sb.WriteString(`<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Dev Cheatsheets</title>
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700;800&display=swap" rel="stylesheet">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.1/css/all.min.css" integrity="sha512-DTOQO9RWCH3ppGqcWaEA1BIZOC6xxalwEsw9c2QQeAIftl+Vegovlnee1c9QX4TctnWMn13TZye+giMm8e2LwA==" crossorigin="anonymous" referrerpolicy="no-referrer" />
<style>
*, *::before, *::after { margin: 0; padding: 0; box-sizing: border-box; }
body {
  font-family: 'Inter', -apple-system, sans-serif;
  background: #F0F7FB;
  color: #1B2A3D;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  -webkit-font-smoothing: antialiased;
}
header {
  width: 100%;
  background: linear-gradient(160deg, #005F73 0%, #00ADD8 45%, #48CAE4 100%);
  color: #fff;
  text-align: center;
  padding: 48px 20px 36px;
}
header h1 { font-size: 2.5rem; font-weight: 800; letter-spacing: -1px; }
header h1 i { margin-right: 12px; }
header p { margin-top: 8px; opacity: 0.85; font-size: 1.05rem; }
.grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 20px;
  max-width: 1000px;
  width: 100%;
  padding: 36px 20px 72px;
}
.card {
  background: #fff;
  border-radius: 16px;
  padding: 28px;
  box-shadow: 0 1px 3px rgba(0,95,115,0.04);
  border: 1px solid rgba(0,173,216,0.08);
  text-decoration: none;
  color: inherit;
  transition: all 0.3s ease;
}
.card:hover {
  box-shadow: 0 12px 40px rgba(0,95,115,0.08);
  transform: translateY(-4px);
}
.card .icon {
  width: 48px; height: 48px;
  display: flex; align-items: center; justify-content: center;
  font-size: 1.3rem;
  border-radius: 12px;
  margin-bottom: 14px;
}
.card .title { font-size: 1.2rem; font-weight: 700; color: #005F73; }
.card .lang { font-size: 0.8rem; color: #5A7089; margin-top: 6px; }
.card .bar {
  height: 3px;
  border-radius: 2px;
  margin-top: 16px;
  background: var(--primary, #00ADD8);
}
footer { padding: 24px; color: #5A7089; font-size: 0.85rem; }
footer a { color: #00ADD8; text-decoration: none; font-weight: 500; }
</style>
</head>
<body>
<header>
  <h1><i class="fas fa-book-open"></i> Dev Cheatsheets</h1>
  <p>快速查阅各种开发技术的速查表</p>
</header>
<div class="grid">
`)
	for _, e := range entries {
		sb.WriteString(fmt.Sprintf(`<a class="card" href="%s/index.html" style="--primary: %s">
  <div class="icon" style="background: %s1A; color: %s"><i class="fas %s"></i></div>
  <div class="title">%s</div>
  <div class="lang">%s</div>
  <div class="bar"></div>
</a>
`, e.Slug, e.Primary, e.Primary, e.Primary, e.Icon, e.Title, e.Lang))
	}
	sb.WriteString(`</div>
<footer>Powered by <a href="https://github.com/Yusuzhan/dev-cheatsheet">cheatsheetgen</a></footer>
</body>
</html>`)
	return sb.String()
}

func detectLang(path string, md []byte) string {
	content := string(md)
	if strings.Contains(content, "```sql") {
		return "sql"
	}
	if strings.Contains(content, "```bash") || strings.Contains(content, "```sh") {
		return "bash"
	}
	if strings.Contains(content, "```python") {
		return "python"
	}
	if strings.Contains(content, "```go") {
		return "go"
	}
	if strings.Contains(content, "```vim") {
		return "vim"
	}
	if strings.Contains(content, "```typescript") {
		return "typescript"
	}
	if strings.Contains(content, "```javascript") {
		return "javascript"
	}
	if strings.Contains(content, "```rust") {
		return "rust"
	}
	ext := path[strings.LastIndex(path, ".")+1:]
	switch ext {
	case "sql":
		return "sql"
	case "go":
		return "go"
	case "py":
		return "python"
	case "rs":
		return "rust"
	case "ts":
		return "typescript"
	case "js":
		return "javascript"
	default:
		return "bash"
	}
}
