package models

import "time"

type User struct {
	ID              int64     `json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	PhoneNumber     *string   `json:"phone_number"`
	Email           string    `json:"email"`
	Gender          *string   `json:"gender"`
	Username        string    `json:"username"`
	ProfileImageUrl *string   `json:"profile_image_url"`
	Type            string    `json:"type"`
	CreatedAt       time.Time `json:"created_at"`
}

type CreateUserRequest struct {
	FirstName       string  `json:"first_name"`
	LastName        string  `json:"last_name"`
	PhoneNumber     *string `json:"phone_number"`
	Email           string  `json:"email"`
	Gender          *string `json:"gender"`
	Username        string  `json:"username"`
	ProfileImageUrl *string `json:"profile_image_url"`
	Type            string  `json:"type"`
	Password        string  `json:"password"`
}
