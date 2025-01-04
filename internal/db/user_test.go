package db

import (
	"context"
	"testing"

	"github.com/litmus-zhang/task-manager/util"
	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {
	hash, err := util.HashPassword(util.RandomString(8))
	require.NoError(t, err)

	args := RegisterUserParams{
		Username:     util.RandomString(6),
		Email:        util.RandomString(6) + "@test.com",
		PasswordHash: hash,
	}

	user, err := testQueries.RegisterUser(context.Background(), args)
	require.NoError(t, err)

	require.NotEmpty(t, user)
	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.Email, user.Email)
	require.Equal(t, args.PasswordHash, user.PasswordHash)
	require.NotZero(t, user.UserID)

}
