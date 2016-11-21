package main

import (

	"gopkg.in/mgo.v2"

"log"


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