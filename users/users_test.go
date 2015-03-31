// Tests for the users package

package users

import (
	"os"
	"reflect"
	"testing"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/njdup/func/db"
	"github.com/njdup/func/settings"
)

// TODO: Add more test users, including malformed users
var validUsers = []User{
	{Username: "user", Firstname: "john", Lastname: "doe", Phonenumber: "+18889991234"},
	{Username: "user2", Firstname: "jane", Lastname: "doe", Phonenumber: "+11111111111"},
	{Username: "user3", Firstname: "Johann", Lastname: "Sebastian Bach", Phonenumber: "+1800OLDDUDE"},
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
			removeUser(user)
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

// Ensures that querying for users behaves correctly
func TestFindingUsers(t *testing.T) {
	for _, user := range validUsers {
		if err := user.Save(); err != nil {
			t.Fatal("Failed to save user in the db: ", user.ToString())
		}
	}

	// Now query for each user with username
	for _, user := range validUsers {
		found, err := FindWithUsername(user.Username)
		if err != nil {
			t.Error("Error encountered querying for user ", user.ToString())
		}

		t.Logf("Found user %s when querying for user %s", found.ToString(), user.ToString())
		if reflect.DeepEqual(found, user) {
			t.Error("Wrong user found when querying for user ", user.ToString())
		}
	}

	// Now try querying with phonenumbers
	for _, user := range validUsers {
		found, err := FindWithPhonenumber(user.Phonenumber)
		if err != nil {
			t.Error("Error encountered querying for user ", user.ToString())
		}

		t.Logf("Found user %s when querying for user %s", found.ToString(), user.ToString())
		if reflect.DeepEqual(found, user) {
			t.Error("Wrong user found when querying for user ", user.ToString())
		}
	}

	// Try querying with nonexistent username and phonenumber
	if _, err := FindWithUsername("BillyBobNonExistent"); err == nil {
		t.Error("No error encountered when querying for nonexistent user")
	}
	if _, err := FindWithPhonenumber("+1800NOTHERE"); err == nil {
		t.Error("No error encountered when querying for nonexistent user")
	}

	for _, user := range validUsers {
		removeUser(user)
	}
}
