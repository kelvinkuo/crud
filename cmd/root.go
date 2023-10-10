package cmd

import (
    "fmt"
    "log"
    "os"
    "strings"
    
    "github.com/kelvinkuo/crud/internal/consts"
    "github.com/kelvinkuo/crud/internal/core/convert/convertfactory"
    "github.com/kelvinkuo/crud/internal/core/convert/filter"
    "github.com/kelvinkuo/crud/internal/db/dbfactory"
    "github.com/kelvinkuo/crud/internal/protocol/protocolfactory"
    "github.com/spf13/cobra"
)

var (
    format       string
    datasource   string
    methodStr    string
    methods      []string
    ignoreColStr string
    ignoreCols   []string
    tableStr     string
    tableNames   []string
    service      string
    goPackage    string
    pkg          string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "crud",
    Short: "generate CRUD defined content from datasource",
    Long: `Generate CRUD defined content from datasource.
Database supported: mysql
Format Support: proto3, zero api
Full example:
crud -f proto3 --source "root:123456@tcp(127.0.0.1:3306)/shop" -m "add,delete,update,info,search" -c "created_at,updated_at,deleted_at" -t * -s shop -go_package shop -package shop

Common example:
crud -f zeroapi --source "root:123456@tcp(127.0.0.1:3306)/shop" -s shop

`,
    Run: func(cmd *cobra.Command, args []string) {
        // flags
        methods = strings.Split(methodStr, ",")
        ignoreCols = strings.Split(ignoreColStr, ",")
        if tableStr != "*" {
            tableNames = strings.Split(tableStr, ",")
        }
        if goPackage == "" {
            goPackage = service
        }
        if pkg == "" {
            pkg = service
        }
        // generate
        generate()
    },
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}

func init() {
    rootCmd.Flags().SortFlags = false
    rootCmd.Flags().StringVarP(&format, "format", "f", "proto3", "output format support proto3, zero")
    rootCmd.Flags().StringVar(&datasource, "source", "", "datasource example: root:123456@tcp(127.0.0.1:3306)/shop")
    rootCmd.Flags().StringVarP(&methodStr, "method", "m", "add,delete,update,info,search", "methods separated by \",\"")
    rootCmd.Flags().StringVarP(&ignoreColStr, "ignore_cols", "c", "create_at,create_time,created_at,update_at,update_time,updated_at",
        "columns ignored separated by \",\"")
    rootCmd.Flags().StringVarP(&tableStr, "table", "t", "*", "tables separated by \",\" or \"*\" for all tables")
    rootCmd.Flags().StringVarP(&service, "service", "s", "", "service name")
    rootCmd.Flags().StringVar(&goPackage, "go_package", "", "go_package used in proto3 (default the same as service)")
    rootCmd.Flags().StringVarP(&pkg, "package", "p", "", "package used in proto3 (default the same as service)")
}

func generate() {
    // init db
    dbInstance := dbfactory.NewDB(consts.MYSQL)
    err := dbInstance.Init(datasource)
    if err != nil {
        log.Fatal("db init error ", err)
    }
    if len(tableNames) == 0 {
        tableNames = dbInstance.AllTableNames()
    }
    tables, err := dbInstance.GetTables(tableNames)
    if err != nil {
        log.Fatal(err)
    }
    
    // init protocol
    protocolType := format
    p := protocolfactory.NewProtocol(protocolType, pkg, goPackage)
    convertor := convertfactory.NewConverter(protocolType)
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
