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

func login(resp http.ResponseWriter, req *http.Request, sessions *sessions.CookieStore) {
	RenderTemplate(resp, "login", pongo2.Context{
		"title": "Func",
	})
}

// Configures handling of static files
// NOTE: With this, all files in the public dir will be
// accessible through HTTP requests - be sure to not put anything
// private in there
func setupFileServer(router *mux.Router) {
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./site/public/")))
}

func InitializeRoutes(router *mux.Router, sessions *sessions.CookieStore) {
	router.Handle("/", HomePage(sessions)).Methods("GET")

	signupPage := web.ConfigureHandler(signup, sessions, web.Options{})
	router.Handle("/signup", signupPage).Methods("GET")

	loginPage := web.ConfigureHandler(login, sessions, web.Options{})
	router.Handle("/login", loginPage).Methods("GET")

	// Static file serving must be configured last
	setupFileServer(router)
}
