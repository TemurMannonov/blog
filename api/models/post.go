package models

import "time"

type Post struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ImageUrl    *string    `json:"image_url"`
	UserID      int64      `json:"user_id"`
	CategoryID  int64      `json:"category_id"`
	UpdatedAt   *time.Time `json:"updated_at"`
	ViewsCount  int32      `json:"views_count"`
	CreatedAt   time.Time  `json:"created_at"`
}

type CreatePostRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageUrl    *string `json:"image_url"`
	UserID      int64   `json:"user_id"`
	CategoryID  int64   `json:"category_id"`
}
