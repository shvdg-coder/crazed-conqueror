### Project Architecture Summary: Crazed Conqueror Backend

#### 1. Core Architectural Strategy

The backend follows a **Domain-Driven Design (DDD)** philosophy implemented using a **Clean Architecture** pattern. The core principle is the **Dependency Rule**: all dependencies must point inwards, towards the `domain` packages.

#### 2. High-Level Directory Structure

*   `crazed-conqueror/`: The project root.
*   `go.mod`: Defines the project as a single Go module.
*   `internal/`: A standard Go directory containing all private application code. The distinct business domains reside directly within this directory.
*   `internal/<domain_name>/`: Each directory at this level represents a single, self-contained business domain (e.g., `user`, `character`).

#### 3. Anatomy of a Single Domain Directory

Each `<domain_name>` directory contains a strict three-layer structure:

*   **`domain/` - The Core (The "Why")**
    *   **Purpose:** This directory holds the heart of the business domain. It contains pure, technology-agnostic business logic.
    *   **What to expect inside:**
        *   Definitions for core domain **Entities** and **Value Objects**.
        *   Business rules and logic encapsulated within the domain models.
        *   **Interfaces** that define contracts for external dependencies required by the domain, such as repositories (these are the "Ports").

*   **`application/` - The Orchestrator (The "How")**
    *   **Purpose:** This directory coordinates the domain models and repositories to execute specific application use cases or workflows that are contained within a single domain.
    *   **What to expect inside:**
        *   **Application Services** that are called by external agents (like an HTTP handler). These services contain the high-level logic for a use case but delegate the core business rules to the domain models.

*   **`infrastructure/` - The Implementation (The "With What")**
    *   **Purpose:** This directory provides the concrete implementations of the interfaces defined in the `domain/` layer. It is the only layer coupled to specific technologies.
    *   **What to expect inside:**
        *   **Repository implementations** that interact with a specific database (e.g., PostgreSQL, Redis).
        *   Clients for other external services (e.g., payment gateways, messaging queues). These are the "Adapters" that connect technology to the domain's ports.

#### 4. Cross-Domain Communication: The Event Bus

*   **Purpose:** To enable communication between different domains without creating direct dependencies. One domain publishes an event, and other domains can react to it, ensuring loose coupling.

*   **Key Directories:**
    *   `internal/shared/events/`: Defines the **contracts**. This is a shared, technology-agnostic package containing the `EventBus` and `Event` interfaces that every part of the application can safely depend on.

#### 5. Use Case Example

This example outlines the entire lifecycle of a single business process: A customer successfully orders a product, which
in turn updates the inventory stock level. This demonstrates both a direct Query for immediate data and an asynchronous
Event/Subscribe model for later reactions.

1. **The Request (Entry Point)**: An HTTP API handler receives a request to order 2 units of "Product X". It creates a
   PlaceOrderCommand object with the product ID and quantity.

2. **Orchestration (ordering/application)**: The handler calls the PlaceOrder method on the Ordering Service. This
   service
   is now in charge of the entire order placement process.

3. **The Synchronous Query**: To proceed, the Ordering Service must know if "Product X" is in stock. It makes a direct,
   synchronous call to the FindByID method on the Inventory Repository interface. It needs this data now, so an event is
   not appropriate.

4. **The Decision**: The repository returns the Product data object. The Ordering Service inspects it and confirms that
   the
   stock level is sufficient.

5. **Domain Logic (ordering/domain)**: The service now calls its own Order aggregate to create the new order, passing
   only
   primitive data (product ID, quantity, current price). The Order aggregate enforces its own business rules and, as
   part
   of its creation, records an OrderPlacedEvent in its internal list of events.

6. **The Pivot - Commit and Publish (ordering/infrastructure)**:
    * The Ordering Service tells its Order Repository to save the new Order.
    * The repository saves the order to the database. Immediately after the transaction commits successfully, it
      retrieves
      the OrderPlacedEvent from the aggregate and publishes it to the Event Bus.
    * The Ordering domain's job is now complete. It has fulfilled its responsibility.

7. **The Asynchronous Subscription (inventory/application)**:
    * Elsewhere, an Inventory Event Listener is subscribed to the Event Bus. It hears the OrderPlacedEvent and
      recognizes
      it as relevant.
    * The listener translates the event into a new, internal DecreaseStockCommand. It then calls the DecreaseStock
      method
      on its own Inventory Service.

8. **The Reaction (inventory/domain)**:
    * The Inventory Service receives the command. It uses its Inventory Repository to load the Product aggregate for
      "Product X".
    * It then calls the DecreaseStock(2) method on the Product aggregate itself. The aggregate updates its internal
      state,
      lowering its stock count.

9. **The Final Commit (inventory/infrastructure)**: The Inventory Service tells its repository to save the updated
   Product.
   The new, lower stock count is persisted to the database. The entire business process is now complete, and the system
   is
   in a new, consistent state.

#### 6. How to Extend This Architecture

*   **To Add a Feature to an Existing Domain:**
    1.  Modify or add business logic and models within the `domain/` directory.
    2.  If storage needs change, update the repository interface in the `domain/` directory.
    3.  Implement any interface changes in the `infrastructure/` directory.
    4.  Expose the new functionality by adding or updating methods in an application service within the `application/` directory.

*   **To Add a New Domain (e.g., "Guilds"):**
    1.  Create a new directory inside `internal/` named after the new domain (e.g., `internal/guild/`).
    2.  Replicate the `domain/`, `application/`, and `infrastructure/` directory structure within it.
    3.  Define the new domain's entities, interfaces, services, and infrastructure implementations following the established pattern.