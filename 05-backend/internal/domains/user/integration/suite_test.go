package integration

import (
	"fmt"
	"shvdg/crazed-conquerer/internal/shared/testing/shared"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInfrastructure(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Domain Integration Tests")
}

// Executes the first block before and the second block after all the tests are run.
var _ = SynchronizedBeforeSuite(func() []byte {
	return shared.InitializeTestSuite()
}, func(data []byte) {
	if err := shared.SetupLocalSuite(data); err != nil {
		Fail(fmt.Sprintf("Failed to setup local test suite: %v", err))
	}
})

// Executes the first block before and the second block after the teardown.
var _ = SynchronizedAfterSuite(func() {
	// N.A
}, func() {
	shared.CleanupSharedSuite()
})
