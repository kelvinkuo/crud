package protocol

type CommonItem struct {
    name     string
    comment  string
    service  string
    request  Message
    response Message
}

func NewCommonItem(name string, comment string, service string, request Message, response Message) CommonItem {
    return CommonItem{name: name, comment: comment, service: service, request: request, response: response}
}

func (i *CommonItem) Name() string {
    return i.name
}

func (i *CommonItem) Comment() string {
    return i.comment
}

func (i *CommonItem) Service() string {
    return i.service
}

func (i *CommonItem) Request() Message {
    return i.request
}

func (i *CommonItem) Response() Message {
    return i.response
}
