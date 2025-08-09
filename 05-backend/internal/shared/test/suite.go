package helpers

import (
	"context"
	"log"
	"net/http"
	"shvdg/crazed-conquerer/internal/schemas"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/environment"

	"github.com/jackc/pgx/v5"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
)

// TestSuite provides common functionality for integration testing.
type TestSuite struct {
	Context    context.Context
	Network    *testcontainers.DockerNetwork
	HttpClient *http.Client

	Postgres *PostgresContainer
	Server   *ServerContainer

	Database *database.Service
	Schemas  *schemas.Service
}

// SetupTestSuite prepares a suite to be used throughout multiple tests.
func SetupTestSuite() *TestSuite {
	ctx := context.Background()

	net, err := network.New(ctx)
	if err != nil {
		log.Fatalf("failed to create network: %s", err)
	}

	postgres, err := SetupPostgresContainer(ctx, net.Name)
	if err != nil {
		log.Fatalf("failed to setup Postgres container: %s", err.Error())
	}

	server, err := SetupServerContainer(ctx, net.Name)
	if err != nil {
		log.Fatalf("failed to setup API container: %s", err.Error())
	}

	dsn := database.CreateDsn(environment.EnvStr(environment.KeyDbUser), environment.EnvStr(environment.KeyDbPassword), environment.EnvStr(environment.KeyDbName), postgres.Host, postgres.Port)

	db, err := database.NewService(DriverName, dsn, database.WithConnection())
	if err != nil {
		log.Fatalf("failed to create database service: %s", err.Error())
	}

	sch := schemas.NewDefaultService(db)

	return &TestSuite{
		Context:    ctx,
		Network:    net,
		HttpClient: &http.Client{},
		Postgres:   postgres,
		Server:     server,
		Database:   db,
		Schemas:    sch,
	}
}

// Terminate clears up resources.
func (s *TestSuite) Terminate() {
	if s.Server != nil {
		s.Server.Terminate()
	}
	if s.Postgres != nil {
		_ = s.Schemas.DropAllTables(s.Context)
		s.Postgres.Terminate()
	}
	if s.Network != nil {
		_ = s.Network.Remove(s.Context)
	}
}

// StartTransaction starts and returns a new transaction
func (s *TestSuite) StartTransaction() (pgx.Tx, error) {
	return s.Database.GetPool().BeginTx(s.Context, pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadWrite,
		DeferrableMode: pgx.NotDeferrable,
	})
}
