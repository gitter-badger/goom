package repo

import (
	"log"
	"os"
	"testing"

	"github.com/boltdb/bolt"
)

func TestCreateItem(t *testing.T) {
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		db.Close()
		os.Remove("test.db")
	}()

	db.Update(func(tx *bolt.Tx) error {
		lists, _ := tx.CreateBucketIfNotExists([]byte("lists"))
		lists.CreateBucket([]byte("gifs"))
		return nil
	})

	item := Item{db, "gifs"}
	err = item.Create("banana", "http://foo.com/banana.gif")

	if err != nil {
		t.Fatalf(`item.Create("banana", "http://foo.com/banana.gif") should not return an error.:%s`, err)
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("lists"))
		gifs := b.Bucket([]byte("gifs"))
		banana := gifs.Get([]byte("banana"))

		if banana == nil {
			t.Fatalf(`the item "banana" has not been found in "gifs" list.`)
		}

		if string(banana) != "http://foo.com/banana.gif" {
			t.Fatalf(`the item "banana" in "gifs" list doesn't have the right value.`)
		}

		return nil
	})

	err = item.Create("banana", "http://foo.com/banana2.gif")

	if err == nil {
		t.Fatalf(`item.Create("banana", "http://foo.com/banana2.gif") should not return an error.:%s`, err)
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("lists"))
		gifs := b.Bucket([]byte("gifs"))
		banana := gifs.Get([]byte("banana"))

		if banana == nil {
			t.Fatalf(`the item "banana" has not been found in "gifs" list.`)
		}

		if string(banana) != "http://foo.com/banana2.gif" {
			t.Fatalf(`the item "banana" in "gifs" list doesn't have the right value.`)
		}

		return nil
	})
}
