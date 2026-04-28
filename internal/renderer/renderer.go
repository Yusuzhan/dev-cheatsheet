package renderer

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"strconv"
	"strings"

	"github.com/Yusuzhan/dev-cheatsheet/internal/model"
)

//go:embed template.html group-template.html
var templateFS embed.FS

var funcMap = template.FuncMap{
	"add": func(a, b int) int { return a + b },
	"mul": func(a, b int) int { return a * b },
}

type templateData struct {
	model.Cheatsheet
	PrimaryDark   string
	PrimaryDeeper string
	PrimaryLight  string
	PrimaryGlow   string
}

type groupTemplateData struct {
	Title         string
	Icon          string
	Primary       string
	PrimaryDark   string
	PrimaryDeeper string
	PrimaryLight  string
	PrimaryGlow   string
	Lang          string
	DefaultLocale string
	Variants      []localeVariant
}

type localeVariant struct {
	Locale   string
	Label    string
	Cheatsheet model.Cheatsheet
}

var localeLabels = map[string]string{
	"en":   "EN",
	"zhs":  "简中",
	"zht":  "繁中",
	"ja":   "日本語",
	"ko":   "한국어",
	"de":   "DE",
	"fr":   "FR",
	"es":   "ES",
	"pt":   "PT",
	"ru":   "RU",
}

func LocaleLabel(locale string) string {
	if label, ok := localeLabels[locale]; ok {
		return label
	}
	return strings.ToUpper(locale)
}

func Render(cs *model.Cheatsheet) (string, error) {
	tmpl, err := template.New("template.html").Funcs(funcMap).ParseFS(templateFS, "template.html")
	if err != nil {
		return "", err
	}

	data := templateData{
		Cheatsheet:    *cs,
		PrimaryDark:   darken(cs.Primary, 30),
		PrimaryDeeper: darken(cs.Primary, 60),
		PrimaryLight:  lighten(cs.Primary, 0.85),
		PrimaryGlow:   cs.Primary + "1F",
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func RenderGroup(variants []*model.Cheatsheet) (string, error) {
	if len(variants) == 0 {
		return "", fmt.Errorf("no variants provided")
	}
	if len(variants) == 1 {
		return Render(variants[0])
	}

	tmpl, err := template.New("group-template.html").Funcs(funcMap).ParseFS(templateFS, "group-template.html")
	if err != nil {
		return "", err
	}

	def := variants[0]
	gdata := groupTemplateData{
		Title:         def.Title,
		Icon:          def.Icon,
		Primary:       def.Primary,
		PrimaryDark:   darken(def.Primary, 30),
		PrimaryDeeper: darken(def.Primary, 60),
		PrimaryLight:  lighten(def.Primary, 0.85),
		PrimaryGlow:   def.Primary + "1F",
		Lang:          def.Lang,
		DefaultLocale: def.Locale,
	}

	if gdata.DefaultLocale == "" {
		gdata.DefaultLocale = "en"
	}

	for _, cs := range variants {
		loc := cs.Locale
		if loc == "" {
			loc = "en"
		}
		gdata.Variants = append(gdata.Variants, localeVariant{
			Locale:     loc,
			Label:      LocaleLabel(loc),
			Cheatsheet: *cs,
		})
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, gdata); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func darken(hex string, amount int) string {
	r, g, b := parseHex(hex)
	r = clamp(r - amount)
	g = clamp(g - amount)
	b = clamp(b - amount)
	return formatHex(r, g, b)
}

func lighten(hex string, factor float64) string {
	r, g, b := parseHex(hex)
	r = clamp(int(float64(r) + (255-float64(r))*factor))
	g = clamp(int(float64(g) + (255-float64(g))*factor))
	b = clamp(int(float64(b) + (255-float64(b))*factor))
	return formatHex(r, g, b)
}

func parseHex(hex string) (r, g, b int) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) == 6 {
	ri, _ := strconv.ParseInt(hex[0:2], 16, 0)
	gi, _ := strconv.ParseInt(hex[2:4], 16, 0)
	bi, _ := strconv.ParseInt(hex[4:6], 16, 0)
	r, g, b = int(ri), int(gi), int(bi)
	}
	return
}

func formatHex(r, g, b int) string {
	return "#" + fmtHex(r) + fmtHex(g) + fmtHex(b)
}

func fmtHex(v int) string {
	s := strconv.FormatInt(int64(v), 16)
	if len(s) < 2 {
		s = "0" + s
	}
	return s
}

func clamp(v int) int {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return v
}
