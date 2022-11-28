package models

import "time"

type Post struct {
	ID          int64         `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	ImageUrl    *string       `json:"image_url"`
	UserID      int64         `json:"user_id"`
	CategoryID  int64         `json:"category_id"`
	UpdatedAt   *time.Time    `json:"updated_at"`
	ViewsCount  int32         `json:"views_count"`
	CreatedAt   time.Time     `json:"created_at"`
	LikeInfo    *PostLikeInfo `json:"like_info"`
}

type PostLikeInfo struct {
	LikesCount    int64 `json:"likes_count"`
	DislikesCount int64 `json:"dislikes_count"`
}

type CreatePostRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageUrl    *string `json:"image_url"`
	CategoryID  int64   `json:"category_id"`
}

type GetAllPostsParams struct {
	Limit      int32  `json:"limit" binding:"required" default:"10"`
	Page       int32  `json:"page" binding:"required" default:"1"`
	Search     string `json:"search"`
	UserID     int64  `json:"user_id"`
	CategoryID int64  `json:"category_id"`
	SortByData string `json:"sort_by_date" enums:"asc,desc" default:"desc"`
}

type GetAllPostsResponse struct {
	Posts []*Post `json:"posts"`
	Count int32   `json:"count"`
}
