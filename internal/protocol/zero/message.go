package zero

import (
    "fmt"
    "strings"
    
    "github.com/kelvinkuo/crud/internal/consts"
    "github.com/kelvinkuo/crud/internal/core/tools"
    "github.com/kelvinkuo/crud/internal/protocol"
)

type Message struct {
    protocol.CommonMessage
}

func NewMessage(name string) *Message {
    return &Message{CommonMessage: protocol.NewCommonMessage(name)}
}

func (m *Message) String(indent int) string {
    b := strings.Builder{}
    b.WriteString(fmt.Sprintf("%s%s {\n", tools.Blank(indent), m.Name()))
    for _, f := range m.Fields() {
        b.WriteString(fmt.Sprintf("%s\n", f.String(indent+consts.IndentZero)))
    }
    b.WriteString(fmt.Sprintf("%s}\n", tools.Blank(indent)))
    
    return b.String()
}

type Field struct {
    protocol.CommonField
    tag string
}

func NewField(name string, dataType string, comment string, tag string) *Field {
    return &Field{CommonField: protocol.NewCommonField(name, dataType, comment), tag: tag}
}

func (f *Field) String(indent int) string {
    builder := strings.Builder{}
    builder.WriteString(tools.Blank(indent))
    builder.WriteString(f.Name)
    if f.DataType != "" {
        builder.WriteString(fmt.Sprintf(" %s", f.DataType))
    }
    if f.tag != "" {
        builder.WriteString(fmt.Sprintf(" `%s`", f.tag))
    }
    if f.Comment != "" {
        builder.WriteString(fmt.Sprintf(" // %s", f.Comment))
    }
    
    return builder.String()
}
