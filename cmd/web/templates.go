package main

import (
	"fmt"
	"html/template"
	"path/filepath"

	"snippetbox.harshasv.net/internal/datamodels"
)

// Using type to define a new Struct type
type templateData struct {
	Snippet        *datamodels.Snippet
	LatestSnippets []*datamodels.Snippet
}

func templateCache() (map[string]*template.Template, error) {
	// new map to simulate Cache
	cache := map[string]*template.Template{}
	// returns slice of all filepaths that match the pattern
	pages, filepathErr := filepath.Glob("../../ui/html/pages/*.html")
	if filepathErr != nil {
		return nil, filepathErr
	}
	// for each pages
	for _, page := range pages {
		name := filepath.Base(page)
		// parse base template into a template set
		template, err := template.ParseFiles("../../ui/html/base.html")
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
