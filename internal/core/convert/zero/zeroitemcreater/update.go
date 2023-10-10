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

type Update struct {
}

func (c *Update) ItemCreate(table db.Table, service string, filters []convert.ColumnFilter) (protocol.Item, error) {
    req := factory.NewMessage(consts.ZeroApi, fmt.Sprintf("%sUpdateReq", tools.UpperCamelCase(table.Name())))
    err := req.AddField(zero.NewField(tools.UpperCamelCase(table.Name()), "", "", ""))
    if err != nil {
        return nil, err
    }
    
    resp := factory.NewMessage(consts.ZeroApi, fmt.Sprintf("%sUpdateResp", tools.UpperCamelCase(table.Name())))
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
    
    // @handler HiolabsAdUpdate
    // post /hiolabsad/update (HiolabsAdUpdateReq) returns (HiolabsAdUpdateResp)
    return zero.NewItem(fmt.Sprintf("%sUpdate", tools.UpperCamelCase(table.Name())), "", service, req, resp, "post",
        fmt.Sprintf("/%s/update", strings.ToLower(tools.LowerCamelCase(table.Name())))), nil
}
