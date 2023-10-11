package pb

import (
	"regexp"
	"strings"

	"github.com/kelvinkuo/crud/internal/consts"
	"github.com/kelvinkuo/crud/internal/core/convert"
	"github.com/kelvinkuo/crud/internal/core/tools"
	"github.com/kelvinkuo/crud/internal/db"
	"github.com/kelvinkuo/crud/internal/protocol"
	"github.com/kelvinkuo/crud/internal/protocol/pb"
	factory "github.com/kelvinkuo/crud/internal/protocol/protocolfactory"
)

type Converter struct {
	convert.CommonConverter
}

func NewConverter(style string) *Converter {
	return &Converter{CommonConverter: convert.NewCommonConverter(style)}
}

func (p *Converter) CreateMetaMessage(table db.Table) ([]protocol.Message, error) {
	msg := factory.NewMessage(consts.ProtoBuf, tools.UpperCamelCase(table.Name()))
	msgList := []protocol.Message{msg}
	i := 1
	for _, col := range table.Cols() {
		if convert.FilterOut(col, p.Filters) {
			continue
		}

		var dataType string
		if col.IsEnum() {
			enumList := regexp.MustCompile(`[enum|set]\((.+?)\)`).FindStringSubmatch(col.ColumnType())
			enumFields := strings.FieldsFunc(enumList[1], func(c rune) bool {
				cs := string(c)
				return "," == cs || "'" == cs
			})
			enumName := tools.UpperCamelCase(table.Name()) + tools.UpperCamelCase(col.Name())
			enum := pb.NewEnum(enumName)
			for fieldIndex, field := range enumFields {
				err := enum.AddField(pb.NewEnumField(strings.ToUpper(field), fieldIndex))
				if err != nil {
					return nil, err
				}
			}
			msgList = append(msgList, enum)
			dataType = enumName
		} else {
			dataType = PbType(col.DataType())
		}

		err := msg.AddField(pb.NewField(StyleString(col.Name(), p.Style), dataType, i, col.Comment(), false))
		if err != nil {
			return nil, err
		}

		i++
	}

	return msgList, nil
}

func PbType(t string) string {
	m := map[string]string{
		// "tinyint", "smallint", "int", "mediumint", "bigint":
		"tinyint":   "int64",
		"smallint":  "int64",
		"int":       "int64",
		"mediumint": "int64",
		"bigint":    "int64",
		// "float", "decimal", "double"
		"float":   "double",
		"decimal": "double",
		"double":  "double",
		// "char", "varchar", "text", "longtext", "mediumtext", "tinytext"
		"char":       "string",
		"varchar":    "string",
		"text":       "string",
		"longtext":   "string",
		"mediumtext": "string",
		"tinytext":   "string",
		// "blob", "mediumblob", "longblob", "varbinary", "binary"
		"blob":       "bytes",
		"mediumblob": "bytes",
		"longblob":   "bytes",
		"varbinary":  "bytes",
		"binary":     "bytes",
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

func StyleString(str, style string) string {
	switch style {
	case consts.StyleCamelCase:
		return tools.LowerCamelCase(str)
	case consts.StyleUnderline:
		return tools.LowerUnderline(str)
	}

	return str
}
