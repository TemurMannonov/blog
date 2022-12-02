package repo

import "time"

type Category struct {
	ID        int64
	Title     string
	CreatedAt time.Time
}

type GetAllCategoriesParams struct {
	Limit  int32
	Page   int32
	Search string
}

type GetAllCategoriesResult struct {
	Categories []*Category
	Count      int32
}

type CategoryStorageI interface {
	Create(c *Category) (*Category, error)
	Get(id int64) (*Category, error)
	GetAll(params *GetAllCategoriesParams) (*GetAllCategoriesResult, error)
	Update(c *Category) (*Category, error)
	Delete(id int64) error
}
