// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: query.sql

package models

import (
	"context"
	"database/sql"
)

const createCategory = `-- name: CreateCategory :one

INSERT INTO categories(category_id, category_name, description, picture) VALUES ($1, $2, $3, $4)
RETURNING category_id
`

type CreateCategoryParams struct {
	CategoryID   int16          `db:"category_id" json:"categoryId"`
	CategoryName string         `db:"category_name" json:"categoryName"`
	Description  sql.NullString `db:"description" json:"description"`
	Picture      []byte         `db:"picture" json:"picture"`
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

const createProduct = `-- name: CreateProduct :one

INSERT INTO products 
(product_id, product_name, supplier_id, category_id, 
quantity_per_unit, unit_price, units_in_stock, 
units_on_order, reorder_level, discontinued)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
RETURNING product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued
`

type CreateProductParams struct {
	ProductID       int16           `db:"product_id" json:"productId"`
	ProductName     string          `db:"product_name" json:"productName"`
	SupplierID      sql.NullInt16   `db:"supplier_id" json:"supplierId"`
	CategoryID      sql.NullInt16   `db:"category_id" json:"categoryId"`
	QuantityPerUnit sql.NullString  `db:"quantity_per_unit" json:"quantityPerUnit"`
	UnitPrice       sql.NullFloat64 `db:"unit_price" json:"unitPrice"`
	UnitsInStock    sql.NullInt16   `db:"units_in_stock" json:"unitsInStock"`
	UnitsOnOrder    sql.NullInt16   `db:"units_on_order" json:"unitsOnOrder"`
	ReorderLevel    sql.NullInt16   `db:"reorder_level" json:"reorderLevel"`
	Discontinued    int32           `db:"discontinued" json:"discontinued"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.ProductID,
		arg.ProductName,
		arg.SupplierID,
		arg.CategoryID,
		arg.QuantityPerUnit,
		arg.UnitPrice,
		arg.UnitsInStock,
		arg.UnitsOnOrder,
		arg.ReorderLevel,
		arg.Discontinued,
	)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.ProductName,
		&i.SupplierID,
		&i.CategoryID,
		&i.QuantityPerUnit,
		&i.UnitPrice,
		&i.UnitsInStock,
		&i.UnitsOnOrder,
		&i.ReorderLevel,
		&i.Discontinued,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM Categories
WHERE category_id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, categoryID int16) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, categoryID)
	return err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products
WHERE product_id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, productID int16) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, productID)
	return err
}

const getCategory = `-- name: GetCategory :one
SELECT category_id, category_name, description, picture FROM categories
WHERE category_id = $1
`

func (q *Queries) GetCategory(ctx context.Context, categoryID int16) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, categoryID)
	var i Category
	err := row.Scan(
		&i.CategoryID,
		&i.CategoryName,
		&i.Description,
		&i.Picture,
	)
	return i, err
}

const getProducts = `-- name: GetProducts :one

SELECT product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued FROM products
WHERE category_id = $1
`

// products
func (q *Queries) GetProducts(ctx context.Context, categoryID sql.NullInt16) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProducts, categoryID)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.ProductName,
		&i.SupplierID,
		&i.CategoryID,
		&i.QuantityPerUnit,
		&i.UnitPrice,
		&i.UnitsInStock,
		&i.UnitsOnOrder,
		&i.ReorderLevel,
		&i.Discontinued,
	)
	return i, err
}

const listCategories = `-- name: ListCategories :many
SELECT category_id, category_name, description, picture FROM Categories
ORDER BY category_name
`

func (q *Queries) ListCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.CategoryID,
			&i.CategoryName,
			&i.Description,
			&i.Picture,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProducts = `-- name: ListProducts :many
SELECT product_id, product_name, supplier_id, category_id, quantity_per_unit, unit_price, units_in_stock, units_on_order, reorder_level, discontinued FROM products
ORDER BY product_name
`

func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.SupplierID,
			&i.CategoryID,
			&i.QuantityPerUnit,
			&i.UnitPrice,
			&i.UnitsInStock,
			&i.UnitsOnOrder,
			&i.ReorderLevel,
			&i.Discontinued,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE categories
  set category_name = $2,
  description = $3
WHERE category_id = $1
`

type UpdateCategoryParams struct {
	CategoryID   int16          `db:"category_id" json:"categoryId"`
	CategoryName string         `db:"category_name" json:"categoryName"`
	Description  sql.NullString `db:"description" json:"description"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateCategory, arg.CategoryID, arg.CategoryName, arg.Description)
	return err
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products
  set product_name = $2,
  unit_price = $3
WHERE product_id = $1
`

type UpdateProductParams struct {
	ProductID   int16           `db:"product_id" json:"productId"`
	ProductName string          `db:"product_name" json:"productName"`
	UnitPrice   sql.NullFloat64 `db:"unit_price" json:"unitPrice"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.ExecContext(ctx, updateProduct, arg.ProductID, arg.ProductName, arg.UnitPrice)
	return err
}
