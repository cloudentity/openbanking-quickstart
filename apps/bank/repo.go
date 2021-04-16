package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/cloudentity/openbanking-quickstart/models"
	paymentModels "github.com/cloudentity/openbanking-quickstart/openbanking/paymentinitiation/models"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
)

var bucketName = []byte(`users`)

type UserRepo struct {
	*bolt.DB
}

type Data struct {
	Accounts     []models.OBAccount6                      `json:"accounts"`
	Balances     []models.OBReadBalance1DataBalanceItems0 `json:"balances"`
	Transactions []models.OBTransaction6                  `json:"transactions"`
	Payments     []paymentModels.OBWriteDomesticResponse5 `json:"payments"`
}

type UserToDataFile map[string]Data

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

func NewUserRepo() (UserRepo, error) {
	var (
		userRepo UserRepo
		u2df     UserToDataFile
		err      error
	)

	// create db
	if userRepo.DB, err = bolt.Open("data/tppdb.db", os.FileMode(0644), &bolt.Options{Timeout: 3 * time.Second}); err != nil {
		return userRepo, errors.Wrapf(err, "failed to open db")
	}

	// read init data from file
	if u2df, err = readUserToDataFile(); err != nil {
		return userRepo, errors.Wrapf(err, "failed to read data file")
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
		return userRepo, err
	}

	return userRepo, nil
}

func (ur *UserRepo) GetAccounts(sub string) ([]models.OBAccount6, error) {
	var (
		data Data
		err  error
	)

	if err = ur.loadUser(sub, &data); err != nil {
		return data.Accounts, err
	}

	return data.Accounts, nil
}

func (ur *UserRepo) GetBalances(sub string) ([]models.OBReadBalance1DataBalanceItems0, error) {
	var (
		data Data
		err  error
	)

	if err = ur.loadUser(sub, &data); err != nil {
		return data.Balances, err
	}

	return data.Balances, nil
}

func (ur *UserRepo) GetTransactions(sub string) ([]models.OBTransaction6, error) {
	var (
		data Data
		err  error
	)

	if err = ur.loadUser(sub, &data); err != nil {
		return data.Transactions, err
	}

	return data.Transactions, nil
}

func (ur *UserRepo) CreateDomesticPayment(sub string, payment paymentModels.OBWriteDomesticResponse5) error {
	var (
		data Data
		err  error
	)

	if err = ur.loadUser(sub, &data); err != nil {
		return errors.Wrapf(err, fmt.Sprintf("failed to load user %s from database", sub))
	}

	for _, p := range data.Payments {
		if p.Data.DomesticPaymentID == payment.Data.DomesticPaymentID {
			return ErrAlreadyExists{fmt.Sprintf("/domestic-payments/%s", *p.Data.DomesticPaymentID)}
		}
	}

	data.Payments = append(data.Payments, payment)

	return ur.writeData(bucketName, []byte(sub), data)
}

func (ur *UserRepo) GetDomesticPayment(sub, domesticPaymentID string) (paymentModels.OBWriteDomesticResponse5, error) {
	var (
		data    Data
		payment paymentModels.OBWriteDomesticResponse5
		err     error
	)

	if err = ur.loadUser(sub, &data); err != nil {
		return payment, errors.Wrapf(err, "failed to load data from db")
	}

	for _, p := range data.Payments {
		if *p.Data.DomesticPaymentID == domesticPaymentID {
			return p, nil
		}
	}

	return payment, ErrNotFound{fmt.Sprintf("domestic-payment with id %s", domesticPaymentID)}
}

func (ur *UserRepo) SetDomesticPaymentStatus(domesticPaymentID string, status DomesticPaymentStatus) error {
	var (
		data = make(map[string]Data)
		err  error
	)

	if err = ur.loadAll(data); err != nil {
		return errors.Wrapf(err, "failed to load data from db")
	}

	for k, v := range data {
		for i, payment := range v.Payments {
			if *payment.Data.DomesticPaymentID == domesticPaymentID {
				*data[k].Payments[i].Data.Status = string(status)
				*data[k].Payments[i].Data.StatusUpdateDateTime = strfmt.DateTime(time.Now())
				return ur.writeData(bucketName, []byte(k), data[k])
			}
		}
	}

	return fmt.Errorf("unable to find domestic payment id %s", domesticPaymentID)
}

func (ur *UserRepo) loadUser(sub string, data *Data) error {
	return ur.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		v := b.Get([]byte(sub))
		if err := json.Unmarshal(v, data); err != nil {
			return errors.Wrapf(err, fmt.Sprintf("failed to unmarshal data for user %s", sub))
		}
		return nil
	})
}

func (ur *UserRepo) writeData(bucket, key []byte, data Data) error {
	var (
		dataBytes []byte
		err       error
	)

	if dataBytes, err = json.Marshal(data); err != nil {
		return err
	}

	return ur.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if err = b.Put(key, dataBytes); err != nil {
			return errors.Wrapf(err, "failed to put value into database")
		}
		return nil
	})
}

func (ur *UserRepo) loadAll(m map[string]Data) error {
	return ur.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var (
				data Data
				err  error
			)
			if err = json.Unmarshal(v, &data); err != nil {
				return errors.Wrapf(err, fmt.Sprintf("failed to unmarshal data for user %s", string(k)))
			}
			m[string(k)] = data
		}

		return nil
	})
}
