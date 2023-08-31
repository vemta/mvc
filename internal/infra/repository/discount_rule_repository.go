package repository

import (
	"database/sql"

	"github.com/vemta/mvc/internal/infra/db"
)

type DiscountRepository struct {
	dbConn *sql.DB
	*db.Queries
	Repository
}

func NewDiscountRepository(dbConn *sql.DB) *DiscountRepository {
	return &DiscountRepository{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}
