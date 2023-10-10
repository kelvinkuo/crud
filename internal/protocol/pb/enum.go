package pb

import (
    "fmt"
    "strings"
    
    "github.com/kelvinkuo/crud/internal/consts"
    "github.com/kelvinkuo/crud/internal/core/tools"
    "github.com/kelvinkuo/crud/internal/protocol"
)

type Enum struct {
    protocol.CommonMessage
}

func NewEnum(name string) *Enum {
    return &Enum{CommonMessage: protocol.NewCommonMessage(name)}
}

func (e *Enum) String(indent int) string {
    b := strings.Builder{}
    b.WriteString(fmt.Sprintf("%senum %s {\n", tools.Blank(indent), e.Name()))
    for _, f := range e.Fields() {
        b.WriteString(fmt.Sprintf("%s\n", f.String(indent+consts.IndentProto3)))
    }
    b.WriteString(fmt.Sprintf("%s}\n", tools.Blank(indent)))
    
    return b.String()
}

type EnumField struct {
    name   string
    number int
}

func NewEnumField(name string, number int) *EnumField {
    return &EnumField{name: name, number: number}
}

func (e *EnumField) String(indent int) string {
    return fmt.Sprintf("%s%s = %d;", tools.Blank(indent), e.name, e.number)
}
