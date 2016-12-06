package db

import (
	"gopkg.in/mgo.v2"
)

const URL = "mongodb://10.86.5.230"

type mgoCollection func(*mgo.Collection) error

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

func WitchCollection(collection string, mc mgoCollection) error {
	session := getMongoSession()
	defer session.Close()
	c := session.DB(dataBase).C(collection)
	return mc(c)
}

func Add(collection string, obj interface{}) bool {
	query := func(c *mgo.Collection) error{
		return c.Insert(obj)
	}

	err := WitchCollection(collection, query)

	if err != nil {
		return false
	}
	return true
}

func Delete(collection string, obj interface{}) bool {
	query := func(c *mgo.Collection) error{
		return c.Remove(obj)
	}

	err := WitchCollection(collection, query)

	if err != nil {
		return false
	}
	return true
}



//func Store(user1 *User) {
//	session, err := mgo.Dial("mongodb://10.86.5.230")
//
//	// Check if connection error, is mongo running?
//	if err != nil {
//		panic(err)
//	}
//	defer session.Close()
//
//	// Optional. Switch the session to a monotonic behavior.
//	session.SetMode(mgo.Monotonic, true)
//
//	c := session.DB("test").C("people")
//	err = c.Insert(user1)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}
//
//func GetbyUser(username, password string) bool {
//	var result []User
//	session, err := mgo.Dial("mongodb://10.86.5.230")
//	if err != nil {
//		panic(err)
//	}
//	defer session.Close()
//
//	c := session.DB("test").C("people")
//	err = c.Find(bson.M{"username":username}).All(&result)
//	if (result != nil && result[0].Password == password) {
//		return true
//	} else {
//		return false
//	}
//}
//
//func GetbyName(username string) bool {
//	var result []User
//	session, err := mgo.Dial("mongodb://10.86.5.230")
//	if err != nil {
//		panic(err)
//	}
//	defer session.Clone()
//
//	c := session.DB("test").C("people")
//	err = c.Find(bson.M{"username":username}).All(&result)
//	if (result != nil) {
//		return true
//	} else {
//		return false
//	}
//}