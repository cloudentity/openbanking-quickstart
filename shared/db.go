package shared

import (
	"os"
	"time"

	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
)

type DBOptions struct {
	Buckets [][]byte
}

type DBOption func(*DBOptions)

func WithBuckets(buckets ...[]byte) DBOption {
	return func(o *DBOptions) {
		o.Buckets = buckets
	}
}

var mode = os.FileMode(0o600)

func InitDB(dbFilePath string, options ...DBOption) (*bolt.DB, error) {
	var (
		db   *bolt.DB
		err  error
		opts = DBOptions{}
	)

	for _, o := range options {
		o(&opts)
	}

	if db, err = bolt.Open(dbFilePath, mode, &bolt.Options{Timeout: 3 * time.Second}); err != nil {
		return nil, errors.Wrapf(err, "failed to open db")
	}

	if err = initBuckets(db, opts.Buckets); err != nil {
		return nil, errors.Wrapf(err, "failed to init buckets")
	}

	return db, nil
}

func initBuckets(db *bolt.DB, buckets [][]byte) error {
	var err error

	for _, b := range buckets {
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
