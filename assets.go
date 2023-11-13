package main

import (
	"embed"
	"html/template"
	"io"
	"io/fs"
	"strings"

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

//go:embed templates/*.go.html
//go:embed templates/**/*.go.html
var templateFS embed.FS

var parsedTemplates *template.Template

func parseTemplates() error {
	if parsedTemplates == nil {
		err := fs.WalkDir(templateFS, ".", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.Type().IsRegular() {
				name, _ := strings.CutPrefix(path, "templates/")

				// The following code is borrowed liberally from the source of the html/template package's parseFS function.
				// It has been altered to support using parts of the entire path to name templates.
				var tmpl *template.Template
				if parsedTemplates == nil {
					parsedTemplates = template.New(name)
				}
				if name == parsedTemplates.Name() {
					tmpl = parsedTemplates
				} else {
					tmpl = parsedTemplates.New(name)
				}

				s, err := templateFS.ReadFile(path)
				if err != nil {
					return err
				}

				_, err = tmpl.Parse(string(s))
				if err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			return err
		}

		log.Logger.Trace().Func(func(e *zerolog.Event) {
			all := parsedTemplates.Templates()
			names := make([]string, len(all))
			for i, t := range all {
				names[i] = t.Name()
			}

			e.Strs("templates", names)
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
