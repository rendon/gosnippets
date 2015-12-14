package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func create(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE people(id SERIAL PRIMARY KEY, data json)")
	if err != nil {
		return err
	}
	return nil
}

func insert(db *sql.DB) error {
	data := `{"twitter_id": 1235, "name": "John Doe"}`
	_, err := db.Exec("INSERT INTO people(data) VALUES($1)", data)
	if err != nil {
		return err
	}
	return nil
}

func read(db *sql.DB) error {
	rows, err := db.Query(`SELECT * FROM people`)
	if err != nil {
		return err
	}

	for rows.Next() {
		var id int
		var data string
		err = rows.Scan(&id, &data)
		if err != nil {
			return err
		}
		fmt.Printf("%d: %s\n", id, data)
	}
	return nil
}

func main() {
	user := "postgres"
	pass := "mysecretpassword"
	dbname := "cache"
	sslmode := "disable"
	connstr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		user, pass, dbname, sslmode)
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}

	if err := create(db); err != nil {
		log.Fatalf("Failed to create table: %s", err)
	}

	if err := insert(db); err != nil {
		log.Fatalf("Failed to insert record: %s", err)
	}

	if err := read(db); err != nil {
		log.Fatalf("Failed to read records: %s", err)
	}
}
