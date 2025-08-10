package infrastructure

// Names
const (
	TableName = "formations"

	FieldId        = "id"
	FieldTiles     = "tiles"
	FieldCreatedAt = "created_at"
	FieldUpdatedAt = "updated_at"
)

// SQL query constants
const (
	CreateTableQuery = `
		CREATE TABLE IF NOT EXISTS ` + TableName + ` (
			` + FieldId + ` VARCHAR(255) PRIMARY KEY,
			` + FieldTiles + ` JSONB NOT NULL DEFAULT '[]'::jsonb,
			` + FieldCreatedAt + ` TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			` + FieldUpdatedAt + ` TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`

	DropTableQuery = `DROP TABLE IF EXISTS ` + TableName + ` CASCADE;`
)
