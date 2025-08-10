// --- CONFIGURATION ---
// Set the name of the new domain you want to create.
const NewDomain = "XXX"
// ---------------------

Hello Junie,

Based on the `NewDomain` constant defined above, please scaffold a new domain in my project.

Your task is to analyze the existing **`user`** domain as a blueprint and generate a complete set of boilerplate files for the new domain. The generated code must mirror the architectural patterns, file structure, and conceptual components found within `internal/user/`.

**Core Instructions:**

1.  **Discover the Entity and Base Path:** Infer the new entity's structure from its protobuf definition, located at `internal/domains/{NewDomain}/domain/{NewDomain}_entity.proto`. Within this file, examine the `go_package` option. The directory path specified here will be the foundation for the new domain's structure. For instance, if the option is `go_package = "05-backend/internal/domains/character/domain;domain";`, the base path for scaffolding is `05-backend/internal/domains/character`.

2.  **Replicate the Structure:** Within the base path identified in the prior step, create the standard `domain/`, `application/`, and `infrastructure/` directories. You must create these directories if they don't already exist.

3.  **Generate Key Files:** For the new domain, please create all the essential files that you see in the `user` domain. This includes, but is not limited to:
    *   **Domain Layer:** A `builder.go` for the entity, a `counters.go` for test data, and a `repository.go` for the interface definition. The repository must only contain a single `GetByXXX` and then the identifier(s) of that entity, as it is unique per domain.
    *   **Application Layer:** A boilerplate `service.go`.
    *   **Infrastructure Layer:** Files for the repository implementation, database schema (`schema.go`, `queries.go`), a row scanner, and a full test suite (`repository_test.go`, `suite_test.go`).

4.  **Update Shared Test Suite:** Modify the `internal/shared/testing/shared/suite_shared.go` file. You must add the necessary import for the new domain's infrastructure package and register its schema within the `GetSharedSuite` function. This will ensure the new domain's database tables are automatically created when the test suite runs.

5.  **Adapt the Content:**
    *   The file contents must be adapted for the new `{NewDomain}` entity, not just a direct copy of the `user` files.
    *   For complex logic like the repository implementation and specific SQL queries, provide clean boilerplate and method stubs. I will implement the business logic later.
    *   For highly reusable components like the builder and test suites, please generate them as completely as possible based on the established pattern.

The primary goal is to produce a consistent, foundational set of files for the new domain that compiles and is ready for me to fill in the specific implementation details.