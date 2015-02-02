package repo

import (
	"fmt"
	"log"

	"github.com/atotto/clipboard"
	"github.com/boltdb/bolt"
)

type Item struct {
	Db   *bolt.DB
	Name string
}

func (i Item) Create(name, value string) error {
	err := i.Db.Update(func(tx *bolt.Tx) error {
		lists := tx.Bucket([]byte("lists"))
		list := lists.Bucket([]byte(i.Name))

		if list == nil {
			return fmt.Errorf("List %s doesn't exist", i.Name)
		}

		err := list.Put([]byte(name), []byte(value))
		if err != nil {
			return fmt.Errorf("Error creating item %s in bucket %s", name, i.Name)
		}
		return nil
	})
	return err
}

func (i Item) Get(name string) (val []byte, err error) {
	err = i.Db.View(func(tx *bolt.Tx) error {
		lists := tx.Bucket([]byte("lists"))
		list := lists.Bucket([]byte(i.Name))

		if list == nil {
			return fmt.Errorf("List %s doesn't exist", i.Name)
		}

		val = list.Get([]byte(name))

		if err := clipboard.WriteAll(string(val)); err != nil {
			return fmt.Errorf("Error copying item value %s to clipboard", val)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return
}
