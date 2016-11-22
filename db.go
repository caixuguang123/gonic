package main

import (

	"gopkg.in/mgo.v2"

"log"

	"gopkg.in/mgo.v2/bson"
)



func store(user1 *user) {
	session, err := mgo.Dial("mongodb://localhost")

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

func GetbyUser(username,password string) bool {
	var result []user
	session, err := mgo.Dial("mongodb://localhost")
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

func GetbyName(username string) bool{
	var result []user
	session,err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	defer session.Clone()

	c := session.DB("test").C("people")
	err = c.Find(bson.M{"username":username}).All(&result)
	if (result != nil) {
		return  true
	} else {
		return false
	}


}
