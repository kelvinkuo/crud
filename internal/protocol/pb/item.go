package pb

import (
    "fmt"
    
    "github.com/kelvinkuo/crud/protocol"
)

type Item struct {
    name     string
    comment  string
    service  string
    request  protocol.Message
    response protocol.Message
}

func NewItem(name string, comment string, service string, request protocol.Message, response protocol.Message) *Item {
    return &Item{name: name, comment: comment, service: service, request: request, response: response}
}

// Item example:
// in file ad.proto
// rpc GetAdById(GetAdByIdReq) returns (GetAdByIdResp);
func (i *Item) String() string {
    return fmt.Sprintf("rpc %s(%s) returns (%s)", i.name, i.request.Name(), i.response.Name())
}

func (i *Item) Name() string {
    return i.name
}

func (i *Item) Comment() string {
    return i.comment
}

func (i *Item) Service() string {
    return i.service
}

func (i *Item) Request() protocol.Message {
    return i.request
}

func (i *Item) Response() protocol.Message {
    return i.response
}
