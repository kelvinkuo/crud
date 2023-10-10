package pb

import (
    "fmt"
    
    "github.com/kelvinkuo/crud/internal/core/tools"
    "github.com/kelvinkuo/crud/internal/protocol"
)

type Item struct {
    protocol.CommonItem
}

func NewItem(name string, comment string, service string, request protocol.Message, response protocol.Message) *Item {
    return &Item{CommonItem: protocol.NewCommonItem(name, comment, service, request, response)}
}

func (i *Item) String(indent int) string {
    return fmt.Sprintf("%srpc %s(%s) returns (%s);", tools.Blank(indent), i.Name(), i.Request().Name(), i.Response().Name())
}
