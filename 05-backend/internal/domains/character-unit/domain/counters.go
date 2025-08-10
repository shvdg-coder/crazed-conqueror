package domain

import "sync/atomic"

var characterUnitCounter uint64 = 1

// NextCharacterUnitNumber generates and returns the next unique character unit number.
func NextCharacterUnitNumber() uint64 { return atomic.AddUint64(&characterUnitCounter, 1) }
