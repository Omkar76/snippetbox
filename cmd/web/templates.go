package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/Omkar76/snippetbox/internal/models"
)

type templateData struct {
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	CurrentYear int
	Form        any
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.go.tpl")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		files := []string{
			"./ui/html/base.go.tpl",
			"./ui/html/partials/nav.go.tpl",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFiles(files...)

		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
