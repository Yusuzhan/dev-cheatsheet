package parser

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
	"gopkg.in/yaml.v3"

	"github.com/Yusuzhan/dev-cheatsheet/internal/model"
)

var faPattern = regexp.MustCompile(`^(fa-[a-z0-9-]+)\s+(.+)$`)

type Frontmatter struct {
	Title   string `yaml:"title"`
	Icon    string `yaml:"icon"`
	Primary string `yaml:"primary"`
	Lang    string `yaml:"lang"`
}

func Parse(md []byte, lang string, primary string) (*model.Cheatsheet, error) {
	body, fm := parseFrontmatter(md)

	if fm.Lang != "" {
		lang = fm.Lang
	}
	if fm.Primary != "" {
		primary = fm.Primary
	}

	gm := goldmark.New()
	reader := text.NewReader(body)
	doc := gm.Parser().Parse(reader)

	title := fm.Title
	icon := fm.Icon
	var sections []model.Section
	var currentSection *model.Section
	var contentBuf bytes.Buffer

	for child := doc.FirstChild(); child != nil; child = child.NextSibling() {
		heading, ok := child.(*ast.Heading)
		if !ok {
			if currentSection != nil {
				renderNodeContent(body, child, &contentBuf, lang)
			}
			continue
		}

		headingText := string(heading.Text(body))

		if heading.Level == 1 {
			if title == "" {
				title = headingText
			}
			continue
		}

		if heading.Level == 2 {
			if currentSection != nil {
				currentSection.Content = template.HTML(contentBuf.String())
				sections = append(sections, *currentSection)
			}

			secIcon, secTitle := extractIcon(headingText)
			id := secTitle
			currentSection = &model.Section{
				ID:    id,
				Icon:  secIcon,
				Title: secTitle,
			}
			contentBuf.Reset()
			continue
		}
	}

	if currentSection != nil {
		currentSection.Content = template.HTML(contentBuf.String())
		sections = append(sections, *currentSection)
	}

	return &model.Cheatsheet{
		Title:    title,
		Icon:     icon,
		Primary:  primary,
		Lang:     lang,
		Sections: sections,
	}, nil
}

func parseFrontmatter(md []byte) (body []byte, fm Frontmatter) {
	text := string(md)
	if !strings.HasPrefix(text, "---\n") {
		return md, fm
	}
	end := strings.Index(text[4:], "\n---\n")
	if end < 0 {
		return md, fm
	}
	fence := text[4 : 4+end]
	yaml.Unmarshal([]byte(fence), &fm)
	return []byte(text[4+end+5:]), fm
}

func extractIcon(s string) (icon, rest string) {
	s = strings.TrimSpace(s)
	m := faPattern.FindStringSubmatch(s)
	if m != nil {
		return m[1], strings.TrimSpace(m[2])
	}
	return "", s
}

func renderNodeContent(source []byte, node ast.Node, buf *bytes.Buffer, defaultLang string) {
	switch n := node.(type) {
	case *ast.FencedCodeBlock:
		lang := defaultLang
		if langStr := string(n.Language(source)); langStr != "" {
			lang = langStr
		}
		buf.WriteString(fmt.Sprintf("<pre><code class=\"language-%s\">", lang))
		lines := n.Lines()
		for i := 0; i < lines.Len(); i++ {
			line := lines.At(i)
			buf.Write(htmlEscape(line.Value(source)))
		}
		buf.WriteString("</code></pre>\n")
	case *ast.Paragraph:
		buf.WriteString("<p>")
		buf.Write(htmlEscape(n.Text(source)))
		buf.WriteString("</p>\n")
	case *ast.List:
		buf.WriteString("<ul>\n")
		for child := n.FirstChild(); child != nil; child = child.NextSibling() {
			li, ok := child.(*ast.ListItem)
			if !ok {
				continue
			}
			buf.WriteString("<li>")
			buf.Write(htmlEscape(li.Text(source)))
			buf.WriteString("</li>\n")
		}
		buf.WriteString("</ul>\n")
	case *ast.Blockquote:
		buf.WriteString("<blockquote>")
		buf.Write(htmlEscape(n.Text(source)))
		buf.WriteString("</blockquote>\n")
	}
}

func htmlEscape(b []byte) []byte {
	s := string(b)
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	return []byte(s)
}
