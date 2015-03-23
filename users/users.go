// Handling of users (creation, management, etc) will
// be done in here
// Defines the user model and interactions/utilities for
// working with users

package users

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/njdup/serve/db"
)

// The User struct defines the database fields associated
// with a registered user
type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Inserted time.Time     `bson:"inserted" json"-"`

	Username     string `bson:"userName" json:"userName"`
	Firstname    string `bson:"firstName" json:"firstName"`
	Lastname     string `bson:"lastName" json:"lastName"`
	Phonenumber  string `bson:"phoneNumber" json:"phoneNumber`
	PasswordHash     string `bson:"password" json:"-"`
	PasswordSalt string `bson:"passwordSalt" json"-"`

	// TODO: Add fields for passwords and handling of passwords
	// Add usernames? @njdup: I think yes

}

var (
	CollectionName = "users" // Name of the collection in mongo
)

// Returns a string representation of the user object
func (user *User) ToString() string {
	return fmt.Sprintf(
		"User %s (%s): %s %s",
		user.Username,
		user.Phonenumber,
		user.Firstname,
		user.Lastname,
	)
}

// Inserts the receiver User into the database
// Returns an error if any are encountered, including
// validation errors
// TODO: Add validations lol
func (user *User) Save() error {
	if emptyFields := checkEmptyFields(user); len(emptyFields) != 0 {
		invalid := strings.Join(emptyFields, " ")
		return errors.New("The following fields cannot be empty: " + invalid)
	}

	insertQuery := func(col *mgo.Collection) error {
		// TODO: Add in existence check for given phone number
		user.Inserted = time.Now()
		return col.Insert(user) // Inserts the user, returning nil or an error
	}

	return db.ExecWithCol(CollectionName, insertQuery)
}

// Checks whether the required fields of a user object are set
// Returns a splice of all required fields that are empty
func checkEmptyFields(user *User) []string {
	result := make([]string, 0)

	if user.Username == "" {
		result = append(result, "Username")
	}

	if user.Phonenumber == "" {
		result = append(result, "Phonenumber")
	}

	return result
}
