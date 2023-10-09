package pb

import (
    "errors"
    "fmt"
    "strings"
    
    "github.com/kelvinkuo/crud/protocol"
)

type Enum struct {
    name   string
    fields []protocol.Field
}

func NewEnum(name string) *Enum {
    return &Enum{name: name}
}

func (e *Enum) Name() string {
    return e.name
}

func (e *Enum) AddField(field protocol.Field) error {
    for _, f := range e.fields {
        if f.StringLine() == field.StringLine() {
            return errors.New("enum field already exists")
        }
    }
    
    e.fields = append(e.fields, field)
    return nil
}

func (e *Enum) String() string {
    b := strings.Builder{}
    b.WriteString(fmt.Sprintf("enum %s {\n", e.name))
    for _, f := range e.fields {
        b.WriteString(fmt.Sprintf("  %s\n", f.StringLine()))
    }
    b.WriteString("}\n")
    
    return b.String()
}

type EnumField struct {
    name   string
    number int
}

func NewEnumField(name string, number int) *EnumField {
    return &EnumField{name: name, number: number}
}

func (e *EnumField) StringLine() string {
    return fmt.Sprintf("%s = %d;", e.name, e.number)
}
