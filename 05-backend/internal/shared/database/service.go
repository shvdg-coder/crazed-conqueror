package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var maxAttempts = 12
var retryDelay = time.Second * 5

// DatabaseSvc manages database connections and configurations using pgx
type DatabaseSvc struct {
	driverName, dsn string
	*pgxpool.Pool
}

// DatabaseSvcOpt configures the DatabaseSvc during initialization
type DatabaseSvcOpt func(*DatabaseSvc) error

// NewDatabaseSvc initializes a new DatabaseSvc with optional configurations
func NewDatabaseSvc(driverName, dsn string, options ...DatabaseSvcOpt) (*DatabaseSvc, error) {
	service := &DatabaseSvc{
		driverName: driverName,
		dsn:        dsn,
	}

	for _, option := range options {
		if err := option(service); err != nil {
			return nil, err
		}
	}

	return service, nil
}

// WithConnection establishes an initial database connection using pgx
func WithConnection() DatabaseSvcOpt {
	return func(s *DatabaseSvc) error {
		return s.Connect()
	}
}

// Connect establishes a database connection using only pgx
func (d *DatabaseSvc) Connect() error {
	if d.driverName != "postgres" {
		return fmt.Errorf("unsupported driver: %s; pgx only supports postgres", d.driverName)
	}

	var lastErr error
	for i := 0; i < maxAttempts; i++ {
		if i > 0 {
			log.Printf("retrying database connection in %s...", retryDelay)
			time.Sleep(retryDelay)
		}

		pgxConf, err := pgxpool.ParseConfig(d.dsn)

		if err != nil {
			lastErr = fmt.Errorf("failed to parse pgx config: %w", err)
			log.Print(lastErr)
			continue
		}

		pool, err := pgxpool.NewWithConfig(context.Background(), pgxConf)
		if err != nil {
			lastErr = fmt.Errorf("failed to create pgx pool: %w", err)
			log.Print(lastErr)
			continue
		}

		if err = pool.Ping(context.Background()); err != nil {
			lastErr = fmt.Errorf("failed to ping database: %w", err)
			log.Print(lastErr)
			pool.Close()
			continue
		}

		d.Pool = pool
		log.Print("successfully connected to database")
		return nil
	}

	return fmt.Errorf("failed to connect to database after %d attempts: %w", maxAttempts, lastErr)
}

// Disconnect cleans up resources
func (d *DatabaseSvc) Disconnect() error {
	if d.Pool != nil {
		d.Pool.Close()
	}
	return nil
}

// GetPool returns the pgx pool connection
func (d *DatabaseSvc) GetPool() *pgxpool.Pool {
	return d.Pool
}

// ExecuteQueries executes a list of SQL queries
func (d *DatabaseSvc) ExecuteQueries(queries []string) error {
	for _, query := range queries {
		if _, err := d.Pool.Exec(context.Background(), query); err != nil {
			log.Printf("Failed to execute query: %s, error: %v", query, err)
			return err
		}
	}
	return nil
}
