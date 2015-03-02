// This will eventually define the main logic
// for the overall Serve server, which both exports
// a RESTful API for creating Serve users and user
// scripts, and handles processing incoming texts
// + running desired commands for users

package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
    "github.com/gorilla/securecookie"

    //"github.com/njdup/settings"
    //"github.com/njdup/utils"
    //"github.com/njdup/users"
    //"github.com/njdup/commands"
    //"github.com/njdup/texts"
)

// TODO: Export this to global settings package
var (
    keyLength = 16
)

// configureRoutes will eventually initialize all application and API routes
// TODO: Fill this in as components of the application are created
// All new components must export an InitializeRoutes func, which takes in
// the app router and session store as params
func configureRoutes(router *mux.Router, sessionStore *sessions.CookieStore) {
    // Example:
    // users.InitializeRoutes(router, sessionStore)
    router.HandleFunc("/", dummyFunc)
}


func dummyFunc(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Welcome to Serve!")
}

func main() {
    router := mux.NewRouter()
    // TODO: Check the security of this - using the generated random key should
    // be sufficient, but better safe than sorry
    sessionStore := sessions.NewCookieStore([]byte(securecookie.GenerateRandomKey(keyLength)))
    configureRoutes(router, sessionStore)
    //router.HandleFunc("/", dummyFunc)

    http.Handle("/", router)
    fmt.Println("Listening on port 8080")
    http.ListenAndServe(":8080", nil)
}
