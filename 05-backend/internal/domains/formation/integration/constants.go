package integration

// smallRowsJson is a small JSON sample for rows.
var smallRowsJson = []byte(`[
				{
					"columns": [
						{"position_x": 0, "position_y": 0, "unit_id": "unit_1"},
						{"position_x": 1, "position_y": 0, "unit_id": "unit_2"}
					]
				}
			]`)

// mediumRowsJson is a medium JSON sample for rows.
var mediumRowsJson = []byte(`[
				{
					"columns": [
						{"position_x": 0, "position_y": 0, "unit_id": "unit_1"},
						{"position_x": 1, "position_y": 0, "unit_id": "unit_2"}
					]
				},
				{
					"columns": [
						{"position_x": 0, "position_y": 1, "unit_id": "unit_4"},
						{"position_x": 1, "position_y": 1, "unit_id": "unit_5"}
					]
				}
			]`)
