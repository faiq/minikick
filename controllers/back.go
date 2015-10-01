package controllers

import (
	"errors"
	"github.com/faiq/minikick/models"
	"strconv"
)

func Back(givenName string, projectName string, card string, amount string) error {
	if !models.ValidateName(givenName) {
		return errors.New("Given name should be no shorter than 4 chars and no longer than 20")
	}
	cardArr := models.ParseCard(card)
	if !models.LuhnCheck(cardArr) {
		return errors.New("Looks like this card is invalid")
	}
	backAmount, err := strconv.ParseFloat(amount, 64)
	if backAmount < 0 {
		return errors.New("you entered a negative amount!")
	}
	if err != nil {
		return err
	}
	result, err := models.FindProjectByName(projectName)
	if err != nil {
		return err
	}
	if result.HasCard(cardArr) {
		return errors.New("Looks like this card is already being used")
	}
	result.UpdateCard(cardArr, backAmount)
	if err != nil {
		return err
	}
	User, err := models.FindUserByName(givenName)
	if err != nil {
		return err
	}
	err = User.AddBacking(result.Id, backAmount)
	if err != nil {
		return err
	}
	err = User.Save()
	if err != nil {
		return err
	}
	//	err = result.Save({}mgo.Database)
	if err != nil {
		return err
	}
	return nil
}
