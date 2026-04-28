package renderer

import (
	"html/template"
	"strings"
	"testing"

	"github.com/Yusuzhan/dev-cheatsheet/internal/model"
)

func TestRender_BasicCheatsheet(t *testing.T) {
	cs := &model.Cheatsheet{
		Title:   "SQL",
		Icon:    "fa-database",
		Primary: "#FF6B35",
		Lang:    "sql",
		Sections: []model.Section{
			{
				ID:      "查询",
				Icon:    "fa-magnifying-glass",
				Title:   "查询",
				Content: template.HTML("<pre><code class=\"language-sql\">SELECT * FROM users;</code></pre>\n"),
			},
		},
	}

	html, err := Render(cs)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	if !strings.Contains(html, "<title>SQL Cheatsheet</title>") {
		t.Error("HTML missing title")
	}
	if !strings.Contains(html, "fa-database") {
		t.Error("HTML missing FA icon")
	}
	if !strings.Contains(html, "#FF6B35") {
		t.Error("HTML missing primary color")
	}
	if !strings.Contains(html, "SELECT * FROM users") {
		t.Error("HTML missing code content")
	}
	if !strings.Contains(html, `id="查询"`) {
		t.Error("HTML missing section id")
	}
	if !strings.Contains(html, "prism-sql.min.js") {
		t.Error("HTML missing Prism.js language script")
	}
	if !strings.Contains(html, "fa-magnifying-glass") {
		t.Error("HTML missing section FA icon")
	}
}

func TestRender_MultipleSections(t *testing.T) {
	cs := &model.Cheatsheet{
		Title:   "Vim",
		Icon:    "fa-vim",
		Primary: "#008000",
		Lang:    "vim",
		Sections: []model.Section{
			{ID: "移动", Icon: "fa-arrows-up-down-left-right", Title: "移动", Content: template.HTML("<p>move</p>")},
			{ID: "编辑", Icon: "fa-pen-to-square", Title: "编辑", Content: template.HTML("<p>edit</p>")},
		},
	}

	html, err := Render(cs)
	if err != nil {
		t.Fatalf("Render failed: %v", err)
	}

	if strings.Count(html, `class="card"`) != 2 {
		t.Errorf("Expected 2 cards, got %d", strings.Count(html, `class="card"`))
	}
	if !strings.Contains(html, "移动") || !strings.Contains(html, "编辑") {
		t.Error("HTML missing section titles")
	}
}
