# Ginkgo Testing

Ginkgo is a testing framework for the Go programming language that helps you write expressive and comprehensive tests [[1]](https://onsi.github.io/ginkgo/).

## Running Tests

To run your tests, you can use the `ginkgo` command-line tool. A common way to execute all tests within your project is to run the following command from the root of your project:

```textmate
ginkgo -v -p ./...
```

*   `ginkgo`: This is the command to invoke the Ginkgo test runner.
*   `-v`: This flag stands for "verbose". It provides more detailed output about the tests that are running.
*   `-p`: This flag stands for "parallel". It tells Ginkgo to run tests in parallel.
*   `./...`: This is a recursive path specifier. It tells Ginkgo to look for tests in the current directory (`.`) and all of its subdirectories (`/...`).

If you only want to run tests in a specific directory, you can navigate to that directory and run the command from there. For example, to run the tests in the `user` directory, you would do the following:

```textmate
cd internal/user
ginkgo -v -p
```

## Focusing Tests

You can run a specific subset of tests by using the focus flags. This is useful for concentrating on a specific area of your application during development without running the entire test suite.

### Focus by Test Description

Use the `ginkgo --focus="<text>"` flag to run tests whose descriptions match the provided text (which can be a regular expression).

For example, to run only the tests related to the "User Repository", you can use the following command:

```textmate
ginkgo --focus="User Repository" ./...
```

### Focus by File

Use the `ginkgo --focus-file="<file_path>"` flag to run only the tests in files that match the filter [[1]](https://onsi.github.io/ginkgo/). You can repeat the flag to specify multiple files.

For example, to run only the tests in `repository_test.go`:

```textmate
ginkgo --focus-file="repository_test.go" ./...
```

## A Note on Integration Tests

Please be aware that our integration tests rely on a test suite that spins up external dependencies, such as a database, using containers. The initial setup might take a moment, so please be patient when running these tests, especially for the first time.

## Comprehensive Test Execution

For comprehensive testing with detailed reporting, coverage analysis, and robust execution options, you can use the following command:
```

ginkgo --randomize-all --randomize-suites --fail-on-pending --fail-on-empty --keep-going --cover --coverprofile=ginkgo-test-cover.profile --trace --json-report=ginkgo-test-report.json --poll-progress-after=120s --poll-progress-interval=30s ./...
```
This command includes several important flags:

*   `--randomize-all`: Randomizes the order of specs within each suite and the order of suites
*   `--randomize-suites`: Randomizes the order in which test suites are run
*   `--fail-on-pending`: Causes Ginkgo to fail if there are pending specs
*   `--fail-on-empty`: Causes Ginkgo to fail if a test suite contains no specs
*   `--keep-going`: Continues running other suites even if one suite fails
*   `--cover`: Enables Go's built-in code coverage analysis
*   `--coverprofile=ginkgo-test-cover.profile`: Writes coverage data to the specified file
*   `--trace`: Enables stack trace output for failures
*   `--json-report=ginkgo-test-report.json`: Generates a JSON test report
*   `--poll-progress-after=120s`: Shows progress updates after 120 seconds
*   `--poll-progress-interval=30s`: Shows progress updates every 30 seconds after the initial delay
