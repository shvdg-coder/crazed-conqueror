package domain

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
)

// fromRowsJsonToRowsEntity converts a JSON array of formation rows to a slice of FormationRowEntity's.
func fromRowsJsonToRowsEntity(rowsJson []byte) ([]*FormationRowEntity, error) {
	var rowsRaw []json.RawMessage
	if err := json.Unmarshal(rowsJson, &rowsRaw); err != nil {
		return nil, fmt.Errorf("failed to unmarshal formation rows: %w", err)
	}

	rows := make([]*FormationRowEntity, len(rowsRaw))
	for i, rowJson := range rowsRaw {
		row := &FormationRowEntity{}
		if err := protojson.Unmarshal(rowJson, row); err != nil {
			return rows, fmt.Errorf("failed to unmarshal formation row: %w", err)
		}
		rows[i] = row
	}

	return rows, nil
}
