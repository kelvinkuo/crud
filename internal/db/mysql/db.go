package mysql

import (
    "database/sql"
    "log"
    
    _ "github.com/go-sql-driver/mysql"
    "github.com/kelvinkuo/crud/internal/db"
)

type DB struct {
    dataSource string
    db         *sql.DB
    schema     string
}

func (d *DB) Init(datasource string) error {
    dbConn, err := sql.Open("mysql", datasource)
    if err != nil {
        log.Println(err)
        return err
    }
    
    d.dataSource = datasource
    d.db = dbConn
    d.schema, err = d.getSchema()
    if err != nil {
        log.Println(err)
        return err
    }
    
    return nil
}

func (d *DB) Close() error {
    return d.db.Close()
}

func (d *DB) AllTableNames() []string {
    q := "select TABLE_NAME from INFORMATION_SCHEMA.TABLES where TABLE_SCHEMA = ? order by TABLE_NAME;"
    
    rows, err := d.db.Query(q, d.schema)
    defer rows.Close()
    if err != nil {
        log.Fatal(err)
    }
    
    names := make([]string, 0)
    for rows.Next() {
        tableName := ""
        err = rows.Scan(&tableName)
        if err != nil {
            log.Fatal(err)
        }
        
        names = append(names, tableName)
    }
    
    return names
}

func (d *DB) getSchema() (string, error) {
    if d.schema != "" {
        return d.schema, nil
    }
    
    err := d.db.QueryRow("SELECT SCHEMA()").Scan(&d.schema)
    return d.schema, err
}

func (d *DB) GetTable(table string) (db.Table, error) {
    q := "select c.COLUMN_NAME, c.DATA_TYPE, c.COLUMN_TYPE, c.COLUMN_COMMENT, c.COLUMN_KEY, t.TABLE_COMMENT " +
        "from information_schema.COLUMNS as c " +
        "left join information_schema.TABLES as t on " +
        "c.TABLE_NAME = t.TABLE_NAME and c.TABLE_SCHEMA = t.TABLE_SCHEMA " +
        "where c.TABLE_SCHEMA = ? and c.TABLE_NAME = ? " +
        "ORDER BY c.TABLE_NAME, c.ORDINAL_POSITION"
    
    rows, err := d.db.Query(q, d.schema, table)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    t := newTable(table)
    
    for rows.Next() {
        col := column{}
        key := ""
        err = rows.Scan(&col.name, &col.dataType, &col.columnType, &col.comment, &key, &t.comment)
        if err != nil {
            log.Println(err)
            return nil, err
        }
        if key == "PRI" {
            col.isPrimary = true
        }
        t.addColumn(&col)
    }
    
    return t, nil
}

func (d *DB) GetTables(tables []string) ([]db.Table, error) {
    tableArr := make([]db.Table, 0)
    for _, t := range tables {
        table, err := d.GetTable(t)
        if err != nil {
            return nil, err
        }
        tableArr = append(tableArr, table)
    }
    
    return tableArr, nil
}
