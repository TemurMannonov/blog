package postgres

import (
	"fmt"

	"github.com/TemurMannonov/blog/storage/repo"
	"github.com/jmoiron/sqlx"
)

type postRepo struct {
	db *sqlx.DB
}

func NewPost(db *sqlx.DB) repo.PostStorageI {
	return &postRepo{
		db: db,
	}
}

func (pr *postRepo) Create(post *repo.Post) (*repo.Post, error) {
	query := `
		INSERT INTO posts(
			title,
			description,
			image_url,
			user_id,
			category_id
		) VALUES($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`

	row := pr.db.QueryRow(
		query,
		post.Title,
		post.Description,
		post.ImageUrl,
		post.UserID,
		post.CategoryID,
	)

	err := row.Scan(
		&post.ID,
		&post.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (pr *postRepo) Get(id int64) (*repo.Post, error) {
	var result repo.Post

	_, err := pr.db.Exec("UPDATE posts SET views_count=views_count+1 WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT
			id,
			title,
			description,
			image_url,
			user_id,
			category_id,
			created_at,
			updated_at,
			views_count
		FROM posts
		WHERE id=$1
	`

	row := pr.db.QueryRow(query, id)
	err = row.Scan(
		&result.ID,
		&result.Title,
		&result.Description,
		&result.ImageUrl,
		&result.UserID,
		&result.CategoryID,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.ViewsCount,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (pr *postRepo) GetAll(params *repo.GetAllPostsParams) (*repo.GetAllPostsResult, error) {
	result := repo.GetAllPostsResult{
		Posts: make([]*repo.Post, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := "WHERE true"
	if params.Search != "" {
		filter += " AND title ilike '%" + params.Search + "%' "
	}

	if params.UserID != 0 {
		filter += fmt.Sprintf(" AND user_id=%d ", params.UserID)
	}

	if params.CategoryID != 0 {
		filter += fmt.Sprintf(" AND category_id=%d ", params.CategoryID)
	}

	orderBy := " ORDER BY created_at desc "
	if params.SortByData != "" {
		orderBy = fmt.Sprintf(" ORDER BY created_at %s ", params.SortByData)
	}

	query := `
		SELECT
			id,
			title,
			description,
			image_url,
			user_id,
			category_id,
			created_at,
			updated_at,
			views_count
		FROM posts
		` + filter + orderBy + limit

	rows, err := pr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var p repo.Post

		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Description,
			&p.ImageUrl,
			&p.UserID,
			&p.CategoryID,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.ViewsCount,
		)
		if err != nil {
			return nil, err
		}

		result.Posts = append(result.Posts, &p)
	}

	queryCount := `SELECT count(1) FROM posts ` + filter
	err = pr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
