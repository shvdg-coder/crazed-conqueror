# TODO: Battle System Domain-Driven Design Implementation

## Overview
Implement a tick-based battle system using Domain-Driven Design principles with clean separation of concerns and event-driven communication between domains.

## Domain Architecture

### 1. Battlefield Domain
**Purpose**: Spatial and environmental management

**Responsibilities**:
- Grid creation and layout management
- Terrain properties and effects
- Movement validation and pathfinding
- Line-of-sight calculations
- Spawner positioning and area management
- Battlefield boundaries enforcement
- Environmental hazards and interactive elements

**Key Entities**: Battlefield, Grid, Tile, Terrain, Spawner

**Events Published**:
- TileOccupied, TileVacated, TerrainEffectTriggered, SpawnerActivated

### 2. Unit Domain (Already Exists)
**Purpose**: Unit lifecycle and state management

**Responsibilities**:
- Unit statistics and properties
- Health, mana, and status tracking
- Unit abilities and skills
- Cooldown management
- Equipment and inventory
- Unit death and revival states
- Unit type definitions and behaviors

**Key Entities**: Unit, UnitStats, Ability, Equipment, StatusEffect

**Events Published**:
- UnitCreated, UnitDied, UnitRevived, StatsChanged, AbilityUsed, StatusApplied

### 3. Combat Domain
**Purpose**: Combat mechanics and calculations

**Responsibilities**:
- Damage calculation formulas
- Hit chance and critical hit determination
- Status effect applications and interactions
- Weapon and ability mechanics
- Damage type interactions (physical, magical, elemental)
- Combat result evaluation
- Healing and restoration calculations

**Key Entities**: Attack, Damage, CombatResult, WeaponType, DamageType

**Events Published**:
- AttackExecuted, DamageDealt, CriticalHit, StatusInflicted, HealingApplied

### 4. Battle Domain
**Purpose**: Battle orchestration and tick management

**Responsibilities**:
- Tick system management and sequencing
- Battle state progression and flow control
- Event ordering within ticks (Movement → Ranged → Melee)
- Victory and defeat condition evaluation
- Multi-tick action coordination
- Battle log generation and management
- Turn order and initiative handling
- Battle phases (setup, combat, resolution)
- Enemy spawning timing and triggers

**Key Entities**: Battle, Tick, BattleLog, BattleState, TurnOrder

**Events Published**:
- BattleStarted, TickProcessed, BattleEnded, VictoryAchieved, DefeatOccurred, PhaseChanged

## Event Bus Integration

### Purpose of Event Bus
- Eliminate direct dependencies between domains
- Enable reactive programming patterns
- Support loose coupling and high cohesion
- Allow for easy extension and modification
- Enable cross-domain workflows without tight binding

### Event Flow Examples

**Battle Initialization**:
1. Battle Domain publishes "BattleStarted" event
2. Battlefield Domain reacts: initializes grid and terrain
3. Unit Domain reacts: prepares unit states and positions
4. Combat Domain reacts: initializes combat systems and caches

**Tick Processing**:
1. Battle Domain publishes "TickStarted" event
2. All domains prepare for the tick processing
3. Battle Domain coordinates action resolution
4. Individual domains publish specific events (AttackExecuted, UnitMoved)
5. Battle Domain collects and logs all tick events
6. Battle Domain publishes "TickCompleted" event

**Unit Death Scenario**:
1. Combat Domain calculates lethal damage
2. Combat Domain publishes "DamageDealt" event
3. Unit Domain reacts: updates unit health to zero
4. Unit Domain publishes "UnitDied" event
5. Battlefield Domain reacts: clears unit from tile
6. Battle Domain reacts: checks victory conditions
7. Potentially triggers "BattleEnded" event

## Cross-Domain Workflows

### Battle Simulation Workflow
1. Battle Domain receives battle initiation request
2. Battle Domain orchestrates setup through events
3. Battle Domain enters tick loop
4. Each tick involves querying and coordinating all domains
5. Battle Domain aggregates results into battle log
6. Battle Domain determines battle conclusion
7. Battle Domain publishes final results

### Multi-Tick Action Handling
- Battle Domain tracks ongoing actions across ticks
- Combat Domain defines multi-tick ability requirements
- Unit Domain maintains action state between ticks
- Event bus coordinates state synchronization

## Data Flow Patterns

### Command Flow (Inward)
- External requests enter through Battle Domain
- Battle Domain delegates to appropriate domains
- Each domain handles its specific responsibilities
- Results bubble back to Battle Domain for coordination

### Event Flow (Outward)
- Domains publish events when state changes occur
- Other domains react to relevant events
- Event bus prevents direct coupling
- Battle Domain subscribes to coordination events

## Implementation Priorities

### Phase 1: Core Domain Setup
- Establish domain directory structures
- Define core entities for each domain
- Implement basic event bus infrastructure
- Create fundamental interfaces and contracts

### Phase 2: Basic Battle Flow
- Implement simple tick processing
- Basic unit movement and positioning
- Simple combat calculations
- Victory/defeat condition checking

### Phase 3: Advanced Features
- Multi-tick actions and abilities
- Complex terrain interactions
- Status effects and buffs/debuffs
- Enemy spawning systems

### Phase 4: Optimization and Enhancement
- Battle replay and logging
- Performance optimization for large battles
- Advanced AI and tactical behaviors
- Statistical analysis and reporting

## Event Bus Considerations

### Event Naming Conventions
- Use past tense for completed actions (UnitMoved, DamageDealt)
- Use present tense for state changes (BattleStarting)
- Include domain context in event names
- Maintain consistency across all domains

### Event Payload Design
- Include all necessary data for reactions
- Avoid large payload objects when possible
- Include correlation IDs for related events
- Timestamp all events for audit trails

### Error Handling
- Events should not fail silently
- Implement retry mechanisms for critical events
- Log all event processing for debugging
- Graceful degradation when domains are unavailable

This structure provides a solid foundation for implementing a scalable, maintainable battle system that adheres to Domain-Driven Design principles while leveraging event-driven architecture for loose coupling.