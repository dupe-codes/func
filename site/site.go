// Logic for the functionality of the main website (routes, generating views,
// etc.) should go here

package site

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"github.com/njdup/func/utils/web"
)

type Page struct {
	Title string
}

func HomePage(sessionStore *sessions.CookieStore) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Printf("Serving up main page.\n\n\n")
		page := &Page{Title: "Func"}
		RenderTemplate(res, "home", page)
	})
}

func signup(resp http.ResponseWriter, req *http.Request, sessions *sessions.CookieStore) {
	page := &Page{Title: "Signup"}
	RenderTemplate(resp, "signup", page)
}

func InitializeRoutes(router *mux.Router, sessionStore *sessions.CookieStore) {
	router.Handle("/", HomePage(sessionStore)).Methods("GET")
	router.Handle("/signup", web.ConfigureHandler(signup, sessionStore)).Methods("GET")
}
