// Logic for the functionality of the main website (routes, generating views,
// etc.) should go here

package site

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type Page struct {
	Title string
}

func HomePage(sessionStore *sessions.CookieStore) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Printf("Serving up main page.\n\n\n")
		page := &Page{Title: "Func"}
		t, err := template.ParseFiles("./site/views/home.html") // TODO: Add path to templates as a project setting
		if err != nil {
			fmt.Fprintf(res, "Error!")
			return
		}
		t.Execute(res, page)
	})
}

func InitializeRoutes(router *mux.Router, sessionStore *sessions.CookieStore) {
	router.Handle("/", HomePage(sessionStore)).Methods("GET")
}
