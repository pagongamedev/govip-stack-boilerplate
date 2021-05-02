package income

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // "postgres driver for sqlx"
	income "github.com/pagongamedev/govip-stack-boilerplate/backend/domain/income/domain"
)

// NewRepository func
func NewRepository(database interface{}) income.Repository {
	sqlxDB := sqlx.NewDb(database.(*sql.DB), "postgres")
	r := repo{db: sqlxDB}
	return &r
}

type repo struct {
	db *sqlx.DB
}
