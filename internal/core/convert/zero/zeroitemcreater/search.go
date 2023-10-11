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

type Search struct {
}

func (c *Search) ItemCreate(table db.Table, service, style string, filters []convert.ColumnFilter) (protocol.Item, error) {
	req := factory.NewMessage(consts.ZeroApi, fmt.Sprintf("%sSearchReq", tools.UpperCamelCase(table.Name())))
	for _, col := range table.Cols() {
		if convert.FilterOut(col, filters) {
			continue
		}
		tag := fmt.Sprintf("form:\"%s,optional\"", tools.LowerUnderline(col.Name()))
		err := req.AddField(zero.NewField(tools.UpperCamelCase(col.Name()), zero2.ZeroType(col.DataType()), col.Comment(), tag))
		if err != nil {
			return nil, err
		}
	}

	resp := factory.NewMessage(consts.ZeroApi, fmt.Sprintf("%sSearchResp", tools.UpperCamelCase(table.Name())))
	// type OrderSearchResp {
	//     List []Order `json:"list"`
	// }
	err := resp.AddField(zero.NewField("List", fmt.Sprintf("[]%s", tools.UpperCamelCase(table.Name())), "", "json:\"list\""))
	if err != nil {
		return nil, err
	}
	return zero.NewItem(fmt.Sprintf("%sSearch", tools.UpperCamelCase(table.Name())), "", service, req, resp, "get",
		fmt.Sprintf("/%s/search", strings.ToLower(tools.LowerCamelCase(table.Name())))), nil
}
