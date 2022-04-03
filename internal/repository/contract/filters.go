package contract

// FilterBy holds a column key from database table and a value related to that key.
type FilterBy struct {
	Field string
	Value any
}
