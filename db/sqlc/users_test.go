package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/a3hi3h/aperno/util"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		ID:        uuid.New(),
		FirstName: util.RandomString(6),
		LastName:  util.RandomString(6),
		Email:     util.RandomEmail(),
		Hashedpwd: util.RandomString(6),
		Type:      1,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.ID, user.ID)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
