package main

import (
	"fmt"
	"html/template"
	"path/filepath"
	"time"

	"snippetbox.harshasv.net/internal/datamodels"
)

// Using type to define a new Struct type
type templateData struct {
	CurrentYear    int
	Snippet        *datamodels.Snippet
	LatestSnippets []*datamodels.Snippet
}

// custom format for the dateTime taken from database
func humanDate(time time.Time) string {
	return time.Format("02 Jan 2006 at 15:04")
}

// A map for custom template and their functions
var customFunctions = template.FuncMap{
	// reffering to the function not invoking i.e treating function as value
	"humanDate": humanDate,
}

func templateCache() (map[string]*template.Template, error) {
	// new map to simulate Cache
	cache := make(map[string]*template.Template)
	// returns slice of all filepaths that match the pattern
	pages, filepathErr := filepath.Glob("../../ui/html/pages/*.html")
	if filepathErr != nil {
		return nil, filepathErr
	}
	// for each pages
	for _, page := range pages {
		name := filepath.Base(page)
		// create a new template and the register the custom template functions and the proceed to parse the content
		template, err := template.New(name).Funcs(customFunctions).ParseFiles("../../ui/html/base.html")
		if err != nil {
			return nil, err
		}

		// parse all partials files
		// ParseGlob is equivalent of calling parseFiles with the list of files matching the pattern
		template, err = template.ParseGlob("../../ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		template, err = template.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		fmt.Println("File name", name, "and template", template)
		cache[name] = template

	}
	return cache, nil
}
