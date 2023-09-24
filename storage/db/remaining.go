package db

import (
	"context"
	"fmt"
	"main/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type remainingRepo struct {
	db *pgxpool.Pool
}

func NewRemainingRepo(db *pgxpool.Pool) *remainingRepo {

	return &remainingRepo{
		db: db,
	}

}

func (r *remainingRepo) Create(req models.CreateUpdateRemaining) (string, error) {

	id := uuid.NewString()

	query := `
	INSERT INTO 
		remaining(id,branch_id,category_id,name,price,barcode,count,total_price) 
	VALUES($1,$2,$3,$4,$5,$6,$7,$8)`

	_, err := r.db.Exec(context.Background(), query,
		id,
		req.BranchId,
		req.CategoryId,
		req.Name,
		req.Price,
		req.Barcode,
		req.Count,
		req.Count*req.Price,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "", err
	}

	return id, nil

}

func (r *remainingRepo) Update(req models.CreateUpdateRemaining, id string) (string, error) {

	query := `
	UPDATE 
		remaining
	SET 
		name=$2,price=$3,barcode=$4,category_id=$5,count=$6,total_price=$7,branch_id=$8,updated_at=NOW()
	WHERE 
		id=$1`

	resp, err := r.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Price,
		req.Barcode,
		req.CategoryId,
		req.Count,
		req.Count*req.Price,
		req.BranchId,
	)
	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}

func (r *remainingRepo) Get(req models.IdRequest) (models.Remaining, error) {

	query := `
	SELECT 
		id,
		name,
		price,
		barcode,
		category_id,
		count,
		total_price,
		branch_id,
		created_at::text,
		updated_at::text
	FROM 
		remaining 
	WHERE 
		id = $1`

	resp := r.db.QueryRow(context.Background(), query, req.Id)

	var remaining models.Remaining

	err := resp.Scan(
		&remaining.Id,
		&remaining.Name,
		&remaining.Price,
		&remaining.Barcode,
		&remaining.CategoryId,
		&remaining.Count,
		&remaining.TotalPrice,
		&remaining.BranchId,
		&remaining.Created_at,
		&remaining.Updated_at,
	)

	if err != nil {
		return models.Remaining{}, err
	}

	return remaining, nil
}

func (r *remainingRepo) GetAll(req models.GetAllRemainingRequest) (models.GetAllRemaining, error) {

	var (
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
		branch_id,
		created_at::text,
		updated_at::text
	FROM 
		remaining
	`
	var count int

	countQuery := `
	SELECT
		COUNT(*)
	FROM
		remaining`

	if req.BranchId != "" {
		filter += ` AND branch_id = '` + req.BranchId + "'"
	}

	if req.Barcode != "" {
		filter += ` AND barcode = '` + req.Barcode + "'"
	}

	if req.CategoryId != "" {
		filter += ` AND category_id = '` + req.CategoryId + "'"
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf(" OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return models.GetAllRemaining{}, err
	}

	err = r.db.QueryRow(context.Background(), countQuery).Scan(&count)

	if err != nil {
		return models.GetAllRemaining{}, err
	}

	defer rows.Close()

	result := []models.Remaining{}

	for rows.Next() {

		var remaining models.Remaining

		err := rows.Scan(
			&remaining.Id,
			&remaining.Name,
			&remaining.Price,
			&remaining.Barcode,
			&remaining.CategoryId,
			&remaining.Count,
			&remaining.TotalPrice,
			&remaining.BranchId,
			&remaining.Created_at,
			&remaining.Updated_at,
		)
		if err != nil {
			return models.GetAllRemaining{}, err
		}

		result = append(result, remaining)

	}

	return models.GetAllRemaining{Remainings: result, Count: count}, nil

}

func (r *remainingRepo) Delete(req models.IdRequest) (string, error) {

	query := `
	DELETE FROM 
		remaining 
	WHERE 
		id = $1`

	resp, err := r.db.Exec(context.Background(), query, req.Id)

	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}

	return "Deleted suc", nil
}
