// Tests for the users package

package users

import (
    "fmt"
	"os"
	"testing"
)

// TODO: Add more test users, including malformed users
var testUser = User{
	Username:    "user",
	Firstname:   "john",
	Lastname:    "doe",
	Phonenumber: "+18889991234",
}

var emptyUsername = User{
	Firstname:   "empty",
	Lastname:    "Username",
	Phonenumber: "+18889991234",
}

var emptyPhone = User{
	Username:  "MrEmptyPhone",
	Firstname: "empty",
	Lastname:  "phone",
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

	if err := emptyUsername.Save(); err == nil {
		t.Error("Error not returned for invalid user: ", emptyUsername.ToString())
	}

	if err := emptyPhone.Save(); err == nil {
		t.Error("Error not returns for invalid user: ", emptyPhone.ToString())
	}

	// TODO: Check if user now in database
}

// Test for adding and confirming passwords
func TestPasswords(t *testing.T) {
    testPasswords := []string{"password", "passwords", "123456", "supersecure"}
    for _, password := range testPasswords {
        if err := testUser.SetPassword(password); err != nil {
            t.Error("Error encountered setting password ", password)
        }

        fmt.Printf("Password hashed and stored as: %s\n", testUser.PasswordHash)
        if !testUser.PasswordsMatch(password) {
            t.Error("Error encountered confirming password ", password)
        }
    }
}
