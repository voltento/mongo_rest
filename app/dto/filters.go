package dto

type Filters struct {
	Limit  int64
	Found  *bool
	Number *float64
	Type   *string
}

func NewDefaultFilters() *Filters {
	return &Filters{Limit: 10}
}
