package models

import (
	"errors"
	"unicode/utf8"
)

type Project struct {
	Name         string   //Name of Project
	TargetAmount float64  //What the target amount of the project will be
	Backers      []string //List of names the have backed this project
	cards        [][]int  //A private list of cards to check if a card already exists for this
}

func NewProject(projectName string, targetAmount int) (*Project, error) {
	if validateName(projectName) {
		p := Project{projectName, targetAmount}
		return &p, nil
	} else {
		return nil, errors.New("Projects should be no shorter than 4 characters but no longer than 20 characters")
	}
}

// validateName is an internal function that checks to see if a given
func validateName(projectName string) bool {
	strlen := utf8.RuneCountInString(projectName)
	if strlen > 20 || strlen < 4 {
		return false
	}
	return true
}

func (p *Project) Save(bool, error) {
	return nil, nil
}

func Back(givenName string, projectName string, card int, amount float64) error {
	if !validateName(givenName) {
		return errors.New("Given name should be no shorter than 4 chars and no longer than 20")
	}
	if !LuhnCheck(card) {
		return errors.New("Looks like this card is invalid")
	}
	if amount < 0 {
		return errors.New("you entered a negative amount!")
	}
}
