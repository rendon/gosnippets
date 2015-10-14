// Basic example using the mgo driver (https://labix.org/mgo): Insert and Find.
package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type User struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func main() {
	var session, err = mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("Error connectiong server: %s\n", err.Error())
	}
	defer session.Close()

	var db = session.DB("test")
	var col = db.C("users")
	var user = User{
		Name: "John",
		Age:  26,
	}
	if err = col.Insert(user); err != nil {
		log.Fatalf("Error inserting document: %s\n", err)
	}

	var u User
	err = col.Find(bson.M{}).One(&u)
	if err != nil {
		log.Fatalf("Error retrieving user: %s\n", err.Error())
	}

	fmt.Printf("User name: %s\n", u.Name)
	fmt.Printf("User age: %d\n", u.Age)
}
