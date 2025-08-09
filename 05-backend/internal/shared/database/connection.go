package database

import "context"

// DatabaseConn defines the interface for database connections.
type DatabaseConn interface {
	Connect() error
	Disconnect() error
	GetExecutor(ctx context.Context) (DatabaseExec, func(), error)
}
