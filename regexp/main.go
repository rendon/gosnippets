package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	fmt.Printf("%s\n", re.ReplaceAllString("my desc_ñdsf con id 03 ", ""))
}
