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

	category, err := strg.Category().Get(c.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category)
}

func TestGetAllCategories(t *testing.T) {
	createCategory(t)

	result, err := strg.Category().GetAll(&repo.GetAllCategoriesParams{
		Limit: 10,
		Page:  1,
	})

	require.NoError(t, err)
	require.GreaterOrEqual(t, int(result.Count), 1)
}

func TestCreateCategory(t *testing.T) {
	createCategory(t)
}

func TestUpdateCategory(t *testing.T) {
	c := createCategory(t)

	c.Title = faker.Sentence()

	category, err := strg.Category().Update(c)
	require.NoError(t, err)
	require.NotEmpty(t, category)
	require.Equal(t, category.Title, c.Title)
}

func TestDeleteCategory(t *testing.T) {
	c := createCategory(t)

	err := strg.Category().Delete(c.ID)
	require.NoError(t, err)
}
