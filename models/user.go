package models

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id             bson.ObjectId `bson:"_id,omitempty"`
	Name           string        `bson:"name"` //Name of User
	BackedProjects []BackProject `bson:"backedProjects"`
}

type BackProject struct {
	Project bson.ObjectId `bson:"project"`
	Amount  float64       `bson:"amount"`
}

func NewUser(name string) User {
	return User{Name: name}
}

// Find a User By that name or create a new User if its not found
func FindUserByName(name string) (*User, error) {
	uri := "mongodb://localhost/"
	sess, err := mgo.Dial(uri)
	defer sess.Close()
	if err != nil {
		return &User{}, err
	}
	c := sess.DB("minikick").C("users")
	var result User
	err = c.Find(bson.M{"name": name}).One(&result)
	if err == mgo.ErrNotFound {
		//Make a New User
		u := NewUser(name)
		u.Id = bson.NewObjectId()
		return &u, nil
	}
	return &result, nil
}

func (u *User) AddBacking(backedProject bson.ObjectId, amount float64) error {
	if u.DidBack(backedProject) {
		return errors.New("you already backed this project")
	}
	u.BackedProjects = append(u.BackedProjects, BackProject{backedProject, amount})
	return nil
}

func (u User) DidBack(backedProject bson.ObjectId) bool {
	for _, backed := range u.BackedProjects {
		if backed.Project.Hex() == backedProject.Hex() {
			return true
		}
	}
	return false
}

func (u *User) Save() error {
	uri := "mongodb://localhost/"
	sess, err := mgo.Dial(uri)
	defer sess.Close()
	if err != nil {
		return err
	}
	c := sess.DB("minikick").C("users")
	if len(u.Id) == 0 {
		u.Id = bson.NewObjectId()
	}
	_, err = c.Upsert(bson.M{"_id": u.Id}, u)
	if err != nil {
		return err
	}
	return nil
}
