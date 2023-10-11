package pb

import (
	"testing"

	"github.com/kelvinkuo/crud/internal/protocol"
)

func Test_field_StringLine(t *testing.T) {
	type fields struct {
		name     string
		dataType string
		number   int
		comment  string
		repeated bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{
				name:     "id",
				dataType: "int64",
				number:   1,
				comment:  "主键",
				repeated: false,
			},
			want: "int64 id = 1; // 主键",
		},
		{
			fields: fields{
				name:     "hiolabsShowSettings",
				dataType: "HiolabsShowSettings",
				number:   1,
				comment:  "hiolabsShowSettings",
				repeated: true,
			},
			want: "repeated HiolabsShowSettings hiolabsShowSettings = 1; // hiolabsShowSettings",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Field{
				CommonField: protocol.NewCommonField(tt.fields.name, tt.fields.dataType, tt.fields.comment),
				number:      tt.fields.number,
				repeated:    tt.fields.repeated,
			}
			if got := f.String(0); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_String(t *testing.T) {
	type fields struct {
		name   string
		fields []protocol.Field
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{
				name: "User",
				fields: []protocol.Field{
					&Field{
						CommonField: protocol.NewCommonField("id", "int64", "id"),
						number:      1,
						repeated:    false,
					},
					&Field{
						CommonField: protocol.NewCommonField("username", "string", "username"),
						number:      2,
						repeated:    false,
					},
					&Field{
						CommonField: protocol.NewCommonField("mobile", "string", "mobile"),
						number:      3,
						repeated:    false,
					},
				},
			},
			want: `message User {
  int64 id = 1; // id
  string username = 2; // username
  string mobile = 3; // mobile
}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Message{
				CommonMessage: protocol.NewCommonMessage(tt.fields.name),
			}
			for _, field := range tt.fields.fields {
				_ = m.AddField(field)
			}
			if got := m.String(0); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
