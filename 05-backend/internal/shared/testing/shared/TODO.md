**Prompt for Junie:**

Refactor the test suite infrastructure to use a single global database container with proper error handling and resource management for parallel test execution.

**Requirements:**

1. **Modify `suite_shared.go`** with robust cross-process coordination:
    - Add `InitializeGlobalSuite() []byte` with error handling - return error-prefixed bytes if setup fails
    - Add `SetupLocalSuite(data []byte) error` with connection validation and error return
    - Add `CleanupGlobalSuite()` with proper resource cleanup and defer statements
    - Add logging with process identification for debugging parallel execution
    - Handle container startup failures gracefully

2. **Update all `suite_test.go` files** with proper error handling:
    - `SynchronizedBeforeSuite` first function: handle returned errors from `InitializeGlobalSuite()`
    - `SynchronizedBeforeSuite` second function: check for errors in `SetupLocalSuite()` and fail fast with descriptive messages
    - Add consistent suite naming for better parallel debugging

3. **Add resilience features:**
    - Connection retry logic in `SetupLocalSuite()`
    - Proper signal handling for interrupted tests
    - Validate shared container is accessible before proceeding
    - Resource cleanup even on test failures

**Expected outcome:**
- Robust parallel test execution with proper error propagation
- Clean resource management even during failures
- Better debugging information for parallel test issues
- Maintained performance benefits with added reliability

This approach follows Ginkgo's best practices [[1]](https://onsi.github.io/ginkgo/) while adding the reliability needed for production test suites.