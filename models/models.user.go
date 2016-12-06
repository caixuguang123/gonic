package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"github.com/gonic/db"
	"gopkg.in/mgo.v2"
	"fmt"
)

type User struct {
	ID       bson.ObjectId `json:"id" bson:"id"`
	UserName string `json:"user_name" bson:"userName"`
	Password string `json:"password" bson:"passWord"`
}

func RegisterUser(userName string, password string) (*User, error) {
	if !IsUserNameAvailable(userName)  {
		return nil, errors.New("user dumplicate")
	}
	if password == "" {
		return nil, errors.New("pwd can't be empty")
	}
	id := bson.NewObjectId();
	fmt.Print(id.Hex())
	var user = User{id, userName, password}
	db.Add("user", user)
	return &user, nil
}

func IsUserNameAvailable(userName string) bool {
	var user []User
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"userName":userName}).All(&user)
	}

	error := db.WitchCollection("user", query)
	if error != nil {
		return false
	}
	if user != nil {
		return false
	}
	return true
}

func IsUserValid(userName, password string) bool {
	var user[]User
	query := func(c *mgo.Collection) error{
		return c.Find(bson.M{"userName":userName}).All(&user)
	}
	error := db.WitchCollection("user", query)
	if error != nil {
		return false
	}
	if user != nil && user[0].Password == password {
		return true
	}
	return false
}