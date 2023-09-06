package main

import (
	"context"

	acpclient "github.com/cloudentity/acp-client-go"
	"github.com/cloudentity/acp-client-go/clients/identity/client/users"
	"github.com/cloudentity/acp-client-go/clients/identity/models"
	"github.com/pkg/errors"
)

type IdentityPoolConsentStorage struct {
	PoolID string
	Client acpclient.Client
}

func NewIdentityPoolConsentStorage(config IdentityPoolConsentStorageConfig) (ConsentStorage, error) {
	var (
		storage = IdentityPoolConsentStorage{
			PoolID: config.PoolID,
		}
		err error
	)

	if storage.Client, err = acpclient.New(acpclient.Config{
		IssuerURL:    config.IssuerURL,
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RootCA:       config.RootCA,
	}); err != nil {
		return nil, errors.Wrapf(err, "failed to create acp client")
	}

	return &storage, nil
}

var _ ConsentStorage = &IdentityPoolConsentStorage{}

func (s *IdentityPoolConsentStorage) Store(ctx context.Context, sub Subject, data Data) (ConsentID, error) {
	var (
		resp *users.CreateUserCreated
		err  error
	)

	accountIDs := []string{} // TODOX

	if resp, err = s.Client.Identity.Acp.Users.CreateUser(
		users.NewCreateUserParamsWithContext(ctx).
			WithIPID(s.PoolID).
			WithNewUser(&models.NewUser{
				Metadata: map[string]interface{}{},
				Payload: map[string]interface{}{
					"subject":        data.Subject,
					"client_id":      data.ClientID,
					"status":         "authorized",
					"granted_scopes": data.GrantedScopes,
					"account_ids":    accountIDs,
				},
			}),
		nil,
	); err != nil {
		return "", errors.Wrapf(err, "failed to create consent")
	}

	return ConsentID(resp.Payload.ID), nil
}
