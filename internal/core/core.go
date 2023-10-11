package core

import (
    "fmt"
    "log"
    
    "github.com/kelvinkuo/crud/internal/consts"
    "github.com/kelvinkuo/crud/internal/core/convert/convertfactory"
    "github.com/kelvinkuo/crud/internal/core/convert/filter"
    "github.com/kelvinkuo/crud/internal/db/dbfactory"
    "github.com/kelvinkuo/crud/internal/protocol/protocolfactory"
)

func Generate(datasource, protocolType, service, pkg, goPackage, style string, tableNames, ignoreCols, methods []string) {
    // init db
    dbInstance := dbfactory.NewDB(consts.MYSQL)
    err := dbInstance.Init(datasource)
    if err != nil {
        log.Fatal("db init error ", err)
    }
    defer dbInstance.Close()
    
    if len(tableNames) == 0 {
        tableNames = dbInstance.AllTableNames()
    }
    tables, err := dbInstance.GetTables(tableNames)
    if err != nil {
        log.Fatal(err)
    }
    
    // init protocol
    p := protocolfactory.NewProtocol(protocolType, pkg, goPackage)
    convertor := convertfactory.NewConverter(protocolType, style)
    for _, method := range methods {
        convertor.AddItemCreator(convertfactory.NewItemCreator(protocolType, method))
    }
    convertor.AddColumnFilter(filter.NewStringFilter(ignoreCols))
    
    for _, table := range tables {
        msgList, err := convertor.CreateMetaMessage(table)
        if err != nil {
            log.Fatal(err)
        }
        for _, msg := range msgList {
            err = p.AddMessage(msg)
            if err != nil {
                log.Fatal(err)
            }
        }
        
        items, err := convertor.CreateItems(table, service)
        if err != nil {
            log.Fatal(err)
        }
        for _, item := range items {
            err = p.AddItem(item)
            if err != nil {
                log.Fatal(err)
            }
            err = p.AddMessage(item.Request())
            if err != nil {
                log.Fatal(err)
            }
            err = p.AddMessage(item.Response())
            if err != nil {
                log.Fatal(err)
            }
        }
    }
    
    fmt.Print(p.String())
}
