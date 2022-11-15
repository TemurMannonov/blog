package storage

import (
	"github.com/TemurMannonov/blog/storage/postgres"
	"github.com/TemurMannonov/blog/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	User() repo.UserStorageI
	Category() repo.CategoryStorageI
	Post() repo.PostStorageI
}

type storagePg struct {
	userRepo     repo.UserStorageI
	categoryRepo repo.CategoryStorageI
	postRepo     repo.PostStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		userRepo:     postgres.NewUser(db),
		categoryRepo: postgres.NewCategory(db),
		postRepo:     postgres.NewPost(db),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}

func (s *storagePg) Category() repo.CategoryStorageI {
	return s.categoryRepo
}

func (s *storagePg) Post() repo.PostStorageI {
	return s.postRepo
}
