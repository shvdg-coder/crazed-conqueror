package infrastructure

import (
	"encoding/json"
	"fmt"
	"shvdg/crazed-conquerer/internal/domains/formation/domain"
	"shvdg/crazed-conquerer/internal/shared/converters"
	"shvdg/crazed-conquerer/internal/shared/database"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/encoding/protojson"
)

// ScanFormationEntity scans database row data into a FormationEntity
func ScanFormationEntity(scanner database.RowScanner) (*domain.FormationEntity, error) {
	var formation domain.FormationEntity
	var tilesJSON []byte
	var createdAt, updatedAt pgtype.Timestamp

	if err := scanner.Scan(&formation.Id, &tilesJSON, &createdAt, &updatedAt); err != nil {
		return nil, fmt.Errorf("failed to scan formation entity: %w", err)
	}

	formation.Tiles = []*domain.FormationTileEntity{}

	if len(tilesJSON) > 0 {
		var tilesData []json.RawMessage
		if err := json.Unmarshal(tilesJSON, &tilesData); err != nil {
			return nil, fmt.Errorf("failed to unmarshal tiles JSON array: %w", err)
		}

		formation.Tiles = make([]*domain.FormationTileEntity, len(tilesData))
		for i, tileJSON := range tilesData {
			tile := &domain.FormationTileEntity{}
			if err := protojson.Unmarshal(tileJSON, tile); err != nil {
				return nil, fmt.Errorf("failed to unmarshal tile %d using protojson: %w", i, err)
			}
			formation.Tiles[i] = tile
		}
	}

	if createdAt.Valid {
		formation.CreatedAt = converters.TimeToTimestamp(createdAt.Time)
	}
	if updatedAt.Valid {
		formation.UpdatedAt = converters.TimeToTimestamp(updatedAt.Time)
	}

	return &formation, nil
}
