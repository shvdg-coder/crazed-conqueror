package sql

// QueryField represents a database field with optional type casting suffix
type QueryField struct {
	Name   string
	Suffix string
}

// String returns the field name with its suffix for use in SQL queries
func (f QueryField) String() string {
	if f.Suffix == "" {
		return f.Name
	}
	return f.Name + f.Suffix
}

// NewQueryField creates a QueryField with name and optional suffix
func NewQueryField(name string, suffix ...string) QueryField {
	if len(suffix) > 0 {
		return QueryField{Name: name, Suffix: suffix[0]}
	}
	return QueryField{Name: name}
}
