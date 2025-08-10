package domain

import (
	"github.com/google/uuid"
)

// CharacterFormationEntityBuilder helps build and configure a CharacterFormationEntity object.
type CharacterFormationEntityBuilder struct {
	characterFormationEntity *CharacterFormationEntity
}

// NewCharacterFormationEntity initializes a new CharacterFormationEntityBuilder with empty values.
func NewCharacterFormationEntity() *CharacterFormationEntityBuilder {
	return &CharacterFormationEntityBuilder{characterFormationEntity: &CharacterFormationEntity{}}
}

// WithCharacterId sets the character ID of the character formation entity.
func (b *CharacterFormationEntityBuilder) WithCharacterId(characterId string) *CharacterFormationEntityBuilder {
	b.characterFormationEntity.CharacterId = characterId
	return b
}

// WithRandomCharacterId sets a random character ID (UUID) for the character formation entity.
func (b *CharacterFormationEntityBuilder) WithRandomCharacterId() *CharacterFormationEntityBuilder {
	b.characterFormationEntity.CharacterId = uuid.New().String()
	return b
}

// WithId sets the id of the character formation entity.
func (b *CharacterFormationEntityBuilder) WithId(id string) *CharacterFormationEntityBuilder {
	b.characterFormationEntity.FormationId = id
	return b
}

// WithRandomId sets a random id for the character formation entity.
func (b *CharacterFormationEntityBuilder) WithRandomId() *CharacterFormationEntityBuilder {
	b.characterFormationEntity.FormationId = uuid.New().String()
	return b
}

// WithDefaults populates all fields with random default values.
func (b *CharacterFormationEntityBuilder) WithDefaults() *CharacterFormationEntityBuilder {
	return b.WithRandomCharacterId().
		WithRandomId()
}

// Build returns the configured CharacterFormationEntity object.
func (b *CharacterFormationEntityBuilder) Build() *CharacterFormationEntity {
	return b.characterFormationEntity
}
