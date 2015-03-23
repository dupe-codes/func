// Tests for the users package

package users

import (
    "testing"
    "os"
)

// TODO: Add more test users, including malformed users
var testUser = User{
    Username: "user",
    Firstname: "john",
    Lastname: "doe",
    Phonenumber: "+18889991234",
}

// Handles setup/teardown of database for tests
func TestMain(m *testing.M) {
    // TODO: Add in clearing of database for tests.
    // DB should be clear before and after tests run
    os.Exit(m.Run())
}

// Tests saving a new user into the DB
func TestUserCreation(t *testing.T) {
    if err := testUser.Save(); err != nil {
        t.Error("Error encountered saving user: ", testUser.ToString())
    }

    // TODO: Check if user now in database
}
