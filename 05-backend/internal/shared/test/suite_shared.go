package test

import (
	"log"
	"shvdg/crazed-conquerer/internal/shared/paths"
	"sync"

	"github.com/joho/godotenv"
)

var (
	sharedTestSuite *Suite
	setupOnce       sync.Once
	mutex           sync.RWMutex
)

// GetSharedTestSuite returns the shared test suite instance, creating it if necessary
func GetSharedTestSuite() *Suite {
	setupOnce.Do(func() {
		err := godotenv.Load(paths.ResolvePath(RootDirectory, ".tst.env"))
		if err != nil {
			log.Fatalf("Failed to read .tst.env: %v", err)
		}

		sharedTestSuite = NewTestSuite()
	})
	return sharedTestSuite
}

// CleanupSharedTestSuite terminates the shared test suite
func CleanupSharedTestSuite() {
	mutex.Lock()
	defer mutex.Unlock()

	if sharedTestSuite != nil {
		sharedTestSuite.Terminate()
		sharedTestSuite = nil
	}
}
