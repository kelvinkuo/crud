package zero

import (
    "fmt"
    "sort"
    "strings"
    "time"
    
    "github.com/kelvinkuo/crud/internal/consts"
    "github.com/kelvinkuo/crud/internal/core/tools"
    "github.com/kelvinkuo/crud/internal/protocol"
)

type Protocol struct {
    protocol.CommonProtocol
    author string
    date   time.Time
}

func NewProtocol(syntax, author string, date time.Time) *Protocol {
    return &Protocol{CommonProtocol: protocol.NewCommonProtocol(syntax), author: author, date: date}
}

func (p *Protocol) String() string {
    builder := strings.Builder{}
    builder.WriteString(fmt.Sprintf("syntax = \"%s\"\n", p.Syntax))
    builder.WriteString(fmt.Sprintf("info (\n"))
    builder.WriteString(fmt.Sprintf("%sauthor: %s\n", tools.Blank(consts.IndentZero), p.author))
    builder.WriteString(fmt.Sprintf("%sdate: %s\n", tools.Blank(consts.IndentZero), p.date.Format(time.DateOnly)))
    builder.WriteString(fmt.Sprintf(")\n\n"))
    builder.WriteString("\n//----------------------types define----------------------\n")
    msgs := make([]protocol.Message, 0, len(p.Messages))
    for _, msg := range p.Messages {
        msgs = append(msgs, msg)
    }
    sort.Slice(msgs, func(i, j int) bool {
        return msgs[i].Name() < msgs[j].Name()
    })
    
    builder.WriteString("type (\n\n")
    for _, msg := range msgs {
        builder.WriteString(fmt.Sprintf("%s\n", msg.String(consts.IndentZero)))
    }
    builder.WriteString(")\n")
    
    builder.WriteString("\n//----------------------http apis----------------------\n")
    for service, itemSlice := range p.Items {
        builder.WriteString(fmt.Sprintf("service %s {\n\n", service))
        for _, it := range itemSlice {
            builder.WriteString(fmt.Sprintf("%s\n", it.String(consts.IndentZero)))
        }
        builder.WriteString("}\n")
    }
    
    return builder.String()
}
