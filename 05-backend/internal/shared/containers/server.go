package containers

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

// NewServerContainer starts the API Server container for testing.
func NewServerContainer(ctx context.Context, config *ContainerConfig) (*ServerContainer, error) {
	req := createServerContainerRequest(config)

	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to start server container: %w", err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get container host: %w", err)
	}

	port, err := container.MappedPort(ctx, nat.Port(environment.EnvStr(environment.KeyApiPort)))
	if err != nil {
		return nil, fmt.Errorf("failed to get container port: %w", err)
	}

	url := fmt.Sprintf("http://%s:%s", host, port.Port())

	return &ServerContainer{
		Container: container,
		Url:       url,
		Host:      host,
		Port:      port.Port(),
	}, nil
}

// createServerContainerRequest creates a request to run the API Server container.
func createServerContainerRequest(config *ContainerConfig) testcontainers.GenericContainerRequest {
	envMap, err := godotenv.Read(paths.ResolvePath(config.GetRootDir(), config.GetEnvFilePath()))
	if err != nil {
		panic(fmt.Sprintf("Failed to read %s: %v", config.GetEnvFilePath(), err))
	}

	apiContainer := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Dockerfile: config.GetDockerfilePath(),
			Context:    paths.ResolvePath(config.GetRootDir(), ""),
		},
		Name:           "test-server",
		Networks:       []string{config.GetNetwork()},
		NetworkAliases: map[string][]string{config.GetNetwork(): {NetworkAliasApi}},
		ExposedPorts:   []string{environment.EnvStr(environment.KeyApiPort) + "/tcp"},
		Env:            envMap,
		WaitingFor: wait.ForLog("http server started on").
			WithStartupTimeout(1 * time.Minute),
	}

	return testcontainers.GenericContainerRequest{
		ProviderType:     testcontainers.ProviderDocker,
		ContainerRequest: apiContainer,
		Started:          true,
		Reuse:            true,
	}
}
