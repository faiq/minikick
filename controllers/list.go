package controllers

import (
	"fmt"
	"github.com/faiq/minikick/models"
)

func List(name string) error {
	project, err := models.FindProjectByName(name)
	if err != nil {
		return err
	}

	users, err := models.FindUsersForProject(project.Id)
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
