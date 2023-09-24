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

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *branchRepo {

	return &branchRepo{
		db: db,
	}

}

func (b *branchRepo) Create(req models.CreateUpdateBranch) (string, error) {

	id := uuid.NewString()

	query := `
	INSERT INTO 
		branch(id,name,address,phone_number) 
	VALUES($1,$2,$3,$4)`

	_, err := b.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Address,
		req.PhoneNumber,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "", err
	}

	return id, nil

}

func (b *branchRepo) Update(req models.CreateUpdateBranch, id string) (string, error) {

	query := `
	UPDATE 
		branch
	SET 
		name=$2,address=$3,phone_number=$4,updated_at=NOW()
	WHERE 
		id=$1`

	resp, err := b.db.Exec(context.Background(), query,
		id,
		req.Name,
		req.Address,
		req.PhoneNumber,
	)
	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}

func (b *branchRepo) Get(req models.IdRequest) (models.Branch, error) {

	query := `
	SELECT 
		id,
		name,
		address,
		phone_number,
		created_at::text,
		updated_at::text
	FROM 
		branch 
	WHERE 
		id = $1`

	resp := b.db.QueryRow(context.Background(), query, req.Id)

	var branch models.Branch

	err := resp.Scan(
		&branch.Id,
		&branch.Name,
		&branch.Address,
		&branch.PhoneNumber,
		&branch.Created_at,
		&branch.Updated_at,
	)

	if err != nil {
		return models.Branch{}, err
	}

	return branch, nil
}

func (b *branchRepo) GetAll(req models.GetAllBranchRequest) (models.GetAllBranch, error) {

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
		address,
		phone_number,
		created_at::text,
		updated_at::text
	FROM 
		branch 
	`
	var count int

	countQuery := `
	SELECT
		COUNT(*)
	FROM
		branch`

	if req.Search != "" {
		filter += ` AND name ILIKE '%' || @name || '%' `
		params["name"] = req.Search
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf(" OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	q, pArr := helper.ReplaceQueryParams(query, params)
	rows, err := b.db.Query(context.Background(), q, pArr...)
	if err != nil {
		return models.GetAllBranch{}, err
	}

	err = b.db.QueryRow(context.Background(), countQuery).Scan(&count)

	if err != nil {
		return models.GetAllBranch{}, err
	}

	defer rows.Close()

	result := []models.Branch{}

	for rows.Next() {

		var branch models.Branch

		err := rows.Scan(
			&branch.Id,
			&branch.Name,
			&branch.Address,
			&branch.PhoneNumber,
			&branch.Created_at,
			&branch.Updated_at,
		)
		if err != nil {
			return models.GetAllBranch{}, err
		}

		result = append(result, branch)

	}

	return models.GetAllBranch{Branches: result, Count: count}, nil

}

func (b *branchRepo) Delete(req models.IdRequest) (string, error) {

	query := `
	DELETE FROM 
		branch 
	WHERE 
		id = $1`

	resp, err := b.db.Exec(context.Background(), query, req.Id)

	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}

	return "Deleted suc", nil
}
