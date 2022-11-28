package models

type Like struct {
	ID     int64 `json:"id"`
	PostID int64 `json:"post_id"`
	UserID int64 `json:"user_id"`
	Status bool  `json:"status"`
}

type CreateOrUpdateLikeRequest struct {
	PostID int64 `json:"post_id" binding:"required"`
	Status bool  `json:"status"`
}
