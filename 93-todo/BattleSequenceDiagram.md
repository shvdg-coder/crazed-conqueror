participant API as HTTP API
participant Battle as Battle Domain
participant Formation as Formation Domain
participant Zone as Zone Domain
participant Unit as Unit Domain
participant Combat as Combat Domain
participant Tick as Tick Domain

API calls:
StartBattle(characterId string, formationId string, zoneId string, coordinates Coordinates) (battleId string, error)

Battle Domain publishes events and handles responses:

1. Formation Request
   PublishFormationRequested(formationId string)
   HandleFormationRetrieved(formationId string, unitIds []string)

2. Zone Request  
   PublishZoneRequested(zoneId string)
   HandleZoneRetrieved(zoneId string, difficultyLevel int)

3. Player Units Request
   PublishPlayerUnitsRequested(characterId string, unitIds []string)
   HandlePlayerUnitsRetrieved(unitIds []string)

4. Enemy Units Request
   PublishEnemyUnitsRequested(difficultyLevel int, coordinates Coordinates, zoneId string)
   HandleEnemyUnitsRetrieved(unitIds []string)

5. Combatants Creation Request
   PublishCombatantsCreationRequested(playerUnitIds []string, enemyUnitIds []string)
   HandleCombatantsCreated(combatantIds []string)

6. Tick System Creation Request
   PublishTickSystemCreationRequested(combatantIds []string)
   HandleTickSystemCreated(tickSystemId string)

7. Tick System Start
   PublishTickSystemStartRequested(tickSystemId string)
   HandleTickSystemStarted(tickSystemId string)

BATTLE EXECUTION

Battle waits for tick system to complete:

HandleTickSystemCompleted(tickSystemId string, winner string, tickEvents []TickEvent)

BATTLE COMPLETION

EndBattle(battleId string, winner string, tickEvents []TickEvent) error

DOMAIN RESPONSIBILITIES

Formation Domain: Manages formation data, returns unit IDs
Zone Domain: Manages zone data, returns difficulty levels
Unit Domain: Manages unit data, keeps unit details in memory
Combat Domain: Manages combatant data, keeps combatant details in memory
Tick Domain: Manages tick processing, keeps tick system state in memory

Battle Domain only holds IDs and coordinates the flow between domains