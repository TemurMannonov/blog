package postgres_test

import (
	"testing"

	"github.com/TemurMannonov/blog/storage/repo"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T) *repo.User {
	u, err := strg.User().Create(&repo.User{
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Email:     faker.Email(),
		Password:  faker.Password(),
		Type:      repo.UserTypeUser,
	})
	require.NoError(t, err)
	require.NotEmpty(t, u)

	return u
}

func TestCreateUser(t *testing.T) {
	createUser(t)
}
