package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"github.com/Cod3ddy/snippet-box/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     models.Snippet   `json:"snippet,omitempty"`
	Snippets    []models.Snippet `json:"snippets,omitempty"`
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache
	cache := map[string]*template.Template{}

	// Use the filepath.Glob() function to get a slice of all filepaths that
	// match the pattern "./ui/html/pages/*.tmpl". This will essentially gives
	// us a slice of all the filepaths for our application 'page' templates
	// like: [ui/html/pages/home.tmpl ui/html/pages/view.tmpl]

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	// Loop through the page Filepaths one-by-one.
	for _, page := range pages {
		// Extract the file name (like 'home.html or tmpl') from the full filepath
		//and assign it to the name variable

		name := filepath.Base(page)

		//Parse the base template file into a template set.
		ts, err := template.ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		// Call ParseGlob() *on this template set* to add any partials.

		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		// Call Parse files *on this template set* to add the page template.

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		//Add the template set to the map, using the name of the page
		//(like 'home.html') as the key.

		cache[name] = ts
	}

	return cache, nil
}

func (app *application) newTemplateData(r *http.Request) templateData {
	return templateData{
		CurrentYear: time.Now().Year(),
	}
}
