package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegisterUser(t *testing.T) {
	_ = CreateTestUser(t)

}

func TestGetUser(t *testing.T) {
	u := CreateTestUser(t)

	u2, err := testQueries.GetUserByEmail(context.Background(), u.Email)
	require.NoError(t, err)
	require.Equal(t, u, u2)

}
