// Defines various utility functions for interacting
// with the database
// TODO: Look over this, confirm best approach

package db

import (
	"gopkg.in/mgo.v2"

	"github.com/njdup/serve/settings"
)

type queryFunc func(*mgo.Collection) error

var (
	mgoSession *mgo.Session
)

// Returns a running session with the database
// Creates a new connection if a current session doesn't
// already exist
func getDbSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(settings.Database.Url)
		if err != nil {
			panic(err) // TODO: Add better error handling
		}
	}
	return mgoSession.Clone()
}

// Executes the given query function on the desired database collection
// Returns any error encountered during execution of the queryFunc
func ExecWithCol(collection string, fn queryFunc) error {
	session := getDbSession()
	defer session.Close()
	col := session.DB(settings.Database.Name).C(collection)
	return fn(col)
}
