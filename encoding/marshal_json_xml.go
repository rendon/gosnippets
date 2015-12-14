package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
)

type User struct {
	Name  string `json:"name"  xml:"name"`
	Email string `json:"email" xml:"email"`
}

func main() {
	u := User{Name: "John", Email: "john@doe.com"}
	// JSON
	body, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)

	// XML
	body, err = xml.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)
}
