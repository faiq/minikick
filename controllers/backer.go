package controllers

import (
	"fmt"
	"github.com/faiq/minikick/models"
)

func Backer(name string) error {
	User, err := models.FindUserByName(name)
	if err != nil {
		return err
	}
	for _, backed := range User.BackedProjects {
		p, err := models.FindProjectById(backed.Project)
		if err != nil {
			return err
		}
		fmt.Printf("backed %s for %f \n", p.Name, backed.Amount)
	}
	return nil
}
