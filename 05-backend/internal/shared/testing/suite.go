package testing

import (
	"context"
	"log"
	"net/http"
	"shvdg/crazed-conquerer/internal/shared/containers"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/environment"
	"shvdg/crazed-conquerer/internal/shared/paths"
	"shvdg/crazed-conquerer/internal/shared/schemas"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
)

// Suite provides common functionality for integration testing.
type Suite struct {
	Context    context.Context
	Network    *testcontainers.DockerNetwork
	HttpClient *http.Client

	Postgres *containers.PostgresContainer
	Server   *containers.ServerContainer

	Database *database.Service
	Schemas  *schemas.Service
}

// init loads environment variables from the .tst.env file.
func init() {
	err := godotenv.Load(paths.ResolvePath(RootDirectory, EnvFilePath))
	if err != nil {
		log.Fatalf("Failed to read .tst.env: %v", err)
	}
}

// NewTestSuite prepares a suite to be used throughout multiple tests.
func NewTestSuite() *Suite {
	ctx := context.Background()

	net, err := network.New(ctx)
	if err != nil {
		log.Fatalf("failed to create network: %s", err)
	}

	config := createDefaultConfig(net.Name)

	postgres, err := containers.NewPostgresContainer(ctx, config)
	if err != nil {
		log.Fatalf("failed to setup Postgres container: %s", err.Error())
	}

	server, err := containers.NewServerContainer(ctx, config)
	if err != nil {
		log.Fatalf("failed to setup server container: %s", err.Error())
	}

	dsn := database.CreateDsn(environment.EnvStr(environment.KeyDbUser), environment.EnvStr(environment.KeyDbPassword), environment.EnvStr(environment.KeyDbName), postgres.Host, postgres.Port)

	db, err := database.NewService(environment.EnvStr(environment.KeyDbDriver), dsn, database.WithConnection(ctx))
	if err != nil {
		log.Fatalf("failed to create database service: %s", err.Error())
	}

	sch := schemas.NewService(db)

	return &Suite{
		Context:    ctx,
		Network:    net,
		HttpClient: &http.Client{},
		Postgres:   postgres,
		Server:     server,
		Database:   db,
		Schemas:    sch,
	}
}

// createDefaultConfig creates a default container configuration.
func createDefaultConfig(network string) *containers.ContainerConfig {
	return &containers.ContainerConfig{
		Network:        network,
		RootDirectory:  RootDirectory,
		EnvFilePath:    EnvFilePath,
		DockerfilePath: DockerfileServer,
	}
}

// AddSchema adds a schema to the suite.
func (s *Suite) AddSchema(schema database.DomainSchema) {
	s.Schemas.AddSchema(schema)
}

// CreateAllTables creates all registered domain tables.
func (s *Suite) CreateAllTables(ctx context.Context) error {
	return s.Schemas.CreateAllTables(ctx)
}

// DropAllTables drops all registered domain tables in reverse order.
func (s *Suite) DropAllTables(ctx context.Context) error {
	return s.Schemas.DropAllTables(ctx)
}

// Terminate clears up resources.
func (s *Suite) Terminate() {
	if s.Server != nil {
		s.Server.Terminate()
	}
	if s.Postgres != nil {
		_ = s.DropAllTables(s.Context)
		s.Postgres.Terminate()
	}
	if s.Network != nil {
		_ = s.Network.Remove(s.Context)
	}
}

// StartTransaction starts and returns a new transaction
func (s *Suite) StartTransaction() (pgx.Tx, error) {
	return s.Database.GetPool().BeginTx(s.Context, pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadWrite,
		DeferrableMode: pgx.NotDeferrable,
	})
}
