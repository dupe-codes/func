// Helpful functions and settings for working with templates are
// defined here

package site

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	templatesDir = "./site/views/"
	templates    = loadTemplates()
)

// Renders the template matching the given template name with the desired template
// data.
// tmpl must be given as just the name of the template without the .html file extension
// Example:
// RenderTemplate(resp, "home", data)
func RenderTemplate(resp http.ResponseWriter, tmpl string, tmplData interface{}) {
	err := templates.ExecuteTemplate(resp, tmpl+".html", tmplData)
	if err != nil {
		http.Error(resp, "Error encountered generating page", http.StatusInternalServerError)
	}
}

/*
 * Helper Functions
 */

// Walks through the templates directory and loads all template files for use
func loadTemplates() *template.Template {
	result := template.New("")
	err := filepath.Walk(templatesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".html") {
			_, err = result.ParseFiles(path)
			return err
		}
		return nil
	})

	// If an error was generated while loading templates, the site won't
	// function properly - so we panic and bring down the server
	if err != nil {
		panic(err)
	}
	return result
}