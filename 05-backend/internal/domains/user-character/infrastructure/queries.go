package infrastructure

// Names
const (
	TableName = "user_characters"

	FieldUserID      = "user_id"
	FieldCharacterID = "character_id"
)

// SQL query constants
const (
	CreateTableQuery = `
		CREATE TABLE IF NOT EXISTS ` + TableName + ` (
			` + FieldUserID + ` VARCHAR(255) NOT NULL,
			` + FieldCharacterID + ` VARCHAR(255) NOT NULL,
			PRIMARY KEY (` + FieldUserID + `, ` + FieldCharacterID + `),
			CONSTRAINT fk_user FOREIGN KEY (` + FieldUserID + `) REFERENCES users(id) ON DELETE CASCADE,
			CONSTRAINT fk_character FOREIGN KEY (` + FieldCharacterID + `) REFERENCES characters(id) ON DELETE CASCADE
		);
	`

	DropTableQuery = `DROP TABLE IF EXISTS ` + TableName + ` CASCADE;`
)
