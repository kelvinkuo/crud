package protocol

import (
    "errors"
)

type CommonMessage struct {
    name   string
    fields []Field
}

func NewCommonMessage(name string) CommonMessage {
    return CommonMessage{name: name}
}

func (m *CommonMessage) Name() string {
    return m.name
}

func (m *CommonMessage) AddField(field Field) error {
    for _, f := range m.fields {
        if f.String(0) == field.String(0) {
            return errors.New("field already exists")
        }
    }
    
    m.fields = append(m.fields, field)
    return nil
}

func (m *CommonMessage) Fields() []Field {
    return m.fields
}

type CommonField struct {
    Name     string
    DataType string
    Comment  string
}

func NewCommonField(name string, dataType string, comment string) CommonField {
    return CommonField{Name: name, DataType: dataType, Comment: comment}
}
