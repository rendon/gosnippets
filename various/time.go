package main

import (
	"fmt"
	"time"
)

func main() {
	var t = 5 * time.Minute
	fmt.Printf("%v\n", t)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%v", time.Now())
}
