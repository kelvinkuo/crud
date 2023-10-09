package convert

import (
    "github.com/kelvinkuo/crud/db"
    "github.com/kelvinkuo/crud/protocol"
)

type ItemCreator interface {
    ItemCreate(table db.Table, service string, filters []ColumnFilter) (protocol.Item, error)
}

type Converter interface {
    AddColumnFilter(filter ColumnFilter)
    CreateMetaMessage(table db.Table) ([]protocol.Message, error)
    AddItemCreator(creator ItemCreator)
    CreateItems(table db.Table, service string) ([]protocol.Item, error)
}

type ColumnFilter interface {
    FilterOut(col db.Column) bool
}
