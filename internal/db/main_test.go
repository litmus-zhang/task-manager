package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/litmus-zhang/momentum-backend/internal/config"
	"github.com/litmus-zhang/momentum-backend/util"
	"github.com/stretchr/testify/require"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("cannot setup a new config:", err)
	}

	testDB, err = sql.Open(cfg.DbDriver, cfg.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
	testDB.Close()
}

func CreateTestUser(t *testing.T) User {
	hash, err := util.HashPassword(util.RandomString(8))
	require.NoError(t, err)

	args := RegisterUserParams{
		Username:     util.RandomString(6),
		Email:        util.RandomString(6) + "@test.com",
		PasswordHash: hash,
		FullName:     util.RandomString(6) + " " + util.RandomString(6),
	}

	user, err := testQueries.RegisterUser(context.Background(), args)
	require.NoError(t, err)

	require.NotEmpty(t, user)
	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.Email, user.Email)
	require.Equal(t, args.PasswordHash, user.PasswordHash)
	require.NotZero(t, user.UserID)
	return user
}

func CreateTestCategory(t *testing.T, u User) Category {
	args := CreateCategoryParams{
		CategoryName: util.RandomString(10),
		Description:  sql.NullString{Valid: true, String: util.RandomString(20)},
		UserID:       u.UserID,
	}

	category, err := testQueries.CreateCategory(context.Background(), args)

	require.NoError(t, err)
	require.NotNil(t, category)
	require.Equal(t, u.UserID, category.UserID)
	return category
}
