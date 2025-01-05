package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/litmus-zhang/task-manager/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTask(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCategory(t, u)
	duration := 24 * time.Hour * time.Duration(util.RandomInt(1, 10))
	durationStr := fmt.Sprintf("%d seconds", int64(duration.Seconds()))
	fmt.Printf("durationStr: %v", durationStr)

	arg := CreateTaskParams{
		TaskName:     fmt.Sprintf("Task %v", util.RandomString(8)),
		CategoryID:   c.CategoryID,
		Description:  sql.NullString{Valid: true, String: util.RandomString(15)},
		TimelineType: durationStr,
	}
	task, err := testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, task)
	require.Equal(t, task.CategoryID, c.CategoryID)
	require.Equal(t, task.IsCompleted, sql.NullBool{Bool: false, Valid: true})

}
