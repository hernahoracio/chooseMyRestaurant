package repo

import (
	"context"
	"database/sql"
)

func Connect(ctx context.Context) (*sql.DB, error) {
	dsn := ""
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, err
}

type Repo interface {
	CreateRestaurant(ctx context.Context, ID string, Name string, Address string, Latitude float64, Longitude float64, cuisine string, rating float64) (string, error)
	FindRestaurants(ctx context.Context, ID string, Name string, Address string, Latitude float64, Longitude float64, cuisine string, rating float64) ([]string, error)
	FetchRestaurant(ctx context.Context, ID string, Name string, Address string, Latitude float64, Longitude float64, cuisine string, rating float64) (string, error)
}

type Repository struct {
	db *sql.DB
}

func (r *Repository) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, opts)
}
