package pb

import (
    "errors"
    "fmt"
    "sort"
    "strings"
    
    "github.com/kelvinkuo/crud/protocol"
)

type Protocol struct {
    syntax    string
    pbPackage string
    goPackage string
    messages  map[string]protocol.Message
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
    items map[string][]protocol.Item
}

func NewProtocol(syntax string, pbPackage string, goPackage string) *Protocol {
    return &Protocol{
        syntax: syntax, pbPackage: pbPackage, goPackage: goPackage,
        messages: map[string]protocol.Message{},
        items:    map[string][]protocol.Item{},
    }
}

func (p *Protocol) AddItem(item protocol.Item) error {
    service := item.Service()
    items, ok := p.items[service]
    if !ok {
        p.items[service] = make([]protocol.Item, 0)
        items = p.items[service]
    }
    
    for _, it := range items {
        if it.Name() == item.Name() {
            return errors.New("pb Item already exists")
        }
    }
    
    p.items[service] = append(p.items[service], item)
    return nil
}

func (p *Protocol) AddMessage(message protocol.Message) error {
    _, ok := p.messages[message.Name()]
    if ok {
        return errors.New("message already exists")
    }
    
    p.messages[message.Name()] = message
    return nil
}

func (p *Protocol) String() string {
    builder := strings.Builder{}
    builder.WriteString(fmt.Sprintf("syntax = \"%s\";\n", p.syntax))
    builder.WriteString(fmt.Sprintf("package %s;\n", p.pbPackage))
    builder.WriteString(fmt.Sprintf("option go_package = \"%s\";\n\n", p.goPackage))
    builder.WriteString("\n----------------------messages----------------------\n")
    msgs := make([]protocol.Message, 0, len(p.messages))
    for _, msg := range p.messages {
        msgs = append(msgs, msg)
    }
    sort.Slice(msgs, func(i, j int) bool {
        return msgs[i].Name() < msgs[j].Name()
    })
    for _, msg := range msgs {
        builder.WriteString(msg.String())
        builder.WriteString("\n")
    }
    builder.WriteString("\n----------------------rpc func----------------------\n")
    for service, itemSlice := range p.items {
        builder.WriteString(fmt.Sprintf("service %s {\n", service))
        for _, it := range itemSlice {
            builder.WriteString(fmt.Sprintf("  %s\n", it.String()))
        }
        builder.WriteString("}\n")
    }
    
    return builder.String()
}
