package pbitemcreater

import (
    "fmt"
    
    "github.com/kelvinkuo/crud/internal/consts"
    "github.com/kelvinkuo/crud/internal/core/convert"
    "github.com/kelvinkuo/crud/internal/core/tools"
    "github.com/kelvinkuo/crud/internal/db"
    "github.com/kelvinkuo/crud/internal/protocol"
    "github.com/kelvinkuo/crud/internal/protocol/pb"
    factory "github.com/kelvinkuo/crud/internal/protocol/protocolfactory"
)

type List struct {
}

func (c *List) ItemCreate(table db.Table, service string, filters []convert.ColumnFilter) (protocol.Item, error) {
    req := factory.NewMessage(consts.ProtoBuf, fmt.Sprintf("%sListReq", tools.UpperCamelCase(table.Name())))
    err := req.AddField(pb.NewField("page", "int32", 1, "", false))
    if err != nil {
        return nil, err
    }
    err = req.AddField(pb.NewField("pageSize", "int32", 2, "", false))
    if err != nil {
        return nil, err
    }
    
    resp := factory.NewMessage(consts.ProtoBuf, fmt.Sprintf("%sListResp", tools.UpperCamelCase(table.Name())))
    // message HiolabsOrderInfoResp {
    //   repeated HiolabsOrder list = 1;
    // }
    err = resp.AddField(pb.NewField("list", tools.UpperCamelCase(table.Name()), 1, "", true))
    if err != nil {
        return nil, err
    }
    
    return pb.NewItem(fmt.Sprintf("%sList", tools.UpperCamelCase(table.Name())), "", service, req, resp), nil
}
