package parser

import (
	"testing"
)

func TestParse_SingleSection(t *testing.T) {
	input := "---\ntitle: Go Cheatsheet\nicon: fa-golang\nprimary: \"#00ADD8\"\nlang: go\n---\n\n## fa-box 基础语法\n\n```go\npackage main\n\nfunc main() {\n    fmt.Println(\"Hello\")\n}\n```\n"

	cs, err := Parse([]byte(input), "go", "#00ADD8")
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if cs.Title != "Go Cheatsheet" {
		t.Errorf("Title = %q, want %q", cs.Title, "Go Cheatsheet")
	}
	if cs.Icon != "fa-golang" {
		t.Errorf("Icon = %q, want %q", cs.Icon, "fa-golang")
	}
	if cs.Primary != "#00ADD8" {
		t.Errorf("Primary = %q, want %q", cs.Primary, "#00ADD8")
	}
	if cs.Lang != "go" {
		t.Errorf("Lang = %q, want %q", cs.Lang, "go")
	}
	if len(cs.Sections) != 1 {
		t.Fatalf("Sections count = %d, want 1", len(cs.Sections))
	}
	s := cs.Sections[0]
	if s.Icon != "fa-box" {
		t.Errorf("Section Icon = %q, want %q", s.Icon, "fa-box")
	}
	if s.Title != "基础语法" {
		t.Errorf("Section Title = %q, want %q", s.Title, "基础语法")
	}
	if s.ID != "基础语法" {
		t.Errorf("Section ID = %q, want %q", s.ID, "基础语法")
	}
}

func TestParse_MultipleSections(t *testing.T) {
	input := "---\ntitle: Go\nicon: fa-golang\n---\n\n## fa-box 基础语法\n\n```go\nfmt.Println(\"hello\")\n```\n\n## fa-font 变量\n\nSome text here.\n\n```go\nx := 42\n```\n"

	cs, err := Parse([]byte(input), "go", "#00ADD8")
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(cs.Sections) != 2 {
		t.Fatalf("Sections count = %d, want 2", len(cs.Sections))
	}
	if cs.Sections[0].Title != "基础语法" {
		t.Errorf("Section[0] Title = %q, want %q", cs.Sections[0].Title, "基础语法")
	}
	if cs.Sections[1].Title != "变量" {
		t.Errorf("Section[1] Title = %q, want %q", cs.Sections[1].Title, "变量")
	}
	if cs.Sections[1].Icon != "fa-font" {
		t.Errorf("Section[1] Icon = %q, want %q", cs.Sections[1].Icon, "fa-font")
	}
}

func TestParse_MixedContentInSection(t *testing.T) {
	input := "---\ntitle: SQL\nicon: fa-database\nprimary: \"#FF6B35\"\nlang: sql\n---\n\n## fa-magnifying-glass 查询\n\n```sql\nSELECT * FROM users;\n```\n\nThis is a description paragraph.\n\n```sql\nSELECT name FROM users WHERE age > 18;\n```\n"

	cs, err := Parse([]byte(input), "sql", "#FF6B35")
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}

	if len(cs.Sections) != 1 {
		t.Fatalf("Sections count = %d, want 1", len(cs.Sections))
	}
	if len(cs.Sections[0].Content) == 0 {
		t.Error("Section Content is empty")
	}
}

func TestExtractIcon(t *testing.T) {
	tests := []struct {
		input string
		icon  string
		rest  string
	}{
		{"fa-box 基础语法", "fa-box", "基础语法"},
		{"fa-database SQL查询", "fa-database", "SQL查询"},
		{"NoIcon", "", "NoIcon"},
	}
	for _, tt := range tests {
		icon, rest := extractIcon(tt.input)
		if icon != tt.icon {
			t.Errorf("extractIcon(%q) icon = %q, want %q", tt.input, icon, tt.icon)
		}
		if rest != tt.rest {
			t.Errorf("extractIcon(%q) rest = %q, want %q", tt.input, rest, tt.rest)
		}
	}
}

func TestParse_FrontmatterOverrides(t *testing.T) {
	input := "---\ntitle: My SQL\nicon: fa-database\nprimary: \"#FF6B35\"\nlang: sql\n---\n\n## fa-table 查询\n\n```sql\nSELECT 1;\n```\n"

	cs, err := Parse([]byte(input), "bash", "#000000")
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	if cs.Lang != "sql" {
		t.Errorf("Lang = %q, want %q (frontmatter override)", cs.Lang, "sql")
	}
	if cs.Primary != "#FF6B35" {
		t.Errorf("Primary = %q, want %q (frontmatter override)", cs.Primary, "#FF6B35")
	}
}

func TestParse_NoFrontmatter(t *testing.T) {
	input := "# Plain Title\n\n## fa-code Section\n\n```go\nhello\n```\n"

	cs, err := Parse([]byte(input), "go", "#00ADD8")
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	if cs.Title != "Plain Title" {
		t.Errorf("Title = %q, want %q", cs.Title, "Plain Title")
	}
	if cs.Icon != "" {
		t.Errorf("Icon = %q, want empty", cs.Icon)
	}
}
