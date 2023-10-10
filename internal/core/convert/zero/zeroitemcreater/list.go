package zeroitemcreater

import (
    "fmt"
    "strings"
    
    "github.com/kelvinkuo/crud/internal/consts"
    "github.com/kelvinkuo/crud/internal/core/convert"
    "github.com/kelvinkuo/crud/internal/core/tools"
    "github.com/kelvinkuo/crud/internal/db"
    "github.com/kelvinkuo/crud/internal/protocol"
    factory "github.com/kelvinkuo/crud/internal/protocol/protocolfactory"
    "github.com/kelvinkuo/crud/internal/protocol/zero"
)

type List struct {
}

func (c *List) ItemCreate(table db.Table, service, style string, filters []convert.ColumnFilter) (protocol.Item, error) {
    req := factory.NewMessage(consts.ZeroApi, fmt.Sprintf("%sListReq", tools.UpperCamelCase(table.Name())))
    err := req.AddField(zero.NewField("page", "int", "", "form:\"page\""))
    if err != nil {
        return nil, err
    }
    err = req.AddField(zero.NewField("pageSize", "int", "", "form:\"page_size\""))
    if err != nil {
        return nil, err
    }
    
    resp := factory.NewMessage(consts.ZeroApi, fmt.Sprintf("%sListResp", tools.UpperCamelCase(table.Name())))
    // type HomestaySearchResp {
    //     List []Homestay `json:"list"`
    // }
    err = resp.AddField(zero.NewField("list", fmt.Sprintf("[]%s", tools.UpperCamelCase(table.Name())), "", "json:\"list\""))
    if err != nil {
        return nil, err
    }
    return zero.NewItem(fmt.Sprintf("%sList", tools.UpperCamelCase(table.Name())), "", service, req, resp, "get",
        fmt.Sprintf("/%s/list", strings.ToLower(tools.LowerCamelCase(table.Name())))), nil
}
