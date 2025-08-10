package infrastructure

// Names
const (
	TableName = "character_formations"

	FieldCharacterId = "character_id"
	FieldFormationId = "formation_id"
)

// SQL query constants
const (
	CreateTableQuery = `
		CREATE TABLE IF NOT EXISTS ` + TableName + ` (
			` + FieldCharacterId + ` VARCHAR(255) NOT NULL,
			` + FieldFormationId + ` VARCHAR(255) NOT NULL,
			PRIMARY KEY (` + FieldCharacterId + `, ` + FieldFormationId + `),
			CONSTRAINT fk_character FOREIGN KEY (` + FieldCharacterId + `) REFERENCES characters(id) ON DELETE CASCADE,
			CONSTRAINT fk_formation FOREIGN KEY (` + FieldFormationId + `) REFERENCES formations(id) ON DELETE CASCADE
		);
	`

	DropTableQuery = `DROP TABLE IF EXISTS ` + TableName + ` CASCADE;`
)
