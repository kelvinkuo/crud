package cmd

import (
    "os"
    "strings"
    
    "github.com/kelvinkuo/crud/internal/core"
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

supported database : mysql
supported format: proto3, zero api

example simple:
crud -f zeroapi --source "root:123456@tcp(127.0.0.1:3306)/shop" -s shop

example full:
crud -f proto3 --source "root:123456@tcp(127.0.0.1:3306)/shop" -m "add,delete,update,info,list,search" -c "created_at,updated_at,deleted_at" -t * -s shop -go_package shop -package shop

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
        core.Generate(datasource, format, service, pkg, goPackage, tableNames, ignoreCols, methods)
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
    rootCmd.Flags().StringVarP(&methodStr, "method", "m", "add,delete,update,info,list,search", "methods separated by \",\"")
    rootCmd.Flags().StringVarP(&ignoreColStr, "ignore_cols", "c", "create_at,create_time,created_at,update_at,update_time,updated_at",
        "columns ignored separated by \",\"")
    rootCmd.Flags().StringVarP(&tableStr, "table", "t", "*", "tables separated by \",\" or \"*\" for all tables")
    rootCmd.Flags().StringVarP(&service, "service", "s", "", "service name")
    rootCmd.Flags().StringVar(&goPackage, "go_package", "", "go_package used in proto3 (default the same as service)")
    rootCmd.Flags().StringVarP(&pkg, "package", "p", "", "package used in proto3 (default the same as service)")
}
