package infrastructure

// Names
const (
	// Table name
	tableName = "users"

	// Field names
	fieldId          = "id"
	fieldEmail       = "email"
	fieldPassword    = "password"
	fieldDisplayName = "display_name"
	fieldLastLoginAt = "last_login_at"
	fieldCreatedAt   = "created_at"
	fieldUpdatedAt   = "updated_at"
)

// SQL query constants
const (
	// Table management
	createTableQuery = `
		CREATE TABLE IF NOT EXISTS users (
			` + fieldId + ` VARCHAR(255) PRIMARY KEY,
			` + fieldEmail + ` VARCHAR(255) UNIQUE NOT NULL,
			` + fieldPassword + ` VARCHAR(255) NOT NULL,
			` + fieldDisplayName + ` VARCHAR(255) NOT NULL,
			` + fieldLastLoginAt + ` TIMESTAMPTZ,
			` + fieldCreatedAt + ` TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			` + fieldUpdatedAt + ` TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`

	dropTableQuery = `DROP TABLE IF EXISTS users CASCADE;`
)
