package controllers

import (
	"fmt"
	"github.com/faiq/minikick/models"
	"gopkg.in/mgo.v2"
)

func Backer(name string, db *mgo.Database) error {
	User, err := models.FindUserByName(name, db)
	if err != nil {
		return err
	}
	for _, backed := range User.BackedProjects {
		p, err := models.FindProjectById(backed.Project, db)
		if err != nil {
			return err
		}
		fmt.Printf("backed %s for %f \n", p.Name, backed.Amount)
	}
	return nil
}
