package infrastructure

// SQL query constants
const (
	// Table management
	createTableQuery = `
		CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			display_name VARCHAR(255) NOT NULL,
			last_login_at TIMESTAMPTZ,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
		CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
	`

	dropTableQuery = `DROP TABLE IF EXISTS users CASCADE;`
)
