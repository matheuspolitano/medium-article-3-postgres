package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := InsertUserParams{
		Role:     "Manager",
		Username: "Matheus Politano",
	}
	user, err := testQueries.InsertUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Role, user.Role)
	require.Equal(t, arg.Username, user.Username)
	require.NotZero(t, user.ID)
}

func TestUpdateUser(t *testing.T) {
	arg := InsertUserParams{
		Role:     "Manager",
		Username: "Matheus Politano",
	}
	user, err := testQueries.InsertUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotZero(t, user.ID)
	argUpdate := UpdateRoleParams{Role: "CTO", ID: user.ID}
	err = testQueries.UpdateRole(context.Background(), argUpdate)
	require.NoError(t, err)
	user, err = testQueries.GetUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, argUpdate.Role, user.Role)
	err = testQueries.DeleteItem(context.Background(), user.ID)
	require.NoError(t, err)

}
