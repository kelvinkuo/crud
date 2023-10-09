package pb

import (
    "errors"
    "fmt"
    "strings"
    
    "github.com/kelvinkuo/crud/protocol"
)

// Message of protobuf, example:
// message User {
//  int64 id = 1; // id
//  string nickname = 2; // nickname
//  string username = 3; // username
//  string password = 4; // password
//  int64 gender = 5; // gender
//  int64 birthday = 6; // birthday
//  int64 registerTime = 7; // registerTime
//  int64 lastLoginTime = 8; // lastLoginTime
//  string lastLoginIp = 9; // lastLoginIp
//  string mobile = 10; // mobile
// }
type Message struct {
    name   string
    fields []protocol.Field
}

func NewMessage(name string) *Message {
    return &Message{name: name}
}

func (m *Message) Name() string {
    return m.name
}

func (m *Message) AddField(field protocol.Field) error {
    for _, f := range m.fields {
        if f.StringLine() == field.StringLine() {
            return errors.New("field already exists")
        }
    }
    
    m.fields = append(m.fields, field)
    return nil
}

func (m *Message) String() string {
    b := strings.Builder{}
    b.WriteString(fmt.Sprintf("message %s {\n", m.name))
    for _, f := range m.fields {
        b.WriteString(fmt.Sprintf("  %s\n", f.StringLine()))
    }
    b.WriteString("}\n")
    
    return b.String()
}

type Field struct {
    name     string
    dataType string
    number   int
    comment  string
    repeated bool
}

func NewField(name string, dataType string, number int, comment string, repeated bool) *Field {
    return &Field{name: name, dataType: dataType, number: number, comment: comment, repeated: repeated}
}

// StringLine
// example:   string username = 6; //username
func (f *Field) StringLine() string {
    prefix := ""
    if f.repeated {
        prefix = "repeated "
    }
    return prefix + fmt.Sprintf("%s %s = %d; // %s", f.dataType, f.name, f.number, f.comment)
}
