package postgres

import (
	"fmt"

	"github.com/TemurMannonov/blog/storage/repo"
	"github.com/jmoiron/sqlx"
)

type commentRepo struct {
	db *sqlx.DB
}

func NewComment(db *sqlx.DB) repo.CommentStorageI {
	return &commentRepo{
		db: db,
	}
}

func (pr *commentRepo) Create(comment *repo.Comment) (*repo.Comment, error) {
	query := `
		INSERT INTO comments(
			user_id,
			post_id,
			description
		) VALUES($1, $2, $3)
		RETURNING id, created_at
	`

	row := pr.db.QueryRow(
		query,
		comment.UserID,
		comment.PostID,
		comment.Description,
	)

	err := row.Scan(
		&comment.ID,
		&comment.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (pr *commentRepo) GetAll(params *repo.GetAllCommentsParams) (*repo.GetAllCommentsResult, error) {
	result := repo.GetAllCommentsResult{
		Comments: make([]*repo.Comment, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := " WHERE true "
	if params.UserID != 0 {
		filter += fmt.Sprintf(" AND c.user_id=%d ", params.UserID)
	}

	if params.PostID != 0 {
		filter += fmt.Sprintf(" AND c.post_id=%d ", params.PostID)
	}

	query := `
		SELECT
			c.id,
			c.user_id,
			c.post_id,
			c.description,
			c.created_at,
			c.updated_at,
			u.first_name,
			u.last_name,
			u.email,
			u.profile_image_url
		FROM comments c
		INNER JOIN users u ON u.id=c.user_id
		` + filter + `
		ORDER BY c.created_at desc` + limit

	rows, err := pr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var c repo.Comment

		err := rows.Scan(
			&c.ID,
			&c.UserID,
			&c.PostID,
			&c.Description,
			&c.CreatedAt,
			&c.UpdatedAt,
			&c.User.FirstName,
			&c.User.LastName,
			&c.User.Email,
			&c.User.ProfileImageUrl,
		)
		if err != nil {
			return nil, err
		}

		result.Comments = append(result.Comments, &c)
	}

	queryCount := `
		SELECT count(1) FROM comments c
		INNER JOIN users u ON u.id=c.user_id ` + filter
	err = pr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
