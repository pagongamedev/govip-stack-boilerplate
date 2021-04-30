package auth

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // "postgres driver for sqlx"
	auth "github.com/pagongamedev/govip-stack-boilerplate/backend/domain/auth/domain"
)

// NewRepository func
func NewRepository(database interface{}) (auth.Repository, error) {
	sqlxDB := sqlx.NewDb(database.(*sql.DB), "postgres")
	r := repo{db: sqlxDB}
	return &r, nil
}

type repo struct {
	db *sqlx.DB
}
