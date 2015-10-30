package main

import (
	"fmt"
	"reflect"
)

func f(E interface{}) {
	fmt.Printf("%s", reflect.TypeOf(E))
}

func main() {
	var array = make([]interface{}, 0)
	f(array)
}
