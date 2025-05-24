package hopper

import (
	"fmt"

	"go.etcd.io/bbolt"
)

const (
	defaultDBName = "default"
)

type Hopper struct {
	db *bbolt.DB
}
type Collection struct {
	bucket *bbolt.Bucket
}

func New() (*Hopper, error) {
	dbname := fmt.Sprintf("%s.hopper", defaultDBName)
	db, err := bbolt.Open(dbname, 0666, nil)
	if err != nil {
		return nil, err
	}
	return &Hopper{
		db: db,
	}, nil
}

func (h *Hopper) CreateColledtion(name string) (*Collection, error) {
	coll := Collection{}
	err := h.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte(name))
		if err != nil {
			return err
		}
		coll.bucket = bucket
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &coll, nil
}

// db.Update(func(tx *bbolt.Tx) error {
// 		id := uuid.New()
// 		for k, v := range user {
// 			if err := bucket.Put([]byte(k), []byte(v)); err != nil {
// 				return err
// 			}
// 		}

// 		if err := bucket.Put([]byte("id"), []byte(id.String())); err != nil {
// 			return err
// 		}
// 		return nil
// 	})
