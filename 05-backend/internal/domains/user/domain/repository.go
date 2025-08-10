package domain

import "context"

// UserRepository representation of a user repository
type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*UserEntity, error)
	Authenticate(ctx context.Context, email, password string) (*UserEntity, error)
}
