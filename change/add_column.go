package change

type AddColumn interface {
	CatalogName() string
	SchemaName() string
	TableName() string
	Columns() []Column
}

type AddColumnImpl struct {
	catalogName string
	schemaName  string
	tableName   string
	columns     []Column
}

func (r *AddColumnImpl) CatalogName() string {
	return r.catalogName
}
func (r *AddColumnImpl) SchemaName() string {
	return r.schemaName
}
func (r *AddColumnImpl) TableName() string {
	return r.tableName
}
func (r *AddColumnImpl) Columns() []Column {
	return r.columns
}

func NewAddColumn() AddColumn {
	return &AddColumnImpl{
		catalogName: "",
		schemaName:  "",
		tableName:   "",
		columns:     make([]Column, 0),
	}
}

type AddColumnChange struct {
	AddColumn AddColumn
}

func (r *AddColumnChange) Name() string {
	return "addColumn"
}
func (r *AddColumnChange) Value() interface{} {
	return r.AddColumn
}

func NewAddColumnChange(c AddColumn) Change {
	return &AddColumnChange{
		AddColumn: c,
	}
}
