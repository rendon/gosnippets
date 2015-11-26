package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create table
	q := `CREATE TABLE users(name VARCHAR(32), score FLOAT);`
	_, err = db.Exec(q)
	if err != nil {
		log.Fatal(err)
	}

	// Insert
	statement, err := db.Prepare("INSERT INTO users(name, score) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := statement.Exec("John", 1.5)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %d\n", id)

	// Query
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var name string
		var score float64
		err = rows.Scan(&name, &score)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Name: %s\tScore: %f\n", name, score)
	}
}
