package main

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/cloudentity/acp-client-go/clients/system/models"
)

type ConsentID string

type Data struct {
	AccountIDs []string `json:"account_ids"`
	models.ScopeGrantSessionResponse
}

type ConsentStorage interface {
	Store(ctx context.Context, data Data) (ConsentID, error)
}

type DummyConsentStorage struct{}

var _ ConsentStorage = &DummyConsentStorage{}

func (d DummyConsentStorage) Store(_ context.Context, data Data) (ConsentID, error) {
	logrus.Infof("Store consent for sub: %s with data: %+v", data.Subject, data)

	return "external-consent-id", nil
}
