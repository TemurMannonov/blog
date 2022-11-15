package repo

import "time"

type Post struct {
	ID          int64
	Title       string
	Description string
	ImageUrl    *string
	UserID      int64
	CategoryID  int64
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	ViewsCount  int32
}

type PostStorageI interface {
	Create(u *Post) (*Post, error)
	Get(id int64) (*Post, error)
}
