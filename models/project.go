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
	Cards        [][]int       `bson:"cards"`        //A private list of cards to check if a card already exists for this
	AmountBacked float64       `bson:"amountBacked"`
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
	if ValidateName(projectName) {
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
	if len(p.Id) == 0 {
		p.Id = bson.NewObjectId()
	}
	_, err = c.Upsert(bson.M{"_id": p.Id}, p)
	if err != nil {
		return err
	}
	return nil
}

// validateName is a function that checks to see if a given name is of appropriate len
func ValidateName(projectName string) bool {
	strlen := utf8.RuneCountInString(projectName)
	if strlen > 20 || strlen < 4 {
		return false
	}
	return true
}

//Takes in a new card and backer, and updates mongo
func (p *Project) UpdateCard(newCard []int, backingAmount float64) {
	p.AmountBacked = backingAmount + p.AmountBacked
	p.Cards = append(p.Cards, newCard)
}

func (p Project) HasCard(check []int) bool {
	for _, card := range p.Cards {
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

func FindProjectByName(projectName string) (Project, error) {
	uri := "mongodb://localhost/"
	sess, err := mgo.Dial(uri)
	defer sess.Close()
	if err != nil {
		return Project{}, err
	}
	c := sess.DB("minikick").C("projects")
	var result Project
	err = c.Find(bson.M{"name": projectName}).One(&result)
	if err == mgo.ErrNotFound {
		return Project{}, errors.New("Looks like you're trying to back something that doesnt exist")
	}
	return result, nil
}

func FindProjectById(objID bson.ObjectId) (Project, error) {
	uri := "mongodb://localhost/"
	sess, err := mgo.Dial(uri)
	defer sess.Close()
	if err != nil {
		return Project{}, err
	}
	c := sess.DB("minikick").C("projects")
	var result Project
	err = c.Find(bson.M{"_id": objId}).One(&result)
	if err == mgo.ErrNotFound {
		return Project{}, errors.New("Looks like you're trying to find something that doesn't exist")
	}
	return result, nil
}
