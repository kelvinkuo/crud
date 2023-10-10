package mysql

import (
    "reflect"
    "testing"
    
    "github.com/kelvinkuo/crud/internal/db"
)

func NewDB(datasource string) (*DB, error) {
    newDB := &DB{}
    return newDB, newDB.Init(datasource)
}

func TestNewDB(t *testing.T) {
    type args struct {
        dataSource string
    }
    tests := []struct {
        name    string
        args    args
        want    *DB
        wantErr bool
    }{
        {
            args:    args{dataSource: "root:123456@tcp(127.0.0.1:3306)/shop"}, // correct
            wantErr: false,
        },
        {
            args:    args{dataSource: "root:123456@tcp(127.0.0.1:3307)/shop"}, // error port
            wantErr: true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            _, err := NewDB(tt.args.dataSource)
            if (err != nil) != tt.wantErr {
                t.Errorf("NewDB() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
        })
    }
}

func TestDB_getSchema(t *testing.T) {
    type args struct {
        dataSource string
    }
    tests := []struct {
        name    string
        args    args
        want    string
        wantErr bool
    }{
        {
            args:    args{dataSource: "root:123456@tcp(127.0.0.1:3306)/shop"},
            want:    "shop",
            wantErr: false,
        },
        {
            args:    args{dataSource: "root:123456@tcp(127.0.0.1:3306)/order"},
            want:    "order",
            wantErr: false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            d, err := NewDB(tt.args.dataSource)
            if err != nil {
                t.Fatal(err)
            }
            got, err := d.getSchema()
            if (err != nil) != tt.wantErr {
                t.Errorf("getSchema() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("getSchema() got = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDB_AllTableNames(t *testing.T) {
    type fields struct {
        dataSource string
    }
    tests := []struct {
        name   string
        fields fields
        want   []string
    }{
        {
            fields: fields{
                dataSource: "root:123456@tcp(127.0.0.1:3306)/shop",
            },
            want: []string{
                "user",
                "product",
                "order",
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            d, err := NewDB(tt.fields.dataSource)
            if err != nil {
                t.Error(err)
            }
            if got := d.AllTableNames(); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("AllTableNames() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDB_GetTable(t *testing.T) {
    type fields struct {
        dataSource string
    }
    type args struct {
        table string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    db.Table
        wantErr bool
    }{
        {
            fields: fields{
                dataSource: "root:123456@tcp(127.0.0.1:3306)/shop",
            },
            args: args{table: "user"},
            want: &table{
                name:    "user",
                comment: "用户表",
                cols: []db.Column{
                    &column{
                        name:      "id",
                        comment:   "自增 ID",
                        dataType:  "bigint",
                        isPrimary: true,
                    },
                    &column{
                        name:      "name",
                        comment:   "用户名",
                        dataType:  "varchar",
                        isPrimary: false,
                    },
                    &column{
                        name:      "created_at",
                        comment:   "创建时间",
                        dataType:  "timestamp",
                        isPrimary: false,
                    },
                    &column{
                        name:      "updated_at",
                        comment:   "更新时间",
                        dataType:  "timestamp",
                        isPrimary: false,
                    },
                    &column{
                        name:      "deleted_at",
                        comment:   "软删除",
                        dataType:  "timestamp",
                        isPrimary: false,
                    },
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            d, err := NewDB(tt.fields.dataSource)
            if err != nil {
                t.Fatal(err)
            }
            got, err := d.GetTable(tt.args.table)
            if (err != nil) != tt.wantErr {
                t.Errorf("GetTable() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("GetTable() got = %v, want %v", got, tt.want)
            }
        })
    }
}
