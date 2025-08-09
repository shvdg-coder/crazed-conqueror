package containers

// ContainerConfig holds configuration for server container setup.
type ContainerConfig struct {
	RootDirectory  string
	Network        string
	EnvFilePath    string
	DockerfilePath string
}

// GetRootDir returns the root directory.
func (c *ContainerConfig) GetRootDir() string {
	return c.RootDirectory
}

// GetNetwork returns the network name.
func (c *ContainerConfig) GetNetwork() string {
	return c.Network
}

// GetEnvFilePath returns the path to the environment file.
func (c *ContainerConfig) GetEnvFilePath() string {
	return c.EnvFilePath
}

// GetDockerfilePath returns the path to the dockerfile.
func (c *ContainerConfig) GetDockerfilePath() string {
	return c.DockerfilePath
}
