package domain

import (
	"github.com/google/uuid"
)

// CharacterUnitEntityBuilder helps build and configure a CharacterUnitEntity object.
type CharacterUnitEntityBuilder struct {
	characterUnitEntity *CharacterUnitEntity
	counter             uint64
}

// GetNumber returns the current number of character units.
func (b *CharacterUnitEntityBuilder) GetNumber() uint64 {
	return b.counter
}

// NewCharacterUnitEntity initializes a new CharacterUnitEntityBuilder with empty values.
func NewCharacterUnitEntity() *CharacterUnitEntityBuilder {
	return &CharacterUnitEntityBuilder{characterUnitEntity: &CharacterUnitEntity{}}
}

// WithCharacterId sets the character ID of the character unit entity.
func (b *CharacterUnitEntityBuilder) WithCharacterId(characterId string) *CharacterUnitEntityBuilder {
	b.characterUnitEntity.CharacterId = characterId
	return b
}

// WithRandomCharacterId sets a random character ID (UUID) for the character unit entity.
func (b *CharacterUnitEntityBuilder) WithRandomCharacterId() *CharacterUnitEntityBuilder {
	b.characterUnitEntity.CharacterId = uuid.New().String()
	return b
}

// WithUnitId sets the unit ID of the character unit entity.
func (b *CharacterUnitEntityBuilder) WithUnitId(unitId string) *CharacterUnitEntityBuilder {
	b.characterUnitEntity.UnitId = unitId
	return b
}

// WithRandomUnitId sets a random unit ID (UUID) for the character unit entity.
func (b *CharacterUnitEntityBuilder) WithRandomUnitId() *CharacterUnitEntityBuilder {
	b.characterUnitEntity.UnitId = uuid.New().String()
	return b
}

// WithDefaults populates all fields with random default values.
func (b *CharacterUnitEntityBuilder) WithDefaults() *CharacterUnitEntityBuilder {
	b.counter = NextCharacterUnitNumber()

	return b.WithRandomCharacterId().
		WithRandomUnitId()
}

// Build returns the configured CharacterUnitEntity object.
func (b *CharacterUnitEntityBuilder) Build() *CharacterUnitEntity {
	return b.characterUnitEntity
}
