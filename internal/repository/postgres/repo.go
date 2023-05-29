package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
	//genSQL  squirrel.StatementBuilderType
}

func New(db *pgxpool.Pool) *Repository {
	var r = &Repository{
		//genSQL: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		db: db,
	}

	return r
}
