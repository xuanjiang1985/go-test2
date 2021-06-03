package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte("kes")); err != nil {
			fmt.Println("create failed", err)
			return err
		}

		b := tx.Bucket([]byte("kes"))
		err = b.Put([]byte("answer"), []byte("42"))
		return err

	}); err != nil {
		fmt.Println("update", err)
	}

	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("kes"))
		v := b.Get([]byte("answer"))
		fmt.Printf("the anser is :%s\n", v)

		return nil
	}); err != nil {
		fmt.Println("view", err)
	}

	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("kes"))
		v := b.Get([]byte("answernotexist"))
		fmt.Printf("the anser is :%s\n", v)

		return nil
	}); err != nil {
		fmt.Println("view", err)
	}
}
