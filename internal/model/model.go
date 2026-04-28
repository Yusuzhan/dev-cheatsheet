package model

import "html/template"

type Cheatsheet struct {
	Title    string
	Icon     string
	Primary  string
	Lang     string
	Sections []Section
}

type Section struct {
	ID      string
	Icon    string
	Title   string
	Content template.HTML
}
