package main

import (
	"encoding/json"

	"github.com/cloudentity/openbanking-quickstart/shared"
	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
)

type ConsentRepo struct {
	shared.DB
}

var (
	consentsBucket     = []byte("consents")
	ErrConsentNotFound = errors.New("consent not found")
)

func NewConsentRepo(db shared.DB) (repo ConsentRepo, err error) {
	if err = shared.CreateBucket(db, consentsBucket); err != nil {
		return repo, err
	}

	repo.DB = db

	return repo, nil
}

func (u *ConsentRepo) List() ([]Consent, error) {
	var (
		consents = []Consent{}
		err      error
	)

	if err = u.DB.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(consentsBucket).Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var consent Consent

			if err = json.Unmarshal(v, &consent); err != nil {
				return errors.Wrapf(err, "failed to unmarshal consent")
			}

			consents = append(consents, consent)
		}

		return nil
	}); err != nil {
		return consents, errors.Wrapf(err, "failed to list consents")
	}

	return ConsentsByCreatedDate(consents), nil
}

func (u *ConsentRepo) Create(consent Consent) error {
	var (
		bs  []byte
		err error
	)

	if bs, err = json.Marshal(&consent); err != nil {
		return errors.Wrapf(err, "failed to marshal consent")
	}

	if err = u.DB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(consentsBucket).Put([]byte(consent.ID), bs)
	}); err != nil {
		return errors.Wrapf(err, "failed to update user")
	}

	return nil
}

func (u *ConsentRepo) Get(id ConsentID) (Consent, error) {
	var (
		bs      []byte
		consent Consent
		err     error
	)

	if err = u.DB.View(func(tx *bolt.Tx) error {
		bs = tx.Bucket(consentsBucket).Get([]byte(id))

		if bs == nil {
			return ErrConsentNotFound
		}

		if err = json.Unmarshal(bs, &consent); err != nil {
			return errors.Wrapf(err, "failed to unmarshal consent")
		}

		return nil
	}); err != nil {
		return Consent{}, err
	}

	return consent, nil
}

func (u *ConsentRepo) Update(consent Consent) error {
	var (
		bs  []byte
		err error
	)

	if bs, err = json.Marshal(&consent); err != nil {
		return errors.Wrapf(err, "failed to marshal consent")
	}

	if err = u.DB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(consentsBucket).Put([]byte(consent.ID), bs)
	}); err != nil {
		return errors.Wrapf(err, "failed to update consent")
	}

	return nil
}
