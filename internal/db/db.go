package db

type DB interface {
	Init(datasource string) error
	Close() error
	AllTableNames() []string
	GetTable(table string) (Table, error)
	GetTables(tables []string) ([]Table, error)
}

type Table interface {
	Name() string
	Comment() string
	Cols() []Column
}

type Column interface {
	Name() string
	Comment() string
	DataType() string
	ColumnType() string
	IsPrimary() bool
	IsEnum() bool
}
