// --- CONFIGURATION ---
// Set the name of the new domain you want to create.
const NewDomain = "character"
// ---------------------

Hello Junie,

Based on the `NewDomain` constant defined above, please scaffold a new domain in my project.

Your task is to analyze the existing **`user`** domain as a blueprint and generate a complete set of boilerplate files for the new domain. The generated code must mirror the architectural patterns, file structure, and conceptual components found within `internal/user/`.

**Core Instructions:**

1.  **Discover the Entity:** Infer the new entity's structure (fields, types, etc.) from its protobuf definition, which you can find at the path: `internal/{NewDomain}/domain/{NewDomain}_entity.proto`.

2.  **Replicate the Structure:** Create the standard `domain/`, `application/`, and `infrastructure/` directory layout within `internal/{NewDomain}/`.

3.  **Generate Key Files:** For the new domain, please create all the essential files that you see in the `user` domain. This includes, but is not limited to:
    *   **Domain Layer:** A `builder.go` for the entity, a `counters.go` for test data, and a `repository.go` for the interface definition. The repository must only contain a single `GetByXXX` and then the identifier(s) of that entity, as it is unique per domain.
    *   **Application Layer:** A boilerplate `service.go`.
    *   **Infrastructure Layer:** Files for the repository implementation, database schema (`schema.go`, `queries.go`), a row scanner, and a full test suite (`repository_test.go`, `suite_test.go`).

4.  **Adapt the Content:**
    *   The file contents must be adapted for the new `{NewDomain}` entity, not just a direct copy of the `user` files.
    *   For complex logic like the repository implementation and specific SQL queries, provide clean boilerplate and method stubs. I will implement the business logic later.
    *   For highly reusable components like the builder and test suites, please generate them as completely as possible based on the established pattern.

The primary goal is to produce a consistent, foundational set of files for the new domain that compiles and is ready for me to fill in the specific implementation details.