package db

import (
	"context"
	"fmt"
	"main/config"
	"main/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type store struct {
	db                    *pgxpool.Pool
	branches              *branchRepo
	categories            *categoryRepo
	products              *productRepo
	coming_tables         *comingTableRepo
	coming_table_products *comingProductRepo
	remainings            *remainingRepo
}

func NewStorage(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresDatabase,
		),
	)

	if err != nil {
		fmt.Println("ParseConfig:", err.Error())
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		fmt.Println("ConnectConfig:", err.Error())
		return nil, err
	}
	return &store{
		db: pool,
	}, nil
}

func (s *store) Branch() storage.BranchesI {
	if s.branches == nil {
		s.branches = NewBranchRepo(s.db)
	}
	return s.branches
}

func (s *store) Category() storage.CategoriesI {
	if s.categories == nil {
		s.categories = NewCategoryRepo(s.db)
	}
	return s.categories
}

func (s *store) Product() storage.ProductI {
	if s.products == nil {
		s.products = NewProductRepo(s.db)
	}
	return s.products
}

func (s *store) ComingTable() storage.ComingTableI {
	if s.coming_tables == nil {
		s.coming_tables = NewComingTableRepo(s.db)
	}
	return s.coming_tables
}

func (s *store) ComingTableProduct() storage.ComingTableProductI {
	if s.coming_table_products == nil {
		s.coming_table_products = NewComingProductRepo(s.db)
	}
	return s.coming_table_products
}

func (s *store) Remaining() storage.RemainingI {
	if s.remainings == nil {
		s.remainings = NewRemainingRepo(s.db)
	}
	return s.remainings
}
