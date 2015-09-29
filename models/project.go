package models

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"unicode/utf8"
)

type Project struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	Name         string        `bson:"name"`         //Name of Project
	TargetAmount float64       `bson:"targetAmount"` //What the target amount of the project will be
	Backers      []string      `bson:"backers"`      //List of names the have backed this project
	cards        [][]int       `bson:"cards"`        //A private list of cards to check if a card already exists for this
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

//Takes in a new card and updates mongo
func (p Project) UpdateCard(newCard []int) error {
	change := bson.M{"$push": bson.M{"cards": newCard}}
	err = c.Update(p.Id, change)
	if err != nil {
		return err
	}
}

func (p Project) HasCard(check []int) bool {
	for _, card := range p.cards {
		if compareCards(card, check) {
			return true
		}
	}
	return false
}

func compareCards(card1 []int, card2 []int) bool {
	if len(card1) != len(card2) {
		return false
	}
	for i, _ := range card1 {
		if card1[i] != card2[i] {
			return false
		}
	}
	return true
}

func Back(givenName string, projectName string, card string, amount float64) error {
	if !validateName(givenName) {
		return errors.New("Given name should be no shorter than 4 chars and no longer than 20")
	}
	cardArr := ParseCard(card)
	if !LuhnCheck(cardArr) {
		return errors.New("Looks like this card is invalid")
	}
	if amount < 0 {
		return errors.New("you entered a negative amount!")
	}
	uri := "mongodb://localhost/"
	sess, err := mgo.Dial(uri)
	defer sess.Close()
	if err != nil {
		return err
	}
	c := sess.DB("minikick").C("projects")
	var result Project
	err := collection.Find(bson.M{"name": projectName}).One(&result)
	if err == mgo.ErrNotFound {
		return errors.New("Looks like you're trying to back something that doesnt exist")
	}
	if err != nil {
		return err
	}
	if result.HasCard(cardArr) {
		return errors.New("Looks like this card is already being used")
	}
	err := result.UpdateCard(cardArr)
	if err != nil {
		return err
	}
	return nil
}
