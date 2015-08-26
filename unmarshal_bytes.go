package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var i interface{}
	b := []byte(`{"user_id": 1, "message": "Hello world!"}`)
	json.Unmarshal(b, &i)
	fmt.Printf("%#v\n", i)
}
