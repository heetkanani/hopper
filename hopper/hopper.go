package hopper

import (
	"fmt"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

const (
	defaultDBName = "default"
)

type M map[string]string

type Hopper struct {
	db *bbolt.DB
}
type Collection struct {
	*bbolt.Bucket
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
		var (
			err    error
			bucket *bbolt.Bucket
		)
		bucket = tx.Bucket([]byte(name))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(name))
			if err != nil {
				return err
			}
		}

		coll.Bucket = bucket
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &coll, nil
}

func (h *Hopper) Insert(colName string, data M) (uuid.UUID, error) {
	id := uuid.New()
	coll, err := h.CreateColledtion(colName)

	if err != nil {
		return id, err
	}
	h.db.Update(func(tx *bbolt.Tx) error {
		for k, v := range data {
			if err := coll.Put([]byte(k), []byte(v)); err != nil {
				return err
			}
		}

		if err := coll.Put([]byte("id"), []byte(id.String())); err != nil {
			return err
		}
		return nil
	})

	return id, nil

}

func (h *Hopper) Select(coll string, k string, query any) {

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
