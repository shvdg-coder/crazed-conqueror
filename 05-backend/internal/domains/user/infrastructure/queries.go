package infrastructure

// Names
const (
	TableName = "users"

	FieldId          = "id"
	FieldEmail       = "email"
	FieldPassword    = "password"
	FieldDisplayName = "display_name"
	FieldLastLoginAt = "last_login_at"
	FieldCreatedAt   = "created_at"
	FieldUpdatedAt   = "updated_at"
)

// SQL query constants
const (
	CreateTableQuery = `
		CREATE TABLE IF NOT EXISTS ` + TableName + ` (
			` + FieldId + ` VARCHAR(255) PRIMARY KEY,
			` + FieldEmail + ` VARCHAR(255) UNIQUE NOT NULL,
			` + FieldPassword + ` VARCHAR(255) NOT NULL,
			` + FieldDisplayName + ` VARCHAR(255) NOT NULL,
			` + FieldLastLoginAt + ` TIMESTAMPTZ,
			` + FieldCreatedAt + ` TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			` + FieldUpdatedAt + ` TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`

	DropTableQuery = `DROP TABLE IF EXISTS ` + TableName + ` CASCADE;`
)
