package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id             bson.ObjectId   `bson:"_id,omitempty"`
	Name           string          `bson:"name"` //Name of User
	BackedProjects []bson.ObjectId `bson:"backedProjects"`
}

func NewUser(name string) User {
	return User{Name: name}
}

func (u User) SaveBacking(backedProject bson.ObjectId) error {
	uri := "mongodb://localhost/"
	sess, err := mgo.Dial(uri)
	defer sess.Close()
	if err != nil {
		return err
	}
	c := sess.DB("minikick").C("users")
	u.Id = bson.NewObjectId()
	u.BackedProjects = append(u.BackedProjects, backedProject)
	err = c.Insert(u)
	if err != nil {
		return err
	}
	return nil
}
