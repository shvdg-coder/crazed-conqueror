package database

import "context"

// Repository defines the interface for database repositories.
type Repository[T any] interface {
	Create(ctx context.Context, entities ...T) error
	Update(ctx context.Context, entities ...T) error
	Upsert(ctx context.Context, entities ...T) error
	Delete(ctx context.Context, entities ...T) error

	ReadOne(ctx context.Context, query string, values []any, scan ScannerFunc[T]) (T, error)
	ReadMany(ctx context.Context, query string, values []any, scan ScannerFunc[T]) ([]T, error)
}
