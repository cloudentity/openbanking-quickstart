package shared

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	bolt "go.etcd.io/bbolt"
)

var mode = os.FileMode(0o600)

type DB struct {
	*bolt.DB
}

func InitDB(dbFilePath string) (DB, error) {
	var (
		db  = DB{}
		err error
	)

	if db.DB, err = bolt.Open(dbFilePath, mode, &bolt.Options{Timeout: 3 * time.Second}); err != nil {
		return db, errors.Wrapf(err, "failed to open db")
	}

	return db, nil
}

func CreateBucket(db DB, bucket []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(bucket); err != nil {
			return errors.Wrapf(err, "failed to create bucket: %s", string(bucket))
		}

		return nil
	})
}

func InitTestDB(t *testing.T) DB {
	t.Helper()

	var (
		dir = t.TempDir()
		db  DB
		err error
	)

	db, err = InitDB(fmt.Sprintf("%s/test.db", dir))
	require.NoError(t, err)

	return db
}
