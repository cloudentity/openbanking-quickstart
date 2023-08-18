package main

import (
	"os"
	"time"

	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
)

var (
	mode        = os.FileMode(0o600)
	usersBucket = []byte("users")
	dcrBucket   = []byte("dcr")
)

func InitDB(config Config) (*bolt.DB, error) {
	var (
		db  *bolt.DB
		err error
	)

	if db, err = bolt.Open(config.DBFile, mode, &bolt.Options{Timeout: 3 * time.Second}); err != nil {
		return nil, errors.Wrapf(err, "failed to open db")
	}

	if err = initBuckets(db); err != nil {
		return nil, errors.Wrapf(err, "failed to init buckets")
	}

	return db, nil
}

func initBuckets(db *bolt.DB) error {
	var err error

	for _, b := range [][]byte{usersBucket, dcrBucket} {
		bucket := b
		if err = db.Update(func(tx *bolt.Tx) error {
			if _, err = tx.CreateBucketIfNotExists(bucket); err != nil {
				return errors.Wrapf(err, "failed to create bucket: %s", string(bucket))
			}

			return nil
		}); err != nil {
			return err
		}
	}

	return nil
}
