package test

import (
	"context"
	"log"
	"net/http"
	"shvdg/crazed-conquerer/internal/schemas"
	"shvdg/crazed-conquerer/internal/shared/containers"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/environment"

	"github.com/jackc/pgx/v5"
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

// NewTestSuite prepares a suite to be used throughout multiple tests.
func NewTestSuite() *Suite {
	ctx := context.Background()

	net, err := network.New(ctx)
	if err != nil {
		log.Fatalf("failed to create network: %s", err)
	}

	// Create container configuration
	config := &containers.ContainerConfig{
		RootDirectory:  RootDirectory,
		Network:        net.Name,
		EnvFilePath:    EnvFile,
		DockerfilePath: DockerfileServer,
	}

	postgres, err := containers.NewPostgresContainer(ctx, config)
	if err != nil {
		log.Fatalf("failed to setup Postgres container: %s", err.Error())
	}

	server, err := containers.NewServerContainer(ctx, config)
	if err != nil {
		log.Fatalf("failed to setup API container: %s", err.Error())
	}

	dsn := database.CreateDsn(environment.EnvStr(environment.KeyDbUser), environment.EnvStr(environment.KeyDbPassword), environment.EnvStr(environment.KeyDbName), postgres.Host, postgres.Port)

	db, err := database.NewService(containers.DriverName, dsn, database.WithConnection(ctx))
	if err != nil {
		log.Fatalf("failed to create database service: %s", err.Error())
	}

	sch := schemas.NewDefaultService(db)
	err = sch.CreateAllTables(ctx)
	if err != nil {
		log.Fatalf("failed to create tables: %s", err.Error())
	}

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

// Terminate clears up resources.
func (s *Suite) Terminate() {
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
func (s *Suite) StartTransaction() (pgx.Tx, error) {
	return s.Database.GetPool().BeginTx(s.Context, pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadWrite,
		DeferrableMode: pgx.NotDeferrable,
	})
}
