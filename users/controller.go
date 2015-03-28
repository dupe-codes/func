// Logic for the RESTful management of users (creation, deletion, etc)
// is defined in here

package users

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"github.com/njdup/serve/utils/web"
)

func InitializeRoutes(router *mux.Router, sessionStore *sessions.CookieStore) {
	router.Handle("/users", web.ConfigureHandler(createUser, sessionStore)).Methods("GET")
}

func createUser(resp http.ResponseWriter, req *http.Request, sessions *sessions.CookieStore) {
	// TODO: Add check if user is logged in here
	fmt.Fprintf(resp, "Hello user!\n")
}
