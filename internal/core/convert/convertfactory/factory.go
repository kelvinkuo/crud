package convertfactory

import (
    "github.com/kelvinkuo/crud/consts"
    "github.com/kelvinkuo/crud/core/convert"
    "github.com/kelvinkuo/crud/core/convert/pb"
)

func NewConverter(protocolType string) convert.Converter {
    switch protocolType {
    case consts.PROTOBUF:
        return pb.NewPbConverter()
    }
    return nil
}

var creatorMap = map[string]map[string]convert.ItemCreator{
    consts.PROTOBUF: {
        consts.Add:    &pb.AddItemCreator{},
        consts.Delete: &pb.DeleteItemCreator{},
        consts.Update: &pb.UpdateItemCreator{},
        consts.Get:    &pb.GetItemCreator{},
        consts.Search: &pb.SearchItemCreator{},
    },
}

func NewItemCreator(protocolType string, name string) convert.ItemCreator {
    if _, ok := creatorMap[protocolType]; !ok {
        return nil
    }
    
    return creatorMap[protocolType][name]
}
