package domain

// BattlefieldBuilder helps build and configure a Battlefield object.
type BattlefieldBuilder struct {
	rows []BattlefieldRow
}

// NewBattlefieldBuilder initializes a new BattlefieldBuilder.
func NewBattlefieldBuilder() *BattlefieldBuilder {
	return &BattlefieldBuilder{
		rows: make([]BattlefieldRow, 0),
	}
}

// AddRow adds a new row of columns to the battlefield.
func (b *BattlefieldBuilder) AddRow(columns ...BattlefieldColumn) *BattlefieldBuilder {
	row := BattlefieldRow{
		Columns: columns,
	}
	b.rows = append(b.rows, row)
	return b
}

// Build returns the configured Battlefield object.
func (b *BattlefieldBuilder) Build() *Battlefield {
	return &Battlefield{
		Rows: b.rows,
	}
}
