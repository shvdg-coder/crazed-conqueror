# Development Guidelines

## Architecture Rules
- **Dependency Rule**: All dependencies point inward to domain layer
- **Domain Isolation**: Each domain communicates only via event bus
- **Layer Purity**: Keep domain/ free of external dependencies (no DB, HTTP imports)

## Code Patterns
- **Constructors**: Use `New*` pattern consistently
- **Builders**: Use builder pattern for complex object construction
- **Pointers**: Be consistent with pointer vs value semantics
- **Error Handling**: Create domain-specific errors, use Go 1.13+ error wrapping

## Testing Standards
- **Framework**: Use Ginkgo consistently across all test suites (read 90-documentation/Testing.md)
- **Development**: Use `--focus` flags for targeted testing during development
- **Integration**: Account for container setup time in integration tests

## Domain Modeling
- **Rich Models**: Keep business logic in domain entities, not services
- **Value Objects**: Use value objects for identity-less concepts
- **Immutability**: Apply where business rules require it

## Event Communication
- **Cross-Domain**: Use event bus instead of direct dependencies
- **Contracts**: Keep event definitions in shared/events package

## File Organization
- **Domain Structure**: `domain/` → `application/` → `infrastructure/`
- **Test Location**: Co-locate tests with implementation
- **Shared Code**: Place in `internal/shared/` packages