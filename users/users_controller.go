// Logic for the RESTful management of users (creation, deletion, etc)
// is defined in here

package users

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"github.com/njdup/func/utils/web"
)

func InitializeRoutes(router *mux.Router, sessionStore *sessions.CookieStore) {
	userCreation := web.ConfigureHandler(createUser, sessionStore, web.Options{})
	router.Handle("/users", userCreation).Methods("POST")
}

// Creates a new user from received data
// Expects the data to be embedded in a form in the received http request
func createUser(resp http.ResponseWriter, req *http.Request, sessions *sessions.CookieStore) {
	req.ParseForm()
	newUser := User{
		Username:    req.FormValue("Username"),
		Firstname:   req.FormValue("Firstname"),
		Lastname:    req.FormValue("Lastname"),
		Phonenumber: req.FormValue("Phonenumber"),
	}

	phonenumber, err := parsePhonenumber(req.FormValue("Phonenumber"))
	if err != nil {
		errorMsg := &web.InvalidFieldsError{
			web.GeneralError{"The given phonenumber is invalid"},
			[]string{"Phonenumber"},
		}
		web.SendErrorResponse(resp, errorMsg, http.StatusBadRequest)
		return
	}
	newUser.Phonenumber = phonenumber

	err = newUser.SetPassword(req.FormValue("Password"))
	if err != nil {
		fmt.Printf("Error setting password")
		errorMsg := &web.InvalidFieldsError{
			web.GeneralError{"The given password is invalid"},
			[]string{"Phonenumber"},
		}
		web.SendErrorResponse(resp, errorMsg, http.StatusBadRequest)
		return
	}

	// Attempt to save, and send an error response if an error encountered
	if err = newUser.Save(); err != nil {
		web.SendErrorResponse(resp, err, http.StatusBadRequest)
	} else {
		web.SendSuccessResponse(resp, "User successfully created")
	}
}

/*
 * Helper functions for interacting with users/user data
 */

// Takes the given phone number and converts it into a standardized form
// Returns an error is the given number cannot be correctly converted
// TODO: Figure this out lol
func parsePhonenumber(phonenumber string) (string, error) {
	return phonenumber, nil
}
