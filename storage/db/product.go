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

type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *productRepo {

	return &productRepo{
		db: db,
	}

}

func (p *productRepo) Create(req models.CreateUpdateProduct) (string, error) {

	id := uuid.NewString()

	query := `
	INSERT INTO 
		product(id,name,price,barcode,category_id) 
	VALUES($1,$2,$3,$4,$5)`

	_, err := p.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Price,
		req.Barcode,
		req.CategoryId,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "", err
	}

	return id, nil

}

func (p *productRepo) Update(req models.CreateUpdateProduct, id string) (string, error) {

	query := `
	UPDATE 
		product
	SET 
		name=$2,price=$3,barcode=$4,category_id=$5,updated_at=NOW()
	WHERE 
		id=$1`

	resp, err := p.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Price,
		req.Barcode,
		req.CategoryId,
	)
	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}

func (p *productRepo) Get(req models.IdRequest) (models.Product, error) {

	query := `
	SELECT 
		id,
		name,
		price,
		barcode,
		category_id,
		created_at::text,
		updated_at::text
	FROM 
		product 
	WHERE 
		id = $1`

	resp := p.db.QueryRow(context.Background(), query, req.Id)

	var product models.Product

	err := resp.Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Barcode,
		&product.CategoryId,
		&product.Created_at,
		&product.Updated_at,
	)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (p *productRepo) GetAll(req models.GetAllProductRequest) (models.GetAllProduct, error) {

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
		price,
		barcode,
		category_id,
		created_at::text,
		updated_at::text
	FROM 
		product 
	`
	var count int

	c := `
	SELECT
		COUNT(*)
	FROM
		product`

	if req.Name != "" {
		filter += ` AND name ILIKE '%' || @name || '%' `
		params["name"] = req.Name
	}

	if req.Barcode != "" {
		filter += ` AND barcode = '` + req.Barcode + "'"
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf(" OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ
	countQuery := c + filter

	q, pArr := helper.ReplaceQueryParams(query, params)
	rows, err := p.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return models.GetAllProduct{}, err
	}

	err = p.db.QueryRow(context.Background(), countQuery).Scan(&count)

	if err != nil {
		return models.GetAllProduct{}, err
	}

	defer rows.Close()

	result := []models.Product{}

	for rows.Next() {

		var product models.Product

		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Barcode,
			&product.CategoryId,
			&product.Created_at,
			&product.Updated_at,
		)
		if err != nil {
			return models.GetAllProduct{}, err
		}

		result = append(result, product)

	}

	return models.GetAllProduct{Products: result, Count: count}, nil

}

func (p *productRepo) Delete(req models.IdRequest) (string, error) {

	query := `
	DELETE FROM 
		product 
	WHERE 
		id = $1`

	resp, err := p.db.Exec(context.Background(), query, req.Id)

	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}

	return "Deleted suc", nil
}

func (p *productRepo) GetScanProduct(barcode string) (models.Product, error) {

	query := `
	SELECT 
		id,
		name,
		price,
		barcode,
		category_id,
		created_at::text,
		updated_at::text
	FROM 
		product 
	WHERE 
		barcode = $1`

	resp := p.db.QueryRow(context.Background(), query, barcode)

	var product models.Product

	err := resp.Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Barcode,
		&product.CategoryId,
		&product.Created_at,
		&product.Updated_at,
	)

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}
