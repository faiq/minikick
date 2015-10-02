package controllers

import (
	"github.com/faiq/minikick/models"
	"github.com/faiq/minikick/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// check to see if backing a project works or not
// we should check the databases with queries to see if the appropriate stuff is created or not
//func TestBacking

// test to see if someone can reuse a used credit card
// should error
func TestAlreadyUsedCreditCard(t *testing.T) {
	assert := assert.New(t)
	sess, db := utils.MakeDB("minikick")
	defer sess.Close()
	proj, err := models.NewProject("testproject", "123.45") //set up a new project
	err = proj.Save(db)
	assert.Nil(err)
	err = Back("faiq", "testproject", "4581237932741116", "123.32", db) //lets back this project
	assert.Nil(err)
	err = Back("faiqus", "testproject", "4581237932741116", "123.32", db) //should fail
	assert.NotNil(err)                                                    //check to see if it fails
	db.DropDatabase()                                                     // drop the database to work on the next test
}

// if you back something that doesn't exist you should probably get an error
func TestBackingNonExistantProject(t *testing.T) {
	assert := assert.New(t)
	sess, db := utils.MakeDB("minikick")
	defer sess.Close()
	err := Back("faiqus", "testproject", "4581237932741116", "123.32", db) //should fail
	assert.NotNil(err)                                                     //check to see if it fails (should fail because test project was deleted )
	db.DropDatabase()                                                      // drop the database to work on the next test
}

//kickstarter does not allow you to back a project more than once
//we should get an error here as well

func TestBackingAlready(t *testing.T) {
	assert := assert.New(t)
	sess, db := utils.MakeDB("minikick")
	defer sess.Close()
	proj, err := models.NewProject("testproject", "123.45") //set up a new project
	err = proj.Save(db)
	assert.Nil(err)
	err = Back("faiq", "testproject", "4581237932741116", "123.32", db) //lets back this project
	err = Back("faiq", "testproject", "5119179016239088", "123.32", db) //should fail
	assert.NotNil(err)                                                  //check to see if it fails (should fail)
	db.DropDatabase()                                                   // drop the database to work on the next test
}
