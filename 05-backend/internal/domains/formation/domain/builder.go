package domain

import (
	"encoding/json"
	"shvdg/crazed-conquerer/internal/shared/converters"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
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

// WithRowsFromJSON directly unmarshal FormationRowEntity array from JSON
func (b *FormationEntityBuilder) WithRowsFromJSON(rowsJSON []byte) *FormationEntityBuilder {
	if len(rowsJSON) == 0 {
		return b.WithEmptyRows()
	}

	var rowsData []json.RawMessage
	if err := json.Unmarshal(rowsJSON, &rowsData); err != nil {
		return b.WithEmptyRows()
	}

	rows := make([]*FormationRowEntity, len(rowsData))
	for i, rowJSON := range rowsData {
		row := &FormationRowEntity{}
		if err := protojson.Unmarshal(rowJSON, row); err != nil {
			continue // Skip malformed rows
		}
		rows[i] = row
	}

	b.formationEntity.Rows = rows
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
