package mysql

type column struct {
	name       string
	comment    string
	dataType   string
	columnType string
	isPrimary  bool
}

func (c *column) Name() string {
	return c.name
}

func (c *column) Comment() string {
	return c.comment
}

func (c *column) DataType() string {
	return c.dataType
}

func (c *column) ColumnType() string {
	return c.columnType
}

func (c *column) IsPrimary() bool {
	return c.isPrimary
}

func (c *column) IsEnum() bool {
	switch c.DataType() {
	case "enum", "set":
		return true
	default:
		return false
	}
}
