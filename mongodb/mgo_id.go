// Example to understand how the `_id` field behaves.
package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Exploration struct {
	Id      bson.ObjectId `bson:"_id"`
	Network string        `bson:"network"`
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("Could not connect to server: %s", err)
	}

	var col = session.DB("test").C("test")
	var x = Exploration{
		Id:      bson.NewObjectId(),
		Network: "twitter",
	}
	err = col.Insert(x)
	if err != nil {
		log.Fatalf("Error inserting exploration: %s", err)
	}

	var r Exploration
	err = col.Find(nil).One(&r)
	if err != nil {
		log.Fatalf("Error retrieving exploration: %s", err)
	}
	fmt.Printf("%v\n", r)
}
