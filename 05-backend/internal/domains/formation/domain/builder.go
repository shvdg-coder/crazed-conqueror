package domain

import (
	"shvdg/crazed-conquerer/internal/shared/converters"
	"time"

	"github.com/google/uuid"
)

// FormationEntityBuilder helps build and configure a FormationEntity object.
type FormationEntityBuilder struct {
	formationEntity *FormationEntity
}

// NewFormationEntity initializes a new FormationEntityBuilder with empty values.
func NewFormationEntity() *FormationEntityBuilder {
	return &FormationEntityBuilder{formationEntity: &FormationEntity{}}
}

// WithSlot sets the slot of the formation entity.
func (b *FormationEntityBuilder) WithSlot(slot string) *FormationEntityBuilder {
	b.formationEntity.Id = slot
	return b
}

// WithRandomId sets a random slot for the formation entity.
func (b *FormationEntityBuilder) WithRandomId() *FormationEntityBuilder {
	b.formationEntity.Id = uuid.New().String()
	return b
}

// WithTiles sets the tiles of the formation entity.
func (b *FormationEntityBuilder) WithTiles(tiles []*FormationTileEntity) *FormationEntityBuilder {
	b.formationEntity.Tiles = tiles
	return b
}

// WithEmptyTiles sets an empty tiles array for the formation entity.
func (b *FormationEntityBuilder) WithEmptyTiles() *FormationEntityBuilder {
	b.formationEntity.Tiles = []*FormationTileEntity{}
	return b
}

// WithCreatedAt sets the creation time of the formation entity.
func (b *FormationEntityBuilder) WithCreatedAt(t time.Time) *FormationEntityBuilder {
	b.formationEntity.CreatedAt = converters.TimeToTimestamp(t)
	return b
}

// WithUpdatedAt sets the updated at time of the formation entity.
func (b *FormationEntityBuilder) WithUpdatedAt(t time.Time) *FormationEntityBuilder {
	b.formationEntity.UpdatedAt = converters.TimeToTimestamp(t)
	return b
}

// WithDefaults populates all fields with random default values.
func (b *FormationEntityBuilder) WithDefaults() *FormationEntityBuilder {
	now := time.Now()
	return b.WithRandomId().
		WithEmptyTiles().
		WithCreatedAt(now).
		WithUpdatedAt(now)
}

// Build returns the configured FormationEntity object.
func (b *FormationEntityBuilder) Build() *FormationEntity {
	return b.formationEntity
}
