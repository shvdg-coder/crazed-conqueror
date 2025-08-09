package database

// RowScanner defines the interface for scanning a single row from a database query.
type RowScanner interface {
	Scan(dest ...any) error
}

// ScannerFunc is a function that scans a single row from a query result into a value of type T.
type ScannerFunc[T any] func(scanner RowScanner) (T, error)
