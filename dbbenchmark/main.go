package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
)

type Edge struct {
	S int64 `bson:"s"`
	E int64 `bson:"e"`
}

func fatal(args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s\n", args...)
	os.Exit(1)
}
func fatalf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func testMongoWrite(ms *mgo.Session, n int, ch chan<- int) error {
	if ms == nil {
		return fmt.Errorf("Nil pointer: %T", ms)
	}
	defer func() { ch <- n }()

	col := ms.DB("graph").C("edges")
	for i := 0; i < n; i++ {
		edge := Edge{
			S: rand.Int63(),
			E: rand.Int63(),
		}
		if err := col.Insert(&edge); err != nil {
			return err
		}
	}
	return nil
}

func testMongoRead(ms *mgo.Session, n int) error {
	return nil
}

func testMongo(op string, n int) error {
	ms, err := mgo.Dial("mongodb-server")
	if err != nil {
		return err
	}
	defer ms.Close()
	if op == "write" {
		ch := make(chan int)
		part := n / 8
		for i := 0; i < 8; i++ {
			go testMongoWrite(ms, part, ch)
		}
		go testMongoWrite(ms, n%8, ch)
		for i := 0; i < 9; i++ {
			fmt.Printf("Done: %d\n", <-ch)
		}
	} else {
		return testMongoRead(ms, n)
	}
	return nil
}

func testPostgres(op string, n int) error {
	return nil
}

func main() {
	if len(os.Args) != 4 {
		fatalf("Usage: %s <db> <op> <n>\n", os.Args[0])
	}
	db := os.Args[1]
	op := os.Args[2]
	val, err := strconv.ParseInt(os.Args[3], 10, 32)
	if err != nil {
		fatal(err)
	}
	n := int(val)
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	if db == "mongo" {
		if err := testMongo(op, n); err != nil {
			fatalf("Failed to test mongoDB: %s\n", err)
		}
	} else if db == "postgres" {
		if err := testPostgres(op, n); err != nil {
			fatalf("Failed to test PostgreSQL: %s\n", err)
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Time: %v %fs\n", elapsed, float64(elapsed)/float64(time.Second))
}
