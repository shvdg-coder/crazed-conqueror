package infrastructure

// Names
const (
	TableName = "characters"

	FieldId        = "id"
	FieldName      = "name"
	FieldCreatedAt = "created_at"
	FieldUpdatedAt = "updated_at"
)

// SQL query constants
const (
	CreateTableQuery = `
		CREATE TABLE IF NOT EXISTS ` + TableName + ` (
			` + FieldId + ` VARCHAR(255) PRIMARY KEY,
			` + FieldName + ` VARCHAR(255) NOT NULL,
			` + FieldCreatedAt + ` TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			` + FieldUpdatedAt + ` TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`

	DropTableQuery = `DROP TABLE IF EXISTS ` + TableName + ` CASCADE;`
)
