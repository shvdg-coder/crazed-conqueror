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

// WithId sets the slot of the formation entity.
func (b *FormationEntityBuilder) WithId(slot string) *FormationEntityBuilder {
	b.formationEntity.Id = slot
	return b
}

// WithRandomId sets a random slot for the formation entity.
func (b *FormationEntityBuilder) WithRandomId() *FormationEntityBuilder {
	b.formationEntity.Id = uuid.New().String()
	return b
}

// WithRows sets the rows of the formation entity.
func (b *FormationEntityBuilder) WithRows(rows []*FormationRowEntity) *FormationEntityBuilder {
	b.formationEntity.Rows = rows
	return b
}

// WithEmptyRows sets an empty rows array for the formation entity.
func (b *FormationEntityBuilder) WithEmptyRows() *FormationEntityBuilder {
	b.formationEntity.Rows = []*FormationRowEntity{}
	return b
}

// WithRowsFromJson directly unmarshal FormationRowEntity array from JSON
func (b *FormationEntityBuilder) WithRowsFromJson(rowsJson []byte) *FormationEntityBuilder {
	rows, err := fromRowsJsonToRowsEntity(rowsJson)
	if err != nil {
		b.WithEmptyRows()
	} else {
		b.WithRows(rows)
	}

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
		WithEmptyRows().
		WithCreatedAt(now).
		WithUpdatedAt(now)
}

// Build returns the configured FormationEntity object.
func (b *FormationEntityBuilder) Build() *FormationEntity {
	return b.formationEntity
}
