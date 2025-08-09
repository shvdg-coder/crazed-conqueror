package integration

import (
	"log"
	"shvdg/crazed-conquerer/internal/shared/test"
	"shvdg/crazed-conquerer/internal/user/infrastructure"
	"sync"
)

var (
	sharedTestSuite *test.Suite
	setupOnce       sync.Once
	mutex           sync.RWMutex
)

// GetSharedSuite returns the shared test suite instance with infrastructure schemas, creating it if necessary
func GetSharedSuite() *test.Suite {
	setupOnce.Do(func() {
		sharedTestSuite = test.NewTestSuite()
		sharedTestSuite.AddSchema(infrastructure.NewUserSchema(sharedTestSuite.Database))

		if err := sharedTestSuite.CreateAllTables(sharedTestSuite.Context); err != nil {
			log.Fatalf("failed to create infrastructure tables: %s", err.Error())
		}
	})

	return sharedTestSuite
}

// CleanupSharedSuite terminates the shared test suite
func CleanupSharedSuite() {
	mutex.Lock()
	defer mutex.Unlock()

	if sharedTestSuite != nil {
		sharedTestSuite.Terminate()
		sharedTestSuite = nil
	}
}
