package change

type AddColumn struct {
	CatalogName string
	SchemaName  string
	TableName   string
	Columns     []*Column
}

func NewAddColumn() *AddColumn {
	return &AddColumn{
		CatalogName: "",
		SchemaName:  "",
		TableName:   "",
		Columns:     make([]*Column, 0),
	}
}

type AddColumnChange struct {
	AddColumn *AddColumn
}

func (r *AddColumnChange) Name() string {
	return "addColumn"
}
func (r *AddColumnChange) Value() interface{} {
	return r.AddColumn
}

func NewAddColumnChange(c *AddColumn) Change {
	return &AddColumnChange{
		AddColumn: c,
	}
}
