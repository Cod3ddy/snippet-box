package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/Cod3ddy/snippet-box/internal/models"
	"github.com/Cod3ddy/snippet-box/ui"
)

type templateData struct {
	CurrentYear     int
	Snippet         models.Snippet
	Snippets        []models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// A string-keyed map which acts as a lookup between custom temptale functions and the functions themselves

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache
	cache := map[string]*template.Template{}

	// Use the filepath.Glob() function to get a slice of all filepaths that
	// match the pattern "./ui/html/pages/*.tmpl". This will essentially gives
	// us a slice of all the filepaths for our application 'page' templates
	// like: [ui/html/pages/home.tmpl ui/html/pages/view.tmpl]

	pages, err := fs.Glob(ui.Files, "html/pages/*.html")
	if err != nil {
		return nil, err
	}

	// Loop through the page Filepaths one-by-one.
	for _, page := range pages {
		// Extract the file name (like 'home.html or tmpl') from the full filepath
		// and assign it to the name variable

		name := filepath.Base(page)

		//filepath patterns for the templates we
        // want to parse.
		patterns := []string{
			"html/base.html",
			"html/partials/*.html",
			page,
		}

		// Parse the base template file into a template set.

		// The template.FuncMap must be registered with template set before you call
		// Use ParseFS() instead of ParseFiles() to parse the template files 
        // from the ui.Files embedded filesystem.
		// to Create an empty template set, iuse the Func() method to register the
		// template.FuncMap, and then parse the file as normal
		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
