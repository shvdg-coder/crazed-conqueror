package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connection defines the interface for database connections.
type Connection interface {
	Connect(ctx context.Context) error
	Disconnect() error
	GetPool() *pgxpool.Pool
	GetExecutor(ctx context.Context) (Executor, func(), error)
}

// CreateDsn formats a string fit for a dsn, using the provided values.
func CreateDsn(username, password, dbName, host, port string) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		username, password, dbName, host, port)
}
