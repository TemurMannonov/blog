package postgres

import (
	"github.com/TemurMannonov/blog/storage/repo"
	"github.com/jmoiron/sqlx"
)

type likeRepo struct {
	db *sqlx.DB
}

func NewLike(db *sqlx.DB) repo.LikeStorageI {
	return &likeRepo{
		db: db,
	}
}

func (lr *likeRepo) Create(l *repo.Like) (*repo.Like, error) {
	query := `
		INSERT INTO likes(user_id, post_id, status) 
		VALUES($1, $2, $3) RETURNING id
	`

	row := lr.db.QueryRow(
		query,
		l.UserID,
		l.PostID,
		l.Status,
	)

	err := row.Scan(&l.ID)
	if err != nil {
		return nil, err
	}

	return l, nil
}

func (cr *likeRepo) Get(userID, postID int64) (*repo.Like, error) {
	var result repo.Like

	query := `
		SELECT
			id,
			user_id,
			post_id,
			status
		FROM likes
		WHERE user_id=$1 AND post_id=$2
	`

	row := cr.db.QueryRow(query, userID, postID)
	err := row.Scan(
		&result.ID,
		&result.UserID,
		&result.PostID,
		&result.Status,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (cr *likeRepo) GetLikesDislikesCount(postID int64) (*repo.LikesDislikesCountsResult, error) {
	var result repo.LikesDislikesCountsResult

	query := `
		SELECT
			COUNT(1) FILTER (WHERE status=true) as likes_count,
			COUNT(1) FILTER (WHERE status=false) as dislikes_count
		FROM likes
		WHERE post_id=$1
	`

	row := cr.db.QueryRow(query, postID)
	err := row.Scan(
		&result.LikesCount,
		&result.DislikesCount,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
