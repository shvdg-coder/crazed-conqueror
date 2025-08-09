package helpers

import (
	"context"
	"fmt"
	"shvdg/crazed-conquerer/internal/shared/environment"
	"shvdg/crazed-conquerer/internal/shared/paths"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/joho/godotenv"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// ServerContainer wraps an API Server for testing.
type ServerContainer struct {
	Container testcontainers.Container
	Url       string
	Host      string
	Port      string
}

// Terminate stops the container and frees resources.
func (c *ServerContainer) Terminate() {
	if err := c.Container.Terminate(context.Background()); err != nil {
		fmt.Printf("failed to terminate api container: %v\n", err)
	}
}

// SetupServerContainer starts the API Server container for testing.
func SetupServerContainer(ctx context.Context, network string) (*ServerContainer, error) {
	req := createServerContainerRequest(network)

	apiContainer, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to start API container: %w", err)
	}

	host, err := apiContainer.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get container host: %w", err)
	}

	port, err := apiContainer.MappedPort(ctx, nat.Port(environment.EnvStr(environment.KeyApiPort)))
	if err != nil {
		return nil, fmt.Errorf("failed to get container port: %w", err)
	}

	url := fmt.Sprintf("http://%s:%s", host, port.Port())

	return &ServerContainer{
		Container: apiContainer,
		Url:       url,
		Host:      host,
		Port:      port.Port(),
	}, nil
}

// createServerContainerRequest creates a request to run the API Server container.
func createServerContainerRequest(network string) testcontainers.GenericContainerRequest {
	envMap, err := godotenv.Read(paths.ResolvePath(RootDirectory, ".tst.env"))
	if err != nil {
		panic(fmt.Sprintf("Failed to read .tst.env: %v", err))
	}

	apiContainer := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Dockerfile: "Dockerfile.tst.Server",
			Context:    paths.ResolvePath(RootDirectory, ""),
		},
		Networks:       []string{network},
		NetworkAliases: map[string][]string{network: {NetworkAliasApi}},
		ExposedPorts:   []string{environment.EnvStr(environment.KeyApiPort) + "/tcp"},
		Env:            envMap,
		WaitingFor: wait.ForLog("http Server started on").
			WithStartupTimeout(1 * time.Minute),
	}

	return testcontainers.GenericContainerRequest{
		ProviderType:     testcontainers.ProviderDocker,
		ContainerRequest: apiContainer,
		Started:          true,
		Reuse:            false,
	}
}
