package domain

import (
	"fmt"
	"shvdg/crazed-conquerer/internal/shared/converters"
	"time"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

// CharacterEntityBuilder helps build and configure a CharacterEntity object.
type CharacterEntityBuilder struct {
	characterEntity *CharacterEntity
	counter         uint64
}

// GetNumber returns the current number of characters.
func (b *CharacterEntityBuilder) GetNumber() uint64 {
	return b.counter
}

// NewCharacterEntity initializes a new CharacterEntityBuilder with empty values.
func NewCharacterEntity() *CharacterEntityBuilder {
	return &CharacterEntityBuilder{characterEntity: &CharacterEntity{}}
}

// WithId sets the ID of the characterEntity entity.
func (b *CharacterEntityBuilder) WithId(id string) *CharacterEntityBuilder {
	b.characterEntity.Id = id
	return b
}

// WithRandomId sets a random ID (UUID) for the characterEntity entity.
func (b *CharacterEntityBuilder) WithRandomId() *CharacterEntityBuilder {
	b.characterEntity.Id = uuid.New().String()
	return b
}

// WithName sets the name of the characterEntity entity.
func (b *CharacterEntityBuilder) WithName(name string) *CharacterEntityBuilder {
	b.characterEntity.Name = name
	return b
}

// WithRandomName sets a random name for the characterEntity entity.
func (b *CharacterEntityBuilder) WithRandomName() *CharacterEntityBuilder {
	b.characterEntity.Name = fmt.Sprintf("%s_%d", fake.FirstName(), b.counter)
	return b
}

// WithCreatedAt sets the creation time of the characterEntity entity.
func (b *CharacterEntityBuilder) WithCreatedAt(t time.Time) *CharacterEntityBuilder {
	b.characterEntity.CreatedAt = converters.TimeToTimestamp(t)
	return b
}

// WithUpdatedAt sets the updated at time of the characterEntity entity.
func (b *CharacterEntityBuilder) WithUpdatedAt(t time.Time) *CharacterEntityBuilder {
	b.characterEntity.UpdatedAt = converters.TimeToTimestamp(t)
	return b
}

// WithDefaults populates all fields with random default values.
func (b *CharacterEntityBuilder) WithDefaults() *CharacterEntityBuilder {
	b.counter = NextCharacterNumber()

	now := time.Now()
	return b.WithRandomId().
		WithRandomName().
		WithCreatedAt(now).
		WithUpdatedAt(now)
}

// Build returns the configured CharacterEntity object.
func (b *CharacterEntityBuilder) Build() *CharacterEntity {
	return b.characterEntity
}
