package controllers

import (
	"fmt"
	"github.com/faiq/minikick/models"
	"gopkg.in/mgo.v2"
)

func List(name string, db *mgo.Database) error {
	project, err := models.FindProjectByName(name, db)
	if err != nil {
		return err
	}

	users, err := models.FindUsersForProject(project.Id, db)
	if err != nil {
		return err
	}

	for _, user := range users {
		fmt.Printf("%s backed %s for %f\n", user.Name, name, user.BackedProjects[user.BackIndex(project.Id)].Amount)
	}
	remaining := project.TargetAmount - project.AmountBacked
	if remaining < 0 {
		// we went over!
		fmt.Println("%s is successful!", project.Name)
	} else {
		fmt.Println("%s needs %f more dollars to be successful!", project.Name, remaining)
	}
	return nil
}
