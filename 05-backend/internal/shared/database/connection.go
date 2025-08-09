package database

import "context"

// Connection defines the interface for database connections.
type Connection interface {
	Connect() error
	Disconnect() error
	GetExecutor(ctx context.Context) (Executor, func(), error)
}
