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
	return User{Id: bson.NewObjectId(), Name: name}
}

// Find a User By that name or create a new User if its not found
func FindUserByName(name string) (User, error) {
	uri := "mongodb://localhost/"
	sess, err := mgo.Dial(uri)
	defer sess.Close()
	if err != nil {
		return User{}, err
	}
	c := sess.DB("minikick").C("users")
	var result User
	err = c.Find(bson.M{"name": name}).One(&result)
	if err == mgo.ErrNotFound {
		//Make a New User
		return NewUser(name), nil
	}
	return result, nil
}

func (u User) SaveBacking(backedProject bson.ObjectId) error {
	uri := "mongodb://localhost/"
	sess, err := mgo.Dial(uri)
	defer sess.Close()
	if err != nil {
		return err
	}
	c := sess.DB("minikick").C("users")
	u.BackedProjects = append(u.BackedProjects, backedProject)
	err = c.Insert(u)
	if err != nil {
		return err
	}
	return nil
}
