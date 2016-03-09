package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to /\n")
	fmt.Printf("Welcome to /\n")
}

func Auth(w http.ResponseWriter, r *http.Request, n http.HandlerFunc) {
	key := r.Header.Get("X-Auth-Key")
	secret := r.Header.Get("X-Auth-Secret")
	fmt.Printf("Key: %s Secret: %s\n", key, secret)
	n(w, r)
	fmt.Println("Done!")
}

func main() {
	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(Auth))

	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	n.UseHandler(router)
	n.Run(":3000")
}
