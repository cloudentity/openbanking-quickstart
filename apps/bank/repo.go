package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/cloudentity/openbanking-quickstart/openbanking/obuk/accountinformation/models"
	paymentModels "github.com/cloudentity/openbanking-quickstart/openbanking/obuk/paymentinitiation/models"
	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
)

type BankUserData struct {
	Accounts     Accounts     `json:"accounts"`
	Balances     Balances     `json:"balances"`
	Transactions Transactions `json:"transactions"`
	Payments     Payments     `json:"payments"`
}

type Accounts struct {
	OBUK []models.OBAccount6 `json:"obuk"`
}

type Balances struct {
	OBUK []models.OBReadBalance1DataBalanceItems0 `json:"obuk"`
}

type Transactions struct {
	OBUK []models.OBTransaction6 `json:"obuk"`
}

type Payments struct {
	OBUK []paymentModels.OBWriteDomesticResponse5 `json:"obuk"`
}

type Storage interface {
	Get(string) (BankUserData, error)
	Put(string, BankUserData) error
}

type Has interface {
	Has(interface{}) bool
}

var bucketName = []byte(`users`)

type UserRepo struct {
	*bolt.DB
}

func (u *UserRepo) Get(sub string) (BankUserData, error) {
	var (
		data BankUserData
		err  error
	)

	if err = u.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		v := b.Get([]byte(sub))
		if err = json.Unmarshal(v, &data); err != nil {
			return errors.Wrapf(err, fmt.Sprintf("failed to unmarshal data for user %s", sub))
		}
		return nil
	}); err != nil {
		return data, err
	}

	return data, nil
}

func (u *UserRepo) Put(sub string, data BankUserData) error {
	var (
		dataBytes []byte
		err       error
	)

	if dataBytes, err = json.Marshal(data); err != nil {
		return err
	}

	return u.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		if err = b.Put([]byte(sub), dataBytes); err != nil {
			return errors.Wrapf(err, "failed to put value into database")
		}
		return nil
	})
}

func NewUserRepo() (*UserRepo, error) {
	var (
		userRepo UserRepo
		u2df     UserToDataFile
		err      error
	)

	// create db
	if userRepo.DB, err = bolt.Open("data/my.db", os.FileMode(0644), &bolt.Options{Timeout: 3 * time.Second}); err != nil {
		return nil, errors.Wrapf(err, "failed to open db")
	}

	// read init data from file
	if u2df, err = readUserToDataFile(); err != nil {
		return nil, errors.Wrapf(err, "failed to read data file")
	}

	// setup bucket and default data
	if err = userRepo.Update(func(tx *bolt.Tx) error {
		var bucket *bolt.Bucket
		if bucket, err = tx.CreateBucketIfNotExists(bucketName); err != nil {
			return errors.Wrapf(err, "failed to create bucket")
		}

		for k, v := range u2df {
			var valBytes []byte

			if valBytes, err = json.Marshal(v); err != nil {
				return errors.Wrapf(err, "failed to unmarshal data value from file")
			}

			if err = bucket.Put([]byte(k), valBytes); err != nil {
				return errors.Wrapf(err, "failed to put value in bucket")
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &userRepo, nil
}

type UserToDataFile map[string]BankUserData

func readUserToDataFile() (UserToDataFile, error) {
	var (
		bs   []byte
		u2df UserToDataFile
		err  error
	)

	if bs, err = ioutil.ReadFile("./data/data.json"); err != nil {
		return u2df, errors.Wrapf(err, "failed to read file")
	}

	if err = json.Unmarshal(bs, &u2df); err != nil {
		return u2df, errors.Wrapf(err, "failed to unmarshal data")
	}

	return u2df, nil
}
