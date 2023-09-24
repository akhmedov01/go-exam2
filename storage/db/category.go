package db

import (
	"context"
	"fmt"
	"main/models"
	"main/packages/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type categoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *categoryRepo {

	return &categoryRepo{
		db: db,
	}

}

func (c *categoryRepo) Create(req models.CreateUpdateCategory) (string, error) {

	id := uuid.NewString()

	query := `
	INSERT INTO 
		category(id,name,parent_id) 
	VALUES($1,$2,$3)`

	_, err := c.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.ParentId,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "", err
	}

	return id, nil

}

func (c *categoryRepo) Update(req models.CreateUpdateCategory, id string) (string, error) {

	query := `
	UPDATE 
		category
	SET 
		name=$2,parent_id=$3,updated_at=NOW()
	WHERE 
		id=$1`

	resp, err := c.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.ParentId,
	)

	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}

func (c *categoryRepo) Get(req models.IdRequest) (models.Category, error) {

	query := `
	SELECT 
		id,
		name,
		parent_id,
		created_at::text,
		updated_at::text
	FROM 
		category 
	WHERE 
		id = $1`

	resp := c.db.QueryRow(context.Background(), query, req.Id)

	var category models.Category

	err := resp.Scan(
		&category.Id,
		&category.Name,
		&category.ParentId,
		&category.Created_at,
		&category.Updated_at,
	)

	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (c *categoryRepo) GetAll(req models.GetAllCategoryRequest) (models.GetAllCategory, error) {

	var (
		params  = make(map[string]interface{})
		filter  = "WHERE true"
		offsetQ = " OFFSET 0 "
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
	)
	s := `
	SELECT 
		id,
		name,
		parent_id,
		created_at::text,
		updated_at::text
	FROM 
		category 
	`
	var count int

	countQuery := `
	SELECT
		COUNT(*)
	FROM
		category`

	if req.Name != "" {
		filter += ` AND name ILIKE '%' || @name || '%' `
		params["name"] = req.Name
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf(" OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	q, pArr := helper.ReplaceQueryParams(query, params)
	rows, err := c.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return models.GetAllCategory{}, err
	}

	err = c.db.QueryRow(context.Background(), countQuery).Scan(&count)

	if err != nil {
		return models.GetAllCategory{}, err
	}

	defer rows.Close()

	result := []models.Category{}

	for rows.Next() {

		var category models.Category

		err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.ParentId,
			&category.Created_at,
			&category.Updated_at,
		)
		if err != nil {
			return models.GetAllCategory{}, err
		}

		result = append(result, category)

	}

	return models.GetAllCategory{Categories: result, Count: count}, nil

}

func (c *categoryRepo) Delete(req models.IdRequest) (string, error) {

	query := `
	DELETE FROM 
		category 
	WHERE 
		id = $1`

	resp, err := c.db.Exec(context.Background(), query, req.Id)

	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}

	return "Deleted suc", nil
}
