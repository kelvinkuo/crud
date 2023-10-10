package pbitemcreater

import (
    "fmt"
    
    "github.com/kelvinkuo/crud/internal/consts"
    "github.com/kelvinkuo/crud/internal/core/convert"
    pb2 "github.com/kelvinkuo/crud/internal/core/convert/pb"
    "github.com/kelvinkuo/crud/internal/core/tools"
    "github.com/kelvinkuo/crud/internal/db"
    "github.com/kelvinkuo/crud/internal/protocol"
    "github.com/kelvinkuo/crud/internal/protocol/pb"
    factory "github.com/kelvinkuo/crud/internal/protocol/protocolfactory"
)

type Update struct {
}

func (c *Update) ItemCreate(table db.Table, service string, filters []convert.ColumnFilter) (protocol.Item, error) {
    req := factory.NewMessage(consts.ProtoBuf, fmt.Sprintf("%sUpdateReq", tools.UpperCamelCase(table.Name())))
    i := 1
    for _, col := range table.Cols() {
        if convert.FilterOut(col, filters) {
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
        
        err := req.AddField(pb.NewField(tools.LowerCamelCase(col.Name()), pb2.PbType(col.DataType()), i, col.Comment(), false))
        i++
        if err != nil {
            return nil, err
        }
    }
    
    resp := factory.NewMessage(consts.ProtoBuf, fmt.Sprintf("%sUpdateResp", tools.UpperCamelCase(table.Name())))
    
    // rpc HiolabsOrderUpdate(HiolabsOrderUpdateReq) returns (HiolabsOrderUpdateResp);
    return pb.NewItem(fmt.Sprintf("%sUpdate", tools.UpperCamelCase(table.Name())), "", service, req, resp), nil
}