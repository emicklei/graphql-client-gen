package tweet

// union FilterValue = ValueFilter | ValuesFilter | RangeFilter

type FilterValue struct {
	ValueFilter  *ValueFilter
	ValuesFilter *ValuesFilter
	RangeFilter  *RangeFilter
}
