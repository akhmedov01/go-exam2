package db

import (
	"context"
	"fmt"
	"main/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type comingTableRepo struct {
	db *pgxpool.Pool
}

func NewComingTableRepo(db *pgxpool.Pool) *comingTableRepo {

	return &comingTableRepo{
		db: db,
	}

}

func (c *comingTableRepo) Create(req models.CreateUpdateComingTable) (string, error) {

	id := uuid.NewString()

	query := `
	INSERT INTO 
		coming_table(id,coming_id,branch_id,date_time,status) 
	VALUES($1,$2,$3,$4,$5)`

	_, err := c.db.Exec(context.Background(), query,
		id,
		req.ComingId,
		req.BranchId,
		req.DateTime,
		req.Status,
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return "", err
	}

	return id, nil

}

func (c *comingTableRepo) Update(req models.CreateUpdateComingTable, id string) (string, error) {

	query := `
	UPDATE 
		coming_table
	SET 
		coming_id=$2,branch_id=$3,date_time=$4,status=$5,updated_at=NOW()
	WHERE 
		id=$1`

	resp, err := c.db.Exec(context.Background(), query,
		id,
		req.ComingId,
		req.BranchId,
		req.DateTime,
		req.Status,
	)

	if err != nil {
		return "", err
	}
	if resp.RowsAffected() == 0 {
		return "", pgx.ErrNoRows
	}
	return "OK", nil
}

func (c *comingTableRepo) Get(req models.IdRequest) (models.ComingTable, error) {

	query := `
	SELECT 
		id,
		coming_id,
		branch_id,
		date_time::text,
		status,
		created_at::text,
		updated_at::text
	FROM 
		coming_table 
	WHERE 
		id = $1`

	resp := c.db.QueryRow(context.Background(), query, req.Id)

	var coming_table models.ComingTable

	err := resp.Scan(
		&coming_table.Id,
		&coming_table.ComingId,
		&coming_table.BranchId,
		&coming_table.DateTime,
		&coming_table.Status,
		&coming_table.Created_at,
		&coming_table.Updated_at,
	)

	if err != nil {
		return models.ComingTable{}, err
	}

	return coming_table, nil
}

func (c *comingTableRepo) GetAll(req models.GetAllComingTableRequest) (models.GetAllComingTable, error) {

	var (
		filter  = "WHERE true"
		offsetQ = " OFFSET 0 "
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
	)
	s := `
	SELECT 
		id,
		coming_id,
		branch_id,
		date_time::text,
		status,
		created_at::text,
		updated_at::text
	FROM 
		coming_table 
	`
	var count int

	countQuery := `
	SELECT
		COUNT(*)
	FROM
		coming_table`

	if req.ComingId != "" {
		filter += ` AND coming_id = '` + req.ComingId + "' "
	}
	if req.BranchId != "" {
		filter += ` AND branch_id = '` + req.BranchId + "' "
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf(" OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	rows, err := c.db.Query(context.Background(), query)
	if err != nil {
		return models.GetAllComingTable{}, err
	}

	err = c.db.QueryRow(context.Background(), countQuery).Scan(&count)

	if err != nil {
		return models.GetAllComingTable{}, err
	}

	defer rows.Close()

	result := []models.ComingTable{}

	for rows.Next() {

		var coming_table models.ComingTable

		err := rows.Scan(
			&coming_table.Id,
			&coming_table.ComingId,
			&coming_table.BranchId,
			&coming_table.DateTime,
			&coming_table.Status,
			&coming_table.Created_at,
			&coming_table.Updated_at,
		)
		if err != nil {
			return models.GetAllComingTable{}, err
		}

		result = append(result, coming_table)

	}

	return models.GetAllComingTable{ComingTables: result, Count: count}, nil

}

func (c *comingTableRepo) Delete(req models.IdRequest) (string, error) {

	query := `
	DELETE FROM 
		coming_table 
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
