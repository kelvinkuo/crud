package convertfactory

import (
	"github.com/kelvinkuo/crud/internal/consts"
	"github.com/kelvinkuo/crud/internal/core/convert"
	"github.com/kelvinkuo/crud/internal/core/convert/pb"
	"github.com/kelvinkuo/crud/internal/core/convert/pb/pbitemcreater"
	"github.com/kelvinkuo/crud/internal/core/convert/zero"
	"github.com/kelvinkuo/crud/internal/core/convert/zero/zeroitemcreater"
)

func NewConverter(protocolType, style string) convert.Converter {
	switch protocolType {
	case consts.ProtoBuf:
		return pb.NewConverter(style)
	case consts.ZeroApi:
		return zero.NewConverter()
	}
	return nil
}

var creatorMap = map[string]map[string]convert.ItemCreator{
	consts.ProtoBuf: {
		consts.Add:    &pbitemcreater.Add{},
		consts.Delete: &pbitemcreater.Delete{},
		consts.Update: &pbitemcreater.Update{},
		consts.Info:   &pbitemcreater.Info{},
		consts.List:   &pbitemcreater.List{},
		consts.Search: &pbitemcreater.Search{},
	},
	consts.ZeroApi: {
		consts.Add:    &zeroitemcreater.Add{},
		consts.Delete: &zeroitemcreater.Delete{},
		consts.Update: &zeroitemcreater.Update{},
		consts.Info:   &zeroitemcreater.Info{},
		consts.List:   &zeroitemcreater.List{},
		consts.Search: &zeroitemcreater.Search{},
	},
}

func NewItemCreator(protocolType string, name string) convert.ItemCreator {
	if _, ok := creatorMap[protocolType]; !ok {
		return nil
	}

	return creatorMap[protocolType][name]
}
