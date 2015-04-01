// Logic for the functionality of the main website (routes, generating views,
// etc.) should go here

package site

import (
	"fmt"
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"github.com/njdup/func/utils/web"
)

func HomePage(sessionStore *sessions.CookieStore) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Printf("Serving up main page.\n\n\n")
		RenderTemplate(res, "home", pongo2.Context{
			"title": "Func",
		})
	})
}

func signup(resp http.ResponseWriter, req *http.Request, sessions *sessions.CookieStore) {
	RenderTemplate(resp, "signup", pongo2.Context{
		"title": "Func",
	})
}

func InitializeRoutes(router *mux.Router, sessionStore *sessions.CookieStore) {
	router.Handle("/", HomePage(sessionStore)).Methods("GET")
	router.Handle("/signup", web.ConfigureHandler(signup, sessionStore)).Methods("GET")
}
