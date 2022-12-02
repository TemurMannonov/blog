package postgres

import (
	"database/sql"
	"fmt"

	"github.com/TemurMannonov/blog/storage/repo"
	"github.com/jmoiron/sqlx"
)

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategory(db *sqlx.DB) repo.CategoryStorageI {
	return &categoryRepo{
		db: db,
	}
}

func (cr *categoryRepo) Create(category *repo.Category) (*repo.Category, error) {
	query := `
		INSERT INTO categories(title) VALUES($1)
		RETURNING id, created_at
	`

	row := cr.db.QueryRow(
		query,
		category.Title,
	)

	err := row.Scan(
		&category.ID,
		&category.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (cr *categoryRepo) Get(id int64) (*repo.Category, error) {
	var result repo.Category

	query := `
		SELECT
			id,
			title,
			created_at
		FROM categories
		WHERE id=$1
	`

	row := cr.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.Title,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (cr *categoryRepo) GetAll(params *repo.GetAllCategoriesParams) (*repo.GetAllCategoriesResult, error) {
	result := repo.GetAllCategoriesResult{
		Categories: make([]*repo.Category, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := ""
	if params.Search != "" {
		filter += " WHERE title ilike '%" + params.Search + "%' "
	}

	query := `
		SELECT
			id, 
			title, 
			created_at
		FROM categories
		` + filter + `
		ORDER BY created_at desc
		` + limit

	rows, err := cr.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var c repo.Category

		err := rows.Scan(
			&c.ID,
			&c.Title,
			&c.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		result.Categories = append(result.Categories, &c)
	}

	queryCount := `SELECT count(1) FROM categories ` + filter
	err = cr.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (cr *categoryRepo) Update(category *repo.Category) (*repo.Category, error) {
	query := `
		UPDATE categories SET title=$1 WHERE id=$2
		RETURNING created_at
	`

	err := cr.db.QueryRow(query, category.Title, category.ID).Scan(&category.CreatedAt)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (cr *categoryRepo) Delete(id int64) error {
	query := `DELETE FROM categories WHERE id=$1`

	result, err := cr.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsEffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsEffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
