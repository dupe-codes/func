// Tests for the users package

package users

import (
	"os"
	"testing"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/njdup/serve/db"
	"github.com/njdup/serve/settings"
)

// TODO: Add more test users, including malformed users
var validUsers = []User{
	{Username: "user", Firstname: "john", Lastname: "doe", Phonenumber: "+18889991234"},
}

var invalidUsers = []User{
	{Firstname: "noUserName", Lastname: "doe", Phonenumber: "+18889991234"},
	{Username: "MrEmptyPhone", Firstname: "empty", Lastname: "phone"},
}

// Handles setup/teardown of database for tests
func TestMain(m *testing.M) {
	devDb := settings.Database.Name
	settings.Database.Name = "func-test-db"
	result := m.Run()
	settings.Database.Name = devDb
	os.Exit(result)
}

func removeUser(user User) error {
	return db.ExecWithCol(CollectionName, func(col *mgo.Collection) error {
		return col.Remove(bson.M{"userName": user.Username, "phoneNumber": user.Phonenumber})
	})
}

// Tests saving a new user into the DB
func TestUserCreation(t *testing.T) {
	for _, user := range validUsers {
		if err := user.Save(); err != nil {
			t.Error("Error encountered saving user: ", user.ToString())
		}

		// Adding user with the dup username/phonenumber should be impossible
		dupUsername := User{Username: user.Username, Phonenumber: "UNIQUEPHONENUMBER"}
		if err := dupUsername.Save(); err == nil {
			removeUser(dupUsername)
			t.Error("Error not encountered saving duplicate user: ", dupUsername.ToString())
		}

		dupPhone := User{Username: "UNIQUEUSERNAME", Phonenumber: user.Phonenumber}
		if err := dupPhone.Save(); err == nil {
			removeUser(dupPhone)
			t.Error("Error not encountered saving duplicate user: ", dupPhone.ToString())
		}

		//TODO: Check that user actually is in DB here
		removeUser(user)
	}

	for _, user := range invalidUsers {
		if err := user.Save(); err == nil {
			t.Error("Error not returned for invalid user: ", user.ToString())
		}
	}
}

// Test for adding and confirming passwords
func TestPasswords(t *testing.T) {
	testPasswords := []string{"password", "passwords", "123456", "supersecure"}
	testUser := validUsers[0]

	for _, password := range testPasswords {
		if err := testUser.SetPassword(password); err != nil {
			t.Error("Error encountered setting password ", password)
		}

		t.Logf("Password hashed and stored as: %s\n", testUser.PasswordHash)
		if !testUser.PasswordsMatch(password) {
			t.Error("Error encountered confirming password ", password)
		}
	}
}
