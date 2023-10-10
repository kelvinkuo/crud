package zeroitemcreater

import (
    "fmt"
    "strings"
    
    "github.com/kelvinkuo/crud/internal/consts"
    "github.com/kelvinkuo/crud/internal/core/convert"
    zero2 "github.com/kelvinkuo/crud/internal/core/convert/zero"
    "github.com/kelvinkuo/crud/internal/core/tools"
    "github.com/kelvinkuo/crud/internal/db"
    "github.com/kelvinkuo/crud/internal/protocol"
    factory "github.com/kelvinkuo/crud/internal/protocol/protocolfactory"
    "github.com/kelvinkuo/crud/internal/protocol/zero"
)

type Add struct {
}

func (c *Add) ItemCreate(table db.Table, service, style string, filters []convert.ColumnFilter) (protocol.Item, error) {
    req := factory.NewMessage(consts.ZeroApi, fmt.Sprintf("%sAddReq", tools.UpperCamelCase(table.Name())))
    err := req.AddField(zero.NewField(tools.UpperCamelCase(table.Name()), "", "", ""))
    if err != nil {
        return nil, err
    }
    
    resp := factory.NewMessage(consts.ZeroApi, fmt.Sprintf("%sAddResp", tools.UpperCamelCase(table.Name())))
    for _, col := range table.Cols() {
        if col.IsPrimary() {
            tag := fmt.Sprintf("json:\"%s\"", tools.LowerUnderline(col.Name()))
            err = resp.AddField(zero.NewField(tools.UpperCamelCase(col.Name()), zero2.ZeroType(col.DataType()), col.Comment(), tag))
            if err != nil {
                return nil, err
            }
            break
        }
    }
    
    // @handler HiolabsAdAdd
    // post /hiolabsad/add (HiolabsAdAddReq) returns (HiolabsAdAddResp)
    return zero.NewItem(fmt.Sprintf("%sAdd", tools.UpperCamelCase(table.Name())), "", service, req, resp, "post",
        fmt.Sprintf("/%s/add", strings.ToLower(tools.LowerCamelCase(table.Name())))), nil
}
