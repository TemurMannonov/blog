package repo

import "time"

type User struct {
	ID              int64
	FirstName       string
	LastName        string
	PhoneNumber     *string
	Email           string
	Gender          *string
	Password        string
	Username        string
	ProfileImageUrl *string
	Type            string
	CreatedAt       time.Time
}

type GetAllUsersParams struct {
	Limit  int32
	Page   int32
	Search string
}

type GetAllUsersResult struct {
	Users []*User
	Count int32
}

type UserStorageI interface {
	Create(u *User) (*User, error)
	Get(id int64) (*User, error)
	GetAll(params *GetAllUsersParams) (*GetAllUsersResult, error)
}
