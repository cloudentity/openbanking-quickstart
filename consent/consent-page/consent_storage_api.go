package main

import (
	"context"

	"github.com/cloudentity/acp-client-go/clients/system/models"
	"github.com/sirupsen/logrus"
)

type ConsentID string

type Subject string

type Data models.ScopeGrantSessionResponse

type ConsentStorage interface {
	Store(ctx context.Context, sub Subject, data Data) (ConsentID, error)
}

type DummyConsentStorage struct{}

var _ ConsentStorage = &DummyConsentStorage{}

func (d DummyConsentStorage) Store(ctx context.Context, sub Subject, data Data) (ConsentID, error) {
	logrus.Infof("Store consent for sub: %s with data: %+v", sub, data)

	return "external-consent-id", nil
}
