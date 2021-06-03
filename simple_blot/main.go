package main

import (
	"fmt"
	"log"

	"github.com/xyproto/simplebolt"
)

func main() {
	// New bolt database struct
	db, err := simplebolt.New("bolt.db")
	if err != nil {
		log.Fatalf("Could not create database! %s", err)
	}
	defer db.Close()

	// Create a list named "greetings"
	bucket, err := simplebolt.NewKeyValue(db, "username")
	if err != nil {
		log.Fatalf("Could not create a list! %s", err)
	}

	err = bucket.Set("man", "zhougang")
	if err != nil {
		log.Fatalf("KevValue could not set %s", err)
	}

	v, err := bucket.Get("man")
	if err != nil {
		log.Fatalf("KevValue could not get %s", err)
	}

	fmt.Println(v)

	v2, err := bucket.Get("man2")
	if err != nil {
		log.Fatalf("KevValue could not get key man2 value %s", err)
	}

	fmt.Println(v2)
}
