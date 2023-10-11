package pb

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
	return &Message{protocol.NewCommonMessage(name)}
}

func (m *Message) String(indent int) string {
	b := strings.Builder{}
	b.WriteString(fmt.Sprintf("%smessage %s {\n", tools.Blank(indent), m.Name()))
	for _, f := range m.Fields() {
		b.WriteString(fmt.Sprintf("%s\n", f.String(indent+consts.IndentProto3)))
	}
	b.WriteString(fmt.Sprintf("%s}\n", tools.Blank(indent)))

	return b.String()
}

type Field struct {
	protocol.CommonField
	number   int
	repeated bool
}

func NewField(name string, dataType string, number int, comment string, repeated bool) *Field {
	return &Field{CommonField: protocol.NewCommonField(name, dataType, comment), number: number, repeated: repeated}
}

func (f *Field) String(indent int) string {
	builder := strings.Builder{}
	builder.WriteString(tools.Blank(indent))
	if f.repeated {
		builder.WriteString("repeated ")
	}
	builder.WriteString(fmt.Sprintf("%s %s = %d;", f.DataType, f.Name, f.number))
	if f.Comment != "" {
		builder.WriteString(fmt.Sprintf(" // %s", f.Comment))
	}

	return builder.String()
}
