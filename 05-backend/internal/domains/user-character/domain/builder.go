package domain

import (
	"github.com/google/uuid"
)

// UserCharacterEntityBuilder helps build and configure a UserCharacterEntity object.
type UserCharacterEntityBuilder struct {
	userCharacterEntity *UserCharacterEntity
}

// NewUserCharacterEntity initializes a new UserCharacterEntityBuilder with empty values.
func NewUserCharacterEntity() *UserCharacterEntityBuilder {
	return &UserCharacterEntityBuilder{userCharacterEntity: &UserCharacterEntity{}}
}

// WithUserID sets the user ID of the user character entity.
func (b *UserCharacterEntityBuilder) WithUserID(userID string) *UserCharacterEntityBuilder {
	b.userCharacterEntity.UserId = userID
	return b
}

// WithRandomUserID sets a random user ID (UUID) for the user character entity.
func (b *UserCharacterEntityBuilder) WithRandomUserID() *UserCharacterEntityBuilder {
	b.userCharacterEntity.UserId = uuid.New().String()
	return b
}

// WithCharacterID sets the character ID of the user character entity.
func (b *UserCharacterEntityBuilder) WithCharacterID(characterID string) *UserCharacterEntityBuilder {
	b.userCharacterEntity.CharacterId = characterID
	return b
}

// WithRandomCharacterID sets a random character ID (UUID) for the user character entity.
func (b *UserCharacterEntityBuilder) WithRandomCharacterID() *UserCharacterEntityBuilder {
	b.userCharacterEntity.CharacterId = uuid.New().String()
	return b
}

// WithDefaults populates all fields with random default values.
func (b *UserCharacterEntityBuilder) WithDefaults() *UserCharacterEntityBuilder {
	return b.WithRandomUserID().
		WithRandomCharacterID()
}

// Build returns the configured UserCharacterEntity object.
func (b *UserCharacterEntityBuilder) Build() *UserCharacterEntity {
	return b.userCharacterEntity
}
