package pb

import (
    "fmt"
    "regexp"
    "strings"
    
    "github.com/kelvinkuo/crud/consts"
    "github.com/kelvinkuo/crud/core/convert"
    "github.com/kelvinkuo/crud/core/tools"
    "github.com/kelvinkuo/crud/db"
    "github.com/kelvinkuo/crud/protocol"
    "github.com/kelvinkuo/crud/protocol/pb"
    factory "github.com/kelvinkuo/crud/protocol/protocolfactory"
)

type Converter struct {
    creators []convert.ItemCreator
    filters  []convert.ColumnFilter
}

func NewPbConverter() *Converter {
    return &Converter{
        creators: []convert.ItemCreator{},
    }
}

func (p *Converter) CreateMetaMessage(table db.Table) ([]protocol.Message, error) {
    msg := factory.NewMessage(consts.PROTOBUF, tools.UpperCamelCase(table.Name()))
    msgList := []protocol.Message{msg}
    i := 1
    for _, col := range table.Cols() {
        if filterOut(col, p.filters) {
            continue
        }
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
            err := msg.AddField(pb.NewField(tools.LowerCamelCase(col.Name()), enumName, i, col.Comment(), false))
            if err != nil {
                return nil, err
            }
        } else {
            err := msg.AddField(pb.NewField(tools.LowerCamelCase(col.Name()), pbType(col.DataType()), i, col.Comment(), false))
            if err != nil {
                return nil, err
            }
        }
        i++
    }
    
    return msgList, nil
}

func (p *Converter) AddItemCreator(creator convert.ItemCreator) {
    p.creators = append(p.creators, creator)
}

// func (p *Converter) filterOut(col db.Column) bool {
// 	for _, filter := range p.filters {
// 		if filter.FilterOut(col) {
// 			return true
// 		}
// 	}
//
// 	return false
// }

func (p *Converter) CreateItems(table db.Table, service string) ([]protocol.Item, error) {
    var items []protocol.Item
    for _, creator := range p.creators {
        item, err := creator.ItemCreate(table, service, p.filters)
        if err != nil {
            return nil, err
        }
        items = append(items, item)
    }
    
    return items, nil
}

func (p *Converter) AddColumnFilter(filter convert.ColumnFilter) {
    p.filters = append(p.filters, filter)
}

type AddItemCreator struct {
}

func (c *AddItemCreator) ItemCreate(table db.Table, service string, filters []convert.ColumnFilter) (protocol.Item, error) {
    req := factory.NewMessage(consts.PROTOBUF, fmt.Sprintf("%sAddReq", tools.UpperCamelCase(table.Name())))
    i := 1
    for _, col := range table.Cols() {
        if filterOut(col, filters) {
            continue
        }
        if col.IsPrimary() {
            continue
        }
        if col.IsEnum() {
            err := req.AddField(pb.NewField(
                tools.LowerCamelCase(col.Name()),
                tools.UpperCamelCase(table.Name())+tools.UpperCamelCase(col.Name()),
                i,
                col.Comment(),
                false))
            i++
            if err != nil {
                return nil, err
            }
            continue
        }
        
        err := req.AddField(pb.NewField(tools.LowerCamelCase(col.Name()), pbType(col.DataType()), i, col.Comment(), false))
        i++
        if err != nil {
            return nil, err
        }
    }
    
    resp := factory.NewMessage(consts.PROTOBUF, fmt.Sprintf("%sAddResp", tools.UpperCamelCase(table.Name())))
    
    // rpc HiolabsOrderAdd(HiolabsOrderAddReq) returns (HiolabsOrderAddResp);
    return pb.NewItem(fmt.Sprintf("%sAdd", tools.UpperCamelCase(table.Name())), "", service, req, resp), nil
}

type DeleteItemCreator struct {
}

func (c *DeleteItemCreator) ItemCreate(table db.Table, service string, filters []convert.ColumnFilter) (protocol.Item, error) {
    req := factory.NewMessage(consts.PROTOBUF, fmt.Sprintf("%sDelReq", tools.UpperCamelCase(table.Name())))
    for _, col := range table.Cols() {
        if col.IsPrimary() {
            err := req.AddField(pb.NewField(tools.LowerCamelCase(col.Name()), pbType(col.DataType()), 1, col.Comment(), false))
            if err != nil {
                return nil, err
            }
            break
        }
    }
    
    resp := factory.NewMessage(consts.PROTOBUF, fmt.Sprintf("%sDelResp", tools.UpperCamelCase(table.Name())))
    
    // rpc HiolabsOrderDel(HiolabsOrderDelReq) returns (HiolabsOrderDelResp);
    return pb.NewItem(fmt.Sprintf("%sDel", tools.UpperCamelCase(table.Name())), "", service, req, resp), nil
}

type UpdateItemCreator struct {
}

func (c *UpdateItemCreator) ItemCreate(table db.Table, service string, filters []convert.ColumnFilter) (protocol.Item, error) {
    req := factory.NewMessage(consts.PROTOBUF, fmt.Sprintf("%sUpdateReq", tools.UpperCamelCase(table.Name())))
    i := 1
    for _, col := range table.Cols() {
        if filterOut(col, filters) {
            continue
        }
        if col.IsEnum() {
            err := req.AddField(pb.NewField(
                tools.LowerCamelCase(col.Name()),
                tools.UpperCamelCase(table.Name())+tools.UpperCamelCase(col.Name()),
                i,
                col.Comment(),
                false))
            i++
            if err != nil {
                return nil, err
            }
            continue
        }
        
        err := req.AddField(pb.NewField(tools.LowerCamelCase(col.Name()), pbType(col.DataType()), i, col.Comment(), false))
        i++
        if err != nil {
            return nil, err
        }
    }
    
    resp := factory.NewMessage(consts.PROTOBUF, fmt.Sprintf("%sUpdateResp", tools.UpperCamelCase(table.Name())))
    
    // rpc HiolabsOrderUpdate(HiolabsOrderUpdateReq) returns (HiolabsOrderUpdateResp);
    return pb.NewItem(fmt.Sprintf("%sUpdate", tools.UpperCamelCase(table.Name())), "", service, req, resp), nil
}

type GetItemCreator struct {
}

func (c *GetItemCreator) ItemCreate(table db.Table, service string, filters []convert.ColumnFilter) (protocol.Item, error) {
    req := factory.NewMessage(consts.PROTOBUF, fmt.Sprintf("%sGetReq", tools.UpperCamelCase(table.Name())))
    for _, col := range table.Cols() {
        if col.IsPrimary() {
            err := req.AddField(pb.NewField(tools.LowerCamelCase(col.Name()), pbType(col.DataType()), 1, col.Comment(), false))
            if err != nil {
                return nil, err
            }
            break
        }
    }
    
    resp := factory.NewMessage(consts.PROTOBUF, fmt.Sprintf("%sGetResp", tools.UpperCamelCase(table.Name())))
    // message HiolabsOrderGetResp {
    //   HiolabsOrder hiolabsOrder = 1;
    // }
    err := resp.AddField(pb.NewField(tools.UpperCamelCase(table.Name()), tools.LowerCamelCase(table.Name()), 1, "", false))
    if err != nil {
        return nil, err
    }
    
    return pb.NewItem(fmt.Sprintf("%sGet", tools.UpperCamelCase(table.Name())), "", service, req, resp), nil
}

type SearchItemCreator struct {
}

func (c *SearchItemCreator) ItemCreate(table db.Table, service string, filters []convert.ColumnFilter) (protocol.Item, error) {
    req := factory.NewMessage(consts.PROTOBUF, fmt.Sprintf("%sSearchReq", tools.UpperCamelCase(table.Name())))
    i := 1
    for _, col := range table.Cols() {
        if filterOut(col, filters) {
            continue
        }
        if col.IsEnum() {
            err := req.AddField(pb.NewField(
                tools.LowerCamelCase(col.Name()),
                tools.UpperCamelCase(table.Name())+tools.UpperCamelCase(col.Name()),
                i,
                col.Comment(),
                false))
            i++
            if err != nil {
                return nil, err
            }
            continue
        }
        
        err := req.AddField(pb.NewField(tools.LowerCamelCase(col.Name()), pbType(col.DataType()), i, col.Comment(), false))
        i++
        if err != nil {
            return nil, err
        }
    }
    
    resp := factory.NewMessage(consts.PROTOBUF, fmt.Sprintf("%sSearchResp", tools.UpperCamelCase(table.Name())))
    // message HiolabsOrderSearchResp {
    //   repeated HiolabsOrder hiolabsOrder = 1;
    // }
    err := resp.AddField(pb.NewField(tools.LowerCamelCase(table.Name()), tools.UpperCamelCase(table.Name()), 1, "", true))
    if err != nil {
        return nil, err
    }
    
    return pb.NewItem(fmt.Sprintf("%sSearch", tools.UpperCamelCase(table.Name())), "", service, req, resp), nil
}

func pbType(t string) string {
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

func filterOut(col db.Column, filters []convert.ColumnFilter) bool {
    for _, filter := range filters {
        if filter.FilterOut(col) {
            return true
        }
    }
    
    return false
}
