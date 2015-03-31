// Management of user programs will be done in here
// TODO: Discuss terminology (users create programs to execute
// via text?)

package programs

import (
	"time"

	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	//"github.com/njdup/func/users"
)

// Defines the database field associated with a user's program
type Program struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"-"`
	Inserted time.Time     `bson:"inserted" json:"-"`

	Name        string   `bson:"name" json:"name"`
	Description string   `bson:"description" json:"description"`
	Recipients  []string `bson:"recipients" json:"recipients"`
	Filepath    string   `bson:"filepath" json:"filepath"` //TODO: Decide how to handle storing program code

	Creator bson.ObjectId `bson:"creator" json:"-"`
}
