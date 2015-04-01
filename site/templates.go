// Helpful functions and settings for working with templates are
// defined here

package site

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/flosch/pongo2"
)

var (
	templatesDir = "./site/views/"
	templates    = loadTemplates()
)

// Renders the template matching the given template name with the desired template
// data.
// tmpl must be given as just the name of the template without the .tmpl file extension
// Example:
// RenderTemplate(resp, "home", data)
func RenderTemplate(resp http.ResponseWriter, name string, tmplData pongo2.Context) {
	tmpl, ok := templates[name+".html"]
	if !ok {
		// TODO: Better way of handling template search error - probably log and panic
		http.Error(resp, "Unable to render template.", http.StatusInternalServerError)
		return
	}

	// Now try writing out the template
	err := tmpl.ExecuteWriter(tmplData, resp)
	if err != nil {
		http.Error(resp, "Error encountered writing template", http.StatusInternalServerError)
	}
}

/*
 * Helper Functions
 */

// Walks through the templates directory and loads all template files for use
func loadTemplates() map[string]*pongo2.Template {
	result := make(map[string]*pongo2.Template)
	err := filepath.Walk(templatesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".html") {
			result[filepath.Base(path)] = pongo2.Must(pongo2.FromFile(path))
		}
		return nil
	})

	// if an error was generated while loading templates, the site won't
	// function properly - so we panic and bring down the server
	if err != nil {
		panic(err)
	}
	return result
}
