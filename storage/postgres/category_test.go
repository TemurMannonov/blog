package postgres_test

import (
	"testing"

	"github.com/TemurMannonov/blog/storage/repo"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createCategory(t *testing.T) *repo.Category {
	blog, err := strg.Category().Create(&repo.Category{
		Title: faker.Sentence(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, blog)

	return blog
}

func TestGetCategory(t *testing.T) {
	c := createCategory(t)

	blog, err := strg.Category().Get(c.ID)
	require.NoError(t, err)
	require.NotEmpty(t, blog)
}

func TestCreateCategory(t *testing.T) {
	createCategory(t)
}
