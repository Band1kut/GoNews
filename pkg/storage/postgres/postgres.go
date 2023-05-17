package postgres

import (
	"context"
	"github.com/jackc/pgx/pgxpool"
)

// Store Хранилище данных.
type Store struct {
	db *pgxpool.Pool
}

// New Конструктор объекта хранилища.
func New(dbUri string) (*Store, error) {
	db, err := pgxpool.New(context.Background(), dbUri)
	if err != nil {
		return nil, err
	}
	s := Store{
		db: db,
	}
	return &s, nil
}
