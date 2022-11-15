package repo

import "time"

type Category struct {
	ID        int64
	Title     string
	CreatedAt time.Time
}

type CategoryStorageI interface {
	Create(u *Category) (*Category, error)
	Get(id int64) (*Category, error)
}
