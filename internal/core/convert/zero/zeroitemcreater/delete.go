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

type Delete struct {
}

func (c *Delete) ItemCreate(table db.Table, service string, filters []convert.ColumnFilter) (protocol.Item, error) {
    req := factory.NewMessage(consts.ZeroApi, fmt.Sprintf("%sDeleteReq", tools.UpperCamelCase(table.Name())))
    resp := factory.NewMessage(consts.ZeroApi, fmt.Sprintf("%sDeleteResp", tools.UpperCamelCase(table.Name())))
    
    for _, col := range table.Cols() {
        if col.IsPrimary() {
            tag := fmt.Sprintf("json:\"%s\"", tools.LowerUnderline(col.Name()))
            err := req.AddField(zero.NewField(tools.UpperCamelCase(col.Name()), zero2.ZeroType(col.DataType()), col.Comment(), tag))
            if err != nil {
                return nil, err
            }
            err = resp.AddField(zero.NewField(tools.UpperCamelCase(col.Name()), zero2.ZeroType(col.DataType()), col.Comment(), tag))
            if err != nil {
                return nil, err
            }
            break
        }
    }
    
    // @handler HiolabsAdDelete
    // post /hiolabsad/delete (HiolabsAdDeleteReq) returns (HiolabsAdDeleteResp)
    return zero.NewItem(fmt.Sprintf("%sDelete", tools.UpperCamelCase(table.Name())), "", service, req, resp, "post",
        fmt.Sprintf("/%s/delete", strings.ToLower(tools.LowerCamelCase(table.Name())))), nil
}
