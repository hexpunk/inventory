package main

import (
	"embed"
	"html/template"
	"io"
	"io/fs"

	"github.com/hexpunk/inventory/htmltemplate"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:embed node_modules/bootstrap/dist/js/bootstrap.js
//go:embed node_modules/bootstrap/dist/js/bootstrap.min.js
//go:embed node_modules/bootstrap/dist/css/bootstrap.css
//go:embed node_modules/bootstrap/dist/css/bootstrap.min.css
//go:embed node_modules/htmx.org/dist/htmx.js
//go:embed node_modules/htmx.org/dist/htmx.min.js
var nodeModules embed.FS

//go:embed templates
var templateFS embed.FS

var parsedTemplates *template.Template

func parseTemplates() error {
	if parsedTemplates == nil {
		dirfs, err := fs.Sub(templateFS, "templates")
		if err != nil {
			return err
		}

		parsedTemplates, err = htmltemplate.ParseFS(dirfs, "*.go.html", "*/*.go.html")
		if err != nil {
			return err
		}

		log.Logger.Trace().Func(func(e *zerolog.Event) {
			e.Strs("templates", htmltemplate.Names(parsedTemplates))
		}).Send()
	}

	return nil
}

func RenderTemplate(w io.Writer, templateName string) error {
	err := parseTemplates()
	if err != nil {
		return err
	}

	return parsedTemplates.ExecuteTemplate(w, templateName, nil)
}
