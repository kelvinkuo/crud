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

type Delete struct {
}

func (c *Delete) ItemCreate(table db.Table, service, style string, filters []convert.ColumnFilter) (protocol.Item, error) {
	req := factory.NewMessage(consts.ProtoBuf, fmt.Sprintf("%sDelReq", tools.UpperCamelCase(table.Name())))
	for _, col := range table.Cols() {
		if col.IsPrimary() {
			err := req.AddField(pb.NewField(pb2.StyleString(col.Name(), style), pb2.PbType(col.DataType()), 1, col.Comment(), false))
			if err != nil {
				return nil, err
			}
			break
		}
	}

	resp := factory.NewMessage(consts.ProtoBuf, fmt.Sprintf("%sDelResp", tools.UpperCamelCase(table.Name())))

	// rpc OrderDel(OrderDelReq) returns (OrderDelResp);
	return pb.NewItem(fmt.Sprintf("%sDel", tools.UpperCamelCase(table.Name())), "", service, req, resp), nil
}
