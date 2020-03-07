package change

type Column struct {
	Name string
	Type string
	Value string
	Computed bool
	ValueNumeric *float64
	ValueBoolean *bool
	ValueDate string
	ValueComputed string
	ValueBlobFile string
	ValueClobFile string
	Encoding string
	DefaultValue string
	DefaultValueNumeric *float64
	DefaultValueBoolean *bool
	DefaultValueDate string
	DefaultValueComputed string
	AutoIncrement bool
	StartWith int64
	IncrementBy int64
	Remarks string
	BeforeColumn string
	AfterColumn string
	Position *int64
	Descending bool
	Constraints *Constraints
}

func NewColumn() *Column {
	return &Column{}
}

type Constraints struct {
	Nullable *bool
	NotNullConstraintName string
	PrimaryKey bool
	PrimaryKeyName string
	PrimaryKeyTablespace string
	Unique bool
	UniqueConstraintName string
	References string
	ReferencedTableCatalogName string
	ReferencedTableSchemaName string
	ReferencedTableName string
	ReferencedColumnNames string
	ForeignKeyName string
	DeleteCascade bool
	Deferrable bool
	InitiallyDeferred bool
	ValidateNullable bool
	ValidateUnique bool
	ValidatePrimaryKey bool
	ValidateForeignKey bool
	CheckConstraint bool
}

func NewConstraints() *Constraints {
	return &Constraints{}
}