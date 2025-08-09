package containers

import (
	"context"
	"shvdg/crazed-conquerer/internal/shared/paths"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
)

var _ = Describe("Server Container Tests", Ordered, func() {
	var ctx context.Context
	var net *testcontainers.DockerNetwork

	var config *ContainerConfig
	var server *ServerContainer

	BeforeAll(func() {
		ctx = context.Background()

		var err error
		net, err = network.New(ctx)
		Expect(err).NotTo(HaveOccurred(), "failed to create network")

		config = &ContainerConfig{
			RootDirectory:  paths.ResolvePath("05-backend", ""),
			Network:        net.Name,
			EnvFilePath:    ".tst.env",
			DockerfilePath: "Dockerfile.tst.server",
		}
	})

	It("should start Server container successfully", func() {
		var err error
		server, err = NewServerContainer(ctx, config)
		Expect(err).NotTo(HaveOccurred(), "failed to create server container")
		Expect(server).NotTo(BeNil(), "failed to get server container")

		Expect(server.Host).NotTo(BeEmpty(), "failed to get server container host")
		Expect(server.Port).NotTo(BeEmpty(), "failed to get server container port")
		Expect(server.Url).NotTo(BeEmpty(), "failed to get server container URL")

		GinkgoWriter.Printf("server container is running at %s\n", server.Url)
	})

	AfterAll(func() {
		if server != nil {
			server.Terminate()
		}
		if net != nil {
			_ = net.Remove(ctx)
		}
	})
})
