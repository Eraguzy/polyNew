package storage

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBManager struct {
	Pool *pgxpool.Pool
}

func ConnectToDB() (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
}
