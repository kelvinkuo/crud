package dbfactory

import (
    "github.com/kelvinkuo/crud/consts"
    "github.com/kelvinkuo/crud/db"
    "github.com/kelvinkuo/crud/db/mysql"
)

func NewDB(dbType string) db.DB {
    switch dbType {
    case consts.MYSQL:
        return &mysql.DB{}
    }
    return nil
}
