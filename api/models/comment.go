package models

import "time"

type Comment struct {
	ID          int64        `json:"id"`
	UserID      int64        `json:"user_id"`
	PostID      int64        `json:"post_id"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   *time.Time   `json:"updated_at"`
	User        *CommentUser `json:"user"`
}

type CommentUser struct {
	ID              int64   `json:"id"`
	FirstName       string  `json:"first_name"`
	LastName        string  `json:"last_name"`
	Email           string  `json:"email"`
	ProfileImageUrl *string `json:"profile_image_url"`
}

type CreateCommentRequest struct {
	Description string `json:"description" binding:"required"`
	PostID      int64  `json:"post_id" binding:"required"`
}

type GetAllCommentsParams struct {
	Limit  int32 `json:"limit" binding:"required" default:"10"`
	Page   int32 `json:"page" binding:"required" default:"1"`
	UserID int64 `json:"user_id"`
	PostID int64 `json:"post_id"`
}

type GetAllCommentsResponse struct {
	Comments []*Comment `json:"comments"`
	Count    int32      `json:"count"`
}
