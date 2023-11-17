package htmltemplate

import (
	"fmt"
	"html/template"
	"io/fs"
	"path"
)

// Mimics the ParseFS func from [html/template], but uses the full path for each template name.
func ParseFS(fs fs.FS, patterns ...string) (*template.Template, error) {
	return parseFS(nil, fs, patterns)
}

func parseFiles(t *template.Template, readFile func(string) (string, []byte, error), filenames ...string) (*template.Template, error) {
	for _, filename := range filenames {
		name, b, err := readFile(filename)
		if err != nil {
			return nil, err
		}
		s := string(b)
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err = tmpl.Parse(s)
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}

func parseFS(t *template.Template, fsys fs.FS, patterns []string) (*template.Template, error) {
	var filenames []string
	for _, pattern := range patterns {
		list, err := fs.Glob(fsys, pattern)
		if err != nil {
			return nil, err
		}
		if len(list) == 0 {
			return nil, fmt.Errorf("template: pattern matches no files: %#q", pattern)
		}
		filenames = append(filenames, list...)
	}

	return parseFiles(t, readFileFS(fsys), filenames...)
}

func readFileFS(fsys fs.FS) func(string) (string, []byte, error) {
	return func(file string) (name string, b []byte, err error) {
		// use the full file path as the template name
		name = path.Clean(file)
		b, err = fs.ReadFile(fsys, file)

		return
	}
}
