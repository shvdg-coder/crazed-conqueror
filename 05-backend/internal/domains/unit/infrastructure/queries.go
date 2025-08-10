package infrastructure

// Names
const (
	TableName = "units"

	FieldId        = "id"
	FieldVocation  = "vocation"
	FieldFaction   = "faction"
	FieldName      = "name"
	FieldLevel     = "level"
	FieldCreatedAt = "created_at"
	FieldUpdatedAt = "updated_at"
)

// SQL query constants
const (
	CreateTableQuery = `
		CREATE TABLE IF NOT EXISTS ` + TableName + ` (
			` + FieldId + ` VARCHAR(255) PRIMARY KEY,
			` + FieldVocation + ` VARCHAR(255) NOT NULL,
			` + FieldFaction + ` VARCHAR(255) NOT NULL,
			` + FieldName + ` VARCHAR(255) NOT NULL,
			` + FieldLevel + ` VARCHAR(255) NOT NULL,
			` + FieldCreatedAt + ` TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			` + FieldUpdatedAt + ` TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`

	DropTableQuery = `DROP TABLE IF EXISTS ` + TableName + ` CASCADE;`
)
