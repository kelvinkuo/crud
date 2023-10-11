package zero

import (
	"fmt"
	"strings"

	"github.com/kelvinkuo/crud/internal/core/tools"
	"github.com/kelvinkuo/crud/internal/protocol"
)

type Item struct {
	protocol.CommonItem
	method string
	path   string
}

func NewItem(name string, comment string, service string, request protocol.Message, response protocol.Message, method string, path string) *Item {
	return &Item{CommonItem: protocol.NewCommonItem(name, comment, service, request, response), method: method, path: path}
}

func (i *Item) String(indent int) string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%s@handler %s\n", tools.Blank(indent), i.Name()))
	builder.WriteString(fmt.Sprintf("%s%s %s (%s) returns (%s)\n", tools.Blank(indent), i.method, i.path, i.Request().Name(), i.Response().Name()))

	return builder.String()
}
