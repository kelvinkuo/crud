package mysql

import (
    "github.com/kelvinkuo/crud/db"
)

type table struct {
    name    string
    comment string
    cols    []db.Column
}

func newTable(name string) *table {
    return &table{name: name, cols: []db.Column{}}
}

func (t *table) Name() string {
    return t.name
}

func (t *table) Comment() string {
    return t.comment
}

func (t *table) Cols() []db.Column {
    return t.cols
}

func (t *table) addColumn(column db.Column) {
    t.cols = append(t.cols, column)
}
