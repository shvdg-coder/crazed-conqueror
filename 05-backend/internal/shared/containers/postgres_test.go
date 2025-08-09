package containers

import (
	"context"
	"shvdg/crazed-conquerer/internal/shared/database"
	"shvdg/crazed-conquerer/internal/shared/environment"
	"shvdg/crazed-conquerer/internal/shared/paths"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/network"
)

var _ = Describe("Postgres Container Tests", Ordered, func() {
	var ctx context.Context
	var net *testcontainers.DockerNetwork

	var config *ContainerConfig
	var postgres *PostgresContainer

	BeforeAll(func() {
		ctx = context.Background()

		var err error
		net, err = network.New(ctx)
		Expect(err).NotTo(HaveOccurred(), "failed to create network")

		config = &ContainerConfig{
			RootDirectory: paths.ResolvePath("05-backend", ""),
			Network:       net.Name,
			EnvFilePath:   ".tst.env",
		}
	})

	It("should start Postgres container successfully", func() {
		var err error
		postgres, err = NewPostgresContainer(ctx, config)
		Expect(err).NotTo(HaveOccurred(), "failed to create postgres container")
		Expect(postgres).NotTo(BeNil(), "failed to get postgres container")

		Expect(postgres.Host).NotTo(BeEmpty(), "failed to get postgres container host")
		Expect(postgres.Port).NotTo(BeEmpty(), "failed to get postgres container port")

		GinkgoWriter.Printf("postgres container is running at %s:%s\n", postgres.Host, postgres.Port)
	})

	Context("When the database is pinged", func() {
		It("should not return with an error", func() {
			dsn := database.CreateDsn(environment.EnvStr(environment.KeyDbUser), environment.EnvStr(environment.KeyDbPassword), environment.EnvStr(environment.KeyDbName), postgres.Host, postgres.Port)
			dbSvc, err := database.NewService(environment.EnvStr(environment.KeyDbDriver), dsn, database.WithConnection(ctx))
			Expect(err).NotTo(HaveOccurred(), "failed to create database service")

			err = dbSvc.Pool.Ping(ctx)
			Expect(err).NotTo(HaveOccurred(), "failed to ping database")
		})
	})

	AfterAll(func() {
		if postgres != nil {
			postgres.Terminate()
		}
		if net != nil {
			_ = net.Remove(ctx)
		}
	})
})
