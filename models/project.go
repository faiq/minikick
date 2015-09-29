package models

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"unicode/utf8"
)

type Project struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	Name         string        `bson:"name"`         //Name of Project
	TargetAmount float64       `bson:"targetAmount"` //What the target amount of the project will be
	Backers      []string      `bson:"backers"`      //List of names the have backed this project
	cards        [][]int       `bson:"cards"`        //A private list of cards to check if a card already exists for this
}

// A constructor returns a new project
func NewProject(projectName string, targetAmount string) (Project, error) {
	newTarg, err := strconv.ParseFloat(targetAmount, 64)
	if err != nil {
		return Project{}, err
	}
	if newTarg < 0 {
		return Project{}, errors.New("Looks like you entered a negative backing amount")
	}
	if validateName(projectName) {
		p := Project{Name: projectName, TargetAmount: newTarg}
		return p, nil
	} else {
		return Project{}, errors.New("Projects should be no shorter than 4 characters but no longer than 20 characters")
	}
}

func (p Project) Save() error {
	uri := "mongodb://localhost/"
	sess, err := mgo.Dial(uri)
	defer sess.Close()
	if err != nil {
		return err
	}
	c := sess.DB("minikick").C("projects")
	p.Id = bson.NewObjectId()
	err = c.Insert(p)
	if err != nil {
		return err
	}
	return nil
}

// validateName is an internal function that checks to see if a given
func validateName(projectName string) bool {
	strlen := utf8.RuneCountInString(projectName)
	if strlen > 20 || strlen < 4 {
		return false
	}
	return true
}

//Takes in a new card and backer, and updates mongo
func (p Project) UpdateCard(mongoSession *mgo.Session, newCard []int, backer string) error {
	sessCopy := mongoSession.Copy()
	defer sessCopy.Close()
	c := sessCopy.DB("minikick").C("projects")
	change := bson.M{"$push": bson.M{"cards": newCard}}
	err := c.Update(p.Id, change)
	if err != nil {
		return err
	}

	change = bson.M{"$push": bson.M{"backer": backer}}
	err = c.Update(p.Id, change)
	if err != nil {
		return err
	}
	return nil
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
	err = c.Find(bson.M{"name": projectName}).One(&result)
	if err == mgo.ErrNotFound {
		return errors.New("Looks like you're trying to back something that doesnt exist")
	}
	if err != nil {
		return err
	}
	if result.HasCard(cardArr) {
		return errors.New("Looks like this card is already being used")
	}
	err = result.UpdateCard(sess, cardArr, givenName)
	if err != nil {
		return err
	}
	return nil
}
