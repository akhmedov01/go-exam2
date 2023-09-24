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

type comingProductRepo struct {
	db *pgxpool.Pool
}

func NewComingProductRepo(db *pgxpool.Pool) *comingProductRepo {

	return &comingProductRepo{
		db: db,
	}

}

func (c *comingProductRepo) Create(req models.CreateUpdateComingTableProduct) (string, error) {

	id := uuid.NewString()

	query := `
	INSERT INTO 
		coming_table_product(id,category_id,name,price,barcode,count,total_price,coming_table_id) 
	VALUES($1,$2,$3,$4,$5,$6,$7,$8)`

	_, err := c.db.Exec(context.Background(), query,
		id,
		req.CategoryId,
		req.Name,
		req.Price,
		req.Barcode,
		req.Count,
		req.Count*req.Price,
		req.ComingTableId,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "", err
	}

	return id, nil

}

func (c *comingProductRepo) Update(req models.CreateUpdateComingTableProduct, id string) (string, error) {

	query := `
	UPDATE 
		coming_table_product
	SET 
		name=$2,price=$3,barcode=$4,category_id=$5,count=$6,total_price=$7,coming_table_id=$8,updated_at=NOW()
	WHERE 
		id=$1`

	resp, err := c.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Price,
		req.Barcode,
		req.CategoryId,
		req.Count,
		req.Count*req.Price,
		req.ComingTableId,
	)
	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}

func (c *comingProductRepo) Get(req models.IdRequest) (models.ComingTableProduct, error) {

	query := `
	SELECT 
		id,
		name,
		price,
		barcode,
		category_id,
		count,
		total_price,
		coming_table_id,
		created_at::text,
		updated_at::text
	FROM 
		coming_table_product 
	WHERE 
		id = $1`

	resp := c.db.QueryRow(context.Background(), query, req.Id)

	var comingProduct models.ComingTableProduct

	err := resp.Scan(
		&comingProduct.Id,
		&comingProduct.Name,
		&comingProduct.Price,
		&comingProduct.Barcode,
		&comingProduct.CategoryId,
		&comingProduct.Count,
		&comingProduct.TotalPrice,
		&comingProduct.ComingTableId,
		&comingProduct.Created_at,
		&comingProduct.Updated_at,
	)

	if err != nil {
		return models.ComingTableProduct{}, err
	}

	return comingProduct, nil
}

func (c *comingProductRepo) GetAll(req models.GetAllComingTableProductRequest) (models.GetAllComingTableProduct, error) {

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
		count,
		total_price,
		coming_table_id,
		created_at::text,
		updated_at::text
	FROM 
		coming_table_product
	`
	var count int

	cQ := `
	SELECT
		COUNT(*)
	FROM
		coming_table_product`

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
	countQuery := cQ + filter

	q, pArr := helper.ReplaceQueryParams(query, params)
	rows, err := c.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return models.GetAllComingTableProduct{}, err
	}

	err = c.db.QueryRow(context.Background(), countQuery).Scan(&count)

	if err != nil {
		return models.GetAllComingTableProduct{}, err
	}

	defer rows.Close()

	result := []models.ComingTableProduct{}

	for rows.Next() {

		var comingProduct models.ComingTableProduct

		err := rows.Scan(
			&comingProduct.Id,
			&comingProduct.Name,
			&comingProduct.Price,
			&comingProduct.Barcode,
			&comingProduct.CategoryId,
			&comingProduct.Count,
			&comingProduct.TotalPrice,
			&comingProduct.ComingTableId,
			&comingProduct.Created_at,
			&comingProduct.Updated_at,
		)
		if err != nil {
			return models.GetAllComingTableProduct{}, err
		}

		result = append(result, comingProduct)

	}

	return models.GetAllComingTableProduct{ComingTableProducts: result, Count: count}, nil

}

func (c *comingProductRepo) Delete(req models.IdRequest) (string, error) {

	query := `
	DELETE FROM 
		coming_table_product 
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

func (c *comingProductRepo) GetScanBarcode(req models.GetScanBarcodeRequest) (models.ComingTableProduct, error) {

	query := `
	SELECT 
		id,
		name,
		price,
		barcode,
		category_id,
		count,
		total_price,
		coming_table_id,
		created_at::text,
		updated_at::text
	FROM 
		coming_table_product 
	WHERE 
		barcode=$1 AND coming_table_id=$2`

	resp := c.db.QueryRow(context.Background(), query, req.Barcode, req.ComingTableId)

	var comingProduct models.ComingTableProduct

	err := resp.Scan(
		&comingProduct.Id,
		&comingProduct.Name,
		&comingProduct.Price,
		&comingProduct.Barcode,
		&comingProduct.CategoryId,
		&comingProduct.Count,
		&comingProduct.TotalPrice,
		&comingProduct.ComingTableId,
		&comingProduct.Created_at,
		&comingProduct.Updated_at,
	)

	if err != nil {
		return models.ComingTableProduct{}, err
	}

	return comingProduct, nil
}

func (c *comingProductRepo) ScanUpdate(resp models.ComingTableProduct, count float64) (string, error) {

	query := `
	UPDATE 
		coming_table_product
	SET 
		count=$2,total_price=$3,updated_at=NOW()
	WHERE 
		barcode=$1`

	r, err := c.db.Exec(context.Background(), query,
		resp.Barcode,
		resp.Count+count,
		resp.Price*(resp.Count+count),
	)
	if err != nil {
		return "", err
	}
	if r.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}
