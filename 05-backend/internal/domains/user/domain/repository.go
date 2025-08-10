package domain

// UserRepository representation of a user repository
type UserRepository interface {
	GetByEmail(email string) (*UserEntity, error)
	Authenticate(email, password string) (*UserEntity, error)
}
