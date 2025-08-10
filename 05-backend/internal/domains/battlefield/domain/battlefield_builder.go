package domain

import "shvdg/crazed-conquerer/internal/shared/types"

// BattlefieldBuilder helps build and configure a Battlefield object.
type BattlefieldBuilder struct {
	tiles [][]BattlefieldTile
}

// NewBattlefieldBuilder initializes a new BattlefieldBuilder.
func NewBattlefieldBuilder() *BattlefieldBuilder {
	return &BattlefieldBuilder{
		tiles: make([][]BattlefieldTile, 0),
	}
}

// AddRow adds a new row of tiles to the battlefield.
func (b *BattlefieldBuilder) AddRow(tiles ...BattlefieldTile) *BattlefieldBuilder {
	b.tiles = append(b.tiles, tiles)
	return b
}

// AddEmptyRow adds a row with the specified number of empty tiles.
func (b *BattlefieldBuilder) AddEmptyRow(columns int) *BattlefieldBuilder {
	row := make([]BattlefieldTile, columns)
	currentRow := int32(len(b.tiles))
	for i := 0; i < columns; i++ {
		coordinates := types.NewCoordinates(currentRow, int32(i))
		row[i] = *NewBattlefieldTile(coordinates, []string{})
	}
	b.tiles = append(b.tiles, row)
	return b
}

// AddEmptyRows adds multiple empty rows with the specified number of columns.
func (b *BattlefieldBuilder) AddEmptyRows(rows, columns int) *BattlefieldBuilder {
	for i := 0; i < rows; i++ {
		b.AddEmptyRow(columns)
	}
	return b
}

// Build returns the configured Battlefield object.
func (b *BattlefieldBuilder) Build() *Battlefield {
	return &Battlefield{
		Tiles: b.tiles,
	}
}
