package testing

import (
	"sync"
)

var (
	sharedTestSuite *Suite
	setupOnce       sync.Once
	mutex           sync.RWMutex
)

// GetSharedSuite returns the shared test suite instance with infrastructure schemas, creating it if necessary
func GetSharedSuite() *Suite {
	setupOnce.Do(func() {
		sharedTestSuite = NewTestSuite()
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
