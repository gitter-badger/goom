package repo

import (
	"fmt"

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
