package change

type AddDefaultValue struct {
	CatalogName string
	ColumnDataType string
	ColumnName string
	DefaultValue string
	DefaultValueBoolean *bool
	DefaultValueComputed string
	DefaultValueConstraintName string
	DefaultValueDate string
	DefaultValueNumeric *float64
	DefaultValueSequenceNext string
	SchemaName string
	TableName string
}

func NewAddDefaultValue() *AddDefaultValue {
	return &AddDefaultValue{
		CatalogName: "",
		ColumnDataType: "",
		ColumnName: "",
		DefaultValue: "",
		DefaultValueBoolean: nil,
		DefaultValueComputed: "",
		DefaultValueConstraintName: "",
		DefaultValueDate: "",
		DefaultValueNumeric: nil,
		DefaultValueSequenceNext: "",
		SchemaName: "",
		TableName: "",
	}
}

type AddDefaultValueChange struct {
	AddDefaultValue *AddDefaultValue
}

func (r *AddDefaultValueChange) Name() string {
	return "addDefaultValue"
}
func (r *AddDefaultValueChange) Value() interface{} {
	return r.AddDefaultValue
}

func NewAddDefaultValueChange(c *AddDefaultValue) Change {
	return &AddDefaultValueChange{
		AddDefaultValue: c,
	}
}
