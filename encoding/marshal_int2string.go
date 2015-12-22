package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	ID   uint64 `json:",string"`
	Name string `json:"name"`
}

func main() {
	// Marshal from integer to string
	u := User{
		ID:   12345678910111213,
		Name: "John",
	}

	buf, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", buf)

	// Unmarshal from string to integer
	buf = []byte(`
	{
		"ID": "12345678910111213",
		"name": "John"
	}
	`)
	var v User
	err = json.Unmarshal(buf, &v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User ID: %d\n", v.ID)
}
