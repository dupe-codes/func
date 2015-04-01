// Defines utility functions for interacting with web requests,
// such as creating handler functions

package web

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type Handler func(http.ResponseWriter, *http.Request, *sessions.CookieStore)

// Exports the ability to configure options for a new handler
// Available Options:
// ReqLogin: Require an active session for accessing the handler
type Options struct {
	ReqLogin bool
}

// Configures the given function to be used as a request handler func
// This can be expanded to handle any shared logic between all handler funcs
func ConfigureHandler(fn Handler, sessions *sessions.CookieStore, options Options) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		if options.ReqLogin {
			session, _ := sessions.Get(req, "func-session")
			_, ok := session.Values["user"]
			if !ok {
				http.Redirect(resp, req, "/login", http.StatusFound)
				return
			}
		}
		fn(resp, req, sessions)
	}
}
