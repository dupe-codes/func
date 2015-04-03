// This will eventually define the main logic
// for the overall Func server, which both exports
// a RESTful API for creating Func users and user
// scripts, and handles processing incoming texts
// + running desired commands for users

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"

	"github.com/njdup/func/settings"
	"github.com/njdup/func/site"
	"github.com/njdup/func/users"
)

// configureRoutes will eventually initialize all application and API routes
// TODO: Fill this in as components of the application are created
// All new components must export an InitializeRoutes func, which takes in
// the app router and session store as params
func configureRoutes(router *mux.Router, sessionStore *sessions.CookieStore) {
	users.InitializeRoutes(router, sessionStore)

	// Site routes must be the last initialized
	site.InitializeRoutes(router, sessionStore)
}

func main() {
	router := mux.NewRouter()
	// TODO: Check the security of this - using the generated random key should
	// be sufficient, but better safe than sorry
	keyLen := settings.Security.SessionKeyLen
	sessionStore := sessions.NewCookieStore([]byte(securecookie.GenerateRandomKey(keyLen)))
	configureRoutes(router, sessionStore)

	http.Handle("/", router)
	fmt.Println("Listening on port " + settings.App.Port)
	http.ListenAndServe(settings.App.Port, context.ClearHandler(http.DefaultServeMux))
}
