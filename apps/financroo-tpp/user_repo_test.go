package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUsersRepo(t *testing.T) {
	t.Parallel()
	db, err := InitDB(Config{DBFile: "./test.db"})
	require.NoError(t, err)
	defer db.Close()

	repo, err := NewUserRepo(db)
	require.NoError(t, err)

	u1, err := repo.Get("test")
	require.NoError(t, err)

	require.Equal(t, "test", u1.Sub)

	cb := ConnectedBank{
		BankID:   "b1",
		IntentID: "i1",
	}

	err = repo.Set(User{
		Sub:   "user",
		Banks: []ConnectedBank{cb},
	})
	require.NoError(t, err)

	u, err := repo.Get("user")
	require.NoError(t, err)

	require.Equal(t, "user", u.Sub)
	require.Equal(t, 1, len(u.Banks))
	require.Equal(t, "b1", u.Banks[0].BankID)
	require.Equal(t, "i1", u.Banks[0].IntentID)
}
