package repositories

import (
	"context"
)

const createCategory = `-- name: CreateCategory :one

INSERT INTO categories(category_id, category_name, description, picture) VALUES ($1, $2, $3, $4)
RETURNING category_id
`

type CreateCategoryParams struct {
	CategoryID   int16  `db:"category_id" json:"categoryId"`
	CategoryName string `db:"category_name" json:"categoryName"`
	Description  string `db:"description" json:"description"`
	Picture      []byte `db:"picture" json:"picture"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (int16, error) {
	row := q.db.QueryRowContext(ctx, createCategory,
		arg.CategoryID,
		arg.CategoryName,
		arg.Description,
		arg.Picture,
	)
	var category_id int16
	err := row.Scan(&category_id)
	return category_id, err
}
