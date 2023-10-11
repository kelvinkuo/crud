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

type Info struct {
}

func (c *Info) ItemCreate(table db.Table, service, style string, filters []convert.ColumnFilter) (protocol.Item, error) {
    req := factory.NewMessage(consts.ProtoBuf, fmt.Sprintf("%sInfoReq", tools.UpperCamelCase(table.Name())))
    for _, col := range table.Cols() {
        if col.IsPrimary() {
            err := req.AddField(pb.NewField(pb2.StyleString(col.Name(), style), pb2.PbType(col.DataType()), 1, col.Comment(), false))
            if err != nil {
                return nil, err
            }
            break
        }
    }
    
    resp := factory.NewMessage(consts.ProtoBuf, fmt.Sprintf("%sInfoResp", tools.UpperCamelCase(table.Name())))
    err := resp.AddField(pb.NewField(pb2.StyleString(table.Name(), style), tools.UpperCamelCase(table.Name()), 1, "", false))
    if err != nil {
        return nil, err
    }
    
    return pb.NewItem(fmt.Sprintf("%sInfo", tools.UpperCamelCase(table.Name())), "", service, req, resp), nil
}
