package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func NewDatabasePool() (*DB, error) {
	pool, err := pgxpool.New(context.Background(), "postgres://postgres:passwordPG@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}
	db := &DB{Pool: pool}
	res := db.Pool.Ping(context.Background())
	if res != nil {
		panic(res)
	}
	println("Connected to database!")
	return db, nil
}
