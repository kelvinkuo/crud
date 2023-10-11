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

type Search struct {
}

func (c *Search) ItemCreate(table db.Table, service, style string, filters []convert.ColumnFilter) (protocol.Item, error) {
    req := factory.NewMessage(consts.ProtoBuf, fmt.Sprintf("%sSearchReq", tools.UpperCamelCase(table.Name())))
    i := 1
    for _, col := range table.Cols() {
        if convert.FilterOut(col, filters) {
            continue
        }
        if col.IsEnum() {
            err := req.AddField(pb.NewField(
                pb2.StyleString(col.Name(), style),
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
        
        err := req.AddField(pb.NewField(pb2.StyleString(col.Name(), style), pb2.PbType(col.DataType()), i, col.Comment(), false))
        i++
        if err != nil {
            return nil, err
        }
    }
    
    resp := factory.NewMessage(consts.ProtoBuf, fmt.Sprintf("%sSearchResp", tools.UpperCamelCase(table.Name())))
    // message OrderSearchResp {
    //   repeated Order list = 1;
    // }
    err := resp.AddField(pb.NewField("list", tools.UpperCamelCase(table.Name()), 1, "", true))
    if err != nil {
        return nil, err
    }
    
    return pb.NewItem(fmt.Sprintf("%sSearch", tools.UpperCamelCase(table.Name())), "", service, req, resp), nil
}
