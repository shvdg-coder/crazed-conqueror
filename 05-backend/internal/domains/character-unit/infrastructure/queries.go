package infrastructure

// Names
const (
	TableName = "character_units"

	FieldCharacterId = "character_id"
	FieldUnitId      = "unit_id"
)

// SQL query constants
const (
	CreateTableQuery = `
		CREATE TABLE IF NOT EXISTS ` + TableName + ` (
			` + FieldCharacterId + ` VARCHAR(255) NOT NULL,
			` + FieldUnitId + ` VARCHAR(255) NOT NULL,
			PRIMARY KEY (` + FieldCharacterId + `, ` + FieldUnitId + `)
			CONSTRAINT fk_character FOREIGN KEY (` + FieldCharacterId + `) REFERENCES characters(id) ON DELETE CASCADE,
			CONSTRAINT fk_unit FOREIGN KEY (` + FieldUnitId + `) REFERENCES units(id) ON DELETE CASCADE
		);
	`

	DropTableQuery = `DROP TABLE IF EXISTS ` + TableName + ` CASCADE;`
)
