package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func main() {
	log.SetFlags(0)
	var session, err = mgo.Dial("mongodb-server")
	if err != nil {
		log.Fatalf("Error connectiong server: %s\n", err.Error())
	}
	defer session.Close()
	var db = session.DB("test")
	var c = db.C("users")
	var u map[string]interface{}
	err = c.Find(bson.M{}).One(&u)
	if err != nil {
		log.Fatalf("Something went wrong: %s\n", err.Error())
	}
	b, err := json.Marshal(u)
	if err != nil {
		log.Fatalf("Couldn't marshal data.")
	}

	fmt.Printf("%s\n", b)
}
