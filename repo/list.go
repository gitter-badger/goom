package repo

import (
	"fmt"

	"github.com/boltdb/bolt"
)

type List struct {
	Db *bolt.DB
}

func (r List) Create(list string) error {
	//retrieve the data
	err := r.Db.Update(func(tx *bolt.Tx) error {
		lists, err := tx.CreateBucketIfNotExists([]byte("lists"))
		if err != nil {
			return fmt.Errorf("Error creating 'lists' bucket %v", err)
		}

		_, err = lists.CreateBucket([]byte(list))
		if err != nil {
			return fmt.Errorf("Error creating list %s bucket", list)
		}
		return nil
	})
	return err
}
