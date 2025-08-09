package helpers

import "sync"

var (
	sharedTestSuite *TestSuite
	setupOnce       sync.Once
	mutex           sync.RWMutex
)

// GetSharedTestSuite returns the shared test suite instance, creating it if necessary
func GetSharedTestSuite() *TestSuite {
	setupOnce.Do(func() {
		sharedTestSuite = SetupTestSuite()
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
