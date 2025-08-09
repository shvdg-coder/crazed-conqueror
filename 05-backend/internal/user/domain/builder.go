package domain

import (
	"fmt"
	"shvdg/crazed-conquerer/internal/shared/convertors"
	"time"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

// UserEntityBuilder helps build and configure a UserEntity object.
type UserEntityBuilder struct {
	userEntity *UserEntity
	counter    uint64
}

// GetNumber returns the current number of users.
func (b *UserEntityBuilder) GetNumber() uint64 {
	return b.counter
}

// NewUserEntity initializes a new UserEntityBuilder with empty values.
func NewUserEntity() *UserEntityBuilder {
	return &UserEntityBuilder{userEntity: &UserEntity{}}
}

// WithId sets the ID of the user entity.
func (b *UserEntityBuilder) WithId(id string) *UserEntityBuilder {
	b.userEntity.Id = id
	return b
}

// WithRandomId sets a random ID (UUID) for the user entity.
func (b *UserEntityBuilder) WithRandomId() *UserEntityBuilder {
	b.userEntity.Id = uuid.New().String()
	return b
}

// WithEmail sets the email of the user entity.
func (b *UserEntityBuilder) WithEmail(email string) *UserEntityBuilder {
	b.userEntity.Email = email
	return b
}

// WithRandomEmail sets a random email for the user entity.
func (b *UserEntityBuilder) WithRandomEmail() *UserEntityBuilder {
	b.userEntity.Email = fmt.Sprintf("user%d@%s", b.counter, fake.DomainName())
	return b
}

// WithPassword sets the password of the user entity.
func (b *UserEntityBuilder) WithPassword(password string) *UserEntityBuilder {
	b.userEntity.Password = password
	return b
}

// WithRandomPassword sets a random password for the user entity.
func (b *UserEntityBuilder) WithRandomPassword() *UserEntityBuilder {
	b.userEntity.Password = fake.Password(true, true, true, true, true, 8)
	return b
}

// WithDisplayName sets the display name of the user entity.
func (b *UserEntityBuilder) WithDisplayName(name string) *UserEntityBuilder {
	b.userEntity.DisplayName = name
	return b
}

// WithRandomDisplayName sets a random display name for the user entity.
func (b *UserEntityBuilder) WithRandomDisplayName() *UserEntityBuilder {
	b.userEntity.DisplayName = fmt.Sprintf("%s_%d", fake.Username(), b.counter)
	return b
}

// WithLastLoginAt sets the last login time of the user entity.
func (b *UserEntityBuilder) WithLastLoginAt(t time.Time) *UserEntityBuilder {
	b.userEntity.LastLoginAt = convertors.TimeToTimestamp(t)
	return b
}

// WithCreatedAt sets the creation time of the user entity.
func (b *UserEntityBuilder) WithCreatedAt(t time.Time) *UserEntityBuilder {
	b.userEntity.CreatedAt = convertors.TimeToTimestamp(t)
	return b
}

// WithUpdatedAt sets the updated at time of the user entity.
func (b *UserEntityBuilder) WithUpdatedAt(t time.Time) *UserEntityBuilder {
	b.userEntity.UpdatedAt = convertors.TimeToTimestamp(t)
	return b
}

// WithDefaults populates all fields with random default values.
func (b *UserEntityBuilder) WithDefaults() *UserEntityBuilder {
	b.counter = NextUser()

	now := time.Now()
	return b.WithRandomId().
		WithRandomEmail().
		WithRandomPassword().
		WithRandomDisplayName().
		WithLastLoginAt(now).
		WithCreatedAt(now).
		WithUpdatedAt(now)
}

// Build returns the configured UserEntity object.
func (b *UserEntityBuilder) Build() *UserEntity {
	return b.userEntity
}
