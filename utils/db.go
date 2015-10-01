package utils

import "gopkg.in/mgo.v2"

const uri = "mongodb://localhost/"

// these commands dont happen concurrently, so this method will be fine
func MakeDB(dbName string) *mgo.Database {
	sess, err := mgo.Dial(uri)
	if err != nil {
		panic(err) //might as well panic
	}
	db := sess.DB(dbName)
	return db
}
