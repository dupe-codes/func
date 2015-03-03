// Handling of users (creation, management, etc) will
// be done in here
// Defines the user model and interactions/utilities for
// working with users

package users

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/njdup/serve/db"
)

// The User struct defines the database fields associated
// with a registered user
type User struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Firstname   string        `bson:"firstName" json:"firstName"`
	Lastname    string        `bson:"lastName" json:"lastName"`
	Phonenumber string        `bson:"phoneNumber" json:"phoneNumber`
	// TODO: Add fields for passwords and handling of passwords
	// Add usernames? @njdup: I think yes

	Inserted time.Time `bson:"inserted" json"-"`
}

var (
	CollectionName = "users" // Name of the collection in mongo
)

// Inserts the receiver User into the database
// Returns an error if any are encountered, including
// validation errors
// TODO: Add validations lol
func (user *User) Save() error {
	insertQuery := func(col *mgo.Collection) error {
		// TODO: Add in existence check for given phone number
		user.Inserted = time.Now()
		return col.Insert(user) // Inserts the user, returning nil or an error
	}

	return db.ExecWithCol(CollectionName, insertQuery)
}
