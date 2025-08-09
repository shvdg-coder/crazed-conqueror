package database

import (
	"context"
	"fmt"
	"log"
	"shvdg/crazed-conquerer/internal/shared/contexts"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var maxAttempts = 12
var retryDelay = time.Second * 5

// Service manages database connections and configurations using pgx
type Service struct {
	driverName, dsn string
	*pgxpool.Pool
}

// ServiceOpt configures the Service during initialization
type ServiceOpt func(*Service) error

// NewService initializes a new Service with optional configurations
func NewService(driverName, dsn string, options ...ServiceOpt) (*Service, error) {
	service := &Service{
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
func WithConnection() ServiceOpt {
	return func(s *Service) error {
		return s.Connect()
	}
}

// Connect establishes a database connection using only pgx
func (db *Service) Connect() error {
	if db.driverName != "postgres" {
		return fmt.Errorf("unsupported driver: %s; pgx only supports postgres", db.driverName)
	}

	log.Printf("waiting for database to initialize...")
	time.Sleep(time.Second * 2)

	var lastErr error
	for i := 0; i < maxAttempts; i++ {
		if i > 0 {
			log.Printf("retrying database connection in %s...", retryDelay)
			time.Sleep(retryDelay)
		}

		if pool, err := db.connectAttempt(); err != nil {
			lastErr = err
			log.Print(lastErr)
			continue
		} else {
			db.Pool = pool
			log.Print("successfully connected to database")
			return nil
		}
	}

	return fmt.Errorf("failed to connect to database after %d attempts: %w", maxAttempts, lastErr)
}

// connectAttempt attempts to connect to the database
func (db *Service) connectAttempt() (*pgxpool.Pool, error) {
	pgxConf, err := pgxpool.ParseConfig(db.dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pgx config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), pgxConf)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx pool: %w", err)
	}

	if err = pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}

// Disconnect cleans up resources
func (db *Service) Disconnect() error {
	if db.Pool != nil {
		db.Pool.Close()
	}
	return nil
}

// GetPool returns the pgx pool connection
func (db *Service) GetPool() *pgxpool.Pool {
	return db.Pool
}

// GetExecutor returns either the transaction or connection as the executor for a query.
func (db *Service) GetExecutor(ctx context.Context) (Executor, func(), error) {
	var executor Executor
	var cleanup func()

	if tx := contexts.GetTransaction(ctx); tx != nil {
		executor = tx
		cleanup = func() {}
	} else {
		conn, err := db.GetPool().Acquire(ctx)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to acquire connection: %w", err)
		}
		executor = conn.Conn()
		cleanup = conn.Release
	}

	return executor, cleanup, nil
}
