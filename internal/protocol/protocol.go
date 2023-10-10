package protocol

import (
    "errors"
)

type CommonProtocol struct {
    Syntax   string
    Messages map[string]Message
    // first key is service name, second is Item name
    // example:
    // {
    // 		service-shop : [
    //			UserAdd,
    //          UserDelete,
    //          UserUpdate,
    //          UserInfo,
    //          UserSearch,
    //			ProductAdd,
    //          ProductDelete,
    //          ProductUpdate,
    //          ProductInfo,
    //          ProductSearch,
    //      ]
    // }
    Items map[string][]Item
}

func NewCommonProtocol(syntax string) CommonProtocol {
    return CommonProtocol{
        Syntax:   syntax,
        Messages: map[string]Message{},
        Items:    map[string][]Item{},
    }
}

func (p *CommonProtocol) AddItem(item Item) error {
    service := item.Service()
    items, ok := p.Items[service]
    if !ok {
        p.Items[service] = make([]Item, 0)
        items = p.Items[service]
    }
    
    for _, it := range items {
        if it.Name() == item.Name() {
            return errors.New("item already exists")
        }
    }
    
    p.Items[service] = append(p.Items[service], item)
    return nil
}

func (p *CommonProtocol) AddMessage(message Message) error {
    _, ok := p.Messages[message.Name()]
    if ok {
        return errors.New("message already exists")
    }
    
    p.Messages[message.Name()] = message
    return nil
}
