package helpers

import (
	"context"
	"fmt"
	"log"
	"shvdg/crazed-conquerer/internal/shared/environment"
	"shvdg/crazed-conquerer/internal/shared/paths"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/joho/godotenv"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// PostgresContainer wraps a database for testing.
type PostgresContainer struct {
	Container testcontainers.Container

	Host string
	Port string
}

// Terminate stops the container and frees resources.
func (c *PostgresContainer) Terminate() {
	if err := c.Container.Terminate(context.Background()); err != nil {
		log.Printf("failed to terminate container: %v", err)
	}
}

// SetupPostgresContainer starts a PostgreSQL container for testing.
func SetupPostgresContainer(ctx context.Context, network string) (*PostgresContainer, error) {
	req := createPostgresContainerRequest(network)

	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to start posgres container: %w", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get container host: %w", err)
	}

	port, err := container.MappedPort(ctx, nat.Port(environment.EnvStr(environment.KeyDbPort)+"/tcp"))
	if err != nil {
		return nil, fmt.Errorf("failed to get container port: %w", err)
	}

	return &PostgresContainer{
		Container: container,
		Host:      host,
		Port:      port.Port(),
	}, nil
}

// createPostgresContainerRequest creates a request to run a PostgreSQL container.
func createPostgresContainerRequest(network string) testcontainers.GenericContainerRequest {
	err := godotenv.Load(paths.ResolvePath(RootDirectory, ".tst.env"))
	if err != nil {
		panic(fmt.Sprintf("Failed to read .tst.env: %v", err))
	}

	postgresContainer := testcontainers.ContainerRequest{
		Image:          "postgres:17-alpine",
		Networks:       []string{network},
		NetworkAliases: map[string][]string{network: {NetworkAliasDb}},
		Env: map[string]string{
			"POSTGRES_DB":       environment.EnvStr(environment.KeyDbName),
			"POSTGRES_USER":     environment.EnvStr(environment.KeyDbUser),
			"POSTGRES_PASSWORD": environment.EnvStr(environment.KeyDbPassword),
		},
		ExposedPorts: []string{environment.EnvStr(environment.KeyDbPort) + "/tcp"},
		WaitingFor: wait.ForLog("database system is ready to accept connections").
			WithOccurrence(1).
			WithStartupTimeout(1 * time.Minute),
	}

	return testcontainers.GenericContainerRequest{
		ProviderType:     testcontainers.ProviderDocker,
		ContainerRequest: postgresContainer,
		Started:          true,
		Reuse:            false,
	}
}
