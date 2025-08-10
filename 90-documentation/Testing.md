# Ginkgo Testing

Ginkgo is a testing framework for the Go programming language that helps you write expressive and comprehensive tests [[1]](https://onsi.github.io/ginkgo/).

## Running Tests

To run your tests, you can use the `ginkgo` command-line tool. A common way to execute all tests within your project is to run the following command from the root of your project:

```textmate
ginkgo -v ./...
```


*   `ginkgo`: This is the command to invoke the Ginkgo test runner.
*   `-v`: This flag stands for "verbose". It provides more detailed output about the tests that are running.
*   `./...`: This is a recursive path specifier. It tells Ginkgo to look for tests in the current directory (`.`) and all of its subdirectories (`/...`).

If you only want to run tests in a specific directory, you can navigate to that directory and run the command from there. For example, to run the tests in the `user` directory, you would do the following:

```textmate
cd internal/user
ginkgo -v
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