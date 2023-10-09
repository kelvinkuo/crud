package pb

import (
    "testing"
    
    "github.com/kelvinkuo/crud/protocol"
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
                name:     tt.fields.name,
                dataType: tt.fields.dataType,
                number:   tt.fields.number,
                comment:  tt.fields.comment,
                repeated: tt.fields.repeated,
            }
            if got := f.StringLine(); got != tt.want {
                t.Errorf("StringLine() = %v, want %v", got, tt.want)
            }
        })
    }
}

// Message User {
//  int64 id = 1; //id
//  string nickname = 2; //nickname
//  string username = 3; //username
//  string password = 4; //password
//  int64 gender = 5; //gender
//  int64 birthday = 6; //birthday
//  int64 registerTime = 7; //registerTime
//  int64 lastLoginTime = 8; //lastLoginTime
//  string lastLoginIp = 9; //lastLoginIp
//  string mobile = 10; //mobile
// }
func Test_message_AddField(t *testing.T) {
    type fields struct {
        name   string
        fields []protocol.Field
    }
    type args struct {
        field protocol.Field
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        {
            fields: fields{
                name: "User",
                fields: []protocol.Field{
                    &Field{
                        name:     "id",
                        dataType: "int64",
                        number:   1,
                        comment:  "id",
                        repeated: false,
                    },
                    &Field{
                        name:     "username",
                        dataType: "string",
                        number:   2,
                        comment:  "username",
                        repeated: false,
                    },
                    &Field{
                        name:     "mobile",
                        dataType: "string",
                        number:   3,
                        comment:  "mobile",
                        repeated: false,
                    },
                },
            },
            args: args{field: &Field{
                name:     "gender",
                dataType: "int64",
                number:   4,
                comment:  "gender",
                repeated: false,
            }},
            wantErr: false,
        },
        {
            fields: fields{
                name: "User",
                fields: []protocol.Field{
                    &Field{
                        name:     "id",
                        dataType: "int64",
                        number:   1,
                        comment:  "id",
                        repeated: false,
                    },
                },
            },
            args: args{field: &Field{
                name:     "id",
                dataType: "int64",
                number:   2,
                comment:  "id",
                repeated: false,
            }},
            wantErr: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            m := &Message{
                name:   tt.fields.name,
                fields: tt.fields.fields,
            }
            if err := m.AddField(tt.args.field); (err != nil) != tt.wantErr {
                t.Errorf("AddField() error = %v, wantErr %v", err, tt.wantErr)
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
                        name:     "id",
                        dataType: "int64",
                        number:   1,
                        comment:  "id",
                        repeated: false,
                    },
                    &Field{
                        name:     "username",
                        dataType: "string",
                        number:   2,
                        comment:  "username",
                        repeated: false,
                    },
                    &Field{
                        name:     "mobile",
                        dataType: "string",
                        number:   3,
                        comment:  "mobile",
                        repeated: false,
                    },
                },
            },
            want: `Message User {
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
                name:   tt.fields.name,
                fields: tt.fields.fields,
            }
            if got := m.String(); got != tt.want {
                t.Errorf("String() = %v, want %v", got, tt.want)
            }
        })
    }
}
