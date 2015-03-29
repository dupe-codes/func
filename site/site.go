// Logic for the functionality of the main website (routes, generating views,
// etc.) should go here

package site

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type Page struct {
	Title string
}

// Test handler
func AnotherPage(sessionStore *sessions.CookieStore) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		page := &Page{Title: "Another page"}
		RenderTemplate(resp, "test", page)
	})
}

// Another test handler
func SubdirPage(sessionStore *sessions.CookieStore) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		page := &Page{Title: "subdir"}
		RenderTemplate(resp, "subdir", page)
	})
}

func HomePage(sessionStore *sessions.CookieStore) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Printf("Serving up main page.\n\n\n")
		page := &Page{Title: "Func"}
		RenderTemplate(res, "home", page)
	})
}

func InitializeRoutes(router *mux.Router, sessionStore *sessions.CookieStore) {
	router.Handle("/", HomePage(sessionStore)).Methods("GET")
	router.Handle("/test", AnotherPage(sessionStore)).Methods("GET")
	router.Handle("/test2", SubdirPage(sessionStore)).Methods("GET")
}
