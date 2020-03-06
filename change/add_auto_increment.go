package change

type AddAutoIncrement interface {
	CatalogName() string
	ColumnDataType() string
	ColumnName() string
	DefaultOnNull() bool
	GenerationType() string
	IncrementBy() int64
	SchemaName() string
	StartWith() int64
	TableName() string
}

type AddAutoIncrementImpl struct {
	catalogName    string
	columnDataType string
	columnName     string
	defaultOnNull  bool
	generationType string
	incrementBy    int64
	schemaName     string
	startWith      int64
	tableName      string
}

func (r *AddAutoIncrementImpl) CatalogName() string {
	return r.catalogName
}
func (r *AddAutoIncrementImpl) ColumnDataType() string {
	return r.columnDataType
}
func (r *AddAutoIncrementImpl) ColumnName() string {
	return r.columnName
}
func (r *AddAutoIncrementImpl) DefaultOnNull() bool {
	return r.defaultOnNull
}
func (r *AddAutoIncrementImpl) GenerationType() string {
	return r.generationType
}
func (r *AddAutoIncrementImpl) IncrementBy() int64 {
	return r.incrementBy
}
func (r *AddAutoIncrementImpl) SchemaName() string {
	return r.schemaName
}
func (r *AddAutoIncrementImpl) StartWith() int64 {
	return r.startWith
}
func (r *AddAutoIncrementImpl) TableName() string {
	return r.tableName
}

func NewAddAutoIncrement() AddAutoIncrement {
	return &AddAutoIncrementImpl{
		catalogName: "",
		columnDataType: "int",
		columnName: "id",
		defaultOnNull: false,
		generationType: "ALWAYS",
		incrementBy: 1,
		schemaName: "",
		startWith: 1,
		tableName: "",
	}
}

type AddAutoIncrementChange struct {
	AddAutoIncrement AddAutoIncrement
}

func (r *AddAutoIncrementChange) Name() string {
	return "addAutoIncrement"
}
func (r *AddAutoIncrementChange) Value() interface{} {
	return r.AddAutoIncrement
}

func NewAddAutoIncrementChange(c AddAutoIncrement) Change {
	return &AddAutoIncrementChange{
		AddAutoIncrement: c,
	}
}
