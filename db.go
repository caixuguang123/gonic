package main

import (
	"gopkg.in/mgo.v2"

	"log"

	"gopkg.in/mgo.v2/bson"
)

const URL = "mongodb://10.86.5.230"

type mgoCollection func(*mgo.Session) error

var (
	mgoSession *mgo.Session
	dataBase = "gonicDB"
)

func getMongoSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(URL)
		if err != nil {
			panic(err)
		}
	}
	return mgoSession.Clone()
}

func witchSession(collection string, mc mgoCollection) error {
	session := getMongoSession()
	defer session.Close()
	c := session.DB(dataBase).C(collection)
	return mc(c)
}



func store(user1 *user) {
	session, err := mgo.Dial("mongodb://10.86.5.230")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(user1)
	if err != nil {
		log.Fatal(err)
	}

}

func GetbyUser(username, password string) bool {
	var result []user
	session, err := mgo.Dial("mongodb://10.86.5.230")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("test").C("people")
	err = c.Find(bson.M{"username":username}).All(&result)
	if (result != nil && result[0].Password == password) {
		return true
	} else {
		return false
	}
}

func GetbyName(username string) bool {
	var result []user
	session, err := mgo.Dial("mongodb://10.86.5.230")
	if err != nil {
		panic(err)
	}
	defer session.Clone()

	c := session.DB("test").C("people")
	err = c.Find(bson.M{"username":username}).All(&result)
	if (result != nil) {
		return true
	} else {
		return false
	}
}

