package database

import "context"

// Repository defines the interface for database repositories.
type Repository[T any] interface {
	CreateOne(context.Context, T) error
	ReadOne(context.Context, string, []string, ScannerFunc[T]) (T, error)
	UpdateOne(context.Context, T) error
	DeleteOne(context.Context, T) error

	CreateMany(context.Context, []T) error
	ReadMany(context.Context, string, []string, ScannerFunc[T]) (T, error)
	UpdateMany(context.Context, []T) error
	DeleteMany(context.Context, []T) error

	Count(context.Context, string, []string) (int, error)

	CreateTable(context.Context) error
	DropTable(context.Context) error
}
