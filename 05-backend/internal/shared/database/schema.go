package database

import (
	"context"
)

// DomainSchema defines the interface for database schema creation and deletion.
type DomainSchema interface {
	CreateTable(ctx context.Context) error
	DropTable(ctx context.Context) error
}
