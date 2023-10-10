package dbfactory

import (
    "github.com/kelvinkuo/crud/internal/consts"
    "github.com/kelvinkuo/crud/internal/db"
    "github.com/kelvinkuo/crud/internal/db/mysql"
)

func NewDB(dbType string) db.DB {
    switch dbType {
    case consts.MYSQL:
        return &mysql.DB{}
    }
    return nil
}
