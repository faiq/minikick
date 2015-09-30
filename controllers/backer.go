package controllers

import (
	"github.com/faiq/minikick/models"
	"log"
)

func Backer(name string) error {
	User, err := models.FindUserByName(givenName)
	if err != nil {
		return err
	}
	for _, backed := range User.BackedProjects {
		p, err := models.FindProjectById(backed.Project)
		if err != nil {
			return err
		}
		log.Printf("backed %s for %d \n", p.Name, backed.Amount)
	}
	return nil
}
