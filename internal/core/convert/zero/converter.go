package zero

import (
    "fmt"
    
    "github.com/kelvinkuo/crud/internal/consts"
    "github.com/kelvinkuo/crud/internal/core/convert"
    "github.com/kelvinkuo/crud/internal/core/tools"
    "github.com/kelvinkuo/crud/internal/db"
    "github.com/kelvinkuo/crud/internal/protocol"
    factory "github.com/kelvinkuo/crud/internal/protocol/protocolfactory"
    "github.com/kelvinkuo/crud/internal/protocol/zero"
)

type Converter struct {
    convert.CommonConverter
}

func NewConverter() *Converter {
    return &Converter{CommonConverter: convert.NewCommonConverter(consts.StyleCamelCase)}
}

func (p *Converter) CreateMetaMessage(table db.Table) ([]protocol.Message, error) {
    msg := factory.NewMessage(consts.ZeroApi, tools.UpperCamelCase(table.Name()))
    msgList := []protocol.Message{msg}
    i := 1
    for _, col := range table.Cols() {
        if convert.FilterOut(col, p.Filters) {
            continue
        }
        tag := ""
        if col.IsPrimary() {
            tag = fmt.Sprintf("json:\"%s,optional\"", tools.LowerUnderline(col.Name()))
        } else {
            tag = fmt.Sprintf("json:\"%s\"", tools.LowerUnderline(col.Name()))
        }
        err := msg.AddField(zero.NewField(tools.UpperCamelCase(col.Name()), ZeroType(col.DataType()), col.Comment(), tag))
        if err != nil {
            return nil, err
        }
        i++
    }
    
    return msgList, nil
}

func ZeroType(t string) string {
    m := map[string]string{
        // "tinyint", "smallint", "int", "mediumint", "bigint":
        "tinyint":   "int",
        "smallint":  "int",
        "int":       "int",
        "mediumint": "int",
        "bigint":    "int64",
        // "float", "decimal", "double"
        "float":   "float64",
        "decimal": "float64",
        "double":  "float64",
        // "char", "varchar", "text", "longtext", "mediumtext", "tinytext"
        "char":       "string",
        "varchar":    "string",
        "text":       "string",
        "longtext":   "string",
        "mediumtext": "string",
        "tinytext":   "string",
        // "blob", "mediumblob", "longblob", "varbinary", "binary"
        "blob":       "[]byte",
        "mediumblob": "[]byte",
        "longblob":   "[]byte",
        "varbinary":  "[]byte",
        "binary":     "[]byte",
        // "bool", "bit"
        "bool": "bool",
        "bit":  "bool",
        // json
        "json": "string",
        // "date", "time", "datetime", "timestamp"
        "date":      "int64",
        "time":      "int64",
        "datetime":  "int64",
        "timestamp": "int64",
        // "enum", "set"
        "enum": "string",
        "set":  "string",
    }
    
    if typ, ok := m[t]; ok {
        return typ
    } else {
        return t
    }
}
