package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type MyType struct {
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	var t = MyType{time.Now()}
	var b, err = json.Marshal(t)
	if err != nil {
		fmt.Printf("Error marshaling data.")
	} else {
		fmt.Printf("%s", b)
	}
}
