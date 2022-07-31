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
		UUuid:      uuid.New(),
		UFirstName: util.RandomString(6),
		ULastName:  util.RandomString(6),
		UType:      1,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.UFirstName, user.UFirstName)
	require.Equal(t, arg.ULastName, user.ULastName)
	require.Equal(t, arg.UUuid, user.UUuid)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
