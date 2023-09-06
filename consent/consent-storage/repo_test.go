package main

import (
	"testing"

	"github.com/cloudentity/openbanking-quickstart/shared"
	"github.com/stretchr/testify/require"
)

func TestRepo(t *testing.T) {
	db := shared.InitTestDB(t)

	repo, err := NewConsentRepo(db)
	require.NoError(t, err)

	list, err := repo.List()
	require.NoError(t, err)
	require.Empty(t, list)

	consent := Consent{ID: ConsentID("1"), Subject: "user1", Status: AuthorizedStatus}
	err = repo.Create(consent)
	require.NoError(t, err)

	require.Equal(t, ConsentID("1"), consent.ID)

	list, err = repo.List()
	require.NoError(t, err)
	require.Equal(t, 1, len(list))

	fetchedConsent, err := repo.Get(consent.ID)
	require.NoError(t, err)

	require.Equal(t, "user1", fetchedConsent.Subject)

	consent.Status = RevokedStatus

	err = repo.Update(consent)
	require.NoError(t, err)

	fetchedConsent, err = repo.Get(consent.ID)
	require.NoError(t, err)

	require.Equal(t, RevokedStatus, fetchedConsent.Status)
}
