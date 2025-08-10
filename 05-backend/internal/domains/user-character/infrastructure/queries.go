package infrastructure

// Names
const (
	TableName = "user_characters"

	FieldUserId      = "user_id"
	FieldCharacterId = "character_id"
)

// SQL query constants
const (
	CreateTableQuery = `
		CREATE TABLE IF NOT EXISTS ` + TableName + ` (
			` + FieldUserId + ` VARCHAR(255) NOT NULL,
			` + FieldCharacterId + ` VARCHAR(255) NOT NULL,
			PRIMARY KEY (` + FieldUserId + `, ` + FieldCharacterId + `),
			CONSTRAINT fk_user FOREIGN KEY (` + FieldUserId + `) REFERENCES users(id) ON DELETE CASCADE,
			CONSTRAINT fk_character FOREIGN KEY (` + FieldCharacterId + `) REFERENCES characters(id) ON DELETE CASCADE
		);
	`

	DropTableQuery = `DROP TABLE IF EXISTS ` + TableName + ` CASCADE;`
)
