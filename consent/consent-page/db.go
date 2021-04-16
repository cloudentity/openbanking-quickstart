package main

import (
	"os"
	"time"

	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
)

var mode = os.FileMode(0600)

func InitDB(config Config) (*bolt.DB, error) {
	var (
		db  *bolt.DB
		err error
	)

	if db, err = bolt.Open(config.DBFile, mode, &bolt.Options{Timeout: 3 * time.Second}); err != nil {
		return nil, errors.Wrapf(err, "failed to open db")
	}

	return db, nil
}
