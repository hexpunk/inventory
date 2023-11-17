package htmltemplate

import "html/template"

// Returns all known template names for debugging purposes.
func Names(t *template.Template) (names []string) {
	all := t.Templates()
	names = make([]string, len(all))
	for i, t := range all {
		names[i] = t.Name()
	}

	return
}
