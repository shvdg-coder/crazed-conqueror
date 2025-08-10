package domain

import "sync/atomic"

var characterCounter uint64 = 1

// NextCharacterNumber generates and returns the next unique characterEntity number.
func NextCharacterNumber() uint64 { return atomic.AddUint64(&characterCounter, 1) }
