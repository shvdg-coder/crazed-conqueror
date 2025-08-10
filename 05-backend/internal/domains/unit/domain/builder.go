package domain

import (
	"fmt"
	"shvdg/crazed-conquerer/internal/shared/converters"
	"time"

	fake "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

// UnitEntityBuilder helps build and configure a UnitEntity object.
type UnitEntityBuilder struct {
	unitEntity *UnitEntity
	counter    uint64
}

// GetNumber returns the current number of units.
func (b *UnitEntityBuilder) GetNumber() uint64 {
	return b.counter
}

// NewUnitEntity initializes a new UnitEntityBuilder with empty values.
func NewUnitEntity() *UnitEntityBuilder {
	return &UnitEntityBuilder{unitEntity: &UnitEntity{}}
}

// WithId sets the ID of the unit entity.
func (b *UnitEntityBuilder) WithId(id string) *UnitEntityBuilder {
	b.unitEntity.Id = id
	return b
}

// WithRandomId sets a random ID (UUID) for the unit entity.
func (b *UnitEntityBuilder) WithRandomId() *UnitEntityBuilder {
	b.unitEntity.Id = uuid.New().String()
	return b
}

// WithVocation sets the vocation of the unit entity.
func (b *UnitEntityBuilder) WithVocation(vocation string) *UnitEntityBuilder {
	b.unitEntity.Vocation = vocation
	return b
}

// WithRandomVocation sets a random vocation for the unit entity.
func (b *UnitEntityBuilder) WithRandomVocation() *UnitEntityBuilder {
	vocations := []string{"Warrior", "Mage", "Archer", "Rogue", "Paladin", "Priest", "Necromancer", "Berserker"}
	b.unitEntity.Vocation = fake.RandomString(vocations)
	return b
}

// WithFaction sets the faction of the unit entity.
func (b *UnitEntityBuilder) WithFaction(faction string) *UnitEntityBuilder {
	b.unitEntity.Faction = faction
	return b
}

// WithRandomFaction sets a random faction for the unit entity.
func (b *UnitEntityBuilder) WithRandomFaction() *UnitEntityBuilder {
	factions := []string{"Alliance", "Horde", "Neutral", "Empire", "Republic", "Rebels"}
	b.unitEntity.Faction = fake.RandomString(factions)
	return b
}

// WithName sets the name of the unit entity.
func (b *UnitEntityBuilder) WithName(name string) *UnitEntityBuilder {
	b.unitEntity.Name = name
	return b
}

// WithRandomName sets a random name for the unit entity.
func (b *UnitEntityBuilder) WithRandomName() *UnitEntityBuilder {
	b.unitEntity.Name = fmt.Sprintf("%s_%d", fake.Name(), b.counter)
	return b
}

// WithLevel sets the level of the unit entity.
func (b *UnitEntityBuilder) WithLevel(level string) *UnitEntityBuilder {
	b.unitEntity.Level = level
	return b
}

// WithRandomLevel sets a random level for the unit entity.
func (b *UnitEntityBuilder) WithRandomLevel() *UnitEntityBuilder {
	b.unitEntity.Level = fmt.Sprintf("%d", fake.Number(1, 100))
	return b
}

// WithCreatedAt sets the creation time of the unit entity.
func (b *UnitEntityBuilder) WithCreatedAt(t time.Time) *UnitEntityBuilder {
	b.unitEntity.CreatedAt = converters.TimeToTimestamp(t)
	return b
}

// WithUpdatedAt sets the updated at time of the unit entity.
func (b *UnitEntityBuilder) WithUpdatedAt(t time.Time) *UnitEntityBuilder {
	b.unitEntity.UpdatedAt = converters.TimeToTimestamp(t)
	return b
}

// WithDefaults populates all fields with random default values.
func (b *UnitEntityBuilder) WithDefaults() *UnitEntityBuilder {
	b.counter = NextUnitNumber()

	now := time.Now()
	return b.WithRandomId().
		WithRandomVocation().
		WithRandomFaction().
		WithRandomName().
		WithRandomLevel().
		WithCreatedAt(now).
		WithUpdatedAt(now)
}

// Build returns the configured UnitEntity object.
func (b *UnitEntityBuilder) Build() *UnitEntity {
	return b.unitEntity
}
