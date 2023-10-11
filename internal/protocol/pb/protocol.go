package pb

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kelvinkuo/crud/internal/consts"
	"github.com/kelvinkuo/crud/internal/protocol"
)

type Protocol struct {
	protocol.CommonProtocol
	pbPackage string
	goPackage string
}

func NewProtocol(syntax string, pbPackage string, goPackage string) *Protocol {
	return &Protocol{CommonProtocol: protocol.NewCommonProtocol(syntax), pbPackage: pbPackage, goPackage: goPackage}
}

func (p *Protocol) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("syntax = \"%s\";\n", p.Syntax))
	builder.WriteString(fmt.Sprintf("package %s;\n", p.pbPackage))
	builder.WriteString(fmt.Sprintf("option go_package = \"%s\";\n\n", p.goPackage))
	builder.WriteString("\n//----------------------messages----------------------\n")
	msgs := make([]protocol.Message, 0, len(p.Messages))
	for _, msg := range p.Messages {
		msgs = append(msgs, msg)
	}
	sort.Slice(msgs, func(i, j int) bool {
		return msgs[i].Name() < msgs[j].Name()
	})
	for _, msg := range msgs {
		builder.WriteString(msg.String(0))
		builder.WriteString("\n")
	}
	builder.WriteString("\n//----------------------rpc func----------------------\n")
	for service, itemSlice := range p.Items {
		builder.WriteString(fmt.Sprintf("service %s {\n", service))
		for _, it := range itemSlice {
			builder.WriteString(fmt.Sprintf("%s\n", it.String(consts.IndentProto3)))
		}
		builder.WriteString("}\n")
	}

	return builder.String()
}
