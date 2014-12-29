package main

import (
	"log"
	"os"
	"testing"

	"github.com/boltdb/bolt"
)

func TestCreateList(t *testing.T) {
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		db.Close()
		os.Remove("test.db")
	}()

	repo := listRepo{db}
	err = repo.create("gifs")

	if err != nil {
		t.Fatalf(`repo.create("gifs") should not return an error.:%s`, err)
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("lists"))
		gifs := b.Bucket([]byte("gifs"))

		if gifs == nil {
			t.Fatalf(`"gifs" list has not been found in "lists" bucket.`)
		}

		return nil
	})

	err = repo.create("gifs")

	if err == nil {
		t.Fatalf(`repo.create("gifs") should return an error.`)
	}
}
