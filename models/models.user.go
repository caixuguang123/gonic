package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"github.com/gonic/db"
	"gopkg.in/mgo.v2"
)

type User struct {
	ID       int `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func RegisterUser(userName string, password string) (*User, error) {
	if IsUserNameAvailable(userName) {
		return nil, errors.New("user dumplicate")
	}
	if password == "" {
		return nil, errors.New("pwd can't be empty")
	}
	id := bson.NewObjectId();
	var user = User{id, userName, password}
	db.Add("user", user)
	return &user, nil
}

func IsUserNameAvailable(userName string) bool {
	//return GetbyName(userName)
	var user User
	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"name":userName}).One(&user)
	}

	error := db.WitchCollection("user", query())
	if error != nil {
		return false
	}
	if user != nil {
		return false
	}
	return true
}

func IsUserValid(userName, password string) bool {
	//return GetbyUser(userName,password)
	return false
}