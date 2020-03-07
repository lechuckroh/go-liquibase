package change

type AddAutoIncrement struct {
	CatalogName    string
	ColumnDataType string
	ColumnName     string
	DefaultOnNull  bool
	GenerationType string
	IncrementBy    int64
	SchemaName     string
	StartWith      int64
	TableName      string
}

func NewAddAutoIncrement() *AddAutoIncrement {
	return &AddAutoIncrement{
		CatalogName: "",
		ColumnDataType: "int",
		ColumnName: "id",
		DefaultOnNull: false,
		GenerationType: "",
		IncrementBy: 1,
		SchemaName: "",
		StartWith: 1,
		TableName: "",
	}
}

type AddAutoIncrementChange struct {
	AddAutoIncrement *AddAutoIncrement
}

func (r *AddAutoIncrementChange) Name() string {
	return "addAutoIncrement"
}
func (r *AddAutoIncrementChange) Value() interface{} {
	return r.AddAutoIncrement
}

func NewAddAutoIncrementChange(c *AddAutoIncrement) Change {
	return &AddAutoIncrementChange{
		AddAutoIncrement: c,
	}
}
