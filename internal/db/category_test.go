package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCategory(t *testing.T) {
	u := CreateTestUser(t)
	_ = CreateTestCategory(t, u)

}

func TestGetAllUserCategory(t *testing.T) {
	u := CreateTestUser(t)

	for i := 0; i < 10; i++ {
		_ = CreateTestCategory(t, u)
	}
	args := GetAllUserCategoryParams{
		UserID: u.UserID,
		Limit:  10,
		Offset: 0,
	}

	categories, err := testQueries.GetAllUserCategory(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, 10, len(categories))

}
