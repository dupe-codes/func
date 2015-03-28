// Defines utility functions for interacting with web requests,
// such as creating handler functions

package web

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type SessionedHandler func(http.ResponseWriter, *http.Request, *sessions.CookieStore)

// Configures the given function to be used as a request handler func
// This can be expanded to handle any shared logic between all handler funcs
func ConfigureHandler(fn SessionedHandler, sessionStore *sessions.CookieStore) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		fn(resp, req, sessionStore)
	}
}
